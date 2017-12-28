// Code generated by protoc-gen-go.
// source: mahjon.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	mahjon.proto

It has these top-level messages:
	LackCard
	CardInfo
	PlayerInfo
	GameOverInfo
	C2GSLogin
	GS2CLoginRet
	C2GSEnterGame
	GS2CEnterGameRet
	GS2CUpdateRoomMember
	GS2CBattleStart
	C2GSExchangeCard
	GS2CExchangeCardRet
	C2GSSelectLack
	GS2CSelectLackRet
	GS2CTurnToNext
	GS2CInterruptAction
	C2GSInterruptActionRet
	GS2CBroadcastProc
	GS2CGameOver
*/
package pb

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type GameType int32

const (
	GameType_XueZhan  GameType = 1
	GameType_XueLiu   GameType = 2
	GameType_DouDiZhu GameType = 3
)

var GameType_name = map[int32]string{
	1: "XueZhan",
	2: "XueLiu",
	3: "DouDiZhu",
}
var GameType_value = map[string]int32{
	"XueZhan":  1,
	"XueLiu":   2,
	"DouDiZhu": 3,
}

func (x GameType) Enum() *GameType {
	p := new(GameType)
	*p = x
	return p
}
func (x GameType) String() string {
	return proto.EnumName(GameType_name, int32(x))
}
func (x *GameType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GameType_value, data, "GameType")
	if err != nil {
		return err
	}
	*x = GameType(value)
	return nil
}

type EnterMode int32

const (
	EnterMode_CreateRoom EnterMode = 1
	EnterMode_JoinRoom   EnterMode = 2
	EnterMode_QuickEnter EnterMode = 3
)

var EnterMode_name = map[int32]string{
	1: "CreateRoom",
	2: "JoinRoom",
	3: "QuickEnter",
}
var EnterMode_value = map[string]int32{
	"CreateRoom": 1,
	"JoinRoom":   2,
	"QuickEnter": 3,
}

func (x EnterMode) Enum() *EnterMode {
	p := new(EnterMode)
	*p = x
	return p
}
func (x EnterMode) String() string {
	return proto.EnumName(EnterMode_name, int32(x))
}
func (x *EnterMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnterMode_value, data, "EnterMode")
	if err != nil {
		return err
	}
	*x = EnterMode(value)
	return nil
}

type MahjonSide int32

const (
	MahjonSide_DEFAULT MahjonSide = 1
	MahjonSide_EAST    MahjonSide = 2
	MahjonSide_SOUTH   MahjonSide = 3
	MahjonSide_WEST    MahjonSide = 4
	MahjonSide_NORTH   MahjonSide = 5
)

var MahjonSide_name = map[int32]string{
	1: "DEFAULT",
	2: "EAST",
	3: "SOUTH",
	4: "WEST",
	5: "NORTH",
}
var MahjonSide_value = map[string]int32{
	"DEFAULT": 1,
	"EAST":    2,
	"SOUTH":   3,
	"WEST":    4,
	"NORTH":   5,
}

func (x MahjonSide) Enum() *MahjonSide {
	p := new(MahjonSide)
	*p = x
	return p
}
func (x MahjonSide) String() string {
	return proto.EnumName(MahjonSide_name, int32(x))
}
func (x *MahjonSide) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MahjonSide_value, data, "MahjonSide")
	if err != nil {
		return err
	}
	*x = MahjonSide(value)
	return nil
}

type CardStatus int32

const (
	CardStatus_Wall   CardStatus = 1
	CardStatus_InHand CardStatus = 2
	CardStatus_P      CardStatus = 3
	CardStatus_G      CardStatus = 4
	CardStatus_Dis    CardStatus = 5
	CardStatus_Deal   CardStatus = 6
	CardStatus_Hu     CardStatus = 7
)

var CardStatus_name = map[int32]string{
	1: "Wall",
	2: "InHand",
	3: "P",
	4: "G",
	5: "Dis",
	6: "Deal",
	7: "Hu",
}
var CardStatus_value = map[string]int32{
	"Wall":   1,
	"InHand": 2,
	"P":      3,
	"G":      4,
	"Dis":    5,
	"Deal":   6,
	"Hu":     7,
}

