// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: jayce-broker.proto

package jayce_broker

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
	JayceBrokerService_ObtenerProducto_FullMethodName = "/jaycebroker.JayceBrokerService/ObtenerProducto"
)

// JayceBrokerServiceClient is the client API for JayceBrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio que define las operaciones entre Jayce y el Broker
type JayceBrokerServiceClient interface {
	// Método para obtener la información de un producto en una región específica
	ObtenerProducto(ctx context.Context, in *JayceRequest, opts ...grpc.CallOption) (*JayceResponse, error)
}

type jayceBrokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJayceBrokerServiceClient(cc grpc.ClientConnInterface) JayceBrokerServiceClient {
	return &jayceBrokerServiceClient{cc}
}

func (c *jayceBrokerServiceClient) ObtenerProducto(ctx context.Context, in *JayceRequest, opts ...grpc.CallOption) (*JayceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JayceResponse)
	err := c.cc.Invoke(ctx, JayceBrokerService_ObtenerProducto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JayceBrokerServiceServer is the server API for JayceBrokerService service.
// All implementations must embed UnimplementedJayceBrokerServiceServer
// for forward compatibility.
//
// Servicio que define las operaciones entre Jayce y el Broker
type JayceBrokerServiceServer interface {
	// Método para obtener la información de un producto en una región específica
	ObtenerProducto(context.Context, *JayceRequest) (*JayceResponse, error)
	mustEmbedUnimplementedJayceBrokerServiceServer()
}

// UnimplementedJayceBrokerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJayceBrokerServiceServer struct{}

func (UnimplementedJayceBrokerServiceServer) ObtenerProducto(context.Context, *JayceRequest) (*JayceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerProducto not implemented")
}
func (UnimplementedJayceBrokerServiceServer) mustEmbedUnimplementedJayceBrokerServiceServer() {}
func (UnimplementedJayceBrokerServiceServer) testEmbeddedByValue()                            {}

// UnsafeJayceBrokerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JayceBrokerServiceServer will
// result in compilation errors.
type UnsafeJayceBrokerServiceServer interface {
	mustEmbedUnimplementedJayceBrokerServiceServer()
}

func RegisterJayceBrokerServiceServer(s grpc.ServiceRegistrar, srv JayceBrokerServiceServer) {
	// If the following call pancis, it indicates UnimplementedJayceBrokerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JayceBrokerService_ServiceDesc, srv)
}

func _JayceBrokerService_ObtenerProducto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JayceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JayceBrokerServiceServer).ObtenerProducto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JayceBrokerService_ObtenerProducto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JayceBrokerServiceServer).ObtenerProducto(ctx, req.(*JayceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JayceBrokerService_ServiceDesc is the grpc.ServiceDesc for JayceBrokerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JayceBrokerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaycebroker.JayceBrokerService",
	HandlerType: (*JayceBrokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ObtenerProducto",
			Handler:    _JayceBrokerService_ObtenerProducto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jayce-broker.proto",
}
