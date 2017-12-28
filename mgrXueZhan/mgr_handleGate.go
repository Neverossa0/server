package mgrXueZhan

import (
	"server/mgrPlayer"
	"server/mgrXueZhan/mgrCard"

	"server/pb"

	//	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

//==========================****玩家请求handler处理****============================
func EnterGameHandler(enterMode pb.EnterMode, roomId string, a gate.Agent) {

	player := mgrPlayer.GetPlayerByAgent(a)
	if player.GetRoomId() != "" {
		sendEnterGameRet(pb.GS2CEnterGameRet_FAIL.Enum(), pb.GameType_XueZhan.Enum(), "", player.GetAgent())
		return
	}

	switch enterMode {
	case pb.EnterMode_CreateRoom:
		log.Debug("mgr_handleGate==>玩家[%v]请求创建房间", player.GetPlayerId())
		regNewRoom(player)
	case pb.EnterMode_JoinRoom:
		log.Debug("mgr_handleGate==>请求加入房间[%v]", roomId)
		joinRoom(player, roomId)
	case pb.EnterMode_QuickEnter:
		log.Debug("mgr_handleGate==>请求快速开始游戏")
	}
}

func ExchangeCardsHandler(exCardList []int32, player *mgrPlayer.PlayerInfo) {

	roomid := player.GetRoomId()
	room := getRoomById(roomid)

	room.RoomExchangeCardHandler(exCardList, player.GetPlayerId())
}

func LackTypeHandler(cardType pb.CardType, player *mgrPlayer.PlayerInfo) {

	roomid := player.GetRoomId()
	room := getRoomById(roomid)

	room.RoomLackTypeHandler(cardType, player.GetPlayerId())
}

func ActionRetHandler(drawcard *pb.CardInfo, procType pb.ProcType, player *mgrPlayer.PlayerInfo) {

	roomid := player.GetRoomId()
	room := getRoomById(roomid)
	log.Debug("房间[%v]收到玩家[%v]提交操作行为==>[%v]", roomid, player.GetPlayerId(), procType)

	card := &mgrCard.CardInfo{
		Oid:       drawcard.GetOID(),
		Id:        drawcard.GetID(),
		PlayerId:  drawcard.GetPlayerOID(),
		FromOther: drawcard.GetFromOther(),
	}

	room.RoomActionRetHandler(card, procType, player.GetPlayerId())
}
