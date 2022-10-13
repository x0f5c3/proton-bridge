// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: focus.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FocusClient is the client API for Focus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FocusClient interface {
	Raise(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type focusClient struct {
	cc grpc.ClientConnInterface
}

func NewFocusClient(cc grpc.ClientConnInterface) FocusClient {
	return &focusClient{cc}
}

func (c *focusClient) Raise(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Focus/Raise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *focusClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/proto.Focus/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FocusServer is the server API for Focus service.
// All implementations must embed UnimplementedFocusServer
// for forward compatibility
type FocusServer interface {
	Raise(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	mustEmbedUnimplementedFocusServer()
}

// UnimplementedFocusServer must be embedded to have forward compatible implementations.
type UnimplementedFocusServer struct {
}

func (UnimplementedFocusServer) Raise(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Raise not implemented")
}
func (UnimplementedFocusServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedFocusServer) mustEmbedUnimplementedFocusServer() {}

// UnsafeFocusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FocusServer will
// result in compilation errors.
type UnsafeFocusServer interface {
	mustEmbedUnimplementedFocusServer()
}

func RegisterFocusServer(s grpc.ServiceRegistrar, srv FocusServer) {
	s.RegisterService(&Focus_ServiceDesc, srv)
}

func _Focus_Raise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FocusServer).Raise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Focus/Raise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FocusServer).Raise(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Focus_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FocusServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Focus/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FocusServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Focus_ServiceDesc is the grpc.ServiceDesc for Focus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Focus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Focus",
	HandlerType: (*FocusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Raise",
			Handler:    _Focus_Raise_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Focus_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "focus.proto",
}
