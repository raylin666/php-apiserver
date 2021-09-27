// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: services/link/v1/short_link.proto

package v1

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

type GenerateShortUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GenerateShortUrlRequest) Reset() {
	*x = GenerateShortUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_link_v1_short_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateShortUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateShortUrlRequest) ProtoMessage() {}

func (x *GenerateShortUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_link_v1_short_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateShortUrlRequest.ProtoReflect.Descriptor instead.
func (*GenerateShortUrlRequest) Descriptor() ([]byte, []int) {
	return file_services_link_v1_short_link_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateShortUrlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GenerateShortUrlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GenerateShortUrlReply) Reset() {
	*x = GenerateShortUrlReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_link_v1_short_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateShortUrlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateShortUrlReply) ProtoMessage() {}

func (x *GenerateShortUrlReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_link_v1_short_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateShortUrlReply.ProtoReflect.Descriptor instead.
func (*GenerateShortUrlReply) Descriptor() ([]byte, []int) {
	return file_services_link_v1_short_link_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateShortUrlReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TransformLongUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TransformLongUrlRequest) Reset() {
	*x = TransformLongUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_link_v1_short_link_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformLongUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformLongUrlRequest) ProtoMessage() {}

func (x *TransformLongUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_link_v1_short_link_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformLongUrlRequest.ProtoReflect.Descriptor instead.
func (*TransformLongUrlRequest) Descriptor() ([]byte, []int) {
	return file_services_link_v1_short_link_proto_rawDescGZIP(), []int{2}
}

func (x *TransformLongUrlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TransformLongUrlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TransformLongUrlReply) Reset() {
	*x = TransformLongUrlReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_link_v1_short_link_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformLongUrlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformLongUrlReply) ProtoMessage() {}

func (x *TransformLongUrlReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_link_v1_short_link_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformLongUrlReply.ProtoReflect.Descriptor instead.
func (*TransformLongUrlReply) Descriptor() ([]byte, []int) {
	return file_services_link_v1_short_link_proto_rawDescGZIP(), []int{3}
}

func (x *TransformLongUrlReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_services_link_v1_short_link_proto protoreflect.FileDescriptor

var file_services_link_v1_short_link_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x29, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2b, 0x0a, 0x17, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x29, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x32, 0x9d, 0x02, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x86, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x22, 0x13, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x2e, 0x75, 0x72, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x10, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12,
	0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x6f, 0x6e, 0x67,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x2e, 0x75, 0x72, 0x6c,
	0x3a, 0x01, 0x2a, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x36, 0x36, 0x36, 0x2f, 0x67, 0x6f, 0x2d, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_link_v1_short_link_proto_rawDescOnce sync.Once
	file_services_link_v1_short_link_proto_rawDescData = file_services_link_v1_short_link_proto_rawDesc
)

func file_services_link_v1_short_link_proto_rawDescGZIP() []byte {
	file_services_link_v1_short_link_proto_rawDescOnce.Do(func() {
		file_services_link_v1_short_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_link_v1_short_link_proto_rawDescData)
	})
	return file_services_link_v1_short_link_proto_rawDescData
}

var file_services_link_v1_short_link_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_link_v1_short_link_proto_goTypes = []interface{}{
	(*GenerateShortUrlRequest)(nil), // 0: services.link.v1.GenerateShortUrlRequest
	(*GenerateShortUrlReply)(nil),   // 1: services.link.v1.GenerateShortUrlReply
	(*TransformLongUrlRequest)(nil), // 2: services.link.v1.TransformLongUrlRequest
	(*TransformLongUrlReply)(nil),   // 3: services.link.v1.TransformLongUrlReply
}
var file_services_link_v1_short_link_proto_depIdxs = []int32{
	0, // 0: services.link.v1.ShortLink.GenerateShortUrl:input_type -> services.link.v1.GenerateShortUrlRequest
	2, // 1: services.link.v1.ShortLink.TransformLongUrl:input_type -> services.link.v1.TransformLongUrlRequest
	1, // 2: services.link.v1.ShortLink.GenerateShortUrl:output_type -> services.link.v1.GenerateShortUrlReply
	3, // 3: services.link.v1.ShortLink.TransformLongUrl:output_type -> services.link.v1.TransformLongUrlReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_link_v1_short_link_proto_init() }
func file_services_link_v1_short_link_proto_init() {
	if File_services_link_v1_short_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_link_v1_short_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateShortUrlRequest); i {
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
		file_services_link_v1_short_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateShortUrlReply); i {
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
		file_services_link_v1_short_link_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformLongUrlRequest); i {
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
		file_services_link_v1_short_link_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformLongUrlReply); i {
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
			RawDescriptor: file_services_link_v1_short_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_link_v1_short_link_proto_goTypes,
		DependencyIndexes: file_services_link_v1_short_link_proto_depIdxs,
		MessageInfos:      file_services_link_v1_short_link_proto_msgTypes,
	}.Build()
	File_services_link_v1_short_link_proto = out.File
	file_services_link_v1_short_link_proto_rawDesc = nil
	file_services_link_v1_short_link_proto_goTypes = nil
	file_services_link_v1_short_link_proto_depIdxs = nil
}
