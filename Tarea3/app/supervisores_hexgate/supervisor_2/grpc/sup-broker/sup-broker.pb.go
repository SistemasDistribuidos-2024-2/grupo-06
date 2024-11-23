// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: sup-broker.proto

package sup_broker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Solicitud básica para obtener un servidor
type ServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"` // Región opcional para balanceo basado en región
}

func (x *ServerRequest) Reset() {
	*x = ServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sup_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerRequest) ProtoMessage() {}

func (x *ServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sup_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerRequest.ProtoReflect.Descriptor instead.
func (*ServerRequest) Descriptor() ([]byte, []int) {
	return file_sup_broker_proto_rawDescGZIP(), []int{0}
}

func (x *ServerRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

// Respuesta con la dirección del Servidor Hextech
type ServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerAddress string `protobuf:"bytes,1,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"` // Dirección del Servidor Hextech asignado
}

func (x *ServerResponse) Reset() {
	*x = ServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sup_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResponse) ProtoMessage() {}

func (x *ServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sup_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerResponse.ProtoReflect.Descriptor instead.
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return file_sup_broker_proto_rawDescGZIP(), []int{1}
}

func (x *ServerResponse) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

// Solicitud de inconsistencia enviada por el Supervisor
type InconsistencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region          string       `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`                                          // Región del producto
	ProductName     string       `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`             // Nombre del producto
	SupervisorClock *VectorClock `protobuf:"bytes,3,opt,name=supervisor_clock,json=supervisorClock,proto3" json:"supervisor_clock,omitempty"` // Reloj vectorial del Supervisor
}

func (x *InconsistencyRequest) Reset() {
	*x = InconsistencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sup_broker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InconsistencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InconsistencyRequest) ProtoMessage() {}

func (x *InconsistencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sup_broker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InconsistencyRequest.ProtoReflect.Descriptor instead.
func (*InconsistencyRequest) Descriptor() ([]byte, []int) {
	return file_sup_broker_proto_rawDescGZIP(), []int{2}
}

func (x *InconsistencyRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *InconsistencyRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *InconsistencyRequest) GetSupervisorClock() *VectorClock {
	if x != nil {
		return x.SupervisorClock
	}
	return nil
}

// Reloj vectorial para sincronización entre servidores
type VectorClock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server1 int32 `protobuf:"varint,1,opt,name=server1,proto3" json:"server1,omitempty"` // Reloj de servidor 1
	Server2 int32 `protobuf:"varint,2,opt,name=server2,proto3" json:"server2,omitempty"` // Reloj de servidor 2
	Server3 int32 `protobuf:"varint,3,opt,name=server3,proto3" json:"server3,omitempty"` // Reloj de servidor 3
}

func (x *VectorClock) Reset() {
	*x = VectorClock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sup_broker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClock) ProtoMessage() {}

func (x *VectorClock) ProtoReflect() protoreflect.Message {
	mi := &file_sup_broker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorClock.ProtoReflect.Descriptor instead.
func (*VectorClock) Descriptor() ([]byte, []int) {
	return file_sup_broker_proto_rawDescGZIP(), []int{3}
}

func (x *VectorClock) GetServer1() int32 {
	if x != nil {
		return x.Server1
	}
	return 0
}

func (x *VectorClock) GetServer2() int32 {
	if x != nil {
		return x.Server2
	}
	return 0
}

func (x *VectorClock) GetServer3() int32 {
	if x != nil {
		return x.Server3
	}
	return 0
}

var File_sup_broker_proto protoreflect.FileDescriptor

var file_sup_broker_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x75, 0x70, 0x2d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x73, 0x75, 0x70, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x27, 0x0a,
	0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x94, 0x01, 0x0a, 0x14, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x75, 0x70, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x5b, 0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x33, 0x32, 0xa5, 0x01, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x18, 0x2e, 0x73, 0x75, 0x70, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x75, 0x70, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x1f, 0x2e, 0x73, 0x75, 0x70, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x63, 0x6f,
	0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x75, 0x70, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x75, 0x70, 0x2d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sup_broker_proto_rawDescOnce sync.Once
	file_sup_broker_proto_rawDescData = file_sup_broker_proto_rawDesc
)

func file_sup_broker_proto_rawDescGZIP() []byte {
	file_sup_broker_proto_rawDescOnce.Do(func() {
		file_sup_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_sup_broker_proto_rawDescData)
	})
	return file_sup_broker_proto_rawDescData
}

var file_sup_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sup_broker_proto_goTypes = []any{
	(*ServerRequest)(nil),        // 0: supbroker.ServerRequest
	(*ServerResponse)(nil),       // 1: supbroker.ServerResponse
	(*InconsistencyRequest)(nil), // 2: supbroker.InconsistencyRequest
	(*VectorClock)(nil),          // 3: supbroker.VectorClock
}
var file_sup_broker_proto_depIdxs = []int32{
	3, // 0: supbroker.InconsistencyRequest.supervisor_clock:type_name -> supbroker.VectorClock
	0, // 1: supbroker.BrokerService.GetServer:input_type -> supbroker.ServerRequest
	2, // 2: supbroker.BrokerService.ResolveInconsistency:input_type -> supbroker.InconsistencyRequest
	1, // 3: supbroker.BrokerService.GetServer:output_type -> supbroker.ServerResponse
	1, // 4: supbroker.BrokerService.ResolveInconsistency:output_type -> supbroker.ServerResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sup_broker_proto_init() }
func file_sup_broker_proto_init() {
	if File_sup_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sup_broker_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sup_broker_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ServerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sup_broker_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InconsistencyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sup_broker_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VectorClock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sup_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sup_broker_proto_goTypes,
		DependencyIndexes: file_sup_broker_proto_depIdxs,
		MessageInfos:      file_sup_broker_proto_msgTypes,
	}.Build()
	File_sup_broker_proto = out.File
	file_sup_broker_proto_rawDesc = nil
	file_sup_broker_proto_goTypes = nil
	file_sup_broker_proto_depIdxs = nil
}