func (x CardStatus) Enum() *CardStatus {
	p := new(CardStatus)
	*p = x
	return p
}
func (x CardStatus) String() string {
	return proto.EnumName(CardStatus_name, int32(x))
}
func (x *CardStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CardStatus_value, data, "CardStatus")
	if err != nil {
		return err
	}
	*x = CardStatus(value)
	return nil
}

type ExchangeType int32

const (
	ExchangeType_ClockWise ExchangeType = 1
	ExchangeType_AntiClock ExchangeType = 2
	ExchangeType_Opposite  ExchangeType = 3
)

var ExchangeType_name = map[int32]string{
	1: "ClockWise",
	2: "AntiClock",
	3: "Opposite",
}
var ExchangeType_value = map[string]int32{
	"ClockWise": 1,
	"AntiClock": 2,
	"Opposite":  3,
}

func (x ExchangeType) Enum() *ExchangeType {
	p := new(ExchangeType)
	*p = x
	return p
}
func (x ExchangeType) String() string {
	return proto.EnumName(ExchangeType_name, int32(x))
}
func (x *ExchangeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExchangeType_value, data, "ExchangeType")
	if err != nil {
		return err
	}
	*x = ExchangeType(value)
	return nil
}

type CardType int32

const (
	CardType_Default CardType = 1
	CardType_Wan     CardType = 2
	CardType_Tiao    CardType = 3
	CardType_Tong    CardType = 4
)

var CardType_name = map[int32]string{
	1: "Default",
	2: "Wan",
	3: "Tiao",
	4: "Tong",
}
var CardType_value = map[string]int32{
	"Default": 1,
	"Wan":     2,
	"Tiao":    3,
	"Tong":    4,
}

func (x CardType) Enum() *CardType {
	p := new(CardType)
	*p = x
	return p
}
func (x CardType) String() string {
	return proto.EnumName(CardType_name, int32(x))
}
func (x *CardType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CardType_value, data, "CardType")
	if err != nil {
		return err
	}
	*x = CardType(value)
	return nil
}

type ProcType int32

const (
	ProcType_Proc_Hu      ProcType = 1
	ProcType_Proc_Gang    ProcType = 2
	ProcType_Proc_Peng    ProcType = 3
	ProcType_Proc_Discard ProcType = 4
	ProcType_Proc_Pass    ProcType = 5
)

var ProcType_name = map[int32]string{
	1: "Proc_Hu",
	2: "Proc_Gang",
	3: "Proc_Peng",
	4: "Proc_Discard",
	5: "Proc_Pass",
}
var ProcType_value = map[string]int32{
	"Proc_Hu":      1,
	"Proc_Gang":    2,
	"Proc_Peng":    3,
	"Proc_Discard": 4,
	"Proc_Pass":    5,
}

func (x ProcType) Enum() *ProcType {
	p := new(ProcType)
	*p = x
	return p
}
func (x ProcType) String() string {
	return proto.EnumName(ProcType_name, int32(x))
}
func (x *ProcType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProcType_value, data, "ProcType")
	if err != nil {
		return err
	}
	*x = ProcType(value)
	return nil
}

type TurnSwitchType int32

const (
	TurnSwitchType_Normal         TurnSwitchType = 1
	TurnSwitchType_JustCanDiscard TurnSwitchType = 2
	TurnSwitchType_NotDrawCard    TurnSwitchType = 3
)

var TurnSwitchType_name = map[int32]string{
	1: "Normal",
	2: "JustCanDiscard",
	3: "NotDrawCard",
}
var TurnSwitchType_value = map[string]int32{
	"Normal":         1,
	"JustCanDiscard": 2,
	"NotDrawCard":    3,
}

