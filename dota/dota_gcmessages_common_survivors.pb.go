// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: dota_gcmessages_common_survivors.proto

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

type CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse int32

const (
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_k_eInternalError CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse = 0
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_k_eSuccess       CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse = 1
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_k_eTooBusy       CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse = 2
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_k_eDisabled      CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse = 3
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_k_eTimeout       CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse = 4
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_k_eNotAllowed    CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse = 5
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_k_eInvalidItem   CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse = 6
)

// Enum value maps for CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse.
var (
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse_name = map[int32]string{
		0: "k_eInternalError",
		1: "k_eSuccess",
		2: "k_eTooBusy",
		3: "k_eDisabled",
		4: "k_eTimeout",
		5: "k_eNotAllowed",
		6: "k_eInvalidItem",
	}
	CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse_value = map[string]int32{
		"k_eInternalError": 0,
		"k_eSuccess":       1,
		"k_eTooBusy":       2,
		"k_eDisabled":      3,
		"k_eTimeout":       4,
		"k_eNotAllowed":    5,
		"k_eInvalidItem":   6,
	}
)

func (x CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse) Enum() *CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse {
	p := new(CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse)
	*p = x
	return p
}

func (x CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_dota_gcmessages_common_survivors_proto_enumTypes[0].Descriptor()
}

func (CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse) Type() protoreflect.EnumType {
	return &file_dota_gcmessages_common_survivors_proto_enumTypes[0]
}

func (x CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse(num)
	return nil
}

// Deprecated: Use CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse.Descriptor instead.
func (CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse) EnumDescriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_survivors_proto_rawDescGZIP(), []int{3, 0}
}

type CMsgSurvivorsUserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributeLevels    []*CMsgSurvivorsUserData_AttributeLevelsEntry `protobuf:"bytes,1,rep,name=attribute_levels,json=attributeLevels" json:"attribute_levels,omitempty"`
	UnlockedDifficulty *uint32                                       `protobuf:"varint,2,opt,name=unlocked_difficulty,json=unlockedDifficulty" json:"unlocked_difficulty,omitempty"`
}

func (x *CMsgSurvivorsUserData) Reset() {
	*x = CMsgSurvivorsUserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSurvivorsUserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSurvivorsUserData) ProtoMessage() {}

func (x *CMsgSurvivorsUserData) ProtoReflect() protoreflect.Message {
	mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSurvivorsUserData.ProtoReflect.Descriptor instead.
func (*CMsgSurvivorsUserData) Descriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_survivors_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgSurvivorsUserData) GetAttributeLevels() []*CMsgSurvivorsUserData_AttributeLevelsEntry {
	if x != nil {
		return x.AttributeLevels
	}
	return nil
}

func (x *CMsgSurvivorsUserData) GetUnlockedDifficulty() uint32 {
	if x != nil && x.UnlockedDifficulty != nil {
		return *x.UnlockedDifficulty
	}
	return 0
}

