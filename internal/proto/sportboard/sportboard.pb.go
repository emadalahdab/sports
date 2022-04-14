// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.7
// source: sportboard/sportboard.proto

package sportboard

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled            bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	FavoriteHidden     bool `protobuf:"varint,2,opt,name=favorite_hidden,json=favoriteHidden,proto3" json:"favorite_hidden,omitempty"`
	FavoriteSticky     bool `protobuf:"varint,3,opt,name=favorite_sticky,json=favoriteSticky,proto3" json:"favorite_sticky,omitempty"`
	ScrollEnabled      bool `protobuf:"varint,4,opt,name=scroll_enabled,json=scrollEnabled,proto3" json:"scroll_enabled,omitempty"`
	TightScrollEnabled bool `protobuf:"varint,5,opt,name=tight_scroll_enabled,json=tightScrollEnabled,proto3" json:"tight_scroll_enabled,omitempty"`
	RecordRankEnabled  bool `protobuf:"varint,6,opt,name=record_rank_enabled,json=recordRankEnabled,proto3" json:"record_rank_enabled,omitempty"`
	OddsEnabled        bool `protobuf:"varint,7,opt,name=odds_enabled,json=oddsEnabled,proto3" json:"odds_enabled,omitempty"`
	UseGradient        bool `protobuf:"varint,8,opt,name=use_gradient,json=useGradient,proto3" json:"use_gradient,omitempty"`
	LiveOnly           bool `protobuf:"varint,9,opt,name=live_only,json=liveOnly,proto3" json:"live_only,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sportboard_sportboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_sportboard_sportboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_sportboard_sportboard_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Status) GetFavoriteHidden() bool {
	if x != nil {
		return x.FavoriteHidden
	}
	return false
}

func (x *Status) GetFavoriteSticky() bool {
	if x != nil {
		return x.FavoriteSticky
	}
	return false
}

func (x *Status) GetScrollEnabled() bool {
	if x != nil {
		return x.ScrollEnabled
	}
	return false
}

func (x *Status) GetTightScrollEnabled() bool {
	if x != nil {
		return x.TightScrollEnabled
	}
	return false
}

func (x *Status) GetRecordRankEnabled() bool {
	if x != nil {
		return x.RecordRankEnabled
	}
	return false
}

func (x *Status) GetOddsEnabled() bool {
	if x != nil {
		return x.OddsEnabled
	}
	return false
}

func (x *Status) GetUseGradient() bool {
	if x != nil {
		return x.UseGradient
	}
	return false
}

func (x *Status) GetLiveOnly() bool {
	if x != nil {
		return x.LiveOnly
	}
	return false
}

type SetStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetStatusReq) Reset() {
	*x = SetStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sportboard_sportboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusReq) ProtoMessage() {}

func (x *SetStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_sportboard_sportboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusReq.ProtoReflect.Descriptor instead.
func (*SetStatusReq) Descriptor() ([]byte, []int) {
	return file_sportboard_sportboard_proto_rawDescGZIP(), []int{1}
}

func (x *SetStatusReq) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type StatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusResp) Reset() {
	*x = StatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sportboard_sportboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResp) ProtoMessage() {}

func (x *StatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_sportboard_sportboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResp.ProtoReflect.Descriptor instead.
func (*StatusResp) Descriptor() ([]byte, []int) {
	return file_sportboard_sportboard_proto_rawDescGZIP(), []int{2}
}

func (x *StatusResp) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_sportboard_sportboard_proto protoreflect.FileDescriptor

var file_sportboard_sportboard_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x48, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73,
	0x74, 0x69, 0x63, 0x6b, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x63, 0x72, 0x6f,
	0x6c, 0x6c, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x74, 0x69, 0x67, 0x68, 0x74, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x72,
	0x61, 0x6e, 0x6b, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x64, 0x64, 0x73, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6f, 0x64, 0x64, 0x73,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x5f, 0x67,
	0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x69,
	0x76, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6c,
	0x69, 0x76, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x38, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x36, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x7f, 0x0a, 0x05, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x39, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x62, 0x62, 0x79, 0x64, 0x79,
	0x65, 0x72, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sportboard_sportboard_proto_rawDescOnce sync.Once
	file_sportboard_sportboard_proto_rawDescData = file_sportboard_sportboard_proto_rawDesc
)

func file_sportboard_sportboard_proto_rawDescGZIP() []byte {
	file_sportboard_sportboard_proto_rawDescOnce.Do(func() {
		file_sportboard_sportboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_sportboard_sportboard_proto_rawDescData)
	})
	return file_sportboard_sportboard_proto_rawDescData
}

var file_sportboard_sportboard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sportboard_sportboard_proto_goTypes = []interface{}{
	(*Status)(nil),       // 0: sport.v1.Status
	(*SetStatusReq)(nil), // 1: sport.v1.SetStatusReq
	(*StatusResp)(nil),   // 2: sport.v1.StatusResp
	(*empty.Empty)(nil),  // 3: google.protobuf.Empty
}
var file_sportboard_sportboard_proto_depIdxs = []int32{
	0, // 0: sport.v1.SetStatusReq.status:type_name -> sport.v1.Status
	0, // 1: sport.v1.StatusResp.status:type_name -> sport.v1.Status
	1, // 2: sport.v1.Sport.SetStatus:input_type -> sport.v1.SetStatusReq
	3, // 3: sport.v1.Sport.GetStatus:input_type -> google.protobuf.Empty
	3, // 4: sport.v1.Sport.SetStatus:output_type -> google.protobuf.Empty
	2, // 5: sport.v1.Sport.GetStatus:output_type -> sport.v1.StatusResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sportboard_sportboard_proto_init() }
func file_sportboard_sportboard_proto_init() {
	if File_sportboard_sportboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sportboard_sportboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_sportboard_sportboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStatusReq); i {
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
		file_sportboard_sportboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResp); i {
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
			RawDescriptor: file_sportboard_sportboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sportboard_sportboard_proto_goTypes,
		DependencyIndexes: file_sportboard_sportboard_proto_depIdxs,
		MessageInfos:      file_sportboard_sportboard_proto_msgTypes,
	}.Build()
	File_sportboard_sportboard_proto = out.File
	file_sportboard_sportboard_proto_rawDesc = nil
	file_sportboard_sportboard_proto_goTypes = nil
	file_sportboard_sportboard_proto_depIdxs = nil
}
