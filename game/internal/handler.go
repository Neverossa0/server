package internal

import (
	"reflect"
	"server/mgrPlayer"
	"server/mgrXueZhan"
	"server/pb"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&pb.C2GSEnterGame{}, recvC2GSEnterGame)
	handler(&pb.C2GSExchangeCard{}, recvC2GSExchangeCard)
	handler(&pb.C2GSSelectLack{}, recvC2GSSelectLack)
	handler(&pb.C2GSInterruptActionRet{}, recvC2GSInterruptActionRet)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func recvC2GSEnterGame(args []interface{}) {
	log.Debug("收到请求==>recvC2GSEnterGame")

	m := args[0].(*pb.C2GSEnterGame)
	a := args[1].(gate.Agent)

	switch m.GetType() {
	case pb.GameType_XueZhan:
		mgrXueZhan.EnterGameHandler(m.GetMode(), m.GetRoomId(), a)
	}
}

func recvC2GSExchangeCard(args []interface{}) {
	log.Debug("收到请求==>recvC2GSExchangeCard")

	m := args[0].(*pb.C2GSExchangeCard)
	a := args[1].(gate.Agent)

	mgrXueZhan.ExchangeCardsHandler(m.GetCardOIDList(), mgrPlayer.GetPlayerByAgent(a))
}

func recvC2GSSelectLack(args []interface{}) {
	log.Debug("收到请求==>recvC2GSSelectLack")

	m := args[0].(*pb.C2GSSelectLack)
	a := args[1].(gate.Agent)

	mgrXueZhan.LackTypeHandler(m.GetType(), mgrPlayer.GetPlayerByAgent(a))

}

func recvC2GSInterruptActionRet(args []interface{}) {
	log.Debug("收到请求==>recvC2GSInterruptActionRet")
	m := args[0].(*pb.C2GSInterruptActionRet)
	a := args[1].(gate.Agent)

	mgrXueZhan.ActionRetHandler(m.GetDrawCard(), m.GetProcType(), mgrPlayer.GetPlayerByAgent(a))

}
