// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: grpc/server/myserver/protoc/hi/hi.proto

package hi

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

// HiClient is the client API for Hi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HiClient interface {
	// 定义SayHi方法
	SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error)
}

type hiClient struct {
	cc grpc.ClientConnInterface
}

func NewHiClient(cc grpc.ClientConnInterface) HiClient {
	return &hiClient{cc}
}

func (c *hiClient) SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiResponse, error) {
	out := new(HiResponse)
	err := c.cc.Invoke(ctx, "/hi.Hi/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HiServer is the server API for Hi service.
// All implementations should embed UnimplementedHiServer
// for forward compatibility
type HiServer interface {
	// 定义SayHi方法
	SayHi(context.Context, *HiRequest) (*HiResponse, error)
}

// UnimplementedHiServer should be embedded to have forward compatible implementations.
type UnimplementedHiServer struct {
}

func (UnimplementedHiServer) SayHi(context.Context, *HiRequest) (*HiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}

// UnsafeHiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HiServer will
// result in compilation errors.
type UnsafeHiServer interface {
	mustEmbedUnimplementedHiServer()
}

func RegisterHiServer(s grpc.ServiceRegistrar, srv HiServer) {
	s.RegisterService(&Hi_ServiceDesc, srv)
}

func _Hi_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HiServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hi.Hi/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HiServer).SayHi(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hi_ServiceDesc is the grpc.ServiceDesc for Hi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hi.Hi",
	HandlerType: (*HiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Hi_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/server/myserver/protoc/hi/hi.proto",
}