type CMsgClientToGCSurvivorsPowerUpTelemetryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PowerupId    *uint32 `protobuf:"varint,1,opt,name=powerup_id,json=powerupId" json:"powerup_id,omitempty"`
	Level        *uint32 `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	TimeReceived *uint32 `protobuf:"varint,3,opt,name=time_received,json=timeReceived" json:"time_received,omitempty"`
	TimeHeld     *uint32 `protobuf:"varint,4,opt,name=time_held,json=timeHeld" json:"time_held,omitempty"`
	TotalDamage  *uint64 `protobuf:"varint,5,opt,name=total_damage,json=totalDamage" json:"total_damage,omitempty"`
	Dps          *uint32 `protobuf:"varint,6,opt,name=dps" json:"dps,omitempty"`
	HasScepter   *uint32 `protobuf:"varint,7,opt,name=has_scepter,json=hasScepter" json:"has_scepter,omitempty"`
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) Reset() {
	*x = CMsgClientToGCSurvivorsPowerUpTelemetryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientToGCSurvivorsPowerUpTelemetryData) ProtoMessage() {}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) ProtoReflect() protoreflect.Message {
	mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientToGCSurvivorsPowerUpTelemetryData.ProtoReflect.Descriptor instead.
func (*CMsgClientToGCSurvivorsPowerUpTelemetryData) Descriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_survivors_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) GetPowerupId() uint32 {
	if x != nil && x.PowerupId != nil {
		return *x.PowerupId
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) GetTimeReceived() uint32 {
	if x != nil && x.TimeReceived != nil {
		return *x.TimeReceived
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) GetTimeHeld() uint32 {
	if x != nil && x.TimeHeld != nil {
		return *x.TimeHeld
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) GetTotalDamage() uint64 {
	if x != nil && x.TotalDamage != nil {
		return *x.TotalDamage
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) GetDps() uint32 {
	if x != nil && x.Dps != nil {
		return *x.Dps
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsPowerUpTelemetryData) GetHasScepter() uint32 {
	if x != nil && x.HasScepter != nil {
		return *x.HasScepter
	}
	return 0
}

type CMsgClientToGCSurvivorsGameTelemetryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeSurvived         *uint32                                        `protobuf:"varint,1,opt,name=time_survived,json=timeSurvived" json:"time_survived,omitempty"`
	PlayerLevel          *uint32                                        `protobuf:"varint,2,opt,name=player_level,json=playerLevel" json:"player_level,omitempty"`
	GameResult           *uint32                                        `protobuf:"varint,3,opt,name=game_result,json=gameResult" json:"game_result,omitempty"`
	GoldEarned           *uint32                                        `protobuf:"varint,4,opt,name=gold_earned,json=goldEarned" json:"gold_earned,omitempty"`
	Powerups             []*CMsgClientToGCSurvivorsPowerUpTelemetryData `protobuf:"bytes,5,rep,name=powerups" json:"powerups,omitempty"`
	Difficulty           *uint32                                        `protobuf:"varint,6,opt,name=difficulty" json:"difficulty,omitempty"`
	MetaprogressionLevel *uint32                                        `protobuf:"varint,7,opt,name=metaprogression_level,json=metaprogressionLevel" json:"metaprogression_level,omitempty"`
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) Reset() {
	*x = CMsgClientToGCSurvivorsGameTelemetryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientToGCSurvivorsGameTelemetryData) ProtoMessage() {}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) ProtoReflect() protoreflect.Message {
	mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientToGCSurvivorsGameTelemetryData.ProtoReflect.Descriptor instead.
func (*CMsgClientToGCSurvivorsGameTelemetryData) Descriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_survivors_proto_rawDescGZIP(), []int{2}
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) GetTimeSurvived() uint32 {
	if x != nil && x.TimeSurvived != nil {
		return *x.TimeSurvived
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) GetPlayerLevel() uint32 {
	if x != nil && x.PlayerLevel != nil {
		return *x.PlayerLevel
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) GetGameResult() uint32 {
	if x != nil && x.GameResult != nil {
		return *x.GameResult
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) GetGoldEarned() uint32 {
	if x != nil && x.GoldEarned != nil {
		return *x.GoldEarned
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) GetPowerups() []*CMsgClientToGCSurvivorsPowerUpTelemetryData {
	if x != nil {
		return x.Powerups
	}
	return nil
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) GetDifficulty() uint32 {
	if x != nil && x.Difficulty != nil {
		return *x.Difficulty
	}
	return 0
}

func (x *CMsgClientToGCSurvivorsGameTelemetryData) GetMetaprogressionLevel() uint32 {
	if x != nil && x.MetaprogressionLevel != nil {
		return *x.MetaprogressionLevel
	}
	return 0
}

type CMsgClientToGCSurvivorsGameTelemetryDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse `protobuf:"varint,1,opt,name=response,enum=dota.CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse" json:"response,omitempty"`
}

func (x *CMsgClientToGCSurvivorsGameTelemetryDataResponse) Reset() {
	*x = CMsgClientToGCSurvivorsGameTelemetryDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientToGCSurvivorsGameTelemetryDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientToGCSurvivorsGameTelemetryDataResponse) ProtoMessage() {}

func (x *CMsgClientToGCSurvivorsGameTelemetryDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientToGCSurvivorsGameTelemetryDataResponse.ProtoReflect.Descriptor instead.
func (*CMsgClientToGCSurvivorsGameTelemetryDataResponse) Descriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_survivors_proto_rawDescGZIP(), []int{3}
}

func (x *CMsgClientToGCSurvivorsGameTelemetryDataResponse) GetResponse() CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse {
	if x != nil && x.Response != nil {
		return *x.Response
	}
	return CMsgClientToGCSurvivorsGameTelemetryDataResponse_k_eInternalError
}

type CMsgSurvivorsUserData_AttributeLevelsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *int32  `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	Value *uint32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (x *CMsgSurvivorsUserData_AttributeLevelsEntry) Reset() {
	*x = CMsgSurvivorsUserData_AttributeLevelsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSurvivorsUserData_AttributeLevelsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSurvivorsUserData_AttributeLevelsEntry) ProtoMessage() {}

