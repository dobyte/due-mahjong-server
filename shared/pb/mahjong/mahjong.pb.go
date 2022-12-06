// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: mahjong.proto

package mahjong

import (
	common "due-mahjong-server/shared/pb/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FetchRoomsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchRoomsReq) Reset() {
	*x = FetchRoomsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRoomsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRoomsReq) ProtoMessage() {}

func (x *FetchRoomsReq) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRoomsReq.ProtoReflect.Descriptor instead.
func (*FetchRoomsReq) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{0}
}

type QuickStartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QuickStartReq) Reset() {
	*x = QuickStartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickStartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickStartReq) ProtoMessage() {}

func (x *QuickStartReq) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickStartReq.ProtoReflect.Descriptor instead.
func (*QuickStartReq) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{1}
}

type QuickStartRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     common.Code `protobuf:"varint,1,opt,name=Code,proto3,enum=common.Code" json:"Code,omitempty"` // 错误码
	GameInfo *GameInfo   `protobuf:"bytes,2,opt,name=GameInfo,proto3" json:"GameInfo,omitempty"`           // 游戏信息
}

func (x *QuickStartRes) Reset() {
	*x = QuickStartRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuickStartRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuickStartRes) ProtoMessage() {}

func (x *QuickStartRes) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuickStartRes.ProtoReflect.Descriptor instead.
func (*QuickStartRes) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{2}
}

func (x *QuickStartRes) GetCode() common.Code {
	if x != nil {
		return x.Code
	}
	return common.Code(0)
}

func (x *QuickStartRes) GetGameInfo() *GameInfo {
	if x != nil {
		return x.GameInfo
	}
	return nil
}

type GameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room  *Room  `protobuf:"bytes,1,opt,name=Room,proto3" json:"Room,omitempty"`   // 房间
	Table *Table `protobuf:"bytes,2,opt,name=Table,proto3" json:"Table,omitempty"` // 牌桌
}

func (x *GameInfo) Reset() {
	*x = GameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameInfo) ProtoMessage() {}

func (x *GameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameInfo.ProtoReflect.Descriptor instead.
func (*GameInfo) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{3}
}

func (x *GameInfo) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *GameInfo) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`                       // 房间ID
	Name          string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`                    // 房间名称
	MinEntryLimit int32  `protobuf:"varint,3,opt,name=MinEntryLimit,proto3" json:"MinEntryLimit,omitempty"` // 最小进入限制
	MaxEntryLimit int32  `protobuf:"varint,4,opt,name=MaxEntryLimit,proto3" json:"MaxEntryLimit,omitempty"` // 最大进入限制
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{4}
}

func (x *Room) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetMinEntryLimit() int32 {
	if x != nil {
		return x.MinEntryLimit
	}
	return 0
}

func (x *Room) GetMaxEntryLimit() int32 {
	if x != nil {
		return x.MaxEntryLimit
	}
	return 0
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`      // 牌桌ID
	Seats []*Seat `protobuf:"bytes,2,rep,name=Seats,proto3" json:"Seats,omitempty"` // 座位信息
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{5}
}

func (x *Table) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Table) GetSeats() []*Seat {
	if x != nil {
		return x.Seats
	}
	return nil
}

type Seat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`        // 座位ID
	Player *Player `protobuf:"bytes,2,opt,name=Player,proto3" json:"Player,omitempty"` // 玩家信息
}

func (x *Seat) Reset() {
	*x = Seat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seat) ProtoMessage() {}

func (x *Seat) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seat.ProtoReflect.Descriptor instead.
func (*Seat) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{6}
}

