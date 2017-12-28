package mgrRoom

import (
	"server/mgrPlayer"
	"server/mgrXueZhan/mgrSide"
	"server/pb"

	"math/rand"
	"time"

	"github.com/name5566/leaf/log"
)

func (room *RoomInfo) AddNewPlayerIntoRoom(joinPlayer *mgrPlayer.PlayerInfo) bool {

	if room.playerNum >= 4 {
		return false
	}

	isowner := (len(room.agents) == 0)
	isrobot := true
	agent := joinPlayer.GetAgent()
	if agent != nil {
		room.setAgents(agent)
		isrobot = false
	}

	joinPlayer.SetRoomId(room.roomId)

	newPid := joinPlayer.GetPlayerId()
	newSider := mgrSide.RegSideByPlayer(isrobot, isowner, room.giveSide(), joinPlayer)
	room.sideMap.Store(newPid, newSider)
	room.playerNum++
	log.Debug("mgr_roomPlayer==>玩家[%v]进入房间[%v]成功", newPid, room.roomId)

	room.sendUpdateRoomMemberRet()

	defer room.checkRoomReady()

	return true
}

func (room *RoomInfo) AddRobotToRoom(num int) {
	for i := 0; i < num; i++ {

		pid := int32(20001 + i)
		newrobot := mgrPlayer.GetRobotById(pid)

		room.AddNewPlayerIntoRoom(newrobot)
	}
}

//===========================****房间内座位管理****===============================
func (room *RoomInfo) giveSide() pb.MahjonSide {
	leftSideList := room.getLeftSideList()
	if len(leftSideList) == 0 {
		return pb.MahjonSide_DEFAULT
	}
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(len(leftSideList))
	return leftSideList[rnd]
}

func (room *RoomInfo) getLeftSideList() []pb.MahjonSide {
	sideList := []pb.MahjonSide{pb.MahjonSide_EAST, pb.MahjonSide_SOUTH, pb.MahjonSide_WEST, pb.MahjonSide_NORTH}
	room.sideMap.Range(func(k, v interface{}) bool {
		sider := v.(*mgrSide.SideInfo)
		for i, side := range sideList {
			if side == sider.GetSide() {
				sideList = append(sideList[:i], sideList[i+1:]...)
				break
			}
		}
		return true
	})
	return sideList
}
