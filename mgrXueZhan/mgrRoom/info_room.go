package mgrRoom

import (
	"server/mgrXueZhan/mgrCard"
	"server/mgrXueZhan/mgrSide"

	"math/rand"
	"sync"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type RoomInfo struct {
	sync.RWMutex

	gameType string
	roomId   string

	cardWall  []*mgrCard.CardInfo
	cardIndex int

	sideMap sync.Map //pid:SideInfo

	actPlayers []int32 //pids
	agents     []gate.Agent

	playerNum int
	dealerId  int32
	curPid    int32
}

func NewRoom(rid string) *RoomInfo {
	newRoom := &RoomInfo{
		gameType: "血战到底",
		roomId:   rid,

		cardWall:  mgrCard.LoadInitCards(),
		cardIndex: 0,

		actPlayers: make([]int32, 0), //pids
		agents:     make([]gate.Agent, 0),

		playerNum: 0,
		dealerId:  0,
		curPid:    0,
	}

	defer log.Debug("info_room==>初始化新房间[%v]成功", rid)
	return newRoom
}

//=============================****房间信息设置****================================
func (room *RoomInfo) setAgents(a gate.Agent) {
	room.Lock()
	defer room.Unlock()
	room.agents = append(room.agents, a)
}

func (room *RoomInfo) setActivePlayers() {
	room.Lock()
	defer room.Unlock()
	room.sideMap.Range(func(k, v interface{}) bool {
		room.actPlayers = append(room.actPlayers, v.(*mgrSide.SideInfo).GetPlayer().GetPlayerId())
		return true
	})
}

func (room *RoomInfo) removeActivePlayers(playerId int32) {

	for i, pid := range room.getActivePlayers() {
		if pid == playerId {
			room.Lock()
			room.actPlayers = append(room.actPlayers[:i], room.actPlayers[i+1:]...)
			room.Unlock()
		}
	}
}

func (room *RoomInfo) turnOverActivePlayers() {
	for _, pid := range room.getActivePlayers() {
		sider := room.getSider(pid)
		if sider.GetProcess() != mgrSide.ProcessStatus_WAITING_HU {
			sider.SetProcess(mgrSide.ProcessStatus_TURN_OVER)
		}
	}
}

func (room *RoomInfo) setDealerId() {
	log.Debug("info_room==>测试用setDealerId")

	playerOidList := make([]int32, 0)

	room.sideMap.Range(func(pid, v interface{}) bool {
		playerOidList = append(playerOidList, pid.(int32))
		return true
	})
	count := len(playerOidList)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(count)

	room.dealerId = playerOidList[index]

	room.dealerId = int32(911) //测试设置玩家为庄家时打开

	room.curPid = room.dealerId

}

//=============================****房间信息获取****================================
func (room *RoomInfo) getAgents() []gate.Agent {
	room.RLock()
	defer room.RUnlock()
	return room.agents
}

func (room *RoomInfo) getActivePlayers() []int32 {
	room.RLock()
	defer room.RUnlock()
	return room.actPlayers
}

func (room *RoomInfo) getDealCards(startIndex int, endIndex int) []*mgrCard.CardInfo {
	room.RLock()
	defer room.RUnlock()
	cardSlices := room.cardWall[startIndex:endIndex]
	return cardSlices

}

func (room *RoomInfo) getSider(pid int32) *mgrSide.SideInfo {
	sider, ex := room.sideMap.Load(pid)
	if !ex {
		log.Debug("info_room==>房间[%v]里不存在玩家[%V],或玩家未正确注册到side", room.roomId, pid)
		return nil
	}
	return sider.(*mgrSide.SideInfo)
}

//=============================****房间状态check****================================
func (room *RoomInfo) checkRoomReady() {
	if room.playerNum == 4 {
		log.Debug("info_room==>房间中玩家坐满，准备开始游戏")
		room.setActivePlayers()
		room.gameStartHandler()
	}
}

func (room *RoomInfo) checkRealPlayerActionOver(proc mgrSide.ProcessStatus) bool {

	for _, pid := range room.getActivePlayers() {
		sider := room.getSider(pid)
		if false == sider.GetIsRobot() && proc != sider.GetProcess() {
			return false
		}
	}
	return true
}

func (room *RoomInfo) checkAllPlayerActionOver(proc mgrSide.ProcessStatus) bool {
	log.Debug("info_room==>当前活跃的sider数量为[%v]", len(room.getActivePlayers()))

	for _, pid := range room.getActivePlayers() {
		sider := room.getSider(pid)
		if proc != sider.GetProcess() {
			return false
		}
	}
	return true
}

func (room *RoomInfo) checkIsOtherWaiting(pid int32, proc mgrSide.ProcessStatus) bool {

	for _, pid := range room.getActivePlayers() {
		sider := room.getSider(pid)
		if sider.GetPlayer().GetPlayerId() != pid && proc == sider.GetProcess() {
			return true
		}
	}
	return false
}
