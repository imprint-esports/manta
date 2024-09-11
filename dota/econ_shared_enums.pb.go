// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.2.0
// source: econ_shared_enums.proto

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

type EGCEconBaseMsg int32

const (
	EGCEconBaseMsg_k_EMsgGCGenericResult EGCEconBaseMsg = 2579
)

// Enum value maps for EGCEconBaseMsg.
var (
	EGCEconBaseMsg_name = map[int32]string{
		2579: "k_EMsgGCGenericResult",
	}
	EGCEconBaseMsg_value = map[string]int32{
		"k_EMsgGCGenericResult": 2579,
	}
)

func (x EGCEconBaseMsg) Enum() *EGCEconBaseMsg {
	p := new(EGCEconBaseMsg)
	*p = x
	return p
}

func (x EGCEconBaseMsg) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EGCEconBaseMsg) Descriptor() protoreflect.EnumDescriptor {
	return file_econ_shared_enums_proto_enumTypes[0].Descriptor()
}

func (EGCEconBaseMsg) Type() protoreflect.EnumType {
	return &file_econ_shared_enums_proto_enumTypes[0]
}

func (x EGCEconBaseMsg) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EGCEconBaseMsg) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EGCEconBaseMsg(num)
	return nil
}

// Deprecated: Use EGCEconBaseMsg.Descriptor instead.
func (EGCEconBaseMsg) EnumDescriptor() ([]byte, []int) {
	return file_econ_shared_enums_proto_rawDescGZIP(), []int{0}
}

type EGCMsgResponse int32

const (
	EGCMsgResponse_k_EGCMsgResponseOK           EGCMsgResponse = 0
	EGCMsgResponse_k_EGCMsgResponseDenied       EGCMsgResponse = 1
	EGCMsgResponse_k_EGCMsgResponseServerError  EGCMsgResponse = 2
	EGCMsgResponse_k_EGCMsgResponseTimeout      EGCMsgResponse = 3
	EGCMsgResponse_k_EGCMsgResponseInvalid      EGCMsgResponse = 4
	EGCMsgResponse_k_EGCMsgResponseNoMatch      EGCMsgResponse = 5
	EGCMsgResponse_k_EGCMsgResponseUnknownError EGCMsgResponse = 6
	EGCMsgResponse_k_EGCMsgResponseNotLoggedOn  EGCMsgResponse = 7
	EGCMsgResponse_k_EGCMsgFailedToCreate       EGCMsgResponse = 8
)

// Enum value maps for EGCMsgResponse.
var (
	EGCMsgResponse_name = map[int32]string{
		0: "k_EGCMsgResponseOK",
		1: "k_EGCMsgResponseDenied",
		2: "k_EGCMsgResponseServerError",
		3: "k_EGCMsgResponseTimeout",
		4: "k_EGCMsgResponseInvalid",
		5: "k_EGCMsgResponseNoMatch",
		6: "k_EGCMsgResponseUnknownError",
		7: "k_EGCMsgResponseNotLoggedOn",
		8: "k_EGCMsgFailedToCreate",
	}
	EGCMsgResponse_value = map[string]int32{
		"k_EGCMsgResponseOK":           0,
		"k_EGCMsgResponseDenied":       1,
		"k_EGCMsgResponseServerError":  2,
		"k_EGCMsgResponseTimeout":      3,
		"k_EGCMsgResponseInvalid":      4,
		"k_EGCMsgResponseNoMatch":      5,
		"k_EGCMsgResponseUnknownError": 6,
		"k_EGCMsgResponseNotLoggedOn":  7,
		"k_EGCMsgFailedToCreate":       8,
	}
)

func (x EGCMsgResponse) Enum() *EGCMsgResponse {
	p := new(EGCMsgResponse)
	*p = x
	return p
}

func (x EGCMsgResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EGCMsgResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_econ_shared_enums_proto_enumTypes[1].Descriptor()
}

func (EGCMsgResponse) Type() protoreflect.EnumType {
	return &file_econ_shared_enums_proto_enumTypes[1]
}

func (x EGCMsgResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EGCMsgResponse) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EGCMsgResponse(num)
	return nil
}

// Deprecated: Use EGCMsgResponse.Descriptor instead.
func (EGCMsgResponse) EnumDescriptor() ([]byte, []int) {
	return file_econ_shared_enums_proto_rawDescGZIP(), []int{1}
}

type EGCMsgUseItemResponse int32

