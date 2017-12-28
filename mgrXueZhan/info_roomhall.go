package mgrXueZhan

import (
	"server/mgrPlayer"
	"server/mgrXueZhan/mgrRoom"
	"server/pb"

	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var roomMap sync.Map //roomId:room

func regNewRoom(owner *mgrPlayer.PlayerInfo) {

	roomId := rndRoomId()
	room := mgrRoom.NewRoom(roomId)
	roomMap.Store(roomId, room)
	sendEnterGameRet(pb.GS2CEnterGameRet_SUCCESS.Enum(), pb.GameType_XueZhan.Enum(), owner.GetRoomId(), owner.GetAgent())

	room.AddNewPlayerIntoRoom(owner)

	waitTimer := time.NewTimer(time.Second * 1)
	<-waitTimer.C
	room.AddRobotToRoom(3)
}

func getRoomById(roomid string) *mgrRoom.RoomInfo {
	room, ex := roomMap.Load(roomid)
	if !ex {
		log.Debug("info_roomhall==>房间[%v]不存在,或已被注销", roomid)
		return nil
	}
	return room.(*mgrRoom.RoomInfo)
}

func joinRoom(player *mgrPlayer.PlayerInfo, roomid string) {

	room, ex := roomMap.Load(roomid)
	if !ex {
		log.Debug("info_roomhall==>目标房间[v%]不存在", roomid)
		sendEnterGameRet(pb.GS2CEnterGameRet_ROOM_NOT_EXIST.Enum(), pb.GameType_XueZhan.Enum(), roomid, player.GetAgent())
	}
	ok := room.(*mgrRoom.RoomInfo).AddNewPlayerIntoRoom(player)
	if !ok {
		sendEnterGameRet(pb.GS2CEnterGameRet_PLAYER_COUNT_LIMITE.Enum(), pb.GameType_XueZhan.Enum(), roomid, player.GetAgent())
	}

	sendEnterGameRet(pb.GS2CEnterGameRet_SUCCESS.Enum(), pb.GameType_XueZhan.Enum(), roomid, player.GetAgent())
}

//=============================****生成随机房间号****==============================
func rndRoomId() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	roomId := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	_, ex := roomMap.Load(roomId)
	if ex {
		log.Debug("info_roomhall==>房间[%v]已存在", roomId)
		return rndRoomId()
	}
	return roomId
}

//=============================****大厅消息管理****===============================
func sendEnterGameRet(errCode *pb.GS2CEnterGameRet_ErrorCode, gameType *pb.GameType, roomId string, a gate.Agent) {

	data := &pb.GS2CEnterGameRet{
		ErrorCode: errCode.Enum(),
		Type:      gameType.Enum(),
		RoomId:    proto.String(roomId),
	}

	a.WriteMsg(data)
	log.Debug("info_roomhall==>进入房间[%v]反馈==>SendEnterGameRet", roomId)
}