func (x *CMsgSurvivorsUserData_AttributeLevelsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_dota_gcmessages_common_survivors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSurvivorsUserData_AttributeLevelsEntry.ProtoReflect.Descriptor instead.
func (*CMsgSurvivorsUserData_AttributeLevelsEntry) Descriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_survivors_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CMsgSurvivorsUserData_AttributeLevelsEntry) GetKey() int32 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *CMsgSurvivorsUserData_AttributeLevelsEntry) GetValue() uint32 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

var File_dota_gcmessages_common_survivors_proto protoreflect.FileDescriptor

var file_dota_gcmessages_common_survivors_proto_rawDesc = []byte{
	0x0a, 0x26, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x67, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x72, 0x76, 0x69, 0x76, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x6f, 0x74, 0x61, 0x1a, 0x13,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x64, 0x6f,
	0x74, 0x61, 0x5f, 0x67, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x63, 0x73, 0x64,
	0x6b, 0x5f, 0x67, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x15, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x75, 0x72, 0x76, 0x69,
	0x76, 0x6f, 0x72, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5b, 0x0a, 0x10,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x4d,
	0x73, 0x67, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x6f, 0x72, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x1a, 0x3e, 0x0a, 0x14, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x2b, 0x43,
	0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x47, 0x43, 0x53, 0x75, 0x72,
	0x76, 0x69, 0x76, 0x6f, 0x72, 0x73, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x70, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x75, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x68, 0x65, 0x6c,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x48, 0x65, 0x6c,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x64, 0x70, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x5f, 0x73, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x68, 0x61, 0x73,
	0x53, 0x63, 0x65, 0x70, 0x74, 0x65, 0x72, 0x22, 0xd8, 0x02, 0x0a, 0x28, 0x43, 0x4d, 0x73, 0x67,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x47, 0x43, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76,
	0x6f, 0x72, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x75, 0x72,
	0x76, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x69, 0x6d,
	0x65, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x67, 0x6f, 0x6c, 0x64, 0x5f, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x67, 0x6f, 0x6c, 0x64, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x4d,
	0x0a, 0x08, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x6f, 0x47, 0x43, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x6f, 0x72, 0x73, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x55, 0x70, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x75, 0x70, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x33, 0x0a,
	0x15, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x6d, 0x65,
	0x74, 0x61, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x9c, 0x02, 0x0a, 0x30, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x6f, 0x47, 0x43, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x6f, 0x72, 0x73, 0x47, 0x61,
	0x6d, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e, 0x64, 0x6f, 0x74, 0x61,
	0x2e, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x47, 0x43, 0x53,
	0x75, 0x72, 0x76, 0x69, 0x76, 0x6f, 0x72, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x6b, 0x5f, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x6b, 0x5f, 0x65,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x6b, 0x5f, 0x65,
	0x54, 0x6f, 0x6f, 0x42, 0x75, 0x73, 0x79, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x6b, 0x5f, 0x65,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x6b, 0x5f,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x5f,
	0x65, 0x4e, 0x6f, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x10, 0x05, 0x12, 0x12, 0x0a,
	0x0e, 0x6b, 0x5f, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x10,
	0x06, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x6f, 0x74, 0x61, 0x62, 0x75, 0x66, 0x66, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x61, 0x2f, 0x64,
	0x6f, 0x74, 0x61, 0x3b, 0x64, 0x6f, 0x74, 0x61,
}

