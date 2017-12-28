package mgrSide

import (
	"server/mgrPlayer"
	"server/mgrXueZhan/mgrCard"
	"server/pb"

	"sync"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type SideInfo struct {
	sync.RWMutex
	side pb.MahjonSide

	isOwner bool
	isRobot bool

	isHu     bool
	fan      int32
	lackType pb.CardType

	agent  gate.Agent
	player *mgrPlayer.PlayerInfo

	handCards []*mgrCard.CardInfo
	discCards []*mgrCard.CardInfo
	pengCards []*mgrCard.CardInfo
	gangCards []*mgrCard.CardInfo

	process ProcessStatus
}

type ProcessStatus int32

const (
	ProcessStatus_DEFAULT       ProcessStatus = 1
	ProcessStatus_EXCHANGE_OVER ProcessStatus = 2
	ProcessStatus_LACK_OVER     ProcessStatus = 3
	ProcessStatus_TURN_START    ProcessStatus = 4
	ProcessStatus_TURN_OVER     ProcessStatus = 5
	ProcessStatus_WAITING_HU    ProcessStatus = 6
	ProcessStatus_GAME_OVER     ProcessStatus = 7
)

func RegSideByPlayer(isrobot bool, isowner bool, side pb.MahjonSide, pyr *mgrPlayer.PlayerInfo) *SideInfo {
	log.Debug("info_side==>注册玩家[%v]side信息", pyr.GetPlayerId())
	sideinfo := &SideInfo{
		side: side,

		isOwner: isowner,
		isRobot: isrobot,

		isHu: false,
		fan:  0,

		agent:  pyr.GetAgent(),
		player: pyr,

		handCards: make([]*mgrCard.CardInfo, 0),
		discCards: make([]*mgrCard.CardInfo, 0),
		pengCards: make([]*mgrCard.CardInfo, 0),
		gangCards: make([]*mgrCard.CardInfo, 0),
	}

	return sideinfo
}

//========================****座位信息管理****====================================

func (sider *SideInfo) SetProcess(proc ProcessStatus) {
	sider.process = proc
}

func (sider *SideInfo) SetLackType(lacktype pb.CardType) {

	sider.lackType = lacktype
	log.Debug("info_side==>玩家[%v]提交的换牌列表为[%v]", sider.GetPlayer().GetPlayerId(), lacktype)
	sider.process = ProcessStatus_LACK_OVER
}

func (sider *SideInfo) GetPlayer() *mgrPlayer.PlayerInfo {
	return sider.player
}

func (sider *SideInfo) GetSide() pb.MahjonSide {
	return sider.side
}

func (sider *SideInfo) GetLackType() pb.CardType {
	return sider.lackType
}

func (sider *SideInfo) GetIsOwner() bool {
	return sider.isOwner
}

func (sider *SideInfo) GetIsHu() bool {
	return sider.isHu
}

func (sider *SideInfo) GetFan() int32 {
	return sider.fan
}

func (sider *SideInfo) GetIsRobot() bool {
	return sider.isRobot
}

func (sider *SideInfo) GetProcess() ProcessStatus {
	return sider.process
}

func (sider *SideInfo) GetCardById(cardId int32) *mgrCard.CardInfo {
	for _, card := range sider.GethandCards() {
		if card.Id == cardId {
			return card
		}
	}
	return nil
}

//==========================****checkLackType管理****============================

func (sider *SideInfo) checkCardLackType(drawcid int32) bool {
	var cardLackType pb.CardType
	if drawcid >= 1 && drawcid <= 9 {
		cardLackType = pb.CardType_Wan
	} else if drawcid >= 11 && drawcid <= 19 {
		cardLackType = pb.CardType_Tiao
	} else {
		cardLackType = pb.CardType_Tong
	}

	return sider.lackType == cardLackType
}

func (sider *SideInfo) checkhandCardsLackType() bool {
	for _, card := range sider.GethandCards() {
		if sider.checkCardLackType(card.Id) {
			return true
		}
	}

	return false
}
