// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: proto/service/service.proto

package service

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

type UidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User1 string `protobuf:"bytes,1,opt,name=user1,proto3" json:"user1,omitempty"`
	User2 string `protobuf:"bytes,2,opt,name=user2,proto3" json:"user2,omitempty"`
}

func (x *UidRequest) Reset() {
	*x = UidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UidRequest) ProtoMessage() {}

func (x *UidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UidRequest.ProtoReflect.Descriptor instead.
func (*UidRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_service_proto_rawDescGZIP(), []int{0}
}

func (x *UidRequest) GetUser1() string {
	if x != nil {
		return x.User1
	}
	return ""
}

func (x *UidRequest) GetUser2() string {
	if x != nil {
		return x.User2
	}
	return ""
}

type UidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UidResponse) Reset() {
	*x = UidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UidResponse) ProtoMessage() {}

func (x *UidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UidResponse.ProtoReflect.Descriptor instead.
func (*UidResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_service_proto_rawDescGZIP(), []int{1}
}

func (x *UidResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_service_service_proto protoreflect.FileDescriptor

var file_proto_service_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x0a, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x32,
	0x22, 0x27, 0x0a, 0x0b, 0x55, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x49, 0x0a, 0x0b, 0x55, 0x69, 0x64,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x55, 0x69, 0x64, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_service_proto_rawDescOnce sync.Once
	file_proto_service_service_proto_rawDescData = file_proto_service_service_proto_rawDesc
)

func file_proto_service_service_proto_rawDescGZIP() []byte {
	file_proto_service_service_proto_rawDescOnce.Do(func() {
		file_proto_service_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_service_proto_rawDescData)
	})
	return file_proto_service_service_proto_rawDescData
}

var file_proto_service_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_service_service_proto_goTypes = []interface{}{
	(*UidRequest)(nil),  // 0: service.UidRequest
	(*UidResponse)(nil), // 1: service.UidResponse
}
var file_proto_service_service_proto_depIdxs = []int32{
	0, // 0: service.UidGenerate.UidGenerate:input_type -> service.UidRequest
	1, // 1: service.UidGenerate.UidGenerate:output_type -> service.UidResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_service_proto_init() }
func file_proto_service_service_proto_init() {
	if File_proto_service_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UidRequest); i {
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
		file_proto_service_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UidResponse); i {
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
			RawDescriptor: file_proto_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_service_proto_goTypes,
		DependencyIndexes: file_proto_service_service_proto_depIdxs,
		MessageInfos:      file_proto_service_service_proto_msgTypes,
	}.Build()
	File_proto_service_service_proto = out.File
	file_proto_service_service_proto_rawDesc = nil
	file_proto_service_service_proto_goTypes = nil
	file_proto_service_service_proto_depIdxs = nil
}