func (x TurnSwitchType) Enum() *TurnSwitchType {
	p := new(TurnSwitchType)
	*p = x
	return p
}
func (x TurnSwitchType) String() string {
	return proto.EnumName(TurnSwitchType_name, int32(x))
}
func (x *TurnSwitchType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TurnSwitchType_value, data, "TurnSwitchType")
	if err != nil {
		return err
	}
	*x = TurnSwitchType(value)
	return nil
}

type GS2CLoginRet_ErrorCode int32

const (
	GS2CLoginRet_SUCCESS        GS2CLoginRet_ErrorCode = 1
	GS2CLoginRet_ACCOUNT_ERROR  GS2CLoginRet_ErrorCode = 2
	GS2CLoginRet_PASSWORD_ERROR GS2CLoginRet_ErrorCode = 3
	GS2CLoginRet_FAIL           GS2CLoginRet_ErrorCode = 4
)

var GS2CLoginRet_ErrorCode_name = map[int32]string{
	1: "SUCCESS",
	2: "ACCOUNT_ERROR",
	3: "PASSWORD_ERROR",
	4: "FAIL",
}
var GS2CLoginRet_ErrorCode_value = map[string]int32{
	"SUCCESS":        1,
	"ACCOUNT_ERROR":  2,
	"PASSWORD_ERROR": 3,
	"FAIL":           4,
}

func (x GS2CLoginRet_ErrorCode) Enum() *GS2CLoginRet_ErrorCode {
	p := new(GS2CLoginRet_ErrorCode)
	*p = x
	return p
}
func (x GS2CLoginRet_ErrorCode) String() string {
	return proto.EnumName(GS2CLoginRet_ErrorCode_name, int32(x))
}
func (x *GS2CLoginRet_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GS2CLoginRet_ErrorCode_value, data, "GS2CLoginRet_ErrorCode")
	if err != nil {
		return err
	}
	*x = GS2CLoginRet_ErrorCode(value)
	return nil
}

type GS2CEnterGameRet_ErrorCode int32

const (
	GS2CEnterGameRet_SUCCESS             GS2CEnterGameRet_ErrorCode = 1
	GS2CEnterGameRet_FAIL                GS2CEnterGameRet_ErrorCode = 2
	GS2CEnterGameRet_PLAYER_COUNT_LIMITE GS2CEnterGameRet_ErrorCode = 3
	GS2CEnterGameRet_ROOM_NOT_EXIST      GS2CEnterGameRet_ErrorCode = 4
	GS2CEnterGameRet_NO_EMPTY_ROOM       GS2CEnterGameRet_ErrorCode = 5
)

var GS2CEnterGameRet_ErrorCode_name = map[int32]string{
	1: "SUCCESS",
	2: "FAIL",
	3: "PLAYER_COUNT_LIMITE",
	4: "ROOM_NOT_EXIST",
	5: "NO_EMPTY_ROOM",
}
var GS2CEnterGameRet_ErrorCode_value = map[string]int32{
	"SUCCESS":             1,
	"FAIL":                2,
	"PLAYER_COUNT_LIMITE": 3,
	"ROOM_NOT_EXIST":      4,
	"NO_EMPTY_ROOM":       5,
}

func (x GS2CEnterGameRet_ErrorCode) Enum() *GS2CEnterGameRet_ErrorCode {
	p := new(GS2CEnterGameRet_ErrorCode)
	*p = x
	return p
}
func (x GS2CEnterGameRet_ErrorCode) String() string {
	return proto.EnumName(GS2CEnterGameRet_ErrorCode_name, int32(x))
}
func (x *GS2CEnterGameRet_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GS2CEnterGameRet_ErrorCode_value, data, "GS2CEnterGameRet_ErrorCode")
	if err != nil {
		return err
	}
	*x = GS2CEnterGameRet_ErrorCode(value)
	return nil
}

