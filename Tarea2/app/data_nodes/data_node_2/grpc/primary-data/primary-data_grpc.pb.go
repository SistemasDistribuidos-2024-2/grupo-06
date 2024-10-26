// data_node.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: primary-data.proto

package primary_data

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
	DataNodeService_GuardarDatos_FullMethodName = "/datanode.DataNodeService/GuardarDatos"
)

// DataNodeServiceClient is the client API for DataNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Servicio para los Data Nodes que reciben datos del Primary Node
type DataNodeServiceClient interface {
	// RPC para recibir datos procesados del Primary Node
	GuardarDatos(ctx context.Context, in *DatosParaDataNode, opts ...grpc.CallOption) (*Confirmacion, error)
}

type dataNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeServiceClient(cc grpc.ClientConnInterface) DataNodeServiceClient {
	return &dataNodeServiceClient{cc}
}

func (c *dataNodeServiceClient) GuardarDatos(ctx context.Context, in *DatosParaDataNode, opts ...grpc.CallOption) (*Confirmacion, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Confirmacion)
	err := c.cc.Invoke(ctx, DataNodeService_GuardarDatos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServiceServer is the server API for DataNodeService service.
// All implementations must embed UnimplementedDataNodeServiceServer
// for forward compatibility.
//
// Servicio para los Data Nodes que reciben datos del Primary Node
type DataNodeServiceServer interface {
	// RPC para recibir datos procesados del Primary Node
	GuardarDatos(context.Context, *DatosParaDataNode) (*Confirmacion, error)
	mustEmbedUnimplementedDataNodeServiceServer()
}

// UnimplementedDataNodeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataNodeServiceServer struct{}

func (UnimplementedDataNodeServiceServer) GuardarDatos(context.Context, *DatosParaDataNode) (*Confirmacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuardarDatos not implemented")
}
func (UnimplementedDataNodeServiceServer) mustEmbedUnimplementedDataNodeServiceServer() {}
func (UnimplementedDataNodeServiceServer) testEmbeddedByValue()                         {}

// UnsafeDataNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServiceServer will
// result in compilation errors.
type UnsafeDataNodeServiceServer interface {
	mustEmbedUnimplementedDataNodeServiceServer()
}

func RegisterDataNodeServiceServer(s grpc.ServiceRegistrar, srv DataNodeServiceServer) {
	// If the following call pancis, it indicates UnimplementedDataNodeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataNodeService_ServiceDesc, srv)
}

func _DataNodeService_GuardarDatos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatosParaDataNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).GuardarDatos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataNodeService_GuardarDatos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).GuardarDatos(ctx, req.(*DatosParaDataNode))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNodeService_ServiceDesc is the grpc.ServiceDesc for DataNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datanode.DataNodeService",
	HandlerType: (*DataNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GuardarDatos",
			Handler:    _DataNodeService_GuardarDatos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "primary-data.proto",
}
