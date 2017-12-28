package mgrSide

import (
	"server/mgrXueZhan/mgrCard"

	"github.com/name5566/leaf/log"
)

//============================****allCards管理****===============================
func (sider *SideInfo) SethandCards(card []*mgrCard.CardInfo, status mgrCard.CardStatus) {
	for _, c := range card {
		c.PlayerId = sider.player.GetPlayerId()
		c.Status = status
	}
	sider.Lock()
	sider.handCards = append(sider.handCards, card...)
	sider.Unlock()
}

func (sider *SideInfo) SetpengCards(drawcard *mgrCard.CardInfo, fromOther bool) {
	card := drawcard
	card.Status = mgrCard.CardStatus_Peng
	card.PlayerId = sider.player.GetPlayerId()
	card.FromOther = fromOther

	sider.Lock()
	sider.pengCards = append(sider.pengCards, card)
	sider.Unlock()
}

func (sider *SideInfo) SetgangCards(drawcard *mgrCard.CardInfo, fromOther bool) {
	card := drawcard
	card.Status = mgrCard.CardStatus_Gang
	card.PlayerId = sider.player.GetPlayerId()
	card.FromOther = fromOther

	sider.Lock()
	sider.gangCards = append(sider.gangCards, card)
	sider.Unlock()
}

func (sider *SideInfo) SetdiscCards(card *mgrCard.CardInfo) {
	card.Status = mgrCard.CardStatus_DisCard

	sider.Lock()
	sider.discCards = append(sider.discCards, card)
	sider.Unlock()
}

func (sider *SideInfo) SetexchangeCards(cardOidList []int32) {
	sider.Lock()
	defer sider.Unlock()
	for _, v := range cardOidList {
		for _, c := range sider.handCards {
			if v == c.Oid {
				c.Status = mgrCard.CardStatus_Exchanged
				break
			}
		}
	}

	sider.process = ProcessStatus_EXCHANGE_OVER
	log.Debug("mgr_sideCard==>玩家[%v]提交的换牌列表为[%v]", sider.GetPlayer().GetPlayerId(), cardOidList)
}

func (sider *SideInfo) RemovehandCardsByStatus(cardStatus mgrCard.CardStatus) []*mgrCard.CardInfo {
	sider.Lock()
	defer sider.Unlock()

	removedCardlist := make([]*mgrCard.CardInfo, 0)
	for i := 0; i < len(sider.handCards); i++ {
		if sider.handCards[i].Status == cardStatus {
			card := sider.handCards[i]
			sider.handCards = append(sider.handCards[0:i], sider.handCards[i+1:]...)
			removedCardlist = append(removedCardlist, card)
			i--
		}
	}

	return removedCardlist
}

func (sider *SideInfo) RemoveCardFromDisList(coid int32) {
	sider.Lock()
	defer sider.Unlock()

	for i := 0; i < len(sider.discCards); i++ {
		if sider.discCards[i].Oid == coid {
			sider.discCards = append(sider.discCards[0:i], sider.discCards[i+1:]...)
			i--
		}
	}

}

func (sider *SideInfo) sorthandCards() {
	sider.Lock()
	defer sider.Unlock()
	for _, c := range sider.handCards {
		if c.Status == mgrCard.CardStatus_Deal {
			c.Status = mgrCard.CardStatus_InHand
			return
		}
	}

}

//==============================****获取手牌信息****===============================
func (sider *SideInfo) GethandCards() []*mgrCard.CardInfo {
	sider.RLock()
	defer sider.RUnlock()
	return sider.handCards
}

func (sider *SideInfo) GetpengCards() []*mgrCard.CardInfo {
	sider.RLock()
	defer sider.RUnlock()
	return sider.pengCards
}

func (sider *SideInfo) GetgangCards() []*mgrCard.CardInfo {
	sider.RLock()
	defer sider.RUnlock()
	return sider.gangCards
}

func (sider *SideInfo) GetdiscCards() []*mgrCard.CardInfo {
	sider.RLock()
	defer sider.RUnlock()
	return sider.discCards
}

//===========================****showcards管理****===============================
func (sider *SideInfo) ShowOwnAllcards() []*mgrCard.CardInfo {
	afterCards := make([]*mgrCard.CardInfo, 0)
	afterCards = append(afterCards, sider.GethandCards()...)
	afterCards = append(afterCards, sider.GetpengCards()...)
	afterCards = append(afterCards, sider.GetgangCards()...)
	afterCards = append(afterCards, sider.GetdiscCards()...)

	return afterCards
}
