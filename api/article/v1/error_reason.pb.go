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
	// 未知错误
	ErrorReason_UNKNOWN_ERROR ErrorReason = 0
	// 服务异常
	ErrorReason_SERVER_ERROR ErrorReason = 100001
	// 数据校验失败
	ErrorReason_DATA_VALIDATE_ERROR ErrorReason = 201001
	// 数据查询失败
	ErrorReason_DATA_SELECT_ERROR ErrorReason = 201002
	// 数据已存在
	ErrorReason_DATA_ALREADY_EXISTS ErrorReason = 201003
	// 数据不存在
	ErrorReason_DATA_NOT_FOUND ErrorReason = 201004
	// 新增数据失败
	ErrorReason_DATA_ADD_ERROR ErrorReason = 201005
	// 更新数据失败
	ErrorReason_DATA_UPDATE_ERROR ErrorReason = 201006
	// 数据删除失败
	ErrorReason_DATA_DELETE_ERROR ErrorReason = 201007
	// 数据资源不存在
	ErrorReason_DATA_RESOURCE_NOT_FOUND ErrorReason = 201008
	// 数据属性更新失败
	ErrorReason_DATA_UPDATE_FIELD_ERROR ErrorReason = 201009
	// 无效ID值
	ErrorReason_ID_INVALID_VALUE_ERROR ErrorReason = 202001
	// 无效的执行指令
	ErrorReason_COMMAND_INVALID_NOT_FOUND ErrorReason = 202002
	// 请先登录后再操作
	ErrorReason_NOT_LOGIN_ERROR ErrorReason = 203001
	// 没有访问权限
	ErrorReason_NOT_VISIT_AUTH ErrorReason = 203002
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:      "UNKNOWN_ERROR",
		100001: "SERVER_ERROR",
		201001: "DATA_VALIDATE_ERROR",
		201002: "DATA_SELECT_ERROR",
		201003: "DATA_ALREADY_EXISTS",
		201004: "DATA_NOT_FOUND",
		201005: "DATA_ADD_ERROR",
		201006: "DATA_UPDATE_ERROR",
		201007: "DATA_DELETE_ERROR",
		201008: "DATA_RESOURCE_NOT_FOUND",
		201009: "DATA_UPDATE_FIELD_ERROR",
		202001: "ID_INVALID_VALUE_ERROR",
		202002: "COMMAND_INVALID_NOT_FOUND",
		203001: "NOT_LOGIN_ERROR",
		203002: "NOT_VISIT_AUTH",
	}
	ErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR":             0,
		"SERVER_ERROR":              100001,
		"DATA_VALIDATE_ERROR":       201001,
		"DATA_SELECT_ERROR":         201002,
		"DATA_ALREADY_EXISTS":       201003,
		"DATA_NOT_FOUND":            201004,
		"DATA_ADD_ERROR":            201005,
		"DATA_UPDATE_ERROR":         201006,
		"DATA_DELETE_ERROR":         201007,
		"DATA_RESOURCE_NOT_FOUND":   201008,
		"DATA_UPDATE_FIELD_ERROR":   201009,
		"ID_INVALID_VALUE_ERROR":    202001,
		"COMMAND_INVALID_NOT_FOUND": 202002,
		"NOT_LOGIN_ERROR":           203001,
		"NOT_VISIT_AUTH":            203002,
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
	0x2a, 0xe5, 0x03, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xa1, 0x8d, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1f, 0x0a,
	0x13, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0xa9, 0xa2, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0xa6, 0x03, 0x12, 0x1d,
	0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xaa, 0xa2, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1f, 0x0a,
	0x13, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58,
	0x49, 0x53, 0x54, 0x53, 0x10, 0xab, 0xa2, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1a,
	0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0xac, 0xa2, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1a, 0x0a, 0x0e, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xad, 0xa2, 0x0c,
	0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xae, 0xa2, 0x0c, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xaf, 0xa2, 0x0c, 0x1a, 0x04,
	0xa8, 0x45, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x17, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0xb0, 0xa2, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x17, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0xb1, 0xa2, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x22,
	0x0a, 0x16, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x91, 0xaa, 0x0c, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x12, 0x25, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x92, 0xaa, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1b, 0x0a, 0x0f, 0x4e, 0x4f, 0x54,
	0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf9, 0xb1, 0x0c,
	0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1a, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x56, 0x49,
	0x53, 0x49, 0x54, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x10, 0xfa, 0xb1, 0x0c, 0x1a, 0x04, 0xa8, 0x45,
	0x91, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x13, 0x5a, 0x11, 0x61, 0x70, 0x69, 0x2f,
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
