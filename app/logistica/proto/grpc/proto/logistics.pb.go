// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: logistics.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Mensaje para los paquetes enviados por las facciones (clientes)
type PackageOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete        string `protobuf:"bytes,1,opt,name=id_paquete,json=idPaquete,proto3" json:"id_paquete,omitempty"`
	Faccion          string `protobuf:"bytes,2,opt,name=faccion,proto3" json:"faccion,omitempty"`                            // Ostronitas o Grineer
	TipoPaquete      string `protobuf:"bytes,3,opt,name=tipo_paquete,json=tipoPaquete,proto3" json:"tipo_paquete,omitempty"` // Normal, Prioritario, Ostronitas
	NombreSuministro string `protobuf:"bytes,4,opt,name=nombre_suministro,json=nombreSuministro,proto3" json:"nombre_suministro,omitempty"`
	ValorSuministro  int32  `protobuf:"varint,5,opt,name=valor_suministro,json=valorSuministro,proto3" json:"valor_suministro,omitempty"`
	Destino          string `protobuf:"bytes,6,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *PackageOrder) Reset() {
	*x = PackageOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageOrder) ProtoMessage() {}

func (x *PackageOrder) ProtoReflect() protoreflect.Message {
	mi := &file_logistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageOrder.ProtoReflect.Descriptor instead.
func (*PackageOrder) Descriptor() ([]byte, []int) {
	return file_logistics_proto_rawDescGZIP(), []int{0}
}

func (x *PackageOrder) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *PackageOrder) GetFaccion() string {
	if x != nil {
		return x.Faccion
	}
	return ""
}

func (x *PackageOrder) GetTipoPaquete() string {
	if x != nil {
		return x.TipoPaquete
	}
	return ""
}

func (x *PackageOrder) GetNombreSuministro() string {
	if x != nil {
		return x.NombreSuministro
	}
	return ""
}

func (x *PackageOrder) GetValorSuministro() int32 {
	if x != nil {
		return x.ValorSuministro
	}
	return 0
}

func (x *PackageOrder) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

// Respuesta con el código de seguimiento de una orden
type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodigoSeguimiento string `protobuf:"bytes,1,opt,name=codigo_seguimiento,json=codigoSeguimiento,proto3" json:"codigo_seguimiento,omitempty"`
	Mensaje           string `protobuf:"bytes,2,opt,name=mensaje,proto3" json:"mensaje,omitempty"` // Mensaje de éxito o error
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_logistics_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResponse) GetCodigoSeguimiento() string {
	if x != nil {
		return x.CodigoSeguimiento
	}
	return ""
}

func (x *OrderResponse) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

// Solicitud para obtener el estado de un paquete
type TrackingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodigoSeguimiento string `protobuf:"bytes,1,opt,name=codigo_seguimiento,json=codigoSeguimiento,proto3" json:"codigo_seguimiento,omitempty"`
}

func (x *TrackingRequest) Reset() {
	*x = TrackingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackingRequest) ProtoMessage() {}

func (x *TrackingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logistics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackingRequest.ProtoReflect.Descriptor instead.
func (*TrackingRequest) Descriptor() ([]byte, []int) {
	return file_logistics_proto_rawDescGZIP(), []int{2}
}

func (x *TrackingRequest) GetCodigoSeguimiento() string {
	if x != nil {
		return x.CodigoSeguimiento
	}
	return ""
}

// Respuesta con el estado actual del paquete
type TrackingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado     string `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"` // Ej: En camino, Entregado, No Entregado
	IdCaravana string `protobuf:"bytes,2,opt,name=id_caravana,json=idCaravana,proto3" json:"id_caravana,omitempty"`
	Intentos   int32  `protobuf:"varint,3,opt,name=intentos,proto3" json:"intentos,omitempty"`
}

func (x *TrackingResponse) Reset() {
	*x = TrackingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackingResponse) ProtoMessage() {}

func (x *TrackingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logistics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackingResponse.ProtoReflect.Descriptor instead.
func (*TrackingResponse) Descriptor() ([]byte, []int) {
	return file_logistics_proto_rawDescGZIP(), []int{3}
}

func (x *TrackingResponse) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *TrackingResponse) GetIdCaravana() string {
	if x != nil {
		return x.IdCaravana
	}
	return ""
}

func (x *TrackingResponse) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

// Mensaje para las caravanas que indica los detalles de la entrega
type DeliveryInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete       string `protobuf:"bytes,1,opt,name=id_paquete,json=idPaquete,proto3" json:"id_paquete,omitempty"`
	TipoCaravana    string `protobuf:"bytes,2,opt,name=tipo_caravana,json=tipoCaravana,proto3" json:"tipo_caravana,omitempty"` // Ostronitas, General
	WarframeEscolta string `protobuf:"bytes,3,opt,name=warframe_escolta,json=warframeEscolta,proto3" json:"warframe_escolta,omitempty"`
	Destino         string `protobuf:"bytes,4,opt,name=destino,proto3" json:"destino,omitempty"`
	TipoPaquete     string `protobuf:"bytes,5,opt,name=tipo_paquete,json=tipoPaquete,proto3" json:"tipo_paquete,omitempty"`
	Seguimiento     string `protobuf:"bytes,6,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Valor           int32  `protobuf:"varint,7,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *DeliveryInstruction) Reset() {
	*x = DeliveryInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryInstruction) ProtoMessage() {}

func (x *DeliveryInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_logistics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryInstruction.ProtoReflect.Descriptor instead.
func (*DeliveryInstruction) Descriptor() ([]byte, []int) {
	return file_logistics_proto_rawDescGZIP(), []int{4}
}

func (x *DeliveryInstruction) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *DeliveryInstruction) GetTipoCaravana() string {
	if x != nil {
		return x.TipoCaravana
	}
	return ""
}

