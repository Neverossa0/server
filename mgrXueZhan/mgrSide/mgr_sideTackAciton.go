package mgrSide

import (
	"server/mgrXueZhan/mgrCard"

	"github.com/name5566/leaf/log"
)

func (sider *SideInfo) TackDiscard(drawcard *mgrCard.CardInfo) {

	for _, c := range sider.GethandCards() {
		if c.Oid == drawcard.Oid {
			sider.SetdiscCards(c)
		}
	}
	sider.RemovehandCardsByStatus(mgrCard.CardStatus_DisCard)

	sider.sorthandCards()
	sider.SetProcess(ProcessStatus_TURN_OVER)

	log.Debug("玩家[%v]执行出牌[%v]", sider.GetPlayer().GetPlayerId(), drawcard.Id)
}

func (sider *SideInfo) TackPeng(drawcard *mgrCard.CardInfo) {
	card := drawcard

	sider.SetpengCards(card, true)

	count := 0
	for _, c := range sider.GethandCards() {
		if c.Id == card.Id {
			sider.SetpengCards(c, false)
			count++
		}
		if count == 2 {
			break
		}
	}

	sider.RemovehandCardsByStatus(mgrCard.CardStatus_Peng)
	log.Debug("玩家[%v]执行碰牌[%v]", sider.GetPlayer().GetPlayerId(), drawcard.Id)
}

func (sider *SideInfo) TackGang(drawcard *mgrCard.CardInfo) {

	card := drawcard

	if card.PlayerId != sider.GetPlayer().GetPlayerId() {
		sider.SetgangCards(card, true)

	} else {
		sider.SetgangCards(card, false)

		for _, c := range sider.GetpengCards() {
			if c.Id == card.Id {
				sider.SetgangCards(card, c.FromOther)
			}
		}
	}

	for _, c := range sider.GethandCards() {
		if c.Id == card.Id {
			sider.SetgangCards(card, false)
		}
	}

	sider.RemovehandCardsByStatus(mgrCard.CardStatus_Gang)
	log.Debug("玩家[%v]执行杠牌[%v]", sider.GetPlayer().GetPlayerId(), card.Id)
}
