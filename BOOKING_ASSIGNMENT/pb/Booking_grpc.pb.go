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

// BookingsClient is the client API for Bookings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingsClient interface {
	Booking(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Info, error)
	ViewBooking(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error)
	CancelBooking(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*Cancel, error)
}

type bookingsClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingsClient(cc grpc.ClientConnInterface) BookingsClient {
	return &bookingsClient{cc}
}

func (c *bookingsClient) Booking(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/booking.Bookings/Booking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingsClient) ViewBooking(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*ViewResponse, error) {
	out := new(ViewResponse)
	err := c.cc.Invoke(ctx, "/booking.Bookings/ViewBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingsClient) CancelBooking(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (*Cancel, error) {
	out := new(Cancel)
	err := c.cc.Invoke(ctx, "/booking.Bookings/CancelBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingsServer is the server API for Bookings service.
// All implementations must embed UnimplementedBookingsServer
// for forward compatibility
type BookingsServer interface {
	Booking(context.Context, *Info) (*Info, error)
	ViewBooking(context.Context, *ViewRequest) (*ViewResponse, error)
	CancelBooking(context.Context, *ViewRequest) (*Cancel, error)
	mustEmbedUnimplementedBookingsServer()
}

// UnimplementedBookingsServer must be embedded to have forward compatible implementations.
type UnimplementedBookingsServer struct {
}

func (UnimplementedBookingsServer) Booking(context.Context, *Info) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Booking not implemented")
}
func (UnimplementedBookingsServer) ViewBooking(context.Context, *ViewRequest) (*ViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewBooking not implemented")
}
func (UnimplementedBookingsServer) CancelBooking(context.Context, *ViewRequest) (*Cancel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBooking not implemented")
}
func (UnimplementedBookingsServer) mustEmbedUnimplementedBookingsServer() {}

// UnsafeBookingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingsServer will
// result in compilation errors.
type UnsafeBookingsServer interface {
	mustEmbedUnimplementedBookingsServer()
}

func RegisterBookingsServer(s grpc.ServiceRegistrar, srv BookingsServer) {
	s.RegisterService(&Bookings_ServiceDesc, srv)
}

func _Bookings_Booking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Info)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingsServer).Booking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bookings/Booking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingsServer).Booking(ctx, req.(*Info))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookings_ViewBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingsServer).ViewBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bookings/ViewBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingsServer).ViewBooking(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookings_CancelBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingsServer).CancelBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bookings/CancelBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingsServer).CancelBooking(ctx, req.(*ViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bookings_ServiceDesc is the grpc.ServiceDesc for Bookings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bookings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.Bookings",
	HandlerType: (*BookingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Booking",
			Handler:    _Bookings_Booking_Handler,
		},
		{
			MethodName: "ViewBooking",
			Handler:    _Bookings_ViewBooking_Handler,
		},
		{
			MethodName: "CancelBooking",
			Handler:    _Bookings_CancelBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Booking.proto",
}
