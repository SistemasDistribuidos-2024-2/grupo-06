// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: sup-broker.proto

package sup_broker

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
	BrokerService_GetServer_FullMethodName            = "/supbroker.BrokerService/GetServer"
	BrokerService_ResolveInconsistency_FullMethodName = "/supbroker.BrokerService/ResolveInconsistency"
)

// BrokerServiceClient is the client API for BrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio del Broker para comunicación con Supervisores
type BrokerServiceClient interface {
	// Solicita la dirección de un Servidor Hextech
	GetServer(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	// Resuelve inconsistencias notificadas por el Supervisor
	ResolveInconsistency(ctx context.Context, in *InconsistencyRequest, opts ...grpc.CallOption) (*ServerResponse, error)
}

type brokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &brokerServiceClient{cc}
}

func (c *brokerServiceClient) GetServer(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, BrokerService_GetServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) ResolveInconsistency(ctx context.Context, in *InconsistencyRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, BrokerService_ResolveInconsistency_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServiceServer is the server API for BrokerService service.
// All implementations must embed UnimplementedBrokerServiceServer
// for forward compatibility.
//
// Servicio del Broker para comunicación con Supervisores
type BrokerServiceServer interface {
	// Solicita la dirección de un Servidor Hextech
	GetServer(context.Context, *ServerRequest) (*ServerResponse, error)
	// Resuelve inconsistencias notificadas por el Supervisor
	ResolveInconsistency(context.Context, *InconsistencyRequest) (*ServerResponse, error)
	mustEmbedUnimplementedBrokerServiceServer()
}

// UnimplementedBrokerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBrokerServiceServer struct{}

func (UnimplementedBrokerServiceServer) GetServer(context.Context, *ServerRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (UnimplementedBrokerServiceServer) ResolveInconsistency(context.Context, *InconsistencyRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveInconsistency not implemented")
}
func (UnimplementedBrokerServiceServer) mustEmbedUnimplementedBrokerServiceServer() {}
func (UnimplementedBrokerServiceServer) testEmbeddedByValue()                       {}

// UnsafeBrokerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServiceServer will
// result in compilation errors.
type UnsafeBrokerServiceServer interface {
	mustEmbedUnimplementedBrokerServiceServer()
}

func RegisterBrokerServiceServer(s grpc.ServiceRegistrar, srv BrokerServiceServer) {
	// If the following call pancis, it indicates UnimplementedBrokerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BrokerService_ServiceDesc, srv)
}

func _BrokerService_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrokerService_GetServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).GetServer(ctx, req.(*ServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_ResolveInconsistency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InconsistencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).ResolveInconsistency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BrokerService_ResolveInconsistency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).ResolveInconsistency(ctx, req.(*InconsistencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BrokerService_ServiceDesc is the grpc.ServiceDesc for BrokerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supbroker.BrokerService",
	HandlerType: (*BrokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServer",
			Handler:    _BrokerService_GetServer_Handler,
		},
		{
			MethodName: "ResolveInconsistency",
			Handler:    _BrokerService_ResolveInconsistency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sup-broker.proto",
}