type LackCard struct {
	PlayerOID        *int32    `protobuf:"varint,1,req,name=playerOID" json:"playerOID,omitempty"`
	Type             *CardType `protobuf:"varint,2,req,name=type,enum=pb.CardType" json:"type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *LackCard) Reset()         { *m = LackCard{} }
func (m *LackCard) String() string { return proto.CompactTextString(m) }
func (*LackCard) ProtoMessage()    {}

func (m *LackCard) GetPlayerOID() int32 {
	if m != nil && m.PlayerOID != nil {
		return *m.PlayerOID
	}
	return 0
}

func (m *LackCard) GetType() CardType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CardType_Default
}

type CardInfo struct {
	PlayerOID        *int32      `protobuf:"varint,1,req,name=playerOID" json:"playerOID,omitempty"`
	OID              *int32      `protobuf:"varint,2,req" json:"OID,omitempty"`
	ID               *int32      `protobuf:"varint,3,req" json:"ID,omitempty"`
	Status           *CardStatus `protobuf:"varint,4,req,enum=pb.CardStatus" json:"Status,omitempty"`
	FromOther        *bool       `protobuf:"varint,5,opt,name=fromOther" json:"fromOther,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CardInfo) Reset()         { *m = CardInfo{} }
func (m *CardInfo) String() string { return proto.CompactTextString(m) }
func (*CardInfo) ProtoMessage()    {}

func (m *CardInfo) GetPlayerOID() int32 {
	if m != nil && m.PlayerOID != nil {
		return *m.PlayerOID
	}
	return 0
}

func (m *CardInfo) GetOID() int32 {
	if m != nil && m.OID != nil {
		return *m.OID
	}
	return 0
}

func (m *CardInfo) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *CardInfo) GetStatus() CardStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return CardStatus_Wall
}

func (m *CardInfo) GetFromOther() bool {
	if m != nil && m.FromOther != nil {
		return *m.FromOther
	}
	return false
}

type PlayerInfo struct {
	OID              *int32      `protobuf:"varint,1,req" json:"OID,omitempty"`
	NickName         *string     `protobuf:"bytes,2,req" json:"NickName,omitempty"`
	HeadIcon         *string     `protobuf:"bytes,3,req" json:"HeadIcon,omitempty"`
	Gold             *int32      `protobuf:"varint,4,req" json:"Gold,omitempty"`
	Diamond          *int32      `protobuf:"varint,5,req" json:"Diamond,omitempty"`
	Side             *MahjonSide `protobuf:"varint,6,opt,enum=pb.MahjonSide" json:"Side,omitempty"`
	IsOwner          *bool       `protobuf:"varint,7,opt" json:"IsOwner,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PlayerInfo) Reset()         { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()    {}

func (m *PlayerInfo) GetOID() int32 {
	if m != nil && m.OID != nil {
		return *m.OID
	}
	return 0
}

func (m *PlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PlayerInfo) GetHeadIcon() string {
	if m != nil && m.HeadIcon != nil {
		return *m.HeadIcon
	}
	return ""
}

func (m *PlayerInfo) GetGold() int32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *PlayerInfo) GetDiamond() int32 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *PlayerInfo) GetSide() MahjonSide {
	if m != nil && m.Side != nil {
		return *m.Side
	}
	return MahjonSide_DEFAULT
}

func (m *PlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

type GameOverInfo struct {
	Player           *PlayerInfo `protobuf:"bytes,1,req,name=player" json:"player,omitempty"`
	Fan              *int32      `protobuf:"varint,2,req,name=fan" json:"fan,omitempty"`
	CardList         []*CardInfo `protobuf:"bytes,3,rep,name=cardList" json:"cardList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *GameOverInfo) Reset()         { *m = GameOverInfo{} }
func (m *GameOverInfo) String() string { return proto.CompactTextString(m) }
func (*GameOverInfo) ProtoMessage()    {}

func (m *GameOverInfo) GetPlayer() *PlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *GameOverInfo) GetFan() int32 {
	if m != nil && m.Fan != nil {
		return *m.Fan
	}
	return 0
}

func (m *GameOverInfo) GetCardList() []*CardInfo {
	if m != nil {
		return m.CardList
	}
	return nil
}