func (x *Seat) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Seat) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *common.User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`          // 用户
	IsMyself bool         `protobuf:"varint,2,opt,name=IsMyself,proto3" json:"IsMyself,omitempty"` // 是否是自己
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{7}
}

func (x *Player) GetUser() *common.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Player) GetIsMyself() bool {
	if x != nil {
		return x.IsMyself
	}
	return false
}

// 离线通知
type OfflineNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatID int32 `protobuf:"varint,1,opt,name=SeatID,proto3" json:"SeatID,omitempty"` // 座位ID
}

func (x *OfflineNotify) Reset() {
	*x = OfflineNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineNotify) ProtoMessage() {}

func (x *OfflineNotify) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineNotify.ProtoReflect.Descriptor instead.
func (*OfflineNotify) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{8}
}

func (x *OfflineNotify) GetSeatID() int32 {
	if x != nil {
		return x.SeatID
	}
	return 0
}

// 上线通知
type OnlineNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatID int32 `protobuf:"varint,1,opt,name=SeatID,proto3" json:"SeatID,omitempty"` // 座位ID
}

func (x *OnlineNotify) Reset() {
	*x = OnlineNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mahjong_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineNotify) ProtoMessage() {}

func (x *OnlineNotify) ProtoReflect() protoreflect.Message {
	mi := &file_mahjong_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineNotify.ProtoReflect.Descriptor instead.
func (*OnlineNotify) Descriptor() ([]byte, []int) {
	return file_mahjong_proto_rawDescGZIP(), []int{9}
}

func (x *OnlineNotify) GetSeatID() int32 {
	if x != nil {
		return x.SeatID
	}
	return 0
}

var File_mahjong_proto protoreflect.FileDescriptor

var file_mahjong_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x1a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x22, 0x0f, 0x0a, 0x0d, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x22, 0x60, 0x0a, 0x0d, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x68, 0x6a,
	0x6f, 0x6e, 0x67, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x47, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x53, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x21, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x24, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x76, 0x0a, 0x04, 0x52,
	0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x4d, 0x69, 0x6e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x4d, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x4d, 0x61, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4d, 0x61, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x3c, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x05,
	0x53, 0x65, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61,
	0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x05, 0x53, 0x65, 0x61, 0x74,
	0x73, 0x22, 0x3f, 0x0a, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x27, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x68, 0x6a,
	0x6f, 0x6e, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x22, 0x46, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x49, 0x73, 0x4d, 0x79, 0x73, 0x65, 0x6c, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x49, 0x73, 0x4d, 0x79, 0x73, 0x65, 0x6c, 0x66, 0x22, 0x27, 0x0a, 0x0d, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x65, 0x61,
	0x74, 0x49, 0x44, 0x22, 0x26, 0x0a, 0x0c, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x44, 0x42, 0x26, 0x5a, 0x24, 0x64,
	0x75, 0x65, 0x2d, 0x6d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x68, 0x6a,
	0x6f, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mahjong_proto_rawDescOnce sync.Once
	file_mahjong_proto_rawDescData = file_mahjong_proto_rawDesc
)

func file_mahjong_proto_rawDescGZIP() []byte {
	file_mahjong_proto_rawDescOnce.Do(func() {
		file_mahjong_proto_rawDescData = protoimpl.X.CompressGZIP(file_mahjong_proto_rawDescData)
	})
	return file_mahjong_proto_rawDescData
}

var file_mahjong_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_mahjong_proto_goTypes = []interface{}{
	(*FetchRoomsReq)(nil), // 0: mahjong.FetchRoomsReq
	(*QuickStartReq)(nil), // 1: mahjong.QuickStartReq
	(*QuickStartRes)(nil), // 2: mahjong.QuickStartRes
	(*GameInfo)(nil),      // 3: mahjong.GameInfo
	(*Room)(nil),          // 4: mahjong.Room
	(*Table)(nil),         // 5: mahjong.Table
	(*Seat)(nil),          // 6: mahjong.Seat
	(*Player)(nil),        // 7: mahjong.Player
	(*OfflineNotify)(nil), // 8: mahjong.OfflineNotify
	(*OnlineNotify)(nil),  // 9: mahjong.OnlineNotify
	(common.Code)(0),      // 10: common.Code
	(*common.User)(nil),   // 11: common.User
}
var file_mahjong_proto_depIdxs = []int32{
	10, // 0: mahjong.QuickStartRes.Code:type_name -> common.Code
	3,  // 1: mahjong.QuickStartRes.GameInfo:type_name -> mahjong.GameInfo
	4,  // 2: mahjong.GameInfo.Room:type_name -> mahjong.Room
	5,  // 3: mahjong.GameInfo.Table:type_name -> mahjong.Table
	6,  // 4: mahjong.Table.Seats:type_name -> mahjong.Seat
	7,  // 5: mahjong.Seat.Player:type_name -> mahjong.Player
	11, // 6: mahjong.Player.User:type_name -> common.User
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_mahjong_proto_init() }
func file_mahjong_proto_init() {
	if File_mahjong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mahjong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRoomsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuickStartReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuickStartRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineNotify); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mahjong_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineNotify); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mahjong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mahjong_proto_goTypes,
		DependencyIndexes: file_mahjong_proto_depIdxs,
		MessageInfos:      file_mahjong_proto_msgTypes,
	}.Build()
	File_mahjong_proto = out.File
	file_mahjong_proto_rawDesc = nil
	file_mahjong_proto_goTypes = nil
	file_mahjong_proto_depIdxs = nil
}