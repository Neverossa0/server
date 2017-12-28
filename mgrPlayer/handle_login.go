package mgrPlayer

import (
	"server/pb"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func LoginHandler(acc string, psw string, a gate.Agent) {
	errCode, player := loginProc(acc, psw, a)
	sendLoginRet(errCode, player, a)
}

func loginProc(acc string, psw string, a gate.Agent) (*pb.GS2CLoginRet_ErrorCode, *PlayerInfo) {
	var errorCode *pb.GS2CLoginRet_ErrorCode

	pid, ex := getTestaccInfo(acc, psw)
	if !ex {
		errorCode = pb.GS2CLoginRet_PASSWORD_ERROR.Enum()
		return errorCode, nil
	}

	loginPlayer := getTestuserInfo(pid).transToPlayerInfo()
	ok := bindAgentPlayer(a, loginPlayer)
	if !ok {
		errorCode = pb.GS2CLoginRet_FAIL.Enum()
		return errorCode, nil
	}

	errorCode = pb.GS2CLoginRet_SUCCESS.Enum()
	defer log.Debug("handle_login==>玩家[%v]登录成功", loginPlayer.oid)

	return errorCode, loginPlayer

}

//============================****发送登录反馈信息****==============================
func sendLoginRet(errCode *pb.GS2CLoginRet_ErrorCode, player *PlayerInfo, a gate.Agent) {
	playerInfo := player.TransToPbPlayerInfo()
	data := &pb.GS2CLoginRet{
		ErrorCode:  errCode,
		PlayerInfo: playerInfo,
	}
	a.WriteMsg(data)
}
