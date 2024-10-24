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

type DatosDigimon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre      string `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Atributo    string `protobuf:"bytes,3,opt,name=atributo,proto3" json:"atributo,omitempty"`
	Sacrificado bool   `protobuf:"varint,4,opt,name=sacrificado,proto3" json:"sacrificado,omitempty"`
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

type Confirmacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string `protobuf:"bytes,1,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *Confirmacion) Reset() {
	*x = Confirmacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_node_1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmacion) ProtoMessage() {}

func (x *Confirmacion) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Confirmacion.ProtoReflect.Descriptor instead.
func (*Confirmacion) Descriptor() ([]byte, []int) {
	return file_data_node_1_proto_rawDescGZIP(), []int{1}
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
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x74, 0x0a,
	0x0c, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x64, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x32, 0x51, 0x0a,
	0x0f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x44, 0x61, 0x74, 0x6f, 0x73,
	0x12, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x6f,
	0x73, 0x44, 0x69, 0x67, 0x69, 0x6d, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_data_node_1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_data_node_1_proto_goTypes = []any{
	(*DatosDigimon)(nil), // 0: datanode.DatosDigimon
	(*Confirmacion)(nil), // 1: datanode.Confirmacion
}
var file_data_node_1_proto_depIdxs = []int32{
	0, // 0: datanode.DataNodeService.GuardarDatos:input_type -> datanode.DatosDigimon
	1, // 1: datanode.DataNodeService.GuardarDatos:output_type -> datanode.Confirmacion
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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
			NumMessages:   2,
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
