// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: sup-serv.proto

package sup_serv

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

// Tipos de operación permitidos
type OperationType int32

const (
	OperationType_AGREGAR    OperationType = 0
	OperationType_RENOMBRAR  OperationType = 1
	OperationType_ACTUALIZAR OperationType = 2
	OperationType_BORRAR     OperationType = 3
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "AGREGAR",
		1: "RENOMBRAR",
		2: "ACTUALIZAR",
		3: "BORRAR",
	}
	OperationType_value = map[string]int32{
		"AGREGAR":    0,
		"RENOMBRAR":  1,
		"ACTUALIZAR": 2,
		"BORRAR":     3,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_sup_serv_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_sup_serv_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_sup_serv_proto_rawDescGZIP(), []int{0}
}

// Estado de la respuesta
type ResponseStatus int32

const (
	ResponseStatus_OK    ResponseStatus = 0
	ResponseStatus_ERROR ResponseStatus = 1
)

// Enum value maps for ResponseStatus.
var (
	ResponseStatus_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	ResponseStatus_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x ResponseStatus) Enum() *ResponseStatus {
	p := new(ResponseStatus)
	*p = x
	return p
}

func (x ResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sup_serv_proto_enumTypes[1].Descriptor()
}

func (ResponseStatus) Type() protoreflect.EnumType {
	return &file_sup_serv_proto_enumTypes[1]
}

func (x ResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseStatus.Descriptor instead.
func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_sup_serv_proto_rawDescGZIP(), []int{1}
}

// Solicitudes enviadas por Supervisores
type SupervisorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region           string        `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`                                                                // Región afectada
	ProductName      string        `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`                                   // Producto afectado
	OperationType    OperationType `protobuf:"varint,3,opt,name=operation_type,json=operationType,proto3,enum=supserv.OperationType" json:"operation_type,omitempty"` // Tipo de operación: AGREGAR, RENOMBRAR, ACTUALIZAR, BORRAR
	Value            *int32        `protobuf:"varint,4,opt,name=value,proto3,oneof" json:"value,omitempty"`                                                           // Valor opcional (para agregar/actualizar)
	NewProductName   *string       `protobuf:"bytes,5,opt,name=new_product_name,json=newProductName,proto3,oneof" json:"new_product_name,omitempty"`                  // Nuevo nombre opcional (para renombrar)
	KnownVectorClock *VectorClock  `protobuf:"bytes,6,opt,name=known_vector_clock,json=knownVectorClock,proto3" json:"known_vector_clock,omitempty"`                  // Reloj vectorial conocido por el Supervisor
}

func (x *SupervisorRequest) Reset() {
	*x = SupervisorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sup_serv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupervisorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupervisorRequest) ProtoMessage() {}

func (x *SupervisorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sup_serv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupervisorRequest.ProtoReflect.Descriptor instead.
func (*SupervisorRequest) Descriptor() ([]byte, []int) {
	return file_sup_serv_proto_rawDescGZIP(), []int{0}
}

func (x *SupervisorRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *SupervisorRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *SupervisorRequest) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_AGREGAR
}

func (x *SupervisorRequest) GetValue() int32 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

func (x *SupervisorRequest) GetNewProductName() string {
	if x != nil && x.NewProductName != nil {
		return *x.NewProductName
	}
	return ""
}

func (x *SupervisorRequest) GetKnownVectorClock() *VectorClock {
	if x != nil {
		return x.KnownVectorClock
	}
	return nil
}

// Respuestas enviadas por los Servidores Hextech
type ServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=supserv.ResponseStatus" json:"status,omitempty"` // Estado de la operación (OK o ERROR)
	VectorClock *VectorClock   `protobuf:"bytes,2,opt,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"` // Reloj vectorial actualizado
	Value       *int32         `protobuf:"varint,3,opt,name=value,proto3,oneof" json:"value,omitempty"`                         // Valor del producto procesado (si aplica)
	Message     *string        `protobuf:"bytes,4,opt,name=message,proto3,oneof" json:"message,omitempty"`                      // Mensaje opcional en caso de error
}

func (x *ServerResponse) Reset() {
	*x = ServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sup_serv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResponse) ProtoMessage() {}

func (x *ServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sup_serv_proto_msgTypes[1]
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
	return file_sup_serv_proto_rawDescGZIP(), []int{1}
}

