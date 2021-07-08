// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: todoapp_service.proto

package protopack

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AddPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AddPostRequest) Reset() {
	*x = AddPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoapp_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostRequest) ProtoMessage() {}

func (x *AddPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todoapp_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostRequest.ProtoReflect.Descriptor instead.
func (*AddPostRequest) Descriptor() ([]byte, []int) {
	return file_todoapp_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddPostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddPostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AddPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *AddPostResponse) Reset() {
	*x = AddPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoapp_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostResponse) ProtoMessage() {}

func (x *AddPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todoapp_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostResponse.ProtoReflect.Descriptor instead.
func (*AddPostResponse) Descriptor() ([]byte, []int) {
	return file_todoapp_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddPostResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_todoapp_service_proto protoreflect.FileDescriptor

var file_todoapp_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61,
	0x63, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x31, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x6a, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x63,
	0x6b, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x64, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x61, 0x64, 0x64, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01,
	0x2a, 0x42, 0x24, 0x5a, 0x22, 0x74, 0x6f, 0x64, 0x6f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x20, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todoapp_service_proto_rawDescOnce sync.Once
	file_todoapp_service_proto_rawDescData = file_todoapp_service_proto_rawDesc
)

func file_todoapp_service_proto_rawDescGZIP() []byte {
	file_todoapp_service_proto_rawDescOnce.Do(func() {
		file_todoapp_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_todoapp_service_proto_rawDescData)
	})
	return file_todoapp_service_proto_rawDescData
}

var file_todoapp_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_todoapp_service_proto_goTypes = []interface{}{
	(*AddPostRequest)(nil),  // 0: protopack.AddPostRequest
	(*AddPostResponse)(nil), // 1: protopack.AddPostResponse
}
var file_todoapp_service_proto_depIdxs = []int32{
	0, // 0: protopack.ManagePostService.AddPost:input_type -> protopack.AddPostRequest
	1, // 1: protopack.ManagePostService.AddPost:output_type -> protopack.AddPostResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_todoapp_service_proto_init() }
func file_todoapp_service_proto_init() {
	if File_todoapp_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todoapp_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostRequest); i {
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
		file_todoapp_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostResponse); i {
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
			RawDescriptor: file_todoapp_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todoapp_service_proto_goTypes,
		DependencyIndexes: file_todoapp_service_proto_depIdxs,
		MessageInfos:      file_todoapp_service_proto_msgTypes,
	}.Build()
	File_todoapp_service_proto = out.File
	file_todoapp_service_proto_rawDesc = nil
	file_todoapp_service_proto_goTypes = nil
	file_todoapp_service_proto_depIdxs = nil
}
