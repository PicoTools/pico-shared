// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: listener/v1/listener.proto

package listenerv1

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
	ListenerService_UpdateListener_FullMethodName = "/listener.v1.ListenerService/UpdateListener"
	ListenerService_RegisterAgent_FullMethodName  = "/listener.v1.ListenerService/RegisterAgent"
	ListenerService_GetTask_FullMethodName        = "/listener.v1.ListenerService/GetTask"
	ListenerService_PutResult_FullMethodName      = "/listener.v1.ListenerService/PutResult"
)

// ListenerServiceClient is the client API for ListenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListenerServiceClient interface {
	// Update listener's information. Can be used to set location and meta infromation
	// of where listener works
	UpdateListener(ctx context.Context, in *UpdateListenerRequest, opts ...grpc.CallOption) (*UpdateListenerResponse, error)
	// Register new agent with collected information from compromised OS
	RegisterAgent(ctx context.Context, in *RegisterAgentRequest, opts ...grpc.CallOption) (*RegisterAgentResponse, error)
	// Get task for agent from queue
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	// Save task's results
	PutResult(ctx context.Context, in *PutResultRequest, opts ...grpc.CallOption) (*PutResultResponse, error)
}

type listenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListenerServiceClient(cc grpc.ClientConnInterface) ListenerServiceClient {
	return &listenerServiceClient{cc}
}

func (c *listenerServiceClient) UpdateListener(ctx context.Context, in *UpdateListenerRequest, opts ...grpc.CallOption) (*UpdateListenerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateListenerResponse)
	err := c.cc.Invoke(ctx, ListenerService_UpdateListener_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerServiceClient) RegisterAgent(ctx context.Context, in *RegisterAgentRequest, opts ...grpc.CallOption) (*RegisterAgentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterAgentResponse)
	err := c.cc.Invoke(ctx, ListenerService_RegisterAgent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, ListenerService_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listenerServiceClient) PutResult(ctx context.Context, in *PutResultRequest, opts ...grpc.CallOption) (*PutResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PutResultResponse)
	err := c.cc.Invoke(ctx, ListenerService_PutResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerServiceServer is the server API for ListenerService service.
// All implementations must embed UnimplementedListenerServiceServer
// for forward compatibility.
type ListenerServiceServer interface {
	// Update listener's information. Can be used to set location and meta infromation
	// of where listener works
	UpdateListener(context.Context, *UpdateListenerRequest) (*UpdateListenerResponse, error)
	// Register new agent with collected information from compromised OS
	RegisterAgent(context.Context, *RegisterAgentRequest) (*RegisterAgentResponse, error)
	// Get task for agent from queue
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	// Save task's results
	PutResult(context.Context, *PutResultRequest) (*PutResultResponse, error)
	mustEmbedUnimplementedListenerServiceServer()
}

// UnimplementedListenerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListenerServiceServer struct{}

func (UnimplementedListenerServiceServer) UpdateListener(context.Context, *UpdateListenerRequest) (*UpdateListenerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListener not implemented")
}
func (UnimplementedListenerServiceServer) RegisterAgent(context.Context, *RegisterAgentRequest) (*RegisterAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAgent not implemented")
}
func (UnimplementedListenerServiceServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedListenerServiceServer) PutResult(context.Context, *PutResultRequest) (*PutResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutResult not implemented")
}
func (UnimplementedListenerServiceServer) mustEmbedUnimplementedListenerServiceServer() {}
func (UnimplementedListenerServiceServer) testEmbeddedByValue()                         {}

// UnsafeListenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListenerServiceServer will
// result in compilation errors.
type UnsafeListenerServiceServer interface {
	mustEmbedUnimplementedListenerServiceServer()
}

func RegisterListenerServiceServer(s grpc.ServiceRegistrar, srv ListenerServiceServer) {
	// If the following call pancis, it indicates UnimplementedListenerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListenerService_ServiceDesc, srv)
}

func _ListenerService_UpdateListener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerServiceServer).UpdateListener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerService_UpdateListener_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerServiceServer).UpdateListener(ctx, req.(*UpdateListenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerService_RegisterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerServiceServer).RegisterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerService_RegisterAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerServiceServer).RegisterAgent(ctx, req.(*RegisterAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListenerService_PutResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerServiceServer).PutResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListenerService_PutResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerServiceServer).PutResult(ctx, req.(*PutResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListenerService_ServiceDesc is the grpc.ServiceDesc for ListenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listener.v1.ListenerService",
	HandlerType: (*ListenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateListener",
			Handler:    _ListenerService_UpdateListener_Handler,
		},
		{
			MethodName: "RegisterAgent",
			Handler:    _ListenerService_RegisterAgent_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _ListenerService_GetTask_Handler,
		},
		{
			MethodName: "PutResult",
			Handler:    _ListenerService_PutResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "listener/v1/listener.proto",
}
