// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Frontend

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
	Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResultResponse, error)
}

type frontendClient struct {
	cc grpc.ClientConnInterface
}

func NewFrontendClient(cc grpc.ClientConnInterface) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidResponse, error) {
	out := new(BidResponse)
	err := c.cc.Invoke(ctx, "/Frontend.Frontend/bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, "/Frontend.Frontend/result", in, out, opts...)
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
	Result(context.Context, *Empty) (*ResultResponse, error)
	mustEmbedUnimplementedFrontendServer()
}

// UnimplementedFrontendServer must be embedded to have forward compatible implementations.
type UnimplementedFrontendServer struct {
}

func (UnimplementedFrontendServer) Bid(context.Context, *BidRequest) (*BidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedFrontendServer) Result(context.Context, *Empty) (*ResultResponse, error) {
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
		FullMethod: "/Frontend.Frontend/bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).Bid(ctx, req.(*BidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Frontend.Frontend/result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).Result(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Frontend_ServiceDesc is the grpc.ServiceDesc for Frontend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Frontend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Frontend.Frontend",
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
	Metadata: "Frontend.proto",
}
