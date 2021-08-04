// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.2
// source: protos/svca/service.proto

package svca

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

type GetServiceARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetServiceARequest) Reset() {
	*x = GetServiceARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_svca_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceARequest) ProtoMessage() {}

func (x *GetServiceARequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_svca_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceARequest.ProtoReflect.Descriptor instead.
func (*GetServiceARequest) Descriptor() ([]byte, []int) {
	return file_protos_svca_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetServiceARequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetServiceAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*ServiceAResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *GetServiceAResponse) Reset() {
	*x = GetServiceAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_svca_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceAResponse) ProtoMessage() {}

func (x *GetServiceAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_svca_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceAResponse.ProtoReflect.Descriptor instead.
func (*GetServiceAResponse) Descriptor() ([]byte, []int) {
	return file_protos_svca_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetServiceAResponse) GetResponses() []*ServiceAResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

type ServiceAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Status      string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ServiceAResponse) Reset() {
	*x = ServiceAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_svca_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAResponse) ProtoMessage() {}

func (x *ServiceAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_svca_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAResponse.ProtoReflect.Descriptor instead.
func (*ServiceAResponse) Descriptor() ([]byte, []int) {
	return file_protos_svca_service_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceAResponse) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ServiceAResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_protos_svca_service_proto protoreflect.FileDescriptor

var file_protos_svca_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x76, 0x63, 0x61, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x70, 0x63,
	0x2e, 0x73, 0x76, 0x63, 0x61, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x76, 0x63, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x22, 0x4d, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x57, 0x0a, 0x08, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x12, 0x1c, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x76, 0x63, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x76, 0x63, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x76, 0x63, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_svca_service_proto_rawDescOnce sync.Once
	file_protos_svca_service_proto_rawDescData = file_protos_svca_service_proto_rawDesc
)

func file_protos_svca_service_proto_rawDescGZIP() []byte {
	file_protos_svca_service_proto_rawDescOnce.Do(func() {
		file_protos_svca_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_svca_service_proto_rawDescData)
	})
	return file_protos_svca_service_proto_rawDescData
}

var file_protos_svca_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_svca_service_proto_goTypes = []interface{}{
	(*GetServiceARequest)(nil),  // 0: rpc.svca.GetServiceARequest
	(*GetServiceAResponse)(nil), // 1: rpc.svca.GetServiceAResponse
	(*ServiceAResponse)(nil),    // 2: rpc.svca.ServiceAResponse
}
var file_protos_svca_service_proto_depIdxs = []int32{
	2, // 0: rpc.svca.GetServiceAResponse.responses:type_name -> rpc.svca.ServiceAResponse
	0, // 1: rpc.svca.AService.CallServiceA:input_type -> rpc.svca.GetServiceARequest
	1, // 2: rpc.svca.AService.CallServiceA:output_type -> rpc.svca.GetServiceAResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_svca_service_proto_init() }
func file_protos_svca_service_proto_init() {
	if File_protos_svca_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_svca_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceARequest); i {
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
		file_protos_svca_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceAResponse); i {
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
		file_protos_svca_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAResponse); i {
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
			RawDescriptor: file_protos_svca_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_svca_service_proto_goTypes,
		DependencyIndexes: file_protos_svca_service_proto_depIdxs,
		MessageInfos:      file_protos_svca_service_proto_msgTypes,
	}.Build()
	File_protos_svca_service_proto = out.File
	file_protos_svca_service_proto_rawDesc = nil
	file_protos_svca_service_proto_goTypes = nil
	file_protos_svca_service_proto_depIdxs = nil
}