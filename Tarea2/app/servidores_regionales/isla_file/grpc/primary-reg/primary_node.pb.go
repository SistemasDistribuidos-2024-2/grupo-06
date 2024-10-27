// primary_node.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: primary_node.proto

package primary_reg

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

// Mensaje para enviar datos cifrados de los Digimons desde los servidores regionales
type DatosCifradosDigimon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreCifrado   string `protobuf:"bytes,1,opt,name=nombre_cifrado,json=nombreCifrado,proto3" json:"nombre_cifrado,omitempty"`       // Nombre cifrado del Digimon
	AtributoCifrado string `protobuf:"bytes,2,opt,name=atributo_cifrado,json=atributoCifrado,proto3" json:"atributo_cifrado,omitempty"` // Atributo cifrado del Digimon (Data, Vaccine, Virus)
	EstadoCifrado   string `protobuf:"bytes,3,opt,name=estado_cifrado,json=estadoCifrado,proto3" json:"estado_cifrado,omitempty"`       // Estado cifrado (Sacrificado o No-Sacrificado)
}

func (x *DatosCifradosDigimon) Reset() {
	*x = DatosCifradosDigimon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primary_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatosCifradosDigimon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatosCifradosDigimon) ProtoMessage() {}

func (x *DatosCifradosDigimon) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DatosCifradosDigimon.ProtoReflect.Descriptor instead.
func (*DatosCifradosDigimon) Descriptor() ([]byte, []int) {
	return file_primary_node_proto_rawDescGZIP(), []int{0}
}

func (x *DatosCifradosDigimon) GetNombreCifrado() string {
	if x != nil {
		return x.NombreCifrado
	}
	return ""
}

func (x *DatosCifradosDigimon) GetAtributoCifrado() string {
	if x != nil {
		return x.AtributoCifrado
	}
	return ""
}

func (x *DatosCifradosDigimon) GetEstadoCifrado() string {
	if x != nil {
		return x.EstadoCifrado
	}
	return ""
}

// Mensaje de confirmación de recepción para los servidores regionales
type Confirmacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"` // Mensaje de confirmación del Primary Node
}

func (x *Confirmacion) Reset() {
	*x = Confirmacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primary_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmacion) ProtoMessage() {}

func (x *Confirmacion) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Confirmacion.ProtoReflect.Descriptor instead.
func (*Confirmacion) Descriptor() ([]byte, []int) {
	return file_primary_node_proto_rawDescGZIP(), []int{1}
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
	0x65, 0x22, 0x8f, 0x01, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x43, 0x69, 0x66, 0x72, 0x61,
	0x64, 0x6f, 0x73, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x66, 0x72, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x69, 0x66, 0x72, 0x61, 0x64,
	0x6f, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x69,
	0x66, 0x72, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x6f, 0x43, 0x69, 0x66, 0x72, 0x61, 0x64, 0x6f, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x5f, 0x63, 0x69, 0x66, 0x72, 0x61, 0x64, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x43, 0x69, 0x66, 0x72,
	0x61, 0x64, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x32, 0x6a, 0x0a,
	0x12, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x69, 0x62, 0x69, 0x72, 0x44, 0x61,
	0x74, 0x6f, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x2e, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x43,
	0x69, 0x66, 0x72, 0x61, 0x64, 0x6f, 0x73, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x2d, 0x72, 0x65, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_primary_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_primary_node_proto_goTypes = []any{
	(*DatosCifradosDigimon)(nil), // 0: primarynode.DatosCifradosDigimon
	(*Confirmacion)(nil),         // 1: primarynode.Confirmacion
}
var file_primary_node_proto_depIdxs = []int32{
	0, // 0: primarynode.PrimaryNodeService.RecibirDatosRegional:input_type -> primarynode.DatosCifradosDigimon
	1, // 1: primarynode.PrimaryNodeService.RecibirDatosRegional:output_type -> primarynode.Confirmacion
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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
			switch v := v.(*DatosCifradosDigimon); i {
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
			NumMessages:   2,
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