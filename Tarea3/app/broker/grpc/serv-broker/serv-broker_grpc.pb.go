// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: serv-broker.proto

package serv_broker

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
	HextechServerService_ProcessRequest_FullMethodName = "/servbroker.HextechServerService/ProcessRequest"
)

// HextechServerServiceClient is the client API for HextechServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio que define las operaciones entre el Broker y los Servidores Hextech
type HextechServerServiceClient interface {
	ProcessRequest(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerResponse, error)
}

type hextechServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHextechServerServiceClient(cc grpc.ClientConnInterface) HextechServerServiceClient {
	return &hextechServerServiceClient{cc}
}

func (c *hextechServerServiceClient) ProcessRequest(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, HextechServerService_ProcessRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HextechServerServiceServer is the server API for HextechServerService service.
// All implementations must embed UnimplementedHextechServerServiceServer
// for forward compatibility.
//
// Servicio que define las operaciones entre el Broker y los Servidores Hextech
type HextechServerServiceServer interface {
	ProcessRequest(context.Context, *ServerRequest) (*ServerResponse, error)
	mustEmbedUnimplementedHextechServerServiceServer()
}

// UnimplementedHextechServerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHextechServerServiceServer struct{}

func (UnimplementedHextechServerServiceServer) ProcessRequest(context.Context, *ServerRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessRequest not implemented")
}
func (UnimplementedHextechServerServiceServer) mustEmbedUnimplementedHextechServerServiceServer() {}
func (UnimplementedHextechServerServiceServer) testEmbeddedByValue()                              {}

// UnsafeHextechServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HextechServerServiceServer will
// result in compilation errors.
type UnsafeHextechServerServiceServer interface {
	mustEmbedUnimplementedHextechServerServiceServer()
}

func RegisterHextechServerServiceServer(s grpc.ServiceRegistrar, srv HextechServerServiceServer) {
	// If the following call pancis, it indicates UnimplementedHextechServerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HextechServerService_ServiceDesc, srv)
}

func _HextechServerService_ProcessRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HextechServerServiceServer).ProcessRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HextechServerService_ProcessRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HextechServerServiceServer).ProcessRequest(ctx, req.(*ServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HextechServerService_ServiceDesc is the grpc.ServiceDesc for HextechServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HextechServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "servbroker.HextechServerService",
	HandlerType: (*HextechServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessRequest",
			Handler:    _HextechServerService_ProcessRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serv-broker.proto",
}
