// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: basic/v1/short_link.proto

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

// ShortLinkClient is the client API for ShortLink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortLinkClient interface {
	// 生成短链接
	GenerateShortUrl(ctx context.Context, in *GenerateShortUrlRequest, opts ...grpc.CallOption) (*GenerateShortUrlResponse, error)
	// 短链接拉取长链接
	TransformLongUrl(ctx context.Context, in *TransformLongUrlRequest, opts ...grpc.CallOption) (*TransformLongUrlResponse, error)
}

type shortLinkClient struct {
	cc grpc.ClientConnInterface
}

func NewShortLinkClient(cc grpc.ClientConnInterface) ShortLinkClient {
	return &shortLinkClient{cc}
}

func (c *shortLinkClient) GenerateShortUrl(ctx context.Context, in *GenerateShortUrlRequest, opts ...grpc.CallOption) (*GenerateShortUrlResponse, error) {
	out := new(GenerateShortUrlResponse)
	err := c.cc.Invoke(ctx, "/basic.v1.ShortLink/GenerateShortUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortLinkClient) TransformLongUrl(ctx context.Context, in *TransformLongUrlRequest, opts ...grpc.CallOption) (*TransformLongUrlResponse, error) {
	out := new(TransformLongUrlResponse)
	err := c.cc.Invoke(ctx, "/basic.v1.ShortLink/TransformLongUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortLinkServer is the server API for ShortLink service.
// All implementations must embed UnimplementedShortLinkServer
// for forward compatibility
type ShortLinkServer interface {
	// 生成短链接
	GenerateShortUrl(context.Context, *GenerateShortUrlRequest) (*GenerateShortUrlResponse, error)
	// 短链接拉取长链接
	TransformLongUrl(context.Context, *TransformLongUrlRequest) (*TransformLongUrlResponse, error)
	mustEmbedUnimplementedShortLinkServer()
}

// UnimplementedShortLinkServer must be embedded to have forward compatible implementations.
type UnimplementedShortLinkServer struct {
}

func (UnimplementedShortLinkServer) GenerateShortUrl(context.Context, *GenerateShortUrlRequest) (*GenerateShortUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateShortUrl not implemented")
}
func (UnimplementedShortLinkServer) TransformLongUrl(context.Context, *TransformLongUrlRequest) (*TransformLongUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransformLongUrl not implemented")
}
func (UnimplementedShortLinkServer) mustEmbedUnimplementedShortLinkServer() {}

// UnsafeShortLinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortLinkServer will
// result in compilation errors.
type UnsafeShortLinkServer interface {
	mustEmbedUnimplementedShortLinkServer()
}

func RegisterShortLinkServer(s grpc.ServiceRegistrar, srv ShortLinkServer) {
	s.RegisterService(&ShortLink_ServiceDesc, srv)
}

func _ShortLink_GenerateShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateShortUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortLinkServer).GenerateShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basic.v1.ShortLink/GenerateShortUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortLinkServer).GenerateShortUrl(ctx, req.(*GenerateShortUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortLink_TransformLongUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransformLongUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortLinkServer).TransformLongUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basic.v1.ShortLink/TransformLongUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortLinkServer).TransformLongUrl(ctx, req.(*TransformLongUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortLink_ServiceDesc is the grpc.ServiceDesc for ShortLink service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortLink_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basic.v1.ShortLink",
	HandlerType: (*ShortLinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateShortUrl",
			Handler:    _ShortLink_GenerateShortUrl_Handler,
		},
		{
			MethodName: "TransformLongUrl",
			Handler:    _ShortLink_TransformLongUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic/v1/short_link.proto",
}
