// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UploadClient is the client API for Upload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadClient interface {
	// 数据流方式上传文件
	StreamUploadFile(ctx context.Context, in *StreamUploadFileRequest, opts ...grpc.CallOption) (*StreamUploadFileReply, error)
	// 通过URL资源地址方式上传文件
	UrlUploadFile(ctx context.Context, in *UrlUploadFileRequest, opts ...grpc.CallOption) (*UrlUploadFileReply, error)
}

type uploadClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadClient(cc grpc.ClientConnInterface) UploadClient {
	return &uploadClient{cc}
}

func (c *uploadClient) StreamUploadFile(ctx context.Context, in *StreamUploadFileRequest, opts ...grpc.CallOption) (*StreamUploadFileReply, error) {
	out := new(StreamUploadFileReply)
	err := c.cc.Invoke(ctx, "/services.upload.v1.Upload/StreamUploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadClient) UrlUploadFile(ctx context.Context, in *UrlUploadFileRequest, opts ...grpc.CallOption) (*UrlUploadFileReply, error) {
	out := new(UrlUploadFileReply)
	err := c.cc.Invoke(ctx, "/services.upload.v1.Upload/UrlUploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadServer is the server API for Upload service.
// All implementations must embed UnimplementedUploadServer
// for forward compatibility
type UploadServer interface {
	// 数据流方式上传文件
	StreamUploadFile(context.Context, *StreamUploadFileRequest) (*StreamUploadFileReply, error)
	// 通过URL资源地址方式上传文件
	UrlUploadFile(context.Context, *UrlUploadFileRequest) (*UrlUploadFileReply, error)
	mustEmbedUnimplementedUploadServer()
}

// UnimplementedUploadServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServer struct {
}

func (UnimplementedUploadServer) StreamUploadFile(context.Context, *StreamUploadFileRequest) (*StreamUploadFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamUploadFile not implemented")
}
func (UnimplementedUploadServer) UrlUploadFile(context.Context, *UrlUploadFileRequest) (*UrlUploadFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UrlUploadFile not implemented")
}
func (UnimplementedUploadServer) mustEmbedUnimplementedUploadServer() {}

// UnsafeUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServer will
// result in compilation errors.
type UnsafeUploadServer interface {
	mustEmbedUnimplementedUploadServer()
}

func RegisterUploadServer(s grpc.ServiceRegistrar, srv UploadServer) {
	s.RegisterService(&Upload_ServiceDesc, srv)
}

func _Upload_StreamUploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamUploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).StreamUploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.upload.v1.Upload/StreamUploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).StreamUploadFile(ctx, req.(*StreamUploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upload_UrlUploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlUploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServer).UrlUploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.upload.v1.Upload/UrlUploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServer).UrlUploadFile(ctx, req.(*UrlUploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Upload_ServiceDesc is the grpc.ServiceDesc for Upload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.upload.v1.Upload",
	HandlerType: (*UploadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamUploadFile",
			Handler:    _Upload_StreamUploadFile_Handler,
		},
		{
			MethodName: "UrlUploadFile",
			Handler:    _Upload_UrlUploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/upload/v1/upload.proto",
}
