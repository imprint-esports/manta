// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.2.0
// source: dota_broadcastmessages.proto

package dota

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

type EDotaBroadcastMessages int32

const (
	EDotaBroadcastMessages_DOTA_BM_LANLobbyRequest EDotaBroadcastMessages = 1
	EDotaBroadcastMessages_DOTA_BM_LANLobbyReply   EDotaBroadcastMessages = 2
)

// Enum value maps for EDotaBroadcastMessages.
var (
	EDotaBroadcastMessages_name = map[int32]string{
		1: "DOTA_BM_LANLobbyRequest",
		2: "DOTA_BM_LANLobbyReply",
	}
	EDotaBroadcastMessages_value = map[string]int32{
		"DOTA_BM_LANLobbyRequest": 1,
		"DOTA_BM_LANLobbyReply":   2,
	}
)

func (x EDotaBroadcastMessages) Enum() *EDotaBroadcastMessages {
	p := new(EDotaBroadcastMessages)
	*p = x
	return p
}

func (x EDotaBroadcastMessages) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EDotaBroadcastMessages) Descriptor() protoreflect.EnumDescriptor {
	return file_dota_broadcastmessages_proto_enumTypes[0].Descriptor()
}

func (EDotaBroadcastMessages) Type() protoreflect.EnumType {
	return &file_dota_broadcastmessages_proto_enumTypes[0]
}

func (x EDotaBroadcastMessages) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EDotaBroadcastMessages) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EDotaBroadcastMessages(num)
	return nil
}

// Deprecated: Use EDotaBroadcastMessages.Descriptor instead.
func (EDotaBroadcastMessages) EnumDescriptor() ([]byte, []int) {
	return file_dota_broadcastmessages_proto_rawDescGZIP(), []int{0}
}

type CDOTABroadcastMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *EDotaBroadcastMessages `protobuf:"varint,1,req,name=type,enum=dota.EDotaBroadcastMessages" json:"type,omitempty"`
	Msg  []byte                  `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (x *CDOTABroadcastMsg) Reset() {
	*x = CDOTABroadcastMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_broadcastmessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDOTABroadcastMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDOTABroadcastMsg) ProtoMessage() {}

func (x *CDOTABroadcastMsg) ProtoReflect() protoreflect.Message {
	mi := &file_dota_broadcastmessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDOTABroadcastMsg.ProtoReflect.Descriptor instead.
func (*CDOTABroadcastMsg) Descriptor() ([]byte, []int) {
	return file_dota_broadcastmessages_proto_rawDescGZIP(), []int{0}
}

func (x *CDOTABroadcastMsg) GetType() EDotaBroadcastMessages {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return EDotaBroadcastMessages_DOTA_BM_LANLobbyRequest
}

func (x *CDOTABroadcastMsg) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

type CDOTABroadcastMsg_LANLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CDOTABroadcastMsg_LANLobbyRequest) Reset() {
	*x = CDOTABroadcastMsg_LANLobbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_broadcastmessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDOTABroadcastMsg_LANLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDOTABroadcastMsg_LANLobbyRequest) ProtoMessage() {}

func (x *CDOTABroadcastMsg_LANLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dota_broadcastmessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDOTABroadcastMsg_LANLobbyRequest.ProtoReflect.Descriptor instead.
func (*CDOTABroadcastMsg_LANLobbyRequest) Descriptor() ([]byte, []int) {
	return file_dota_broadcastmessages_proto_rawDescGZIP(), []int{1}
}

type CDOTABroadcastMsg_LANLobbyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               *uint64                                         `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TournamentId     *uint32                                         `protobuf:"varint,2,opt,name=tournament_id,json=tournamentId" json:"tournament_id,omitempty"`
	TournamentGameId *uint32                                         `protobuf:"varint,3,opt,name=tournament_game_id,json=tournamentGameId" json:"tournament_game_id,omitempty"`
	Members          []*CDOTABroadcastMsg_LANLobbyReply_CLobbyMember `protobuf:"bytes,4,rep,name=members" json:"members,omitempty"`
	RequiresPassKey  *bool                                           `protobuf:"varint,5,opt,name=requires_pass_key,json=requiresPassKey" json:"requires_pass_key,omitempty"`
	LeaderAccountId  *uint32                                         `protobuf:"varint,6,opt,name=leader_account_id,json=leaderAccountId" json:"leader_account_id,omitempty"`
	GameMode         *uint32                                         `protobuf:"varint,7,opt,name=game_mode,json=gameMode" json:"game_mode,omitempty"`
	Name             *string                                         `protobuf:"bytes,8,opt,name=name" json:"name,omitempty"`
	Players          *uint32                                         `protobuf:"varint,9,opt,name=players" json:"players,omitempty"`
}