const (
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed                    EGCMsgUseItemResponse = 0
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_GiftNoOtherPlayers          EGCMsgUseItemResponse = 1
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ServerError                 EGCMsgUseItemResponse = 2
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_MiniGameAlreadyStarted      EGCMsgUseItemResponse = 3
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted       EGCMsgUseItemResponse = 4
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted EGCMsgUseItemResponse = 5
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_NotInLowPriorityPool        EGCMsgUseItemResponse = 6
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_NotHighEnoughLevel          EGCMsgUseItemResponse = 7
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EventNotActive              EGCMsgUseItemResponse = 8
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted EGCMsgUseItemResponse = 9
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_MissingRequirement          EGCMsgUseItemResponse = 10
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew        EGCMsgUseItemResponse = 11
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_EmoticonUnlock_Complete     EGCMsgUseItemResponse = 12
	EGCMsgUseItemResponse_k_EGCMsgUseItemResponse_ItemUsed_Compendium         EGCMsgUseItemResponse = 13
)

// Enum value maps for EGCMsgUseItemResponse.
var (
	EGCMsgUseItemResponse_name = map[int32]string{
		0:  "k_EGCMsgUseItemResponse_ItemUsed",
		1:  "k_EGCMsgUseItemResponse_GiftNoOtherPlayers",
		2:  "k_EGCMsgUseItemResponse_ServerError",
		3:  "k_EGCMsgUseItemResponse_MiniGameAlreadyStarted",
		4:  "k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted",
		5:  "k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted",
		6:  "k_EGCMsgUseItemResponse_NotInLowPriorityPool",
		7:  "k_EGCMsgUseItemResponse_NotHighEnoughLevel",
		8:  "k_EGCMsgUseItemResponse_EventNotActive",
		9:  "k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted",
		10: "k_EGCMsgUseItemResponse_MissingRequirement",
		11: "k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew",
		12: "k_EGCMsgUseItemResponse_EmoticonUnlock_Complete",
		13: "k_EGCMsgUseItemResponse_ItemUsed_Compendium",
	}
	EGCMsgUseItemResponse_value = map[string]int32{
		"k_EGCMsgUseItemResponse_ItemUsed":                    0,
		"k_EGCMsgUseItemResponse_GiftNoOtherPlayers":          1,
		"k_EGCMsgUseItemResponse_ServerError":                 2,
		"k_EGCMsgUseItemResponse_MiniGameAlreadyStarted":      3,
		"k_EGCMsgUseItemResponse_ItemUsed_ItemsGranted":       4,
		"k_EGCMsgUseItemResponse_DropRateBonusAlreadyGranted": 5,
		"k_EGCMsgUseItemResponse_NotInLowPriorityPool":        6,
		"k_EGCMsgUseItemResponse_NotHighEnoughLevel":          7,
		"k_EGCMsgUseItemResponse_EventNotActive":              8,
		"k_EGCMsgUseItemResponse_ItemUsed_EventPointsGranted": 9,
		"k_EGCMsgUseItemResponse_MissingRequirement":          10,
		"k_EGCMsgUseItemResponse_EmoticonUnlock_NoNew":        11,
		"k_EGCMsgUseItemResponse_EmoticonUnlock_Complete":     12,
		"k_EGCMsgUseItemResponse_ItemUsed_Compendium":         13,
	}
)

func (x EGCMsgUseItemResponse) Enum() *EGCMsgUseItemResponse {
	p := new(EGCMsgUseItemResponse)
	*p = x
	return p
}

func (x EGCMsgUseItemResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EGCMsgUseItemResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_econ_shared_enums_proto_enumTypes[2].Descriptor()
}

func (EGCMsgUseItemResponse) Type() protoreflect.EnumType {
	return &file_econ_shared_enums_proto_enumTypes[2]
}

func (x EGCMsgUseItemResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EGCMsgUseItemResponse) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EGCMsgUseItemResponse(num)
	return nil
}

// Deprecated: Use EGCMsgUseItemResponse.Descriptor instead.
func (EGCMsgUseItemResponse) EnumDescriptor() ([]byte, []int) {
	return file_econ_shared_enums_proto_rawDescGZIP(), []int{2}
}

type CMsgGenericResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eresult      *uint32 `protobuf:"varint,1,opt,name=eresult" json:"eresult,omitempty"`
	DebugMessage *string `protobuf:"bytes,2,opt,name=debug_message,json=debugMessage" json:"debug_message,omitempty"`
}

func (x *CMsgGenericResult) Reset() {
	*x = CMsgGenericResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_econ_shared_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgGenericResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgGenericResult) ProtoMessage() {}

