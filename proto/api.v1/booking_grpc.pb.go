// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: booking.proto

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
	TrainBooking_BookTicket_FullMethodName       = "/TrainBooking/BookTicket"
	TrainBooking_GetBookingStatus_FullMethodName = "/TrainBooking/GetBookingStatus"
)

// TrainBookingClient is the client API for TrainBooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainBookingClient interface {
	BookTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketResponse, error)
	GetBookingStatus(ctx context.Context, in *BookingStatusRequest, opts ...grpc.CallOption) (*BookingStatusResponse, error)
}

type trainBookingClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainBookingClient(cc grpc.ClientConnInterface) TrainBookingClient {
	return &trainBookingClient{cc}
}

func (c *trainBookingClient) BookTicket(ctx context.Context, in *TicketRequest, opts ...grpc.CallOption) (*TicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketResponse)
	err := c.cc.Invoke(ctx, TrainBooking_BookTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainBookingClient) GetBookingStatus(ctx context.Context, in *BookingStatusRequest, opts ...grpc.CallOption) (*BookingStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingStatusResponse)
	err := c.cc.Invoke(ctx, TrainBooking_GetBookingStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainBookingServer is the server API for TrainBooking service.
// All implementations must embed UnimplementedTrainBookingServer
// for forward compatibility.
type TrainBookingServer interface {
	BookTicket(context.Context, *TicketRequest) (*TicketResponse, error)
	GetBookingStatus(context.Context, *BookingStatusRequest) (*BookingStatusResponse, error)
	mustEmbedUnimplementedTrainBookingServer()
}

// UnimplementedTrainBookingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrainBookingServer struct{}

func (UnimplementedTrainBookingServer) BookTicket(context.Context, *TicketRequest) (*TicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTicket not implemented")
}
func (UnimplementedTrainBookingServer) GetBookingStatus(context.Context, *BookingStatusRequest) (*BookingStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingStatus not implemented")
}
func (UnimplementedTrainBookingServer) mustEmbedUnimplementedTrainBookingServer() {}
func (UnimplementedTrainBookingServer) testEmbeddedByValue()                      {}

// UnsafeTrainBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainBookingServer will
// result in compilation errors.
type UnsafeTrainBookingServer interface {
	mustEmbedUnimplementedTrainBookingServer()
}

func RegisterTrainBookingServer(s grpc.ServiceRegistrar, srv TrainBookingServer) {
	// If the following call pancis, it indicates UnimplementedTrainBookingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrainBooking_ServiceDesc, srv)
}

func _TrainBooking_BookTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainBookingServer).BookTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainBooking_BookTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainBookingServer).BookTicket(ctx, req.(*TicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainBooking_GetBookingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainBookingServer).GetBookingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainBooking_GetBookingStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainBookingServer).GetBookingStatus(ctx, req.(*BookingStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainBooking_ServiceDesc is the grpc.ServiceDesc for TrainBooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainBooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TrainBooking",
	HandlerType: (*TrainBookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookTicket",
			Handler:    _TrainBooking_BookTicket_Handler,
		},
		{
			MethodName: "GetBookingStatus",
			Handler:    _TrainBooking_GetBookingStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
