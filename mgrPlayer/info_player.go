package mgrPlayer

import (
	"server/pb"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type PlayerInfo struct {
	oid      int32
	nickName string
	headIcon string
	gold     int32
	diamond  int32
	roomcard int32

	roomId string
	agent  gate.Agent
}

//=============================****玩家信息设置****================================
func (player *PlayerInfo) SetRoomId(rid string) {
	player.roomId = rid
}

//=============================****玩家信息读取****================================
func (player *PlayerInfo) GetRoomId() string {
	return player.roomId
}

func (player *PlayerInfo) GetPlayerId() int32 {
	return player.oid
}

func (player *PlayerInfo) GetAgent() gate.Agent {
	if player.agent == nil {
		log.Debug("玩家[%v]未登录,或为机器人", player.oid)
	}
	return player.agent
}

//===========================****玩家信息类型转换****==============================
func (player *PlayerInfo) TransToPbPlayerInfo() *pb.PlayerInfo {
	result := &pb.PlayerInfo{
		OID:      proto.Int32(player.oid),
		NickName: proto.String(player.nickName),
		HeadIcon: proto.String(player.headIcon),
		Gold:     proto.Int32(player.gold),
		Diamond:  proto.Int32(player.diamond),
	}
	return result
}