func (x *ServerResponse) GetStatus() ResponseStatus {
	if x != nil {
		return x.Status
	}
	return ResponseStatus_OK
}

func (x *ServerResponse) GetVectorClock() *VectorClock {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

func (x *ServerResponse) GetValue() int32 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

func (x *ServerResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// Representación de un reloj vectorial
type VectorClock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server1 int32 `protobuf:"varint,1,opt,name=server1,proto3" json:"server1,omitempty"`
	Server2 int32 `protobuf:"varint,2,opt,name=server2,proto3" json:"server2,omitempty"`
	Server3 int32 `protobuf:"varint,3,opt,name=server3,proto3" json:"server3,omitempty"`
}

func (x *VectorClock) Reset() {
	*x = VectorClock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sup_serv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClock) ProtoMessage() {}

func (x *VectorClock) ProtoReflect() protoreflect.Message {
	mi := &file_sup_serv_proto_msgTypes[2]
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
	return file_sup_serv_proto_rawDescGZIP(), []int{2}
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

var File_sup_serv_proto protoreflect.FileDescriptor

var file_sup_serv_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x75, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x22, 0xba, 0x02, 0x0a, 0x11, 0x53, 0x75,
	0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0e, 0x6e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x12, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x10, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x75, 0x70, 0x73,
	0x65, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x5b, 0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x33,
	0x2a, 0x47, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x47, 0x52, 0x45, 0x47, 0x41, 0x52, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x52, 0x45, 0x4e, 0x4f, 0x4d, 0x42, 0x52, 0x41, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x41, 0x43, 0x54, 0x55, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x52, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x42, 0x4f, 0x52, 0x52, 0x41, 0x52, 0x10, 0x03, 0x2a, 0x23, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0x56,
	0x0a, 0x0e, 0x48, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x2e, 0x73, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x73, 0x75, 0x70, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x75, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sup_serv_proto_rawDescOnce sync.Once
	file_sup_serv_proto_rawDescData = file_sup_serv_proto_rawDesc
)

func file_sup_serv_proto_rawDescGZIP() []byte {
	file_sup_serv_proto_rawDescOnce.Do(func() {
		file_sup_serv_proto_rawDescData = protoimpl.X.CompressGZIP(file_sup_serv_proto_rawDescData)
	})
	return file_sup_serv_proto_rawDescData
}

var file_sup_serv_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sup_serv_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sup_serv_proto_goTypes = []any{
	(OperationType)(0),        // 0: supserv.OperationType
	(ResponseStatus)(0),       // 1: supserv.ResponseStatus
	(*SupervisorRequest)(nil), // 2: supserv.SupervisorRequest
	(*ServerResponse)(nil),    // 3: supserv.ServerResponse
	(*VectorClock)(nil),       // 4: supserv.VectorClock
}
var file_sup_serv_proto_depIdxs = []int32{
	0, // 0: supserv.SupervisorRequest.operation_type:type_name -> supserv.OperationType
	4, // 1: supserv.SupervisorRequest.known_vector_clock:type_name -> supserv.VectorClock
	1, // 2: supserv.ServerResponse.status:type_name -> supserv.ResponseStatus
	4, // 3: supserv.ServerResponse.vector_clock:type_name -> supserv.VectorClock
	2, // 4: supserv.HextechService.HandleRequest:input_type -> supserv.SupervisorRequest
	3, // 5: supserv.HextechService.HandleRequest:output_type -> supserv.ServerResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sup_serv_proto_init() }
func file_sup_serv_proto_init() {
	if File_sup_serv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sup_serv_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SupervisorRequest); i {
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
		file_sup_serv_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_sup_serv_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
	file_sup_serv_proto_msgTypes[0].OneofWrappers = []any{}
	file_sup_serv_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sup_serv_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sup_serv_proto_goTypes,
		DependencyIndexes: file_sup_serv_proto_depIdxs,
		EnumInfos:         file_sup_serv_proto_enumTypes,
		MessageInfos:      file_sup_serv_proto_msgTypes,
	}.Build()
	File_sup_serv_proto = out.File
	file_sup_serv_proto_rawDesc = nil
	file_sup_serv_proto_goTypes = nil
	file_sup_serv_proto_depIdxs = nil
}
