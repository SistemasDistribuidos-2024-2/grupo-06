// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: primary_node.proto

package proto

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

// Mensaje que contiene los detalles a transferir a un Data Node
type DatosDigimon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                   // ID único del Digimon
	Nombre      string `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`            // Nombre del Digimon
	Atributo    string `protobuf:"bytes,3,opt,name=atributo,proto3" json:"atributo,omitempty"`        // Atributo (Data, Vaccine, Virus)
	Sacrificado bool   `protobuf:"varint,4,opt,name=sacrificado,proto3" json:"sacrificado,omitempty"` // Si fue sacrificado (true/false)
}

func (x *DatosDigimon) Reset() {
	*x = DatosDigimon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primary_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatosDigimon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatosDigimon) ProtoMessage() {}

func (x *DatosDigimon) ProtoReflect() protoreflect.Message {
	mi := &file_primary_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatosDigimon.ProtoReflect.Descriptor instead.
func (*DatosDigimon) Descriptor() ([]byte, []int) {
	return file_primary_node_proto_rawDescGZIP(), []int{0}
}

func (x *DatosDigimon) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DatosDigimon) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *DatosDigimon) GetAtributo() string {
	if x != nil {
		return x.Atributo
	}
	return ""
}

func (x *DatosDigimon) GetSacrificado() bool {
	if x != nil {
		return x.Sacrificado
	}
	return false
}

// Mensaje para solicitar datos acumulados (usado por Nodo Tai)
type SolicitudTai struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"` // Mensaje de solicitud
}

func (x *SolicitudTai) Reset() {
	*x = SolicitudTai{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primary_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudTai) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudTai) ProtoMessage() {}

func (x *SolicitudTai) ProtoReflect() protoreflect.Message {
	mi := &file_primary_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudTai.ProtoReflect.Descriptor instead.
func (*SolicitudTai) Descriptor() ([]byte, []int) {
	return file_primary_node_proto_rawDescGZIP(), []int{1}
}

func (x *SolicitudTai) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

// Mensaje de respuesta al Nodo Tai con la cantidad de datos acumulados
type RespuestaTai struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CantidadDatos float32 `protobuf:"fixed32,1,opt,name=cantidad_datos,json=cantidadDatos,proto3" json:"cantidad_datos,omitempty"` // Cantidad de datos acumulada
}

func (x *RespuestaTai) Reset() {
	*x = RespuestaTai{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primary_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaTai) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaTai) ProtoMessage() {}

func (x *RespuestaTai) ProtoReflect() protoreflect.Message {
	mi := &file_primary_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaTai.ProtoReflect.Descriptor instead.
func (*RespuestaTai) Descriptor() ([]byte, []int) {
	return file_primary_node_proto_rawDescGZIP(), []int{2}
}

func (x *RespuestaTai) GetCantidadDatos() float32 {
	if x != nil {
		return x.CantidadDatos
	}
	return 0
}

// Mensaje de confirmación de que los datos fueron transferidos
type Confirmacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"` // Mensaje de confirmación
}

func (x *Confirmacion) Reset() {
	*x = Confirmacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primary_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmacion) ProtoMessage() {}

func (x *Confirmacion) ProtoReflect() protoreflect.Message {
	mi := &file_primary_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmacion.ProtoReflect.Descriptor instead.
func (*Confirmacion) Descriptor() ([]byte, []int) {
	return file_primary_node_proto_rawDescGZIP(), []int{3}
}

func (x *Confirmacion) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

var File_primary_node_proto protoreflect.FileDescriptor

var file_primary_node_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6e, 0x6f, 0x64,
	0x65, 0x22, 0x74, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x64, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x61, 0x63, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x75, 0x64, 0x54, 0x61, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x22, 0x35, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x54, 0x61,
	0x69, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x74, 0x69,
	0x64, 0x61, 0x64, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x32, 0xa5, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x69, 0x72, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73,
	0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69,
	0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x44,
	0x61, 0x74, 0x6f, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x54, 0x61, 0x69, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x54, 0x61, 0x69, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_primary_node_proto_rawDescOnce sync.Once
	file_primary_node_proto_rawDescData = file_primary_node_proto_rawDesc
)

func file_primary_node_proto_rawDescGZIP() []byte {
	file_primary_node_proto_rawDescOnce.Do(func() {
		file_primary_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_primary_node_proto_rawDescData)
	})
	return file_primary_node_proto_rawDescData
}

var file_primary_node_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_primary_node_proto_goTypes = []any{
	(*DatosDigimon)(nil), // 0: primarynode.DatosDigimon
	(*SolicitudTai)(nil), // 1: primarynode.SolicitudTai
	(*RespuestaTai)(nil), // 2: primarynode.RespuestaTai
	(*Confirmacion)(nil), // 3: primarynode.Confirmacion
}
var file_primary_node_proto_depIdxs = []int32{
	0, // 0: primarynode.PrimaryNodeService.TransferirDatos:input_type -> primarynode.DatosDigimon
	1, // 1: primarynode.PrimaryNodeService.SolicitarDatos:input_type -> primarynode.SolicitudTai
	3, // 2: primarynode.PrimaryNodeService.TransferirDatos:output_type -> primarynode.Confirmacion
	2, // 3: primarynode.PrimaryNodeService.SolicitarDatos:output_type -> primarynode.RespuestaTai
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_primary_node_proto_init() }
func file_primary_node_proto_init() {
	if File_primary_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_primary_node_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DatosDigimon); i {
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
		file_primary_node_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SolicitudTai); i {
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
		file_primary_node_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RespuestaTai); i {
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
		file_primary_node_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Confirmacion); i {
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
			RawDescriptor: file_primary_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_primary_node_proto_goTypes,
		DependencyIndexes: file_primary_node_proto_depIdxs,
		MessageInfos:      file_primary_node_proto_msgTypes,
	}.Build()
	File_primary_node_proto = out.File
	file_primary_node_proto_rawDesc = nil
	file_primary_node_proto_goTypes = nil
	file_primary_node_proto_depIdxs = nil
}