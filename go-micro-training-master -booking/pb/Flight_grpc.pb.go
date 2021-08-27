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

// FlightServiceClient is the client API for FlightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlightServiceClient interface {
	CreateFlight(ctx context.Context, in *Flight, opts ...grpc.CallOption) (*Flight, error)
	UpdateFlight(ctx context.Context, in *Flight, opts ...grpc.CallOption) (*Flight, error)
	SearchFlight(ctx context.Context, in *SearchFlightRequest, opts ...grpc.CallOption) (*ListFlight, error)
	FindFlight(ctx context.Context, in *FindResquest, opts ...grpc.CallOption) (*Flight, error)
}

type flightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightServiceClient(cc grpc.ClientConnInterface) FlightServiceClient {
	return &flightServiceClient{cc}
}

func (c *flightServiceClient) CreateFlight(ctx context.Context, in *Flight, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/booking.FlightService/CreateFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) UpdateFlight(ctx context.Context, in *Flight, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/booking.FlightService/UpdateFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) SearchFlight(ctx context.Context, in *SearchFlightRequest, opts ...grpc.CallOption) (*ListFlight, error) {
	out := new(ListFlight)
	err := c.cc.Invoke(ctx, "/booking.FlightService/SearchFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) FindFlight(ctx context.Context, in *FindResquest, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, "/booking.FlightService/FindFlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightServiceServer is the server API for FlightService service.
// All implementations must embed UnimplementedFlightServiceServer
// for forward compatibility
type FlightServiceServer interface {
	CreateFlight(context.Context, *Flight) (*Flight, error)
	UpdateFlight(context.Context, *Flight) (*Flight, error)
	SearchFlight(context.Context, *SearchFlightRequest) (*ListFlight, error)
	FindFlight(context.Context, *FindResquest) (*Flight, error)
	mustEmbedUnimplementedFlightServiceServer()
}

// UnimplementedFlightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlightServiceServer struct {
}

func (UnimplementedFlightServiceServer) CreateFlight(context.Context, *Flight) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlight not implemented")
}
func (UnimplementedFlightServiceServer) UpdateFlight(context.Context, *Flight) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlight not implemented")
}
func (UnimplementedFlightServiceServer) SearchFlight(context.Context, *SearchFlightRequest) (*ListFlight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFlight not implemented")
}
func (UnimplementedFlightServiceServer) FindFlight(context.Context, *FindResquest) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFlight not implemented")
}
func (UnimplementedFlightServiceServer) mustEmbedUnimplementedFlightServiceServer() {}

// UnsafeFlightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightServiceServer will
// result in compilation errors.
type UnsafeFlightServiceServer interface {
	mustEmbedUnimplementedFlightServiceServer()
}

func RegisterFlightServiceServer(s grpc.ServiceRegistrar, srv FlightServiceServer) {
	s.RegisterService(&FlightService_ServiceDesc, srv)
}

func _FlightService_CreateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Flight)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).CreateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.FlightService/CreateFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).CreateFlight(ctx, req.(*Flight))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_UpdateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Flight)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).UpdateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.FlightService/UpdateFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).UpdateFlight(ctx, req.(*Flight))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_SearchFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).SearchFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.FlightService/SearchFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).SearchFlight(ctx, req.(*SearchFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_FindFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).FindFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.FlightService/FindFlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).FindFlight(ctx, req.(*FindResquest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlightService_ServiceDesc is the grpc.ServiceDesc for FlightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.FlightService",
	HandlerType: (*FlightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFlight",
			Handler:    _FlightService_CreateFlight_Handler,
		},
		{
			MethodName: "UpdateFlight",
			Handler:    _FlightService_UpdateFlight_Handler,
		},
		{
			MethodName: "SearchFlight",
			Handler:    _FlightService_SearchFlight_Handler,
		},
		{
			MethodName: "FindFlight",
			Handler:    _FlightService_FindFlight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Flight.proto",
}
