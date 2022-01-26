// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package message

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

// ChannelGrpcClient is the client API for ChannelGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelGrpcClient interface {
	GetChannels(ctx context.Context, in *GetChannelsRequest, opts ...grpc.CallOption) (*GetChannelsResponse, error)
}

type channelGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelGrpcClient(cc grpc.ClientConnInterface) ChannelGrpcClient {
	return &channelGrpcClient{cc}
}

func (c *channelGrpcClient) GetChannels(ctx context.Context, in *GetChannelsRequest, opts ...grpc.CallOption) (*GetChannelsResponse, error) {
	out := new(GetChannelsResponse)
	err := c.cc.Invoke(ctx, "/dashboard.ChannelGrpc/GetChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelGrpcServer is the server API for ChannelGrpc service.
// All implementations must embed UnimplementedChannelGrpcServer
// for forward compatibility
type ChannelGrpcServer interface {
	GetChannels(context.Context, *GetChannelsRequest) (*GetChannelsResponse, error)
	mustEmbedUnimplementedChannelGrpcServer()
}

// UnimplementedChannelGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedChannelGrpcServer struct {
}

func (UnimplementedChannelGrpcServer) GetChannels(context.Context, *GetChannelsRequest) (*GetChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannels not implemented")
}
func (UnimplementedChannelGrpcServer) mustEmbedUnimplementedChannelGrpcServer() {}

// UnsafeChannelGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelGrpcServer will
// result in compilation errors.
type UnsafeChannelGrpcServer interface {
	mustEmbedUnimplementedChannelGrpcServer()
}

func RegisterChannelGrpcServer(s grpc.ServiceRegistrar, srv ChannelGrpcServer) {
	s.RegisterService(&ChannelGrpc_ServiceDesc, srv)
}

func _ChannelGrpc_GetChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelGrpcServer).GetChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.ChannelGrpc/GetChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelGrpcServer).GetChannels(ctx, req.(*GetChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelGrpc_ServiceDesc is the grpc.ServiceDesc for ChannelGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.ChannelGrpc",
	HandlerType: (*ChannelGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChannels",
			Handler:    _ChannelGrpc_GetChannels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel.proto",
}