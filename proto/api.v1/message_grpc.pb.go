// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: message.proto

package api_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ComService_ProcessMessage_FullMethodName = "/ComService/ProcessMessage"
)

// ComServiceClient is the client API for ComService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComServiceClient interface {
	ProcessMessage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type comServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComServiceClient(cc grpc.ClientConnInterface) ComServiceClient {
	return &comServiceClient{cc}
}

func (c *comServiceClient) ProcessMessage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, ComService_ProcessMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComServiceServer is the server API for ComService service.
// All implementations must embed UnimplementedComServiceServer
// for forward compatibility.
type ComServiceServer interface {
	ProcessMessage(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedComServiceServer()
}

// UnimplementedComServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedComServiceServer struct{}

func (UnimplementedComServiceServer) ProcessMessage(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessMessage not implemented")
}
func (UnimplementedComServiceServer) mustEmbedUnimplementedComServiceServer() {}
func (UnimplementedComServiceServer) testEmbeddedByValue()                    {}

// UnsafeComServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComServiceServer will
// result in compilation errors.
type UnsafeComServiceServer interface {
	mustEmbedUnimplementedComServiceServer()
}

func RegisterComServiceServer(s grpc.ServiceRegistrar, srv ComServiceServer) {
	// If the following call pancis, it indicates UnimplementedComServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ComService_ServiceDesc, srv)
}

func _ComService_ProcessMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComServiceServer).ProcessMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComService_ProcessMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComServiceServer).ProcessMessage(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ComService_ServiceDesc is the grpc.ServiceDesc for ComService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ComService",
	HandlerType: (*ComServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessMessage",
			Handler:    _ComService_ProcessMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}