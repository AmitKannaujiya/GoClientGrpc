// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: trainbook.proto

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
	TrainTicketing_PurchaseTicket_FullMethodName    = "/TrainTicketing/PurchaseTicket"
	TrainTicketing_GetReceipt_FullMethodName        = "/TrainTicketing/GetReceipt"
	TrainTicketing_GetUsersBySection_FullMethodName = "/TrainTicketing/GetUsersBySection"
	TrainTicketing_RemoveUser_FullMethodName        = "/TrainTicketing/RemoveUser"
	TrainTicketing_ModifySeat_FullMethodName        = "/TrainTicketing/ModifySeat"
)

// TrainTicketingClient is the client API for TrainTicketing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainTicketingClient interface {
	PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
	GetReceipt(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (*ReceiptResponse, error)
	GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	RemoveUser(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error)
}

type trainTicketingClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainTicketingClient(cc grpc.ClientConnInterface) TrainTicketingClient {
	return &trainTicketingClient{cc}
}

func (c *trainTicketingClient) PurchaseTicket(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, TrainTicketing_PurchaseTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketingClient) GetReceipt(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (*ReceiptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReceiptResponse)
	err := c.cc.Invoke(ctx, TrainTicketing_GetReceipt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketingClient) GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, TrainTicketing_GetUsersBySection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketingClient) RemoveUser(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, TrainTicketing_RemoveUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketingClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifySeatResponse)
	err := c.cc.Invoke(ctx, TrainTicketing_ModifySeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainTicketingServer is the server API for TrainTicketing service.
// All implementations must embed UnimplementedTrainTicketingServer
// for forward compatibility.
type TrainTicketingServer interface {
	PurchaseTicket(context.Context, *PurchaseRequest) (*PurchaseResponse, error)
	GetReceipt(context.Context, *ReceiptRequest) (*ReceiptResponse, error)
	GetUsersBySection(context.Context, *SectionRequest) (*UsersResponse, error)
	RemoveUser(context.Context, *RemoveRequest) (*RemoveResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error)
	mustEmbedUnimplementedTrainTicketingServer()
}

// UnimplementedTrainTicketingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrainTicketingServer struct{}

func (UnimplementedTrainTicketingServer) PurchaseTicket(context.Context, *PurchaseRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTicket not implemented")
}
func (UnimplementedTrainTicketingServer) GetReceipt(context.Context, *ReceiptRequest) (*ReceiptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedTrainTicketingServer) GetUsersBySection(context.Context, *SectionRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersBySection not implemented")
}
func (UnimplementedTrainTicketingServer) RemoveUser(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTrainTicketingServer) ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTrainTicketingServer) mustEmbedUnimplementedTrainTicketingServer() {}
func (UnimplementedTrainTicketingServer) testEmbeddedByValue()                        {}

// UnsafeTrainTicketingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainTicketingServer will
// result in compilation errors.
type UnsafeTrainTicketingServer interface {
	mustEmbedUnimplementedTrainTicketingServer()
}

func RegisterTrainTicketingServer(s grpc.ServiceRegistrar, srv TrainTicketingServer) {
	// If the following call pancis, it indicates UnimplementedTrainTicketingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrainTicketing_ServiceDesc, srv)
}

func _TrainTicketing_PurchaseTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServer).PurchaseTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketing_PurchaseTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServer).PurchaseTicket(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketing_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketing_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServer).GetReceipt(ctx, req.(*ReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketing_GetUsersBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServer).GetUsersBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketing_GetUsersBySection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServer).GetUsersBySection(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketing_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketing_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServer).RemoveUser(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketing_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketingServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketing_ModifySeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketingServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainTicketing_ServiceDesc is the grpc.ServiceDesc for TrainTicketing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainTicketing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TrainTicketing",
	HandlerType: (*TrainTicketingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PurchaseTicket",
			Handler:    _TrainTicketing_PurchaseTicket_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _TrainTicketing_GetReceipt_Handler,
		},
		{
			MethodName: "GetUsersBySection",
			Handler:    _TrainTicketing_GetUsersBySection_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TrainTicketing_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TrainTicketing_ModifySeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trainbook.proto",
}
