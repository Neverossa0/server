package mgrRoom

import (
	"server/mgrXueZhan/mgrCard"

	"github.com/name5566/leaf/log"

	"math/rand"
	"time"
)

func (room *RoomInfo) shuffle() {
	room.Lock()
	defer room.Unlock()
	k := len(room.cardWall)
	r_seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(room.cardWall); i++ {
		x := r_seed.Intn(k)

		t := room.cardWall[x].Id
		room.cardWall[x].Id = room.cardWall[k-1].Id
		room.cardWall[k-1].Id = t
		k--
	}
	log.Debug("mgr_roomCard==>房间[%v]洗牌完毕", room.roomId)
}

func (room *RoomInfo) dealStartCard() {

	for i := 0; i < 4; i++ {
		room.dealcards(room.curPid, mgrCard.CardStatus_InHand, 13)
		room.turnToNext()
	}
	room.dealcards(room.dealerId, mgrCard.CardStatus_InHand, 1)
}

func (room *RoomInfo) dealcards(playerId int32, cardstatus mgrCard.CardStatus, num int) []*mgrCard.CardInfo {
	defer log.Debug("mgr_roomCard==>发牌给玩家[%v],共[%v]张", playerId, num)
	startId := room.cardIndex
	endId := startId + num
	if endId > 108 {
		log.Debug("mgr_roomCard==>卡牌已经发完:[%v]", startId)
		return nil
	}

	cardSlices := room.getDealCards(startId, endId)

	for _, v := range cardSlices {
		v.PlayerId = playerId
		v.Status = cardstatus

	}

	sider := room.getSider(playerId)
	sider.SethandCards(cardSlices, cardstatus)
	room.cardIndex = endId
	return cardSlices
}
