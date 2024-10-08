// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: data_node_1.proto

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

// Mensaje de datos del Digimon
type DatosDigimon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`            // ID único del Digimon
	Atributo string `protobuf:"bytes,2,opt,name=atributo,proto3" json:"atributo,omitempty"` // Atributo (Data, Vaccine, Virus)
}

func (x *DatosDigimon) Reset() {
	*x = DatosDigimon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_node_1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatosDigimon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatosDigimon) ProtoMessage() {}

func (x *DatosDigimon) ProtoReflect() protoreflect.Message {
	mi := &file_data_node_1_proto_msgTypes[0]
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
	return file_data_node_1_proto_rawDescGZIP(), []int{0}
}

func (x *DatosDigimon) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DatosDigimon) GetAtributo() string {
	if x != nil {
		return x.Atributo
	}
	return ""
}

// Mensaje de solicitud de atributo por ID
type SolicitudAtributo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // ID del Digimon
}

func (x *SolicitudAtributo) Reset() {
	*x = SolicitudAtributo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_node_1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudAtributo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudAtributo) ProtoMessage() {}

func (x *SolicitudAtributo) ProtoReflect() protoreflect.Message {
	mi := &file_data_node_1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudAtributo.ProtoReflect.Descriptor instead.
func (*SolicitudAtributo) Descriptor() ([]byte, []int) {
	return file_data_node_1_proto_rawDescGZIP(), []int{1}
}

func (x *SolicitudAtributo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Mensaje que contiene el atributo de un Digimon
type AtributoDigimon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Atributo string `protobuf:"bytes,1,opt,name=atributo,proto3" json:"atributo,omitempty"` // Atributo (Data, Vaccine, Virus)
}

func (x *AtributoDigimon) Reset() {
	*x = AtributoDigimon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_node_1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtributoDigimon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtributoDigimon) ProtoMessage() {}

func (x *AtributoDigimon) ProtoReflect() protoreflect.Message {
	mi := &file_data_node_1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtributoDigimon.ProtoReflect.Descriptor instead.
func (*AtributoDigimon) Descriptor() ([]byte, []int) {
	return file_data_node_1_proto_rawDescGZIP(), []int{2}
}

func (x *AtributoDigimon) GetAtributo() string {
	if x != nil {
		return x.Atributo
	}
	return ""
}

// Mensaje de confirmación para recibir datos
type Confirmacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"` // Mensaje de confirmación
}

func (x *Confirmacion) Reset() {
	*x = Confirmacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_node_1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmacion) ProtoMessage() {}

func (x *Confirmacion) ProtoReflect() protoreflect.Message {
	mi := &file_data_node_1_proto_msgTypes[3]
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
	return file_data_node_1_proto_rawDescGZIP(), []int{3}
}

func (x *Confirmacion) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

var File_data_node_1_proto protoreflect.FileDescriptor

var file_data_node_1_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a,
	0x0c, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x41, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d,
	0x0a, 0x0f, 0x41, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x22, 0x28, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x32, 0xa0, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61,
	0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x41,
	0x6c, 0x6d, 0x61, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x69,
	0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a,
	0x11, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x41, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x6f, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x41, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x1a,
	0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x6f, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_node_1_proto_rawDescOnce sync.Once
	file_data_node_1_proto_rawDescData = file_data_node_1_proto_rawDesc
)

func file_data_node_1_proto_rawDescGZIP() []byte {
	file_data_node_1_proto_rawDescOnce.Do(func() {
		file_data_node_1_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_node_1_proto_rawDescData)
	})
	return file_data_node_1_proto_rawDescData
}

var file_data_node_1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_node_1_proto_goTypes = []any{
	(*DatosDigimon)(nil),      // 0: datanode.DatosDigimon
	(*SolicitudAtributo)(nil), // 1: datanode.SolicitudAtributo
	(*AtributoDigimon)(nil),   // 2: datanode.AtributoDigimon
	(*Confirmacion)(nil),      // 3: datanode.Confirmacion
}
var file_data_node_1_proto_depIdxs = []int32{
	0, // 0: datanode.DataNodeService.AlmacenarDatos:input_type -> datanode.DatosDigimon
	1, // 1: datanode.DataNodeService.SolicitarAtributo:input_type -> datanode.SolicitudAtributo
	3, // 2: datanode.DataNodeService.AlmacenarDatos:output_type -> datanode.Confirmacion
	2, // 3: datanode.DataNodeService.SolicitarAtributo:output_type -> datanode.AtributoDigimon
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_node_1_proto_init() }
func file_data_node_1_proto_init() {
	if File_data_node_1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_node_1_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_data_node_1_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SolicitudAtributo); i {
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
		file_data_node_1_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AtributoDigimon); i {
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
		file_data_node_1_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
			RawDescriptor: file_data_node_1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_node_1_proto_goTypes,
		DependencyIndexes: file_data_node_1_proto_depIdxs,
		MessageInfos:      file_data_node_1_proto_msgTypes,
	}.Build()
	File_data_node_1_proto = out.File
	file_data_node_1_proto_rawDesc = nil
	file_data_node_1_proto_goTypes = nil
	file_data_node_1_proto_depIdxs = nil
}