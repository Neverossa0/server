package mgrSide

import (
	"server/mgrXueZhan/AI"
	"server/mgrXueZhan/mgrCard"
	"server/pb"

	"time"

	"github.com/name5566/leaf/log"
)

func (sider *SideInfo) SubmitExchangeCardByAI() {
	exCardOidList := AI.SelectexchangeCard(sider.separateCardByType())
	sider.SetexchangeCards(exCardOidList)
}

func (sider *SideInfo) SubmitLackTypeByAI() {
	lackId := AI.SelectlackType(sider.separateCardByType())
	var lacktype pb.CardType
	switch lackId {
	case 0:
		lacktype = pb.CardType_Wan
	case 1:
		lacktype = pb.CardType_Tiao
	case 2:
		lacktype = pb.CardType_Tong
	}
	sider.SetLackType(lacktype)
}

func (sider *SideInfo) separateCardByType() ([]int32, []int32, []int32) {
	wanList := make([]int32, 0)
	tiaoList := make([]int32, 0)
	tongList := make([]int32, 0)

	for _, card := range sider.handCards {
		if card.Id > 0 && card.Id < 10 {
			wanList = append(wanList, card.Oid)
		} else if card.Id > 10 && card.Id < 20 {
			tiaoList = append(tiaoList, card.Oid)
		} else if card.Id > 20 && card.Id < 30 {
			tongList = append(tongList, card.Oid)
		}
	}
	return wanList, tiaoList, tongList
}

func (sider *SideInfo) SubmitAcitonWhenDealByAI(procList []pb.ProcType) (pb.ProcType, *mgrCard.CardInfo) {
	timer := time.NewTimer(time.Second * 1)
	<-timer.C

	proc := procList[0]
	var dcard *mgrCard.CardInfo

	switch proc {
	case pb.ProcType_Proc_Discard:
		for _, card := range sider.GethandCards() {
			if sider.checkCardLackType(card.Id) {
				dcard = card
				break
			}
		}
		if dcard == nil {
			dcardId := AI.GetDiscardID(mgrCard.TransToIntList(sider.GethandCards()), AI.DiscardLevel1)
			dcard = sider.GetCardById(dcardId)
		}
	case pb.ProcType_Proc_Gang:
		for _, card := range sider.GethandCards() {
			for _, card2 := range sider.GetpengCards() {
				if card.Id == card2.Id {
					log.Debug("自摸明杠牌[%v]", card.Id)
					return proc, card
				}

			}
			count := 0
			for _, card2 := range sider.GethandCards() {
				if card.Id == card2.Id {
					count++
				}
				if count == 4 {
					log.Debug("自摸暗杠[%v]", card.Id)
					return proc, card
				}
			}
		}
	case pb.ProcType_Proc_Hu:
		for i, card := range sider.GethandCards() {
			if card.Status == mgrCard.CardStatus_Deal {
				return proc, card
			}
			if i == len(sider.GethandCards())-1 {
				return proc, card
			}

		}

	}

	return proc, dcard

}

func (sider *SideInfo) SubmitAcitonAfterDiscardByAI(procList []pb.ProcType, card *mgrCard.CardInfo) pb.ProcType {

	proc := procList[0]
	return proc
}
