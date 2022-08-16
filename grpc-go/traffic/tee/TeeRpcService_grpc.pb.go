// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tee

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

// TeeRpcServiceClient is the client API for TeeRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeeRpcServiceClient interface {
	// processing transaction message requests
	SendRemoteTeeRequest(ctx context.Context, in *TeeTxRequest, opts ...grpc.CallOption) (*TeeTxResponse, error)
}

type teeRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeeRpcServiceClient(cc grpc.ClientConnInterface) TeeRpcServiceClient {
	return &teeRpcServiceClient{cc}
}

func (c *teeRpcServiceClient) SendRemoteTeeRequest(ctx context.Context, in *TeeTxRequest, opts ...grpc.CallOption) (*TeeTxResponse, error) {
	out := new(TeeTxResponse)
	err := c.cc.Invoke(ctx, "/traffic.tee.TeeRpcService/SendRemoteTeeRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRpcServiceServer is the server API for TeeRpcService service.
// All implementations must embed UnimplementedTeeRpcServiceServer
// for forward compatibility
type TeeRpcServiceServer interface {
	// processing transaction message requests
	SendRemoteTeeRequest(context.Context, *TeeTxRequest) (*TeeTxResponse, error)
	mustEmbedUnimplementedTeeRpcServiceServer()
}

// UnimplementedTeeRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeeRpcServiceServer struct {
}

func (UnimplementedTeeRpcServiceServer) SendRemoteTeeRequest(context.Context, *TeeTxRequest) (*TeeTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRemoteTeeRequest not implemented")
}
func (UnimplementedTeeRpcServiceServer) mustEmbedUnimplementedTeeRpcServiceServer() {}

// UnsafeTeeRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeeRpcServiceServer will
// result in compilation errors.
type UnsafeTeeRpcServiceServer interface {
	mustEmbedUnimplementedTeeRpcServiceServer()
}

func RegisterTeeRpcServiceServer(s grpc.ServiceRegistrar, srv TeeRpcServiceServer) {
	s.RegisterService(&TeeRpcService_ServiceDesc, srv)
}

func _TeeRpcService_SendRemoteTeeRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeeTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeeRpcServiceServer).SendRemoteTeeRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/traffic.tee.TeeRpcService/SendRemoteTeeRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeeRpcServiceServer).SendRemoteTeeRequest(ctx, req.(*TeeTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeeRpcService_ServiceDesc is the grpc.ServiceDesc for TeeRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeeRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "traffic.tee.TeeRpcService",
	HandlerType: (*TeeRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRemoteTeeRequest",
			Handler:    _TeeRpcService_SendRemoteTeeRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TeeRpcService.proto",
}
