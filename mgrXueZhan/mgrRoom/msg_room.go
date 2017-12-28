package mgrRoom

import (
	"server/mgrXueZhan/mgrCard"
	"server/mgrXueZhan/mgrSide"
	"server/pb"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
)

func (room *RoomInfo) broadcastRoom(msg interface{}) {
	for _, a := range room.getAgents() {
		a.WriteMsg(msg)
	}
}

func (room *RoomInfo) sendUpdateRoomMemberRet() {
	players := make([]*pb.PlayerInfo, 0)

	room.sideMap.Range(func(k, v interface{}) bool {
		sider := v.(*mgrSide.SideInfo)
		pbPlayer := sider.GetPlayer().TransToPbPlayerInfo()
		pbPlayer.Side = sider.GetSide().Enum()
		pbPlayer.IsOwner = proto.Bool(sider.GetIsOwner())
		players = append(players, pbPlayer)
		return true
	})
	data := &pb.GS2CUpdateRoomMember{
		Player: players,
	}
	room.broadcastRoom(data)
}

func (room *RoomInfo) sendUpdateStartBattleRet() {
	cardListSum := make([]*mgrCard.CardInfo, 0)

	room.sideMap.Range(func(k, v interface{}) bool {
		sider := v.(*mgrSide.SideInfo)
		cardListSum = append(cardListSum, sider.ShowOwnAllcards()...)
		return true
	})

	data := &pb.GS2CBattleStart{
		DealerId: proto.Int32(room.dealerId),
		CardList: mgrCard.TransToPBCardList(cardListSum),
	}

	room.broadcastRoom(data)
}

func (room *RoomInfo) sendExchangeCardRet(exchangeType pb.ExchangeType) {
	cardListSum := make([]*mgrCard.CardInfo, 0)

	room.sideMap.Range(func(k, v interface{}) bool {
		sider := v.(*mgrSide.SideInfo)
		cardListSum = append(cardListSum, sider.ShowOwnAllcards()...)
		return true
	})
	data := &pb.GS2CExchangeCardRet{
		Type:     exchangeType.Enum(),
		CardList: mgrCard.TransToPBCardList(cardListSum),
	}

	room.broadcastRoom(data)
}

func (room *RoomInfo) sendLackTypeRet() {
	typeList := make([]*pb.LackCard, 0)

	room.sideMap.Range(func(k, v interface{}) bool {
		pid := k.(int32)
		sider := v.(*mgrSide.SideInfo)

		tempData := &pb.LackCard{
			PlayerOID: proto.Int32(pid),
			Type:      sider.GetLackType().Enum(),
		}
		typeList = append(typeList, tempData)
		return true
	})

	data := &pb.GS2CSelectLackRet{
		LackCard: typeList,
	}
	room.broadcastRoom(data)
}

func (room *RoomInfo) sendTurnToNext(drawcard *mgrCard.CardInfo) {
	log.Debug("msg_room==>房间[%v]当前可行动玩家为[%v]", room.roomId, room.curPid)
	data := &pb.GS2CTurnToNext{
		PlayerOID: proto.Int32(room.curPid),
	}
	if drawcard != nil {
		data.DrawCard = drawcard.ToPBCard()
		log.Debug("msg_room==>玩家[%v]收到发牌[%v]", room.curPid, data.DrawCard)
	}
	room.broadcastRoom(data)

	for _, pid := range room.getActivePlayers() {
		sider := room.getSider(pid)
		sider.SetProcess(mgrSide.ProcessStatus_TURN_START)
	}

}

func (room *RoomInfo) sendBroadcastProc(procPlayer int32, proc pb.ProcType, cardList []*mgrCard.CardInfo, drawcard *mgrCard.CardInfo) {

	data := &pb.GS2CBroadcastProc{
		ProcPlayer: proto.Int32(procPlayer),
		ProcType:   proc.Enum(),
		CardList:   mgrCard.TransToPBCardList(cardList),
	}
	room.broadcastRoom(data)
	if proc == pb.ProcType_Proc_Discard {
		room.checkActionByDiscard(drawcard)
	}
}

func (room *RoomInfo) sendGameOverInfo() {

	log.Debug("msg_room==>反馈消息==>牌局已结束")
}
