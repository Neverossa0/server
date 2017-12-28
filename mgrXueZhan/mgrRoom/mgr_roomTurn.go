package mgrRoom

import (
	"server/mgrXueZhan/mgrCard"
	"server/mgrXueZhan/mgrSide"
	"server/pb"
)

func (room *RoomInfo) turnToNext() {
	curSider := room.getSider(room.curPid)
	curSide := curSider.GetSide()
	nextsideId := pb.MahjonSide_value[curSide.String()] + 1
	if nextsideId == 6 {
		nextsideId = 2
	}
	nextSide := pb.MahjonSide_name[nextsideId]

	room.sideMap.Range(func(k, v interface{}) bool {
		pid := k.(int32)
		sider := v.(*mgrSide.SideInfo)
		if nextSide == sider.GetSide().String() {
			room.curPid = pid
			return false
		}

		return true
	})

	newCurSider := room.getSider(room.curPid)
	if newCurSider.GetProcess() == mgrSide.ProcessStatus_GAME_OVER {
		room.turnToNext()
	}
}

func (room *RoomInfo) newTurnWithDeal() {
	dealcards := room.dealcards(room.curPid, mgrCard.CardStatus_Deal, 1)
	if dealcards == nil {
		room.sendGameOverInfo()
		return
	}
	room.sendTurnToNext(dealcards[0])
	room.checkActionByDeal()
}

func (room *RoomInfo) newTurnWithNoDeal() {

	room.sendTurnToNext(nil)
	room.checkActionByDeal()
}

func (room *RoomInfo) newTurnWhenPeng() {

	room.sendTurnToNext(nil)
	room.checkActionByPeng()

}