func (x *DeliveryInstruction) GetWarframeEscolta() string {
	if x != nil {
		return x.WarframeEscolta
	}
	return ""
}

func (x *DeliveryInstruction) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *DeliveryInstruction) GetTipoPaquete() string {
	if x != nil {
		return x.TipoPaquete
	}
	return ""
}

func (x *DeliveryInstruction) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

func (x *DeliveryInstruction) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

// Mensaje de confirmación de la caravana
type DeliveryStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdPaquete string `protobuf:"bytes,1,opt,name=id_paquete,json=idPaquete,proto3" json:"id_paquete,omitempty"`
	Estado    string `protobuf:"bytes,2,opt,name=estado,proto3" json:"estado,omitempty"` // Entregado, No Entregado
	Intentos  int32  `protobuf:"varint,3,opt,name=intentos,proto3" json:"intentos,omitempty"`
}

func (x *DeliveryStatus) Reset() {
	*x = DeliveryStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryStatus) ProtoMessage() {}

func (x *DeliveryStatus) ProtoReflect() protoreflect.Message {
	mi := &file_logistics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryStatus.ProtoReflect.Descriptor instead.
func (*DeliveryStatus) Descriptor() ([]byte, []int) {
	return file_logistics_proto_rawDescGZIP(), []int{5}
}

func (x *DeliveryStatus) GetIdPaquete() string {
	if x != nil {
		return x.IdPaquete
	}
	return ""
}

func (x *DeliveryStatus) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *DeliveryStatus) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

var File_logistics_proto protoreflect.FileDescriptor

var file_logistics_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x0c, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64,
	0x5f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63,
	0x63, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x63,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x70, 0x61, 0x71, 0x75,
	0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x70, 0x6f, 0x50,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x5f, 0x73, 0x75, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x53, 0x75, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x6f, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x5f, 0x73, 0x75, 0x6d,
	0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x76,
	0x61, 0x6c, 0x6f, 0x72, 0x53, 0x75, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0x58, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x64,
	0x69, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67,
	0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x22, 0x40, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x5f,
	0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x61, 0x76, 0x61, 0x6e, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x43, 0x61, 0x72, 0x61, 0x76, 0x61, 0x6e,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x22, 0xf9, 0x01,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x70, 0x61, 0x71, 0x75,
	0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x50, 0x61, 0x71,
	0x75, 0x65, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x63, 0x61, 0x72,
	0x61, 0x76, 0x61, 0x6e, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x70,
	0x6f, 0x43, 0x61, 0x72, 0x61, 0x76, 0x61, 0x6e, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x77, 0x61, 0x72,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x73, 0x63, 0x6f, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x61, 0x72, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x45, 0x73, 0x63,
	0x6f, 0x6c, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x70, 0x6f, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x63, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x64, 0x5f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x64, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x32, 0x9f,
	0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xa8, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x61, 0x76, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x49, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_logistics_proto_rawDescOnce sync.Once
	file_logistics_proto_rawDescData = file_logistics_proto_rawDesc
)

func file_logistics_proto_rawDescGZIP() []byte {
	file_logistics_proto_rawDescOnce.Do(func() {
		file_logistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_logistics_proto_rawDescData)
	})
	return file_logistics_proto_rawDescData
}

var file_logistics_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_logistics_proto_goTypes = []any{
	(*PackageOrder)(nil),        // 0: logistics.PackageOrder
	(*OrderResponse)(nil),       // 1: logistics.OrderResponse
	(*TrackingRequest)(nil),     // 2: logistics.TrackingRequest
	(*TrackingResponse)(nil),    // 3: logistics.TrackingResponse
	(*DeliveryInstruction)(nil), // 4: logistics.DeliveryInstruction
	(*DeliveryStatus)(nil),      // 5: logistics.DeliveryStatus
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_logistics_proto_depIdxs = []int32{
	0, // 0: logistics.LogisticsService.SendOrder:input_type -> logistics.PackageOrder
	2, // 1: logistics.LogisticsService.CheckOrderStatus:input_type -> logistics.TrackingRequest
	4, // 2: logistics.CaravanService.AssignDelivery:input_type -> logistics.DeliveryInstruction
	5, // 3: logistics.CaravanService.ReportDeliveryStatus:input_type -> logistics.DeliveryStatus
	1, // 4: logistics.LogisticsService.SendOrder:output_type -> logistics.OrderResponse
	3, // 5: logistics.LogisticsService.CheckOrderStatus:output_type -> logistics.TrackingResponse
	5, // 6: logistics.CaravanService.AssignDelivery:output_type -> logistics.DeliveryStatus
	6, // 7: logistics.CaravanService.ReportDeliveryStatus:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logistics_proto_init() }
func file_logistics_proto_init() {
	if File_logistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logistics_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PackageOrder); i {
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
		file_logistics_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OrderResponse); i {
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
		file_logistics_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TrackingRequest); i {
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
		file_logistics_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TrackingResponse); i {
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
		file_logistics_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeliveryInstruction); i {
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
		file_logistics_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeliveryStatus); i {
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
			RawDescriptor: file_logistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_logistics_proto_goTypes,
		DependencyIndexes: file_logistics_proto_depIdxs,
		MessageInfos:      file_logistics_proto_msgTypes,
	}.Build()
	File_logistics_proto = out.File
	file_logistics_proto_rawDesc = nil
	file_logistics_proto_goTypes = nil
	file_logistics_proto_depIdxs = nil
}
