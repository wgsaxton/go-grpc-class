// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/config.proto

package proto

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

type LongRunningRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LongRunningRequest) Reset() {
	*x = LongRunningRequest{}
	mi := &file_proto_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LongRunningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongRunningRequest) ProtoMessage() {}

func (x *LongRunningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongRunningRequest.ProtoReflect.Descriptor instead.
func (*LongRunningRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{0}
}

type LongRunningResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LongRunningResponse) Reset() {
	*x = LongRunningResponse{}
	mi := &file_proto_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LongRunningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongRunningResponse) ProtoMessage() {}

func (x *LongRunningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongRunningResponse.ProtoReflect.Descriptor instead.
func (*LongRunningResponse) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{1}
}

type FlakyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlakyRequest) Reset() {
	*x = FlakyRequest{}
	mi := &file_proto_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlakyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlakyRequest) ProtoMessage() {}

func (x *FlakyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlakyRequest.ProtoReflect.Descriptor instead.
func (*FlakyRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{2}
}

type FlakyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlakyResponse) Reset() {
	*x = FlakyResponse{}
	mi := &file_proto_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlakyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlakyResponse) ProtoMessage() {}

func (x *FlakyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlakyResponse.ProtoReflect.Descriptor instead.
func (*FlakyResponse) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{3}
}

type GetServerAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServerAddressRequest) Reset() {
	*x = GetServerAddressRequest{}
	mi := &file_proto_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServerAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerAddressRequest) ProtoMessage() {}

func (x *GetServerAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerAddressRequest.ProtoReflect.Descriptor instead.
func (*GetServerAddressRequest) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{4}
}

type GetServerAddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetServerAddressResponse) Reset() {
	*x = GetServerAddressResponse{}
	mi := &file_proto_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServerAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerAddressResponse) ProtoMessage() {}

func (x *GetServerAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerAddressResponse.ProtoReflect.Descriptor instead.
func (*GetServerAddressResponse) Descriptor() ([]byte, []int) {
	return file_proto_config_proto_rawDescGZIP(), []int{5}
}

func (x *GetServerAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_proto_config_proto protoreflect.FileDescriptor

var file_proto_config_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x14, 0x0a, 0x12,
	0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x46, 0x6c, 0x61,
	0x6b, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x46, 0x6c, 0x61,
	0x6b, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xe4, 0x01, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x0b, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x46, 0x6c, 0x61, 0x6b, 0x79, 0x12, 0x14,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x6c, 0x61, 0x6b, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x6c,
	0x61, 0x6b, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x67, 0x73, 0x61, 0x78, 0x74, 0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x36, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_config_proto_rawDescOnce sync.Once
	file_proto_config_proto_rawDescData []byte
)

func file_proto_config_proto_rawDescGZIP() []byte {
	file_proto_config_proto_rawDescOnce.Do(func() {
		file_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_config_proto_rawDesc), len(file_proto_config_proto_rawDesc)))
	})
	return file_proto_config_proto_rawDescData
}

var file_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_config_proto_goTypes = []any{
	(*LongRunningRequest)(nil),       // 0: config.LongRunningRequest
	(*LongRunningResponse)(nil),      // 1: config.LongRunningResponse
	(*FlakyRequest)(nil),             // 2: config.FlakyRequest
	(*FlakyResponse)(nil),            // 3: config.FlakyResponse
	(*GetServerAddressRequest)(nil),  // 4: config.GetServerAddressRequest
	(*GetServerAddressResponse)(nil), // 5: config.GetServerAddressResponse
}
var file_proto_config_proto_depIdxs = []int32{
	0, // 0: config.ConfigService.LongRunning:input_type -> config.LongRunningRequest
	2, // 1: config.ConfigService.Flaky:input_type -> config.FlakyRequest
	4, // 2: config.ConfigService.GetServerAddress:input_type -> config.GetServerAddressRequest
	1, // 3: config.ConfigService.LongRunning:output_type -> config.LongRunningResponse
	3, // 4: config.ConfigService.Flaky:output_type -> config.FlakyResponse
	5, // 5: config.ConfigService.GetServerAddress:output_type -> config.GetServerAddressResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_config_proto_init() }
func file_proto_config_proto_init() {
	if File_proto_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_config_proto_rawDesc), len(file_proto_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_config_proto_goTypes,
		DependencyIndexes: file_proto_config_proto_depIdxs,
		MessageInfos:      file_proto_config_proto_msgTypes,
	}.Build()
	File_proto_config_proto = out.File
	file_proto_config_proto_goTypes = nil
	file_proto_config_proto_depIdxs = nil
}