// ///////////////////////////////////////////////////////////////////
type C2GSLogin struct {
	Account          *string `protobuf:"bytes,1,req,name=account" json:"account,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *C2GSLogin) Reset()         { *m = C2GSLogin{} }
func (m *C2GSLogin) String() string { return proto.CompactTextString(m) }
func (*C2GSLogin) ProtoMessage()    {}

func (m *C2GSLogin) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *C2GSLogin) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type GS2CLoginRet struct {
	ErrorCode        *GS2CLoginRet_ErrorCode `protobuf:"varint,1,req,name=errorCode,enum=pb.GS2CLoginRet_ErrorCode" json:"errorCode,omitempty"`
	PlayerInfo       *PlayerInfo             `protobuf:"bytes,2,opt,name=playerInfo" json:"playerInfo,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *GS2CLoginRet) Reset()         { *m = GS2CLoginRet{} }
func (m *GS2CLoginRet) String() string { return proto.CompactTextString(m) }
func (*GS2CLoginRet) ProtoMessage()    {}

func (m *GS2CLoginRet) GetErrorCode() GS2CLoginRet_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return GS2CLoginRet_SUCCESS
}

func (m *GS2CLoginRet) GetPlayerInfo() *PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

type C2GSEnterGame struct {
	Type             *GameType  `protobuf:"varint,1,req,name=type,enum=pb.GameType" json:"type,omitempty"`
	Mode             *EnterMode `protobuf:"varint,2,req,name=mode,enum=pb.EnterMode" json:"mode,omitempty"`
	RoomId           *string    `protobuf:"bytes,3,opt,name=roomId" json:"roomId,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *C2GSEnterGame) Reset()         { *m = C2GSEnterGame{} }
func (m *C2GSEnterGame) String() string { return proto.CompactTextString(m) }
func (*C2GSEnterGame) ProtoMessage()    {}

func (m *C2GSEnterGame) GetType() GameType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return GameType_XueZhan
}

func (m *C2GSEnterGame) GetMode() EnterMode {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return EnterMode_CreateRoom
}

func (m *C2GSEnterGame) GetRoomId() string {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return ""
}

type GS2CEnterGameRet struct {
	ErrorCode        *GS2CEnterGameRet_ErrorCode `protobuf:"varint,1,req,name=errorCode,enum=pb.GS2CEnterGameRet_ErrorCode" json:"errorCode,omitempty"`
	Type             *GameType                   `protobuf:"varint,2,req,name=type,enum=pb.GameType" json:"type,omitempty"`
	RoomId           *string                     `protobuf:"bytes,3,opt,name=roomId" json:"roomId,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *GS2CEnterGameRet) Reset()         { *m = GS2CEnterGameRet{} }
func (m *GS2CEnterGameRet) String() string { return proto.CompactTextString(m) }
func (*GS2CEnterGameRet) ProtoMessage()    {}

func (m *GS2CEnterGameRet) GetErrorCode() GS2CEnterGameRet_ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return GS2CEnterGameRet_SUCCESS
}

func (m *GS2CEnterGameRet) GetType() GameType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return GameType_XueZhan
}

func (m *GS2CEnterGameRet) GetRoomId() string {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return ""
}

