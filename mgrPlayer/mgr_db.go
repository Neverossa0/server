package mgrPlayer

import (
	"github.com/name5566/leaf/log"
)

//============================****获取测试用户信息****=============================

func getTestaccInfo(acc string, psw string) (int32, bool) {
	log.Debug("无数据库test信息==>[acc=%v][psw=%v]", acc, psw)
	return int32(911), true
}

func getTestuserInfo(pid int32) *userInfo {
	user := &userInfo{
		oid:      pid,
		nickName: "test911",
		headIcon: "",
		gold:     int32(0),
		diamond:  int32(0),
	}

	return user
}