var (
	file_dota_gcmessages_common_survivors_proto_rawDescOnce sync.Once
	file_dota_gcmessages_common_survivors_proto_rawDescData = file_dota_gcmessages_common_survivors_proto_rawDesc
)

func file_dota_gcmessages_common_survivors_proto_rawDescGZIP() []byte {
	file_dota_gcmessages_common_survivors_proto_rawDescOnce.Do(func() {
		file_dota_gcmessages_common_survivors_proto_rawDescData = protoimpl.X.CompressGZIP(file_dota_gcmessages_common_survivors_proto_rawDescData)
	})
	return file_dota_gcmessages_common_survivors_proto_rawDescData
}

var file_dota_gcmessages_common_survivors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dota_gcmessages_common_survivors_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dota_gcmessages_common_survivors_proto_goTypes = []interface{}{
	(CMsgClientToGCSurvivorsGameTelemetryDataResponse_EResponse)(0), // 0: dota.CMsgClientToGCSurvivorsGameTelemetryDataResponse.EResponse
	(*CMsgSurvivorsUserData)(nil),                                   // 1: dota.CMsgSurvivorsUserData
	(*CMsgClientToGCSurvivorsPowerUpTelemetryData)(nil),             // 2: dota.CMsgClientToGCSurvivorsPowerUpTelemetryData
	(*CMsgClientToGCSurvivorsGameTelemetryData)(nil),                // 3: dota.CMsgClientToGCSurvivorsGameTelemetryData
	(*CMsgClientToGCSurvivorsGameTelemetryDataResponse)(nil),        // 4: dota.CMsgClientToGCSurvivorsGameTelemetryDataResponse
	(*CMsgSurvivorsUserData_AttributeLevelsEntry)(nil),              // 5: dota.CMsgSurvivorsUserData.AttributeLevelsEntry
}
var file_dota_gcmessages_common_survivors_proto_depIdxs = []int32{
	5, // 0: dota.CMsgSurvivorsUserData.attribute_levels:type_name -> dota.CMsgSurvivorsUserData.AttributeLevelsEntry
	2, // 1: dota.CMsgClientToGCSurvivorsGameTelemetryData.powerups:type_name -> dota.CMsgClientToGCSurvivorsPowerUpTelemetryData
	0, // 2: dota.CMsgClientToGCSurvivorsGameTelemetryDataResponse.response:type_name -> dota.CMsgClientToGCSurvivorsGameTelemetryDataResponse.EResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dota_gcmessages_common_survivors_proto_init() }
func file_dota_gcmessages_common_survivors_proto_init() {
	if File_dota_gcmessages_common_survivors_proto != nil {
		return
	}
	file_steammessages_proto_init()
	file_dota_shared_enums_proto_init()
	file_dota_gcmessages_common_proto_init()
	file_gcsdk_gcmessages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dota_gcmessages_common_survivors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgSurvivorsUserData); i {
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
		file_dota_gcmessages_common_survivors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientToGCSurvivorsPowerUpTelemetryData); i {
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
		file_dota_gcmessages_common_survivors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientToGCSurvivorsGameTelemetryData); i {
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
		file_dota_gcmessages_common_survivors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgClientToGCSurvivorsGameTelemetryDataResponse); i {
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
		file_dota_gcmessages_common_survivors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMsgSurvivorsUserData_AttributeLevelsEntry); i {
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
			RawDescriptor: file_dota_gcmessages_common_survivors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dota_gcmessages_common_survivors_proto_goTypes,
		DependencyIndexes: file_dota_gcmessages_common_survivors_proto_depIdxs,
		EnumInfos:         file_dota_gcmessages_common_survivors_proto_enumTypes,
		MessageInfos:      file_dota_gcmessages_common_survivors_proto_msgTypes,
	}.Build()
	File_dota_gcmessages_common_survivors_proto = out.File
	file_dota_gcmessages_common_survivors_proto_rawDesc = nil
	file_dota_gcmessages_common_survivors_proto_goTypes = nil
	file_dota_gcmessages_common_survivors_proto_depIdxs = nil
}