type GS2CUpdateRoomMember struct {
	Player           []*PlayerInfo `protobuf:"bytes,1,rep,name=player" json:"player,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GS2CUpdateRoomMember) Reset()         { *m = GS2CUpdateRoomMember{} }
func (m *GS2CUpdateRoomMember) String() string { return proto.CompactTextString(m) }
func (*GS2CUpdateRoomMember) ProtoMessage()    {}

func (m *GS2CUpdateRoomMember) GetPlayer() []*PlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

type GS2CBattleStart struct {
	DealerId         *int32      `protobuf:"varint,1,req,name=dealerId" json:"dealerId,omitempty"`
	CardList         []*CardInfo `protobuf:"bytes,2,rep,name=cardList" json:"cardList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *GS2CBattleStart) Reset()         { *m = GS2CBattleStart{} }
func (m *GS2CBattleStart) String() string { return proto.CompactTextString(m) }
func (*GS2CBattleStart) ProtoMessage()    {}

func (m *GS2CBattleStart) GetDealerId() int32 {
	if m != nil && m.DealerId != nil {
		return *m.DealerId
	}
	return 0
}

func (m *GS2CBattleStart) GetCardList() []*CardInfo {
	if m != nil {
		return m.CardList
	}
	return nil
}

type C2GSExchangeCard struct {
	CardOIDList      []int32 `protobuf:"varint,1,rep,name=cardOIDList" json:"cardOIDList,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *C2GSExchangeCard) Reset()         { *m = C2GSExchangeCard{} }
func (m *C2GSExchangeCard) String() string { return proto.CompactTextString(m) }
func (*C2GSExchangeCard) ProtoMessage()    {}

func (m *C2GSExchangeCard) GetCardOIDList() []int32 {
	if m != nil {
		return m.CardOIDList
	}
	return nil
}

type GS2CExchangeCardRet struct {
	Type             *ExchangeType `protobuf:"varint,1,req,name=type,enum=pb.ExchangeType" json:"type,omitempty"`
	CardList         []*CardInfo   `protobuf:"bytes,2,rep,name=cardList" json:"cardList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GS2CExchangeCardRet) Reset()         { *m = GS2CExchangeCardRet{} }
func (m *GS2CExchangeCardRet) String() string { return proto.CompactTextString(m) }
func (*GS2CExchangeCardRet) ProtoMessage()    {}

func (m *GS2CExchangeCardRet) GetType() ExchangeType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ExchangeType_ClockWise
}

func (m *GS2CExchangeCardRet) GetCardList() []*CardInfo {
	if m != nil {
		return m.CardList
	}
	return nil
}

type C2GSSelectLack struct {
	Type             *CardType `protobuf:"varint,1,req,name=type,enum=pb.CardType" json:"type,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *C2GSSelectLack) Reset()         { *m = C2GSSelectLack{} }
func (m *C2GSSelectLack) String() string { return proto.CompactTextString(m) }
func (*C2GSSelectLack) ProtoMessage()    {}

func (m *C2GSSelectLack) GetType() CardType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CardType_Default
}

type GS2CSelectLackRet struct {
	LackCard         []*LackCard `protobuf:"bytes,1,rep,name=lackCard" json:"lackCard,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *GS2CSelectLackRet) Reset()         { *m = GS2CSelectLackRet{} }
func (m *GS2CSelectLackRet) String() string { return proto.CompactTextString(m) }
func (*GS2CSelectLackRet) ProtoMessage()    {}

func (m *GS2CSelectLackRet) GetLackCard() []*LackCard {
	if m != nil {
		return m.LackCard
	}
	return nil
}

type GS2CTurnToNext struct {
	PlayerOID        *int32    `protobuf:"varint,1,req,name=playerOID" json:"playerOID,omitempty"`
	DrawCard         *CardInfo `protobuf:"bytes,2,opt,name=drawCard" json:"drawCard,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *GS2CTurnToNext) Reset()         { *m = GS2CTurnToNext{} }
func (m *GS2CTurnToNext) String() string { return proto.CompactTextString(m) }
func (*GS2CTurnToNext) ProtoMessage()    {}

func (m *GS2CTurnToNext) GetPlayerOID() int32 {
	if m != nil && m.PlayerOID != nil {
		return *m.PlayerOID
	}
	return 0
}

func (m *GS2CTurnToNext) GetDrawCard() *CardInfo {
	if m != nil {
		return m.DrawCard
	}
	return nil
}

type GS2CInterruptAction struct {
	ProcList         []ProcType `protobuf:"varint,1,rep,name=procList,enum=pb.ProcType" json:"procList,omitempty"`
	DrawCard         *CardInfo  `protobuf:"bytes,2,opt,name=drawCard" json:"drawCard,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *GS2CInterruptAction) Reset()         { *m = GS2CInterruptAction{} }
func (m *GS2CInterruptAction) String() string { return proto.CompactTextString(m) }
func (*GS2CInterruptAction) ProtoMessage()    {}

func (m *GS2CInterruptAction) GetProcList() []ProcType {
	if m != nil {
		return m.ProcList
	}
	return nil
}

func (m *GS2CInterruptAction) GetDrawCard() *CardInfo {
	if m != nil {
		return m.DrawCard
	}
	return nil
}

type C2GSInterruptActionRet struct {
	ProcType         *ProcType `protobuf:"varint,1,req,name=procType,enum=pb.ProcType" json:"procType,omitempty"`
	DrawCard         *CardInfo `protobuf:"bytes,2,opt,name=drawCard" json:"drawCard,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *C2GSInterruptActionRet) Reset()         { *m = C2GSInterruptActionRet{} }
func (m *C2GSInterruptActionRet) String() string { return proto.CompactTextString(m) }
func (*C2GSInterruptActionRet) ProtoMessage()    {}

func (m *C2GSInterruptActionRet) GetProcType() ProcType {
	if m != nil && m.ProcType != nil {
		return *m.ProcType
	}
	return ProcType_Proc_Hu
}

func (m *C2GSInterruptActionRet) GetDrawCard() *CardInfo {
	if m != nil {
		return m.DrawCard
	}
	return nil
}

type GS2CBroadcastProc struct {
	ProcPlayer       *int32      `protobuf:"varint,1,req,name=procPlayer" json:"procPlayer,omitempty"`
	ProcType         *ProcType   `protobuf:"varint,2,req,name=procType,enum=pb.ProcType" json:"procType,omitempty"`
	CardList         []*CardInfo `protobuf:"bytes,3,rep,name=cardList" json:"cardList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *GS2CBroadcastProc) Reset()         { *m = GS2CBroadcastProc{} }
func (m *GS2CBroadcastProc) String() string { return proto.CompactTextString(m) }
func (*GS2CBroadcastProc) ProtoMessage()    {}

func (m *GS2CBroadcastProc) GetProcPlayer() int32 {
	if m != nil && m.ProcPlayer != nil {
		return *m.ProcPlayer
	}
	return 0
}

func (m *GS2CBroadcastProc) GetProcType() ProcType {
	if m != nil && m.ProcType != nil {
		return *m.ProcType
	}
	return ProcType_Proc_Hu
}

func (m *GS2CBroadcastProc) GetCardList() []*CardInfo {
	if m != nil {
		return m.CardList
	}
	return nil
}

type GS2CGameOver struct {
	List             []*GameOverInfo `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GS2CGameOver) Reset()         { *m = GS2CGameOver{} }
func (m *GS2CGameOver) String() string { return proto.CompactTextString(m) }
func (*GS2CGameOver) ProtoMessage()    {}

func (m *GS2CGameOver) GetList() []*GameOverInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.GameType", GameType_name, GameType_value)
	proto.RegisterEnum("pb.EnterMode", EnterMode_name, EnterMode_value)
	proto.RegisterEnum("pb.MahjonSide", MahjonSide_name, MahjonSide_value)
	proto.RegisterEnum("pb.CardStatus", CardStatus_name, CardStatus_value)
	proto.RegisterEnum("pb.ExchangeType", ExchangeType_name, ExchangeType_value)
	proto.RegisterEnum("pb.CardType", CardType_name, CardType_value)
	proto.RegisterEnum("pb.ProcType", ProcType_name, ProcType_value)
	proto.RegisterEnum("pb.TurnSwitchType", TurnSwitchType_name, TurnSwitchType_value)
	proto.RegisterEnum("pb.GS2CLoginRet_ErrorCode", GS2CLoginRet_ErrorCode_name, GS2CLoginRet_ErrorCode_value)
	proto.RegisterEnum("pb.GS2CEnterGameRet_ErrorCode", GS2CEnterGameRet_ErrorCode_name, GS2CEnterGameRet_ErrorCode_value)
}