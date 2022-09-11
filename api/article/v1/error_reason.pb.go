// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: article/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	// 服务异常
	ErrorReason_SERVER_ERROR ErrorReason = 0
	// 数据校验失败
	ErrorReason_DATA_VALIDATE_ERROR ErrorReason = 100001
	// 数据查询失败
	ErrorReason_DATA_SELECT_ERROR ErrorReason = 100002
	// 数据已存在
	ErrorReason_DATA_ALREADY_EXISTS ErrorReason = 100003
	// 数据不存在
	ErrorReason_DATA_NOT_FOUND ErrorReason = 100004
	// 请先登录后再操作
	ErrorReason_NOT_LOGIN_ERROR ErrorReason = 100005
	// 新增数据失败
	ErrorReason_DATA_ADD_ERROR ErrorReason = 100006
	// 更新数据失败
	ErrorReason_DATA_UPDATE_ERROR ErrorReason = 100007
	// 无效ID值
	ErrorReason_ID_INVALID_VALUE_ERROR ErrorReason = 100008
	// 数据删除失败
	ErrorReason_DATA_DELETE_ERROR ErrorReason = 100009
	// 数据资源不存在
	ErrorReason_DATA_RESOURCE_NOT_FOUND ErrorReason = 100010
	// 无效的执行指令
	ErrorReason_COMMAND_INVALID_NOT_FOUND ErrorReason = 100011
	// 数据属性更新失败
	ErrorReason_DATA_UPDATE_FIELD_ERROR ErrorReason = 100012
	// 请先删除分类下的文章
	ErrorReason_ARTICLE_IN_CATEGORY_ERROR ErrorReason = 200001
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:      "SERVER_ERROR",
		100001: "DATA_VALIDATE_ERROR",
		100002: "DATA_SELECT_ERROR",
		100003: "DATA_ALREADY_EXISTS",
		100004: "DATA_NOT_FOUND",
		100005: "NOT_LOGIN_ERROR",
		100006: "DATA_ADD_ERROR",
		100007: "DATA_UPDATE_ERROR",
		100008: "ID_INVALID_VALUE_ERROR",
		100009: "DATA_DELETE_ERROR",
		100010: "DATA_RESOURCE_NOT_FOUND",
		100011: "COMMAND_INVALID_NOT_FOUND",
		100012: "DATA_UPDATE_FIELD_ERROR",
		200001: "ARTICLE_IN_CATEGORY_ERROR",
	}
	ErrorReason_value = map[string]int32{
		"SERVER_ERROR":              0,
		"DATA_VALIDATE_ERROR":       100001,
		"DATA_SELECT_ERROR":         100002,
		"DATA_ALREADY_EXISTS":       100003,
		"DATA_NOT_FOUND":            100004,
		"NOT_LOGIN_ERROR":           100005,
		"DATA_ADD_ERROR":            100006,
		"DATA_UPDATE_ERROR":         100007,
		"ID_INVALID_VALUE_ERROR":    100008,
		"DATA_DELETE_ERROR":         100009,
		"DATA_RESOURCE_NOT_FOUND":   100010,
		"COMMAND_INVALID_NOT_FOUND": 100011,
		"DATA_UPDATE_FIELD_ERROR":   100012,
		"ARTICLE_IN_CATEGORY_ERROR": 200001,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_article_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_article_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_article_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_article_v1_error_reason_proto protoreflect.FileDescriptor

var file_article_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0xd5, 0x03, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x00, 0x12, 0x1f, 0x0a, 0x13, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa1, 0x8d, 0x06, 0x1a, 0x04, 0xa8,
	0x45, 0xa6, 0x03, 0x12, 0x1d, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x4c, 0x45,
	0x43, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa2, 0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x12, 0x1f, 0x0a, 0x13, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0xa3, 0x8d, 0x06, 0x1a, 0x04, 0xa8,
	0x45, 0x90, 0x03, 0x12, 0x1a, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa4, 0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x1b, 0x0a, 0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0xa5, 0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1a, 0x0a, 0x0e,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa6,
	0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xa7, 0x8d,
	0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x22, 0x0a, 0x16, 0x49, 0x44, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0xa8, 0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a, 0x11, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0xa9, 0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x17, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xaa, 0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x25, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xab, 0x8d, 0x06,
	0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x17, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0xac, 0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x25, 0x0a, 0x19, 0x41,
	0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xc1, 0x9a, 0x0c, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x13, 0x5a, 0x11, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_article_v1_error_reason_proto_rawDescOnce sync.Once
	file_article_v1_error_reason_proto_rawDescData = file_article_v1_error_reason_proto_rawDesc
)

func file_article_v1_error_reason_proto_rawDescGZIP() []byte {
	file_article_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_article_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_v1_error_reason_proto_rawDescData)
	})
	return file_article_v1_error_reason_proto_rawDescData
}

var file_article_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_article_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: article.v1.ErrorReason
}
var file_article_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_article_v1_error_reason_proto_init() }
func file_article_v1_error_reason_proto_init() {
	if File_article_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_article_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_article_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_article_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_article_v1_error_reason_proto_enumTypes,
	}.Build()
	File_article_v1_error_reason_proto = out.File
	file_article_v1_error_reason_proto_rawDesc = nil
	file_article_v1_error_reason_proto_goTypes = nil
	file_article_v1_error_reason_proto_depIdxs = nil
}
