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

// ShortLinkClient is the client API for ShortLink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortLinkClient interface {
	GenerateShortLink(ctx context.Context, in *GenerateShortLinkRequest, opts ...grpc.CallOption) (*GenerateShortLinkReply, error)
}

type shortLinkClient struct {
	cc grpc.ClientConnInterface
}

func NewShortLinkClient(cc grpc.ClientConnInterface) ShortLinkClient {
	return &shortLinkClient{cc}
}

func (c *shortLinkClient) GenerateShortLink(ctx context.Context, in *GenerateShortLinkRequest, opts ...grpc.CallOption) (*GenerateShortLinkReply, error) {
	out := new(GenerateShortLinkReply)
	err := c.cc.Invoke(ctx, "/link.v1.ShortLink/GenerateShortLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortLinkServer is the server API for ShortLink service.
// All implementations must embed UnimplementedShortLinkServer
// for forward compatibility
type ShortLinkServer interface {
	GenerateShortLink(context.Context, *GenerateShortLinkRequest) (*GenerateShortLinkReply, error)
	mustEmbedUnimplementedShortLinkServer()
}

// UnimplementedShortLinkServer must be embedded to have forward compatible implementations.
type UnimplementedShortLinkServer struct {
}

func (UnimplementedShortLinkServer) GenerateShortLink(context.Context, *GenerateShortLinkRequest) (*GenerateShortLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateShortLink not implemented")
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

func _ShortLink_GenerateShortLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateShortLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortLinkServer).GenerateShortLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/link.v1.ShortLink/GenerateShortLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortLinkServer).GenerateShortLink(ctx, req.(*GenerateShortLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortLink_ServiceDesc is the grpc.ServiceDesc for ShortLink service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortLink_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "link.v1.ShortLink",
	HandlerType: (*ShortLinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateShortLink",
			Handler:    _ShortLink_GenerateShortLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "link/v1/short_link.proto",
}
