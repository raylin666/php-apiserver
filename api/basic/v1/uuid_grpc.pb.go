// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: basic/v1/uuid.proto

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

// UuidClient is the client API for Uuid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UuidClient interface {
	// 生成唯一标识
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type uuidClient struct {
	cc grpc.ClientConnInterface
}

func NewUuidClient(cc grpc.ClientConnInterface) UuidClient {
	return &uuidClient{cc}
}

func (c *uuidClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, "/basic.v1.Uuid/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UuidServer is the server API for Uuid service.
// All implementations must embed UnimplementedUuidServer
// for forward compatibility
type UuidServer interface {
	// 生成唯一标识
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	mustEmbedUnimplementedUuidServer()
}

// UnimplementedUuidServer must be embedded to have forward compatible implementations.
type UnimplementedUuidServer struct {
}

func (UnimplementedUuidServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedUuidServer) mustEmbedUnimplementedUuidServer() {}

// UnsafeUuidServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UuidServer will
// result in compilation errors.
type UnsafeUuidServer interface {
	mustEmbedUnimplementedUuidServer()
}

func RegisterUuidServer(s grpc.ServiceRegistrar, srv UuidServer) {
	s.RegisterService(&Uuid_ServiceDesc, srv)
}

func _Uuid_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UuidServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basic.v1.Uuid/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UuidServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Uuid_ServiceDesc is the grpc.ServiceDesc for Uuid service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Uuid_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basic.v1.Uuid",
	HandlerType: (*UuidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Uuid_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic/v1/uuid.proto",
}