func (x *CMsgGenericResult) ProtoReflect() protoreflect.Message {
	mi := &file_econ_shared_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgGenericResult.ProtoReflect.Descriptor instead.
func (*CMsgGenericResult) Descriptor() ([]byte, []int) {
	return file_econ_shared_enums_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgGenericResult) GetEresult() uint32 {
	if x != nil && x.Eresult != nil {
		return *x.Eresult
	}
	return 0
}

func (x *CMsgGenericResult) GetDebugMessage() string {
	if x != nil && x.DebugMessage != nil {
		return *x.DebugMessage
	}
	return ""
}

var File_econ_shared_enums_proto protoreflect.FileDescriptor

var file_econ_shared_enums_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x63, 0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x6f, 0x74, 0x61, 0x22,
	0x52, 0x0a, 0x11, 0x43, 0x4d, 0x73, 0x67, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x62, 0x75, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0x2c, 0x0a, 0x0e, 0x45, 0x47, 0x43, 0x45, 0x63, 0x6f, 0x6e, 0x42, 0x61,
	0x73, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x15, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47,
	0x43, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x10, 0x93,
	0x14, 0x2a, 0x9b, 0x02, 0x0a, 0x0e, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x6b, 0x5f, 0x45, 0x47,
	0x43, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x6b, 0x5f, 0x45,
	0x47, 0x43, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x05,
	0x12, 0x20, 0x0a, 0x1c, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x06, 0x12, 0x1f, 0x0a, 0x1b, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x4f,
	0x6e, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x08, 0x2a,
	0xc5, 0x05, 0x0a, 0x15, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x6b, 0x5f, 0x45,
	0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x2e, 0x0a, 0x2a, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x47, 0x69, 0x66, 0x74, 0x4e,
	0x6f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x10, 0x01, 0x12,
	0x27, 0x0a, 0x23, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x32, 0x0a, 0x2e, 0x6b, 0x5f, 0x45, 0x47,
	0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x4d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x03, 0x12, 0x31, 0x0a, 0x2d,
	0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64,
	0x5f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x10, 0x04, 0x12,
	0x37, 0x0a, 0x33, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x44, 0x72, 0x6f, 0x70, 0x52,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x10, 0x05, 0x12, 0x30, 0x0a, 0x2c, 0x6b, 0x5f, 0x45, 0x47,
	0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x4c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x10, 0x06, 0x12, 0x2e, 0x0a, 0x2a, 0x6b, 0x5f,
	0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x48, 0x69, 0x67, 0x68, 0x45, 0x6e, 0x6f,
	0x75, 0x67, 0x68, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x07, 0x12, 0x2a, 0x0a, 0x26, 0x6b, 0x5f,
	0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x10, 0x08, 0x12, 0x37, 0x0a, 0x33, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d,
	0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x5f, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x10, 0x09, 0x12,
	0x2e, 0x0a, 0x2a, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x0a, 0x12,
	0x30, 0x0a, 0x2c, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x45, 0x6d, 0x6f, 0x74, 0x69,
	0x63, 0x6f, 0x6e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x4e, 0x6f, 0x4e, 0x65, 0x77, 0x10,
	0x0b, 0x12, 0x33, 0x0a, 0x2f, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x45, 0x6d, 0x6f,
	0x74, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x10, 0x0c, 0x12, 0x2f, 0x0a, 0x2b, 0x6b, 0x5f, 0x45, 0x47, 0x43, 0x4d,
	0x73, 0x67, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x5f, 0x43, 0x6f, 0x6d, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x75, 0x6d, 0x10, 0x0d, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2d, 0x65, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x61, 0x2f, 0x64, 0x6f, 0x74, 0x61,
	0x3b, 0x64, 0x6f, 0x74, 0x61,
}

var (
	file_econ_shared_enums_proto_rawDescOnce sync.Once
	file_econ_shared_enums_proto_rawDescData = file_econ_shared_enums_proto_rawDesc
)

func file_econ_shared_enums_proto_rawDescGZIP() []byte {
	file_econ_shared_enums_proto_rawDescOnce.Do(func() {
		file_econ_shared_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_econ_shared_enums_proto_rawDescData)
	})
	return file_econ_shared_enums_proto_rawDescData
}

var file_econ_shared_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_econ_shared_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_econ_shared_enums_proto_goTypes = []any{
	(EGCEconBaseMsg)(0),        // 0: dota.EGCEconBaseMsg
	(EGCMsgResponse)(0),        // 1: dota.EGCMsgResponse
	(EGCMsgUseItemResponse)(0), // 2: dota.EGCMsgUseItemResponse
	(*CMsgGenericResult)(nil),  // 3: dota.CMsgGenericResult
}
var file_econ_shared_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_econ_shared_enums_proto_init() }
func file_econ_shared_enums_proto_init() {
	if File_econ_shared_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_econ_shared_enums_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgGenericResult); i {
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
			RawDescriptor: file_econ_shared_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_econ_shared_enums_proto_goTypes,
		DependencyIndexes: file_econ_shared_enums_proto_depIdxs,
		EnumInfos:         file_econ_shared_enums_proto_enumTypes,
		MessageInfos:      file_econ_shared_enums_proto_msgTypes,
	}.Build()
	File_econ_shared_enums_proto = out.File
	file_econ_shared_enums_proto_rawDesc = nil
	file_econ_shared_enums_proto_goTypes = nil
	file_econ_shared_enums_proto_depIdxs = nil
}
