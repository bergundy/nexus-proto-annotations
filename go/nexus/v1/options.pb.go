// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: nexus/v1/options.proto

package nexusv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperationOptions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Nexus operation name (defaults to proto method name).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional start result for the operation, a proto message name.
	StartResult   string `protobuf:"bytes,2,opt,name=start_result,json=startResult,proto3" json:"start_result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperationOptions) Reset() {
	*x = OperationOptions{}
	mi := &file_nexus_v1_options_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationOptions) ProtoMessage() {}

func (x *OperationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_nexus_v1_options_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationOptions.ProtoReflect.Descriptor instead.
func (*OperationOptions) Descriptor() ([]byte, []int) {
	return file_nexus_v1_options_proto_rawDescGZIP(), []int{0}
}

func (x *OperationOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OperationOptions) GetStartResult() string {
	if x != nil {
		return x.StartResult
	}
	return ""
}

type ServiceOptions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Nexus service name (defaults to proto service full name).
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceOptions) Reset() {
	*x = ServiceOptions{}
	mi := &file_nexus_v1_options_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOptions) ProtoMessage() {}

func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_nexus_v1_options_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOptions.ProtoReflect.Descriptor instead.
func (*ServiceOptions) Descriptor() ([]byte, []int) {
	return file_nexus_v1_options_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var file_nexus_v1_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ServiceOptions)(nil),
		Field:         8233,
		Name:          "nexus.v1.service",
		Tag:           "bytes,8233,opt,name=service",
		Filename:      "nexus/v1/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*OperationOptions)(nil),
		Field:         8234,
		Name:          "nexus.v1.operation",
		Tag:           "bytes,8234,opt,name=operation",
		Filename:      "nexus/v1/options.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional nexus.v1.ServiceOptions service = 8233;
	E_Service = &file_nexus_v1_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional nexus.v1.OperationOptions operation = 8234;
	E_Operation = &file_nexus_v1_options_proto_extTypes[1]
)

var File_nexus_v1_options_proto protoreflect.FileDescriptor

var file_nexus_v1_options_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x24, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x57, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xa9, 0x40, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x78, 0x75, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x5c,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaa, 0x40, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x9e, 0x01, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x72, 0x67, 0x75, 0x6e,
	0x64, 0x79, 0x2f, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x65,
	0x78, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x08, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x4e, 0x65, 0x78,
	0x75, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x09, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nexus_v1_options_proto_rawDescOnce sync.Once
	file_nexus_v1_options_proto_rawDescData = file_nexus_v1_options_proto_rawDesc
)

func file_nexus_v1_options_proto_rawDescGZIP() []byte {
	file_nexus_v1_options_proto_rawDescOnce.Do(func() {
		file_nexus_v1_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_nexus_v1_options_proto_rawDescData)
	})
	return file_nexus_v1_options_proto_rawDescData
}

var file_nexus_v1_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nexus_v1_options_proto_goTypes = []any{
	(*OperationOptions)(nil),            // 0: nexus.v1.OperationOptions
	(*ServiceOptions)(nil),              // 1: nexus.v1.ServiceOptions
	(*descriptorpb.ServiceOptions)(nil), // 2: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 3: google.protobuf.MethodOptions
}
var file_nexus_v1_options_proto_depIdxs = []int32{
	2, // 0: nexus.v1.service:extendee -> google.protobuf.ServiceOptions
	3, // 1: nexus.v1.operation:extendee -> google.protobuf.MethodOptions
	1, // 2: nexus.v1.service:type_name -> nexus.v1.ServiceOptions
	0, // 3: nexus.v1.operation:type_name -> nexus.v1.OperationOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nexus_v1_options_proto_init() }
func file_nexus_v1_options_proto_init() {
	if File_nexus_v1_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nexus_v1_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_nexus_v1_options_proto_goTypes,
		DependencyIndexes: file_nexus_v1_options_proto_depIdxs,
		MessageInfos:      file_nexus_v1_options_proto_msgTypes,
		ExtensionInfos:    file_nexus_v1_options_proto_extTypes,
	}.Build()
	File_nexus_v1_options_proto = out.File
	file_nexus_v1_options_proto_rawDesc = nil
	file_nexus_v1_options_proto_goTypes = nil
	file_nexus_v1_options_proto_depIdxs = nil
}
