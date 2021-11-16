// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package AuctionReplica

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

// ReplicaClient is the client API for Replica service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicaClient interface {
	Bid(ctx context.Context, in *BidReplicaRequest, opts ...grpc.CallOption) (*BidReplicaResponse, error)
	Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResultReplicaResponse, error)
}

type replicaClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicaClient(cc grpc.ClientConnInterface) ReplicaClient {
	return &replicaClient{cc}
}

func (c *replicaClient) Bid(ctx context.Context, in *BidReplicaRequest, opts ...grpc.CallOption) (*BidReplicaResponse, error) {
	out := new(BidReplicaResponse)
	err := c.cc.Invoke(ctx, "/AuctionReplica.Replica/bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicaClient) Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResultReplicaResponse, error) {
	out := new(ResultReplicaResponse)
	err := c.cc.Invoke(ctx, "/AuctionReplica.Replica/result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicaServer is the server API for Replica service.
// All implementations must embed UnimplementedReplicaServer
// for forward compatibility
type ReplicaServer interface {
	Bid(context.Context, *BidReplicaRequest) (*BidReplicaResponse, error)
	Result(context.Context, *Empty) (*ResultReplicaResponse, error)
	mustEmbedUnimplementedReplicaServer()
}

// UnimplementedReplicaServer must be embedded to have forward compatible implementations.
type UnimplementedReplicaServer struct {
}

func (UnimplementedReplicaServer) Bid(context.Context, *BidReplicaRequest) (*BidReplicaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedReplicaServer) Result(context.Context, *Empty) (*ResultReplicaResponse, error) {
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
	in := new(BidReplicaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicaServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuctionReplica.Replica/bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicaServer).Bid(ctx, req.(*BidReplicaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Replica_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicaServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuctionReplica.Replica/result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicaServer).Result(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Replica_ServiceDesc is the grpc.ServiceDesc for Replica service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Replica_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuctionReplica.Replica",
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
	Metadata: "AuctionBackend/AuctionReplica.proto",
}