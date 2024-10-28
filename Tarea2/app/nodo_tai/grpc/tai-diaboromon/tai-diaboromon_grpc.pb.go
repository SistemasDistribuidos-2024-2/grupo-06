// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: tai-diaboromon.proto

package tai_diaboromon

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
	TaiDiaboromonService_AtaqueDiaboromon_FullMethodName = "/taidiaboromon.TaiDiaboromonService/AtaqueDiaboromon"
	TaiDiaboromonService_EstadoVida_FullMethodName       = "/taidiaboromon.TaiDiaboromonService/EstadoVida"
)

// TaiDiaboromonServiceClient is the client API for TaiDiaboromonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio para la comunicación entre Nodo Tai y Diaboromon
type TaiDiaboromonServiceClient interface {
	// RPC para enviar ataque desde Diaboromon a Nodo Tai
	AtaqueDiaboromon(ctx context.Context, in *SolicitudAtaque, opts ...grpc.CallOption) (*ConfirmacionAtaque, error)
	// RPC para que Nodo Tai envíe el resultado de la vida restante
	EstadoVida(ctx context.Context, in *SolicitudEstado, opts ...grpc.CallOption) (*RespuestaEstado, error)
}

type taiDiaboromonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaiDiaboromonServiceClient(cc grpc.ClientConnInterface) TaiDiaboromonServiceClient {
	return &taiDiaboromonServiceClient{cc}
}

func (c *taiDiaboromonServiceClient) AtaqueDiaboromon(ctx context.Context, in *SolicitudAtaque, opts ...grpc.CallOption) (*ConfirmacionAtaque, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmacionAtaque)
	err := c.cc.Invoke(ctx, TaiDiaboromonService_AtaqueDiaboromon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taiDiaboromonServiceClient) EstadoVida(ctx context.Context, in *SolicitudEstado, opts ...grpc.CallOption) (*RespuestaEstado, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RespuestaEstado)
	err := c.cc.Invoke(ctx, TaiDiaboromonService_EstadoVida_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaiDiaboromonServiceServer is the server API for TaiDiaboromonService service.
// All implementations must embed UnimplementedTaiDiaboromonServiceServer
// for forward compatibility.
//
// Servicio para la comunicación entre Nodo Tai y Diaboromon
type TaiDiaboromonServiceServer interface {
	// RPC para enviar ataque desde Diaboromon a Nodo Tai
	AtaqueDiaboromon(context.Context, *SolicitudAtaque) (*ConfirmacionAtaque, error)
	// RPC para que Nodo Tai envíe el resultado de la vida restante
	EstadoVida(context.Context, *SolicitudEstado) (*RespuestaEstado, error)
	mustEmbedUnimplementedTaiDiaboromonServiceServer()
}

// UnimplementedTaiDiaboromonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaiDiaboromonServiceServer struct{}

func (UnimplementedTaiDiaboromonServiceServer) AtaqueDiaboromon(context.Context, *SolicitudAtaque) (*ConfirmacionAtaque, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtaqueDiaboromon not implemented")
}
func (UnimplementedTaiDiaboromonServiceServer) EstadoVida(context.Context, *SolicitudEstado) (*RespuestaEstado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstadoVida not implemented")
}
func (UnimplementedTaiDiaboromonServiceServer) mustEmbedUnimplementedTaiDiaboromonServiceServer() {}
func (UnimplementedTaiDiaboromonServiceServer) testEmbeddedByValue()                              {}

// UnsafeTaiDiaboromonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaiDiaboromonServiceServer will
// result in compilation errors.
type UnsafeTaiDiaboromonServiceServer interface {
	mustEmbedUnimplementedTaiDiaboromonServiceServer()
}

func RegisterTaiDiaboromonServiceServer(s grpc.ServiceRegistrar, srv TaiDiaboromonServiceServer) {
	// If the following call pancis, it indicates UnimplementedTaiDiaboromonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaiDiaboromonService_ServiceDesc, srv)
}

func _TaiDiaboromonService_AtaqueDiaboromon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudAtaque)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiDiaboromonServiceServer).AtaqueDiaboromon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaiDiaboromonService_AtaqueDiaboromon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiDiaboromonServiceServer).AtaqueDiaboromon(ctx, req.(*SolicitudAtaque))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaiDiaboromonService_EstadoVida_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolicitudEstado)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiDiaboromonServiceServer).EstadoVida(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaiDiaboromonService_EstadoVida_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiDiaboromonServiceServer).EstadoVida(ctx, req.(*SolicitudEstado))
	}
	return interceptor(ctx, in, info, handler)
}

// TaiDiaboromonService_ServiceDesc is the grpc.ServiceDesc for TaiDiaboromonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaiDiaboromonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taidiaboromon.TaiDiaboromonService",
	HandlerType: (*TaiDiaboromonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AtaqueDiaboromon",
			Handler:    _TaiDiaboromonService_AtaqueDiaboromon_Handler,
		},
		{
			MethodName: "EstadoVida",
			Handler:    _TaiDiaboromonService_EstadoVida_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tai-diaboromon.proto",
}
