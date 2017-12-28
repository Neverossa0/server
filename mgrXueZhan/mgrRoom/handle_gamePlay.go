package mgrRoom

import (
	"server/mgrXueZhan/mgrCard"
	"server/mgrXueZhan/mgrSide"
	"server/pb"

	"github.com/name5566/leaf/log"
)

func (room *RoomInfo) gameStartHandler() {
	log.Debug("handle_gamePlay==>牌局开始")
	room.shuffle()
	room.setDealerId()
	room.dealStartCard()
	log.Debug("=======1")
	room.sendUpdateStartBattleRet()

}

func (room *RoomInfo) RoomExchangeCardHandler(exCardOidList []int32, pid int32) {

	sider := room.getSider(pid)
	sider.SetexchangeCards(exCardOidList)

	if room.checkRealPlayerActionOver(mgrSide.ProcessStatus_EXCHANGE_OVER) {

		room.sideMap.Range(func(k, v interface{}) bool {
			sider := v.(*mgrSide.SideInfo)
			if sider.GetIsRobot() {
				sider.SubmitExchangeCardByAI()
			}
			return true
		})

		room.exchangeCardProc()
	}
}

func (room *RoomInfo) RoomLackTypeHandler(lackType pb.CardType, pid int32) {

	sider := room.getSider(pid)
	sider.SetLackType(lackType)
	if room.checkRealPlayerActionOver(mgrSide.ProcessStatus_LACK_OVER) {

		room.sideMap.Range(func(k, v interface{}) bool {
			sider := v.(*mgrSide.SideInfo)
			if sider.GetIsRobot() {
				sider.SubmitLackTypeByAI()
			}
			return true
		})
		room.sendLackTypeRet()
		log.Debug("handle_gamePlay==>房间[%v]玩家定缺结束", room.roomId)

		room.newTurnWithNoDeal()
	}
}

func (room *RoomInfo) RoomActionRetHandler(card *mgrCard.CardInfo, proc pb.ProcType, pid int32) {

	sider := room.getSider(pid)
	if sider.GetProcess() == mgrSide.ProcessStatus_TURN_OVER {
		return
	}

	switch proc {
	case pb.ProcType_Proc_Discard:
		log.Debug("handle_gamePlay==>收到玩家[%v]行为请求==>出牌[%v]", pid, card.Id)
		sider.TackDiscard(card)

		afterCards := sider.ShowOwnAllcards()
		room.sendBroadcastProc(pid, proc, afterCards, card)

	case pb.ProcType_Proc_Hu:
		log.Debug("handle_gamePlay==>收到玩家[%v]行为请求==>胡牌[%v]", pid, card.Id)

		if room.curPid == pid {
			log.Debug("自摸胡牌结算")
		} else {
			log.Debug("其他玩家打牌胡牌结算")
		}
		room.curPid = pid
		room.sendTurnToNext(nil)

		sider.SetProcess(mgrSide.ProcessStatus_GAME_OVER)
		room.removeActivePlayers(pid)
		room.turnOverActivePlayers()

	case pb.ProcType_Proc_Pass:
		sider.SetProcess(mgrSide.ProcessStatus_TURN_OVER)
	}

	if !room.checkIsOtherWaiting(pid, mgrSide.ProcessStatus_WAITING_HU) {
		switch proc {
		case pb.ProcType_Proc_Peng:
			log.Debug("handle_gamePlay==>收到玩家[%v]行为请求==>碰牌[%v]", pid, card.Id)

			room.curPid = pid
			room.sendTurnToNext(nil)

			sider.TackPeng(card)
			afterCards := sider.ShowOwnAllcards()

			lastSider := room.getSider(card.PlayerId)
			lastSider.RemoveCardFromDisList(card.Oid)
			afterCards = append(afterCards, lastSider.ShowOwnAllcards()...)

			room.sendBroadcastProc(pid, proc, afterCards, card)

			room.newTurnWhenPeng()

		case pb.ProcType_Proc_Gang:
			log.Debug("handle_gamePlay==>收到玩家[%v]行为请求==>杠牌[%v]", pid, card.Id)

			room.curPid = pid
			room.sendTurnToNext(nil)

			sider.TackGang(card)
			afterCards := sider.ShowOwnAllcards()

			if card.PlayerId != pid {
				lastSider := room.getSider(card.PlayerId)
				lastSider.RemoveCardFromDisList(card.Oid)
				afterCards = append(afterCards, lastSider.ShowOwnAllcards()...)
			}

			room.sendBroadcastProc(pid, proc, afterCards, card)
			room.newTurnWithDeal()
		}
	}
	if len(room.getActivePlayers()) == 1 {
		room.sendGameOverInfo()
		return
	}
	if room.checkAllPlayerActionOver(mgrSide.ProcessStatus_TURN_OVER) {
		room.turnToNext()
		room.newTurnWithDeal()

	}

}
