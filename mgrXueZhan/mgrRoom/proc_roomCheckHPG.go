package mgrRoom

import (
	"server/mgrXueZhan/mgrCard"
	"server/pb"

	"github.com/name5566/leaf/log"
)

func (room *RoomInfo) checkActionByDeal() {

	sider := room.getSider(room.curPid)
	procList := sider.CheckHPGByOwn()
	log.Debug("proc_roomCheckHpg==>发牌时玩家可进行行为列表=%v", procList)
	if sider.GetIsRobot() {
		proc, drawcard := sider.SubmitAcitonWhenDealByAI(procList)
		room.RoomActionRetHandler(drawcard, proc, sider.GetPlayer().GetPlayerId())
	} else {
		sider.SendActionListByDeal(procList)
	}
}

func (room *RoomInfo) checkActionByPeng() {
	sider := room.getSider(room.curPid)
	procList := make([]pb.ProcType, 0)
	procList = append(procList, pb.ProcType_Proc_Discard)
	if sider.GetIsRobot() {
		proc, drawcard := sider.SubmitAcitonWhenDealByAI(procList)
		room.RoomActionRetHandler(drawcard, proc, sider.GetPlayer().GetPlayerId())
	} else {
		sider.SendActionListByDeal(procList)
	}
}

func (room *RoomInfo) checkActionByDiscard(drawcard *mgrCard.CardInfo) {

	for _, pid := range room.getActivePlayers() {
		sider := room.getSider(pid)
		if sider.GetPlayer().GetPlayerId() != room.curPid {
			procList := sider.CheckHPGByOther(drawcard)
			if sider.GetIsRobot() {
				proc := sider.SubmitAcitonAfterDiscardByAI(procList, drawcard)
				room.RoomActionRetHandler(drawcard, proc, sider.GetPlayer().GetPlayerId())
			} else {
				sider.SendActionListByDiscard(procList, drawcard)
			}
		}
	}
}
