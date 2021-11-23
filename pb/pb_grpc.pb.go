// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrontendClient interface {
	Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error)
	Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error)
}

type frontendClient struct {
	cc grpc.ClientConnInterface
}

func NewFrontendClient(cc grpc.ClientConnInterface) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error) {
	out := new(BidResponse)
	err := c.cc.Invoke(ctx, "/pb.Frontend/bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/pb.Frontend/result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontendServer is the server API for Frontend service.
// All implementations must embed UnimplementedFrontendServer
// for forward compatibility
type FrontendServer interface {
	Bid(context.Context, *BidRequest) (*BidResponse, error)
	Result(context.Context, *ResultRequest) (*ResultResponse, error)
	mustEmbedUnimplementedFrontendServer()
}

// UnimplementedFrontendServer must be embedded to have forward compatible implementations.
type UnimplementedFrontendServer struct {
}

func (UnimplementedFrontendServer) Bid(context.Context, *BidRequest) (*BidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedFrontendServer) Result(context.Context, *ResultRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedFrontendServer) mustEmbedUnimplementedFrontendServer() {}

// UnsafeFrontendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrontendServer will
// result in compilation errors.
type UnsafeFrontendServer interface {
	mustEmbedUnimplementedFrontendServer()
}

func RegisterFrontendServer(s grpc.ServiceRegistrar, srv FrontendServer) {
	s.RegisterService(&Frontend_ServiceDesc, srv)
}

func _Frontend_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Frontend/bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).Bid(ctx, req.(*BidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Frontend/result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).Result(ctx, req.(*ResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Frontend_ServiceDesc is the grpc.ServiceDesc for Frontend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Frontend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "bid",
			Handler:    _Frontend_Bid_Handler,
		},
		{
			MethodName: "result",
			Handler:    _Frontend_Result_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

// ReplicaClient is the client API for Replica service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicaClient interface {
	Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidReplicaResponse, error)
	Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error)
}

type replicaClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicaClient(cc grpc.ClientConnInterface) ReplicaClient {
	return &replicaClient{cc}
}

func (c *replicaClient) Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidReplicaResponse, error) {
	out := new(BidReplicaResponse)
	err := c.cc.Invoke(ctx, "/pb.Replica/bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicaClient) Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/pb.Replica/result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicaServer is the server API for Replica service.
// All implementations must embed UnimplementedReplicaServer
// for forward compatibility
type ReplicaServer interface {
	Bid(context.Context, *BidRequest) (*BidReplicaResponse, error)
	Result(context.Context, *ResultRequest) (*ResultResponse, error)
	mustEmbedUnimplementedReplicaServer()
}

// UnimplementedReplicaServer must be embedded to have forward compatible implementations.
type UnimplementedReplicaServer struct {
}

func (UnimplementedReplicaServer) Bid(context.Context, *BidRequest) (*BidReplicaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedReplicaServer) Result(context.Context, *ResultRequest) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedReplicaServer) mustEmbedUnimplementedReplicaServer() {}

// UnsafeReplicaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicaServer will
// result in compilation errors.
type UnsafeReplicaServer interface {
	mustEmbedUnimplementedReplicaServer()
}

func RegisterReplicaServer(s grpc.ServiceRegistrar, srv ReplicaServer) {
	s.RegisterService(&Replica_ServiceDesc, srv)
}

func _Replica_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicaServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Replica/bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicaServer).Bid(ctx, req.(*BidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replica_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicaServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Replica/result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicaServer).Result(ctx, req.(*ResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Replica_ServiceDesc is the grpc.ServiceDesc for Replica service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Replica_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Replica",
	HandlerType: (*ReplicaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "bid",
			Handler:    _Replica_Bid_Handler,
		},
		{
			MethodName: "result",
			Handler:    _Replica_Result_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}
