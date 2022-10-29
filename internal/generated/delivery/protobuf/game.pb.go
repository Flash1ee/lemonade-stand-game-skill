// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: game.proto

package game

import (
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

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *CreateResult) Reset() {
	*x = CreateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResult) ProtoMessage() {}

func (x *CreateResult) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResult.ProtoReflect.Descriptor instead.
func (*CreateResult) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Weather struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WeatherName string `protobuf:"bytes,1,opt,name=WeatherName,proto3" json:"WeatherName,omitempty"`
	RainChance  int64  `protobuf:"varint,2,opt,name=RainChance,proto3" json:"RainChance,omitempty"`
}

func (x *Weather) Reset() {
	*x = Weather{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weather) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weather) ProtoMessage() {}

func (x *Weather) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weather.ProtoReflect.Descriptor instead.
func (*Weather) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *Weather) GetWeatherName() string {
	if x != nil {
		return x.WeatherName
	}
	return ""
}

func (x *Weather) GetRainChance() int64 {
	if x != nil {
		return x.RainChance
	}
	return 0
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance int64 `protobuf:"varint,1,opt,name=Balance,proto3" json:"Balance,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *Balance) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=Dummy,proto3" json:"Dummy,omitempty"`
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *Nothing) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

type GameID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GameID) Reset() {
	*x = GameID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameID) ProtoMessage() {}

func (x *GameID) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameID.ProtoReflect.Descriptor instead.
func (*GameID) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *GameID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CalculateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game        *GameID `protobuf:"bytes,1,opt,name=Game,proto3" json:"Game,omitempty"`
	CupsAmount  int64   `protobuf:"varint,2,opt,name=CupsAmount,proto3" json:"CupsAmount,omitempty"`
	IceAmount   int64   `protobuf:"varint,3,opt,name=IceAmount,proto3" json:"IceAmount,omitempty"`
	StandAmount int64   `protobuf:"varint,4,opt,name=StandAmount,proto3" json:"StandAmount,omitempty"`
	Price       int64   `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{6}
}

func (x *CalculateRequest) GetGame() *GameID {
	if x != nil {
		return x.Game
	}
	return nil
}

func (x *CalculateRequest) GetCupsAmount() int64 {
	if x != nil {
		return x.CupsAmount
	}
	return 0
}

func (x *CalculateRequest) GetIceAmount() int64 {
	if x != nil {
		return x.IceAmount
	}
	return 0
}

func (x *CalculateRequest) GetStandAmount() int64 {
	if x != nil {
		return x.StandAmount
	}
	return 0
}

func (x *CalculateRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance int64 `protobuf:"varint,1,opt,name=Balance,proto3" json:"Balance,omitempty"`
	Day     int64 `protobuf:"varint,2,opt,name=Day,proto3" json:"Day,omitempty"`
	Profit  int64 `protobuf:"varint,3,opt,name=Profit,proto3" json:"Profit,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{7}
}

func (x *CalculateResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *CalculateResponse) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *CalculateResponse) GetProfit() int64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{8}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x07, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x61, 0x69, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x61, 0x69,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x1f, 0x0a, 0x07,
	0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x18, 0x0a,
	0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04,
	0x47, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x52, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x43, 0x75, 0x70, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x43, 0x75, 0x70, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x49, 0x63, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x49, 0x63, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x57, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x44, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x22, 0x22, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xcf, 0x01, 0x0a, 0x0c, 0x4c, 0x65, 0x6d, 0x6f, 0x6e, 0x61, 0x64, 0x65, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x0d,
	0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x0c, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xd6, 0x01, 0x0a, 0x13, 0x42, 0x6f, 0x74, 0x61, 0x6e, 0x69, 0x63, 0x61,
	0x6c, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x1a,
	0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_game_proto_goTypes = []interface{}{
	(*Game)(nil),              // 0: game.Game
	(*CreateResult)(nil),      // 1: game.CreateResult
	(*Weather)(nil),           // 2: game.Weather
	(*Balance)(nil),           // 3: game.Balance
	(*Nothing)(nil),           // 4: game.Nothing
	(*GameID)(nil),            // 5: game.GameID
	(*CalculateRequest)(nil),  // 6: game.CalculateRequest
	(*CalculateResponse)(nil), // 7: game.CalculateResponse
	(*User)(nil),              // 8: game.User
}
var file_game_proto_depIdxs = []int32{
	5, // 0: game.CalculateRequest.Game:type_name -> game.GameID
	8, // 1: game.LemonadeGame.Create:input_type -> game.User
	5, // 2: game.LemonadeGame.RandomWeather:input_type -> game.GameID
	5, // 3: game.LemonadeGame.GetBalance:input_type -> game.GameID
	6, // 4: game.LemonadeGame.Calculate:input_type -> game.CalculateRequest
	8, // 5: game.BotanicalGardenGame.Create:input_type -> game.User
	5, // 6: game.BotanicalGardenGame.RandomWeather:input_type -> game.GameID
	5, // 7: game.BotanicalGardenGame.GetBalance:input_type -> game.GameID
	6, // 8: game.BotanicalGardenGame.Calculate:input_type -> game.CalculateRequest
	1, // 9: game.LemonadeGame.Create:output_type -> game.CreateResult
	2, // 10: game.LemonadeGame.RandomWeather:output_type -> game.Weather
	3, // 11: game.LemonadeGame.GetBalance:output_type -> game.Balance
	7, // 12: game.LemonadeGame.Calculate:output_type -> game.CalculateResponse
	1, // 13: game.BotanicalGardenGame.Create:output_type -> game.CreateResult
	2, // 14: game.BotanicalGardenGame.RandomWeather:output_type -> game.Weather
	3, // 15: game.BotanicalGardenGame.GetBalance:output_type -> game.Balance
	7, // 16: game.BotanicalGardenGame.Calculate:output_type -> game.CalculateResponse
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResult); i {
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
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weather); i {
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
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameID); i {
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
		file_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRequest); i {
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
		file_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse); i {
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
		file_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
