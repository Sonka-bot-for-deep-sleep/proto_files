// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: time_service.proto

package api

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
	TimeService_GetWakeUpTime_FullMethodName            = "/api.TimeService/GetWakeUpTime"
	TimeService_GetFallingAsleep_FullMethodName         = "/api.TimeService/GetFallingAsleep"
	TimeService_SetWakeUpTime_FullMethodName            = "/api.TimeService/SetWakeUpTime"
	TimeService_SetFallingAsleep_FullMethodName         = "/api.TimeService/SetFallingAsleep"
	TimeService_GetAllUsersByTimeInRange_FullMethodName = "/api.TimeService/GetAllUsersByTimeInRange"
)

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeServiceClient interface {
	GetWakeUpTime(ctx context.Context, in *GetWakeUpTimeRequest, opts ...grpc.CallOption) (*GetWakeUpTimeResponse, error)
	GetFallingAsleep(ctx context.Context, in *GetFallingAsleepRequest, opts ...grpc.CallOption) (*GetFallingAsleepResponse, error)
	SetWakeUpTime(ctx context.Context, in *SetWakeUpTimeRequest, opts ...grpc.CallOption) (*SetWakeUpTimeResponse, error)
	SetFallingAsleep(ctx context.Context, in *SetFallingAsleepRequest, opts ...grpc.CallOption) (*SetFallingAsleepResponse, error)
	GetAllUsersByTimeInRange(ctx context.Context, in *GetAllUsersByTimeInRangeRequest, opts ...grpc.CallOption) (*GetAllUsersByTimeInRangeResponse, error)
}

type timeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeServiceClient(cc grpc.ClientConnInterface) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) GetWakeUpTime(ctx context.Context, in *GetWakeUpTimeRequest, opts ...grpc.CallOption) (*GetWakeUpTimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWakeUpTimeResponse)
	err := c.cc.Invoke(ctx, TimeService_GetWakeUpTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeServiceClient) GetFallingAsleep(ctx context.Context, in *GetFallingAsleepRequest, opts ...grpc.CallOption) (*GetFallingAsleepResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFallingAsleepResponse)
	err := c.cc.Invoke(ctx, TimeService_GetFallingAsleep_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeServiceClient) SetWakeUpTime(ctx context.Context, in *SetWakeUpTimeRequest, opts ...grpc.CallOption) (*SetWakeUpTimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetWakeUpTimeResponse)
	err := c.cc.Invoke(ctx, TimeService_SetWakeUpTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeServiceClient) SetFallingAsleep(ctx context.Context, in *SetFallingAsleepRequest, opts ...grpc.CallOption) (*SetFallingAsleepResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetFallingAsleepResponse)
	err := c.cc.Invoke(ctx, TimeService_SetFallingAsleep_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeServiceClient) GetAllUsersByTimeInRange(ctx context.Context, in *GetAllUsersByTimeInRangeRequest, opts ...grpc.CallOption) (*GetAllUsersByTimeInRangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllUsersByTimeInRangeResponse)
	err := c.cc.Invoke(ctx, TimeService_GetAllUsersByTimeInRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServiceServer is the server API for TimeService service.
// All implementations must embed UnimplementedTimeServiceServer
// for forward compatibility.
type TimeServiceServer interface {
	GetWakeUpTime(context.Context, *GetWakeUpTimeRequest) (*GetWakeUpTimeResponse, error)
	GetFallingAsleep(context.Context, *GetFallingAsleepRequest) (*GetFallingAsleepResponse, error)
	SetWakeUpTime(context.Context, *SetWakeUpTimeRequest) (*SetWakeUpTimeResponse, error)
	SetFallingAsleep(context.Context, *SetFallingAsleepRequest) (*SetFallingAsleepResponse, error)
	GetAllUsersByTimeInRange(context.Context, *GetAllUsersByTimeInRangeRequest) (*GetAllUsersByTimeInRangeResponse, error)
	mustEmbedUnimplementedTimeServiceServer()
}

// UnimplementedTimeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTimeServiceServer struct{}

func (UnimplementedTimeServiceServer) GetWakeUpTime(context.Context, *GetWakeUpTimeRequest) (*GetWakeUpTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWakeUpTime not implemented")
}
func (UnimplementedTimeServiceServer) GetFallingAsleep(context.Context, *GetFallingAsleepRequest) (*GetFallingAsleepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFallingAsleep not implemented")
}
func (UnimplementedTimeServiceServer) SetWakeUpTime(context.Context, *SetWakeUpTimeRequest) (*SetWakeUpTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWakeUpTime not implemented")
}
func (UnimplementedTimeServiceServer) SetFallingAsleep(context.Context, *SetFallingAsleepRequest) (*SetFallingAsleepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFallingAsleep not implemented")
}
func (UnimplementedTimeServiceServer) GetAllUsersByTimeInRange(context.Context, *GetAllUsersByTimeInRangeRequest) (*GetAllUsersByTimeInRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsersByTimeInRange not implemented")
}
func (UnimplementedTimeServiceServer) mustEmbedUnimplementedTimeServiceServer() {}
func (UnimplementedTimeServiceServer) testEmbeddedByValue()                     {}

// UnsafeTimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeServiceServer will
// result in compilation errors.
type UnsafeTimeServiceServer interface {
	mustEmbedUnimplementedTimeServiceServer()
}

func RegisterTimeServiceServer(s grpc.ServiceRegistrar, srv TimeServiceServer) {
	// If the following call pancis, it indicates UnimplementedTimeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TimeService_ServiceDesc, srv)
}

func _TimeService_GetWakeUpTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWakeUpTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).GetWakeUpTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeService_GetWakeUpTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).GetWakeUpTime(ctx, req.(*GetWakeUpTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeService_GetFallingAsleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFallingAsleepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).GetFallingAsleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeService_GetFallingAsleep_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).GetFallingAsleep(ctx, req.(*GetFallingAsleepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeService_SetWakeUpTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWakeUpTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).SetWakeUpTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeService_SetWakeUpTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).SetWakeUpTime(ctx, req.(*SetWakeUpTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeService_SetFallingAsleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFallingAsleepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).SetFallingAsleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeService_SetFallingAsleep_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).SetFallingAsleep(ctx, req.(*SetFallingAsleepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeService_GetAllUsersByTimeInRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersByTimeInRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).GetAllUsersByTimeInRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimeService_GetAllUsersByTimeInRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).GetAllUsersByTimeInRange(ctx, req.(*GetAllUsersByTimeInRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeService_ServiceDesc is the grpc.ServiceDesc for TimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWakeUpTime",
			Handler:    _TimeService_GetWakeUpTime_Handler,
		},
		{
			MethodName: "GetFallingAsleep",
			Handler:    _TimeService_GetFallingAsleep_Handler,
		},
		{
			MethodName: "SetWakeUpTime",
			Handler:    _TimeService_SetWakeUpTime_Handler,
		},
		{
			MethodName: "SetFallingAsleep",
			Handler:    _TimeService_SetFallingAsleep_Handler,
		},
		{
			MethodName: "GetAllUsersByTimeInRange",
			Handler:    _TimeService_GetAllUsersByTimeInRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time_service.proto",
}
