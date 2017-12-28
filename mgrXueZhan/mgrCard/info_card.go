package mgrCard

import (
	"server/pb"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
)

type CardInfo struct {
	Oid       int32
	Id        int32
	PlayerId  int32
	Status    CardStatus
	FromOther bool
}

type CardStatus int32

const (
	CardStatus_Wall      = 1
	CardStatus_InHand    = 2
	CardStatus_Peng      = 3
	CardStatus_Gang      = 4
	CardStatus_DisCard   = 5
	CardStatus_Deal      = 6
	CardStatus_Hu        = 7
	CardStatus_Exchanged = 8
)

//================================****  ****====================================
func LoadInitCards() []*CardInfo {
	maxCount := 108
	cardWall := make([]*CardInfo, 0)
	id := int32(0)
	for i := 0; i < maxCount; i++ {
		card := &CardInfo{}
		card.Oid = int32(i)
		if i%4 == 0 {
			id++
			if id%10 == 0 {
				id++
			}
		}
		card.Id = id
		card.Status = CardStatus_Wall

		cardWall = append(cardWall, card)
	}
	log.Debug("装载全部卡牌完毕, 卡牌数量=%v", len(cardWall))
	return cardWall
}

//=============================****卡牌信息管理****===============================
func TransToPBCardList(cardList []*CardInfo) []*pb.CardInfo {
	pbCardList := make([]*pb.CardInfo, 0)
	for _, card := range cardList {
		pbCard := card.ToPBCard()
		pbCardList = append(pbCardList, pbCard)
	}
	return pbCardList
}
func TransToIntList(cardList []*CardInfo) []int {
	intCardList := make([]int, 0)
	for _, card := range cardList {
		intCard := int(card.Id)
		intCardList = append(intCardList, intCard)
	}
	return intCardList
}

func (card *CardInfo) ToPBCard() *pb.CardInfo {
	pbCardInfo := &pb.CardInfo{
		PlayerOID: proto.Int32(card.PlayerId),
		OID:       proto.Int32(card.Oid),
		ID:        proto.Int32(card.Id),
		Status:    card.transCardStatus(),
		FromOther: proto.Bool(false),
	}
	return pbCardInfo
}

func (card *CardInfo) transCardStatus() *pb.CardStatus {
	switch card.Status {
	case CardStatus_Wall:
		return pb.CardStatus_Wall.Enum()
	case CardStatus_InHand:
		return pb.CardStatus_InHand.Enum()
	case CardStatus_Deal:
		return pb.CardStatus_Deal.Enum()
	case CardStatus_Exchanged:
		log.Debug("卡牌服务器状态为[%v]:客户端使用[CardStatus_InHand]", card.Status)
		return pb.CardStatus_InHand.Enum()
	case CardStatus_DisCard:
		return pb.CardStatus_Dis.Enum()
	case CardStatus_Peng:
		return pb.CardStatus_P.Enum()
	case CardStatus_Gang:
		return pb.CardStatus_G.Enum()
	default:
		log.Debug("卡牌状态转换缺失类型:%v", card.Status)
	}
	return nil
}
