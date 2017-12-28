package mgrPlayer

import (
	"strconv"
	"sync"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var a2pMap sync.Map //agent:player

//============================****登录玩家路由绑定****=============================
func bindAgentPlayer(a gate.Agent, newplayer *PlayerInfo) bool {
	log.Debug("mgr_Player==>绑定玩家路由:玩家id=%v", newplayer.oid)
	_, ex := a2pMap.LoadOrStore(a, newplayer)
	if ex {
		log.Debug("玩家[%v]重复登录")
		return false
	}
	newplayer.agent = a
	return true
}

//==============================****玩家信息管理****===============================
func GetPlayerByAgent(a gate.Agent) *PlayerInfo {
	player, ok := a2pMap.Load(a)
	if !ok {
		log.Debug("mgr_Player==>目标路由未连接")
		return nil
	}
	return player.(*PlayerInfo)
}

func GetRobotById(pid int32) *PlayerInfo {
	log.Debug("mgr_Player==>创建机器人[%v]", pid)
	newRobot := &PlayerInfo{
		oid:      pid,
		nickName: "robot" + strconv.Itoa(int(pid)),
		headIcon: "",
		gold:     8888,
		diamond:  20,
	}
	return newRobot
}
