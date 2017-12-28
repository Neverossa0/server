package mgrPlayer

import (
	"time"
)

type accInfo struct {
	acc       string `xorm:"pk"`
	psw       string
	pid       int32
	version   int       `xorm:"version"`
	updatedAt time.Time `xorm:"updated"`
}

type userInfo struct {
	oid       int32 `xorm:"pk autoincr"`
	nickName  string
	headIcon  string
	gold      int32
	diamond   int32
	roomcard  int32
	version   int       `xorm:"version"`
	updatedAt time.Time `xorm:"updated"`
}

//===========================****玩家信息类型转换****==============================
func (user *userInfo) transToPlayerInfo() *PlayerInfo {
	player := &PlayerInfo{
		oid:      user.oid,
		nickName: user.nickName,
		headIcon: user.headIcon,
		gold:     user.gold,
		diamond:  user.diamond,
		roomcard: user.roomcard,
	}
	return player
}
