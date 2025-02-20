// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: dota_usercmd.proto

package dota

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CDota2UserCmdPB struct {
	state                      protoimpl.MessageState `protogen:"open.v1"`
	Base                       *CBaseUserCmdPB        `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	SpectatorQueryUnitEntindex *int32                 `protobuf:"varint,2,opt,name=spectator_query_unit_entindex,json=spectatorQueryUnitEntindex" json:"spectator_query_unit_entindex,omitempty"`
	Crosshairtrace             *CMsgVector            `protobuf:"bytes,3,opt,name=crosshairtrace" json:"crosshairtrace,omitempty"`
	CamerapositionX            *int32                 `protobuf:"varint,4,opt,name=cameraposition_x,json=camerapositionX" json:"cameraposition_x,omitempty"`
	CamerapositionY            *int32                 `protobuf:"varint,5,opt,name=cameraposition_y,json=camerapositionY" json:"cameraposition_y,omitempty"`
	Clickbehavior              *uint32                `protobuf:"varint,6,opt,name=clickbehavior" json:"clickbehavior,omitempty"`
	Statspanel                 *uint32                `protobuf:"varint,7,opt,name=statspanel" json:"statspanel,omitempty"`
	Shoppanel                  *uint32                `protobuf:"varint,8,opt,name=shoppanel" json:"shoppanel,omitempty"`
	StatsDropdown              *uint32                `protobuf:"varint,9,opt,name=stats_dropdown,json=statsDropdown" json:"stats_dropdown,omitempty"`
	StatsDropdownSort          *uint32                `protobuf:"varint,10,opt,name=stats_dropdown_sort,json=statsDropdownSort" json:"stats_dropdown_sort,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *CDota2UserCmdPB) Reset() {
	*x = CDota2UserCmdPB{}
	mi := &file_dota_usercmd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CDota2UserCmdPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDota2UserCmdPB) ProtoMessage() {}

func (x *CDota2UserCmdPB) ProtoReflect() protoreflect.Message {
	mi := &file_dota_usercmd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDota2UserCmdPB.ProtoReflect.Descriptor instead.
func (*CDota2UserCmdPB) Descriptor() ([]byte, []int) {
	return file_dota_usercmd_proto_rawDescGZIP(), []int{0}
}

func (x *CDota2UserCmdPB) GetBase() *CBaseUserCmdPB {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CDota2UserCmdPB) GetSpectatorQueryUnitEntindex() int32 {
	if x != nil && x.SpectatorQueryUnitEntindex != nil {
		return *x.SpectatorQueryUnitEntindex
	}
	return 0
}

func (x *CDota2UserCmdPB) GetCrosshairtrace() *CMsgVector {
	if x != nil {
		return x.Crosshairtrace
	}
	return nil
}

func (x *CDota2UserCmdPB) GetCamerapositionX() int32 {
	if x != nil && x.CamerapositionX != nil {
		return *x.CamerapositionX
	}
	return 0
}

func (x *CDota2UserCmdPB) GetCamerapositionY() int32 {
	if x != nil && x.CamerapositionY != nil {
		return *x.CamerapositionY
	}
	return 0
}

func (x *CDota2UserCmdPB) GetClickbehavior() uint32 {
	if x != nil && x.Clickbehavior != nil {
		return *x.Clickbehavior
	}
	return 0
}

func (x *CDota2UserCmdPB) GetStatspanel() uint32 {
	if x != nil && x.Statspanel != nil {
		return *x.Statspanel
	}
	return 0
}

func (x *CDota2UserCmdPB) GetShoppanel() uint32 {
	if x != nil && x.Shoppanel != nil {
		return *x.Shoppanel
	}
	return 0
}

func (x *CDota2UserCmdPB) GetStatsDropdown() uint32 {
	if x != nil && x.StatsDropdown != nil {
		return *x.StatsDropdown
	}
	return 0
}

func (x *CDota2UserCmdPB) GetStatsDropdownSort() uint32 {
	if x != nil && x.StatsDropdownSort != nil {
		return *x.StatsDropdownSort
	}
	return 0
}

var File_dota_usercmd_proto protoreflect.FileDescriptor

var file_dota_usercmd_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6d, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x6f, 0x74, 0x61, 0x1a, 0x16, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc9, 0x03, 0x0a, 0x0f, 0x43, 0x44, 0x6f, 0x74, 0x61, 0x32, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6d, 0x64, 0x50, 0x42, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x42, 0x61, 0x73, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6d, 0x64, 0x50, 0x42, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x1d, 0x73, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1a, 0x73, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x6f,
	0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x38, 0x0a, 0x0e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x68, 0x61, 0x69, 0x72, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x74,
	0x61, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x63, 0x72,
	0x6f, 0x73, 0x73, 0x68, 0x61, 0x69, 0x72, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x58, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x59, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x61, 0x6e, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f,
	0x64, 0x72, 0x6f, 0x70, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x44, 0x72, 0x6f, 0x70, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x2e, 0x0a,
	0x13, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x64, 0x6f, 0x77, 0x6e, 0x5f,
	0x73, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x44, 0x72, 0x6f, 0x70, 0x64, 0x6f, 0x77, 0x6e, 0x53, 0x6f, 0x72, 0x74, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x2d, 0x65, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x74,
	0x61, 0x2f, 0x64, 0x6f, 0x74, 0x61, 0x3b, 0x64, 0x6f, 0x74, 0x61,
})

var (
	file_dota_usercmd_proto_rawDescOnce sync.Once
	file_dota_usercmd_proto_rawDescData []byte
)

func file_dota_usercmd_proto_rawDescGZIP() []byte {
	file_dota_usercmd_proto_rawDescOnce.Do(func() {
		file_dota_usercmd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_dota_usercmd_proto_rawDesc), len(file_dota_usercmd_proto_rawDesc)))
	})
	return file_dota_usercmd_proto_rawDescData
}

var file_dota_usercmd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dota_usercmd_proto_goTypes = []any{
	(*CDota2UserCmdPB)(nil), // 0: dota.CDota2UserCmdPB
	(*CBaseUserCmdPB)(nil),  // 1: dota.CBaseUserCmdPB
	(*CMsgVector)(nil),      // 2: dota.CMsgVector
}
var file_dota_usercmd_proto_depIdxs = []int32{
	1, // 0: dota.CDota2UserCmdPB.base:type_name -> dota.CBaseUserCmdPB
	2, // 1: dota.CDota2UserCmdPB.crosshairtrace:type_name -> dota.CMsgVector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dota_usercmd_proto_init() }
func file_dota_usercmd_proto_init() {
	if File_dota_usercmd_proto != nil {
		return
	}
	file_networkbasetypes_proto_init()
	file_usercmd_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_dota_usercmd_proto_rawDesc), len(file_dota_usercmd_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dota_usercmd_proto_goTypes,
		DependencyIndexes: file_dota_usercmd_proto_depIdxs,
		MessageInfos:      file_dota_usercmd_proto_msgTypes,
	}.Build()
	File_dota_usercmd_proto = out.File
	file_dota_usercmd_proto_goTypes = nil
	file_dota_usercmd_proto_depIdxs = nil
}