func (x *CDOTABroadcastMsg_LANLobbyReply) Reset() {
	*x = CDOTABroadcastMsg_LANLobbyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_broadcastmessages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDOTABroadcastMsg_LANLobbyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDOTABroadcastMsg_LANLobbyReply) ProtoMessage() {}

func (x *CDOTABroadcastMsg_LANLobbyReply) ProtoReflect() protoreflect.Message {
	mi := &file_dota_broadcastmessages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDOTABroadcastMsg_LANLobbyReply.ProtoReflect.Descriptor instead.
func (*CDOTABroadcastMsg_LANLobbyReply) Descriptor() ([]byte, []int) {
	return file_dota_broadcastmessages_proto_rawDescGZIP(), []int{2}
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetTournamentId() uint32 {
	if x != nil && x.TournamentId != nil {
		return *x.TournamentId
	}
	return 0
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetTournamentGameId() uint32 {
	if x != nil && x.TournamentGameId != nil {
		return *x.TournamentGameId
	}
	return 0
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetMembers() []*CDOTABroadcastMsg_LANLobbyReply_CLobbyMember {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetRequiresPassKey() bool {
	if x != nil && x.RequiresPassKey != nil {
		return *x.RequiresPassKey
	}
	return false
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetLeaderAccountId() uint32 {
	if x != nil && x.LeaderAccountId != nil {
		return *x.LeaderAccountId
	}
	return 0
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetGameMode() uint32 {
	if x != nil && x.GameMode != nil {
		return *x.GameMode
	}
	return 0
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CDOTABroadcastMsg_LANLobbyReply) GetPlayers() uint32 {
	if x != nil && x.Players != nil {
		return *x.Players
	}
	return 0
}

type CDOTABroadcastMsg_LANLobbyReply_CLobbyMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId  *uint32 `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	PlayerName *string `protobuf:"bytes,2,opt,name=player_name,json=playerName" json:"player_name,omitempty"`
}

func (x *CDOTABroadcastMsg_LANLobbyReply_CLobbyMember) Reset() {
	*x = CDOTABroadcastMsg_LANLobbyReply_CLobbyMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_broadcastmessages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDOTABroadcastMsg_LANLobbyReply_CLobbyMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDOTABroadcastMsg_LANLobbyReply_CLobbyMember) ProtoMessage() {}

func (x *CDOTABroadcastMsg_LANLobbyReply_CLobbyMember) ProtoReflect() protoreflect.Message {
	mi := &file_dota_broadcastmessages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDOTABroadcastMsg_LANLobbyReply_CLobbyMember.ProtoReflect.Descriptor instead.
func (*CDOTABroadcastMsg_LANLobbyReply_CLobbyMember) Descriptor() ([]byte, []int) {
	return file_dota_broadcastmessages_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CDOTABroadcastMsg_LANLobbyReply_CLobbyMember) GetAccountId() uint32 {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return 0
}

func (x *CDOTABroadcastMsg_LANLobbyReply_CLobbyMember) GetPlayerName() string {
	if x != nil && x.PlayerName != nil {
		return *x.PlayerName
	}
	return ""
}

var File_dota_broadcastmessages_proto protoreflect.FileDescriptor

var file_dota_broadcastmessages_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x64, 0x6f, 0x74, 0x61, 0x22, 0x57, 0x0a, 0x11, 0x43, 0x44, 0x4f, 0x54, 0x41, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x45,
	0x44, 0x6f, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x23, 0x0a,
	0x21, 0x43, 0x44, 0x4f, 0x54, 0x41, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d,
	0x73, 0x67, 0x5f, 0x4c, 0x41, 0x4e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xc5, 0x03, 0x0a, 0x1f, 0x43, 0x44, 0x4f, 0x54, 0x41, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x5f, 0x4c, 0x41, 0x4e, 0x4c, 0x6f, 0x62, 0x62,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x74,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x64, 0x6f, 0x74,
	0x61, 0x2e, 0x43, 0x44, 0x4f, 0x54, 0x41, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x5f, 0x4c, 0x41, 0x4e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x43, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x50, 0x61, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x1a, 0x4e, 0x0a, 0x0c, 0x43, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x50, 0x0a, 0x16, 0x45, 0x44,
	0x6f, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x42, 0x4d, 0x5f,
	0x4c, 0x41, 0x4e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x42, 0x4d, 0x5f, 0x4c, 0x41, 0x4e,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x10, 0x02, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x2d, 0x65, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x61,
	0x2f, 0x64, 0x6f, 0x74, 0x61, 0x3b, 0x64, 0x6f, 0x74, 0x61,
}

var (
	file_dota_broadcastmessages_proto_rawDescOnce sync.Once
	file_dota_broadcastmessages_proto_rawDescData = file_dota_broadcastmessages_proto_rawDesc
)

func file_dota_broadcastmessages_proto_rawDescGZIP() []byte {
	file_dota_broadcastmessages_proto_rawDescOnce.Do(func() {
		file_dota_broadcastmessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_dota_broadcastmessages_proto_rawDescData)
	})
	return file_dota_broadcastmessages_proto_rawDescData
}

var file_dota_broadcastmessages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dota_broadcastmessages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dota_broadcastmessages_proto_goTypes = []any{
	(EDotaBroadcastMessages)(0),                          // 0: dota.EDotaBroadcastMessages
	(*CDOTABroadcastMsg)(nil),                            // 1: dota.CDOTABroadcastMsg
	(*CDOTABroadcastMsg_LANLobbyRequest)(nil),            // 2: dota.CDOTABroadcastMsg_LANLobbyRequest
	(*CDOTABroadcastMsg_LANLobbyReply)(nil),              // 3: dota.CDOTABroadcastMsg_LANLobbyReply
	(*CDOTABroadcastMsg_LANLobbyReply_CLobbyMember)(nil), // 4: dota.CDOTABroadcastMsg_LANLobbyReply.CLobbyMember
}
var file_dota_broadcastmessages_proto_depIdxs = []int32{
	0, // 0: dota.CDOTABroadcastMsg.type:type_name -> dota.EDotaBroadcastMessages
	4, // 1: dota.CDOTABroadcastMsg_LANLobbyReply.members:type_name -> dota.CDOTABroadcastMsg_LANLobbyReply.CLobbyMember
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dota_broadcastmessages_proto_init() }
func file_dota_broadcastmessages_proto_init() {
	if File_dota_broadcastmessages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dota_broadcastmessages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CDOTABroadcastMsg); i {
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
		file_dota_broadcastmessages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CDOTABroadcastMsg_LANLobbyRequest); i {
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
		file_dota_broadcastmessages_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CDOTABroadcastMsg_LANLobbyReply); i {
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
		file_dota_broadcastmessages_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CDOTABroadcastMsg_LANLobbyReply_CLobbyMember); i {
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
			RawDescriptor: file_dota_broadcastmessages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dota_broadcastmessages_proto_goTypes,
		DependencyIndexes: file_dota_broadcastmessages_proto_depIdxs,
		EnumInfos:         file_dota_broadcastmessages_proto_enumTypes,
		MessageInfos:      file_dota_broadcastmessages_proto_msgTypes,
	}.Build()
	File_dota_broadcastmessages_proto = out.File
	file_dota_broadcastmessages_proto_rawDesc = nil
	file_dota_broadcastmessages_proto_goTypes = nil
	file_dota_broadcastmessages_proto_depIdxs = nil
}
