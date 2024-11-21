// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: sup-serv.proto

package sup_serv

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
	HextechService_HandleRequest_FullMethodName = "/supserv.HextechService/HandleRequest"
)

// HextechServiceClient is the client API for HextechService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio implementado por los Servidores Hextech
type HextechServiceClient interface {
	HandleRequest(ctx context.Context, in *SupervisorRequest, opts ...grpc.CallOption) (*ServerResponse, error)
}

type hextechServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHextechServiceClient(cc grpc.ClientConnInterface) HextechServiceClient {
	return &hextechServiceClient{cc}
}

func (c *hextechServiceClient) HandleRequest(ctx context.Context, in *SupervisorRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, HextechService_HandleRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HextechServiceServer is the server API for HextechService service.
// All implementations must embed UnimplementedHextechServiceServer
// for forward compatibility.
//
// Servicio implementado por los Servidores Hextech
type HextechServiceServer interface {
	HandleRequest(context.Context, *SupervisorRequest) (*ServerResponse, error)
	mustEmbedUnimplementedHextechServiceServer()
}

// UnimplementedHextechServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHextechServiceServer struct{}

func (UnimplementedHextechServiceServer) HandleRequest(context.Context, *SupervisorRequest) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleRequest not implemented")
}
func (UnimplementedHextechServiceServer) mustEmbedUnimplementedHextechServiceServer() {}
func (UnimplementedHextechServiceServer) testEmbeddedByValue()                        {}

// UnsafeHextechServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HextechServiceServer will
// result in compilation errors.
type UnsafeHextechServiceServer interface {
	mustEmbedUnimplementedHextechServiceServer()
}

func RegisterHextechServiceServer(s grpc.ServiceRegistrar, srv HextechServiceServer) {
	// If the following call pancis, it indicates UnimplementedHextechServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HextechService_ServiceDesc, srv)
}

func _HextechService_HandleRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupervisorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HextechServiceServer).HandleRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HextechService_HandleRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HextechServiceServer).HandleRequest(ctx, req.(*SupervisorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HextechService_ServiceDesc is the grpc.ServiceDesc for HextechService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HextechService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supserv.HextechService",
	HandlerType: (*HextechServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleRequest",
			Handler:    _HextechService_HandleRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sup-serv.proto",
}
