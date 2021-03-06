package internal

import (
	"reflect"

	"server/mgrPlayer"
	"server/pb"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&pb.C2GSLogin{}, recvC2GSLogin)
}

func recvC2GSLogin(args []interface{}) {
	m := args[0].(*pb.C2GSLogin)
	a := args[1].(gate.Agent)

	mgrPlayer.LoginHandler(m.GetAccount(), m.GetPassword(), a)
}
