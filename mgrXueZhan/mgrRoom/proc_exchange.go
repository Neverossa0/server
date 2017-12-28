package mgrRoom

import (
	"server/mgrXueZhan/mgrCard"
	"server/mgrXueZhan/mgrSide"
	"server/pb"

	"math/rand"
	"time"

	"github.com/name5566/leaf/log"
)

func (room *RoomInfo) exchangeCardProc() {
	log.Debug("proc_exchange==>房间[%v]处理玩家间换牌==>proc_exchangecard", room.roomId)

	exchangeAllMap := make(map[string][]*mgrCard.CardInfo)

	room.sideMap.Range(func(k, v interface{}) bool {
		sider := v.(*mgrSide.SideInfo)
		exCardList := sider.RemovehandCardsByStatus(mgrCard.CardStatus_Exchanged)
		exchangeAllMap[sider.GetSide().String()] = exCardList
		return true
	})
	exchangeType := getExchangeType()

	room.sideMap.Range(func(k, v interface{}) bool {
		sider := v.(*mgrSide.SideInfo)
		fromSideStr := getExchangeCardSide(exchangeType, sider.GetSide())
		sider.SethandCards(exchangeAllMap[fromSideStr], mgrCard.CardStatus_InHand)
		return true
	})
	room.sendExchangeCardRet(exchangeType)
}

func getExchangeType() pb.ExchangeType {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(3)
	if rnd == 0 {
		log.Debug("proc_exchange==>换牌方式为==>[顺时针]")
		return pb.ExchangeType_ClockWise
	} else if rnd == 1 {
		log.Debug("proc_exchange==>换牌方式为==>[逆时针]")
		return pb.ExchangeType_AntiClock
	} else {
		log.Debug("proc_exchange==>换牌方式为==>[对换]")
		return pb.ExchangeType_Opposite
	}
}

func getExchangeCardSide(exchangeType pb.ExchangeType, curSide pb.MahjonSide) string {
	var fromSideStr string
	switch exchangeType {
	case pb.ExchangeType_AntiClock: //逆时针
		fromsideId := pb.MahjonSide_value[curSide.String()] - 1
		if fromsideId == 1 {
			fromsideId = 5
		}
		fromSideStr = pb.MahjonSide_name[fromsideId]
	case pb.ExchangeType_ClockWise: //顺时针
		fromsideId := pb.MahjonSide_value[curSide.String()] + 1
		if fromsideId == 6 {
			fromsideId = 2
		}
		fromSideStr = pb.MahjonSide_name[fromsideId]
	case pb.ExchangeType_Opposite:
		fromsideId := pb.MahjonSide_value[curSide.String()] + 2
		if fromsideId == 7 {
			fromsideId = 3
		} else if fromsideId == 6 {
			fromsideId = 2
		}
		fromSideStr = pb.MahjonSide_name[fromsideId]
	}
	return fromSideStr
}
