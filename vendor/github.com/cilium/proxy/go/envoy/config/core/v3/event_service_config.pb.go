// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: envoy/config/core/v3/event_service_config.proto

package envoy_config_core_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [#not-implemented-hide:]
// Configuration of the event reporting service endpoint.
type EventServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigSourceSpecifier:
	//	*EventServiceConfig_GrpcService
	ConfigSourceSpecifier isEventServiceConfig_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
}

func (x *EventServiceConfig) Reset() {
	*x = EventServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_event_service_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventServiceConfig) ProtoMessage() {}

func (x *EventServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_event_service_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventServiceConfig.ProtoReflect.Descriptor instead.
func (*EventServiceConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_event_service_config_proto_rawDescGZIP(), []int{0}
}

func (m *EventServiceConfig) GetConfigSourceSpecifier() isEventServiceConfig_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (x *EventServiceConfig) GetGrpcService() *GrpcService {
	if x, ok := x.GetConfigSourceSpecifier().(*EventServiceConfig_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

type isEventServiceConfig_ConfigSourceSpecifier interface {
	isEventServiceConfig_ConfigSourceSpecifier()
}

type EventServiceConfig_GrpcService struct {
	// Specifies the gRPC service that hosts the event reporting service.
	GrpcService *GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

func (*EventServiceConfig_GrpcService) isEventServiceConfig_ConfigSourceSpecifier() {}

var File_envoy_config_core_v3_event_service_config_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_event_service_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x12,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x46, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x67,
	0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x2b, 0x9a, 0xc5, 0x88, 0x1e,
	0x26, 0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x1e, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x47, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x17, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_event_service_config_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_event_service_config_proto_rawDescData = file_envoy_config_core_v3_event_service_config_proto_rawDesc
)

func file_envoy_config_core_v3_event_service_config_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_event_service_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_event_service_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_event_service_config_proto_rawDescData)
	})
	return file_envoy_config_core_v3_event_service_config_proto_rawDescData
}

var file_envoy_config_core_v3_event_service_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_core_v3_event_service_config_proto_goTypes = []interface{}{
	(*EventServiceConfig)(nil), // 0: envoy.config.core.v3.EventServiceConfig
	(*GrpcService)(nil),        // 1: envoy.config.core.v3.GrpcService
}
var file_envoy_config_core_v3_event_service_config_proto_depIdxs = []int32{
	1, // 0: envoy.config.core.v3.EventServiceConfig.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_event_service_config_proto_init() }
func file_envoy_config_core_v3_event_service_config_proto_init() {
	if File_envoy_config_core_v3_event_service_config_proto != nil {
		return
	}
	file_envoy_config_core_v3_grpc_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_event_service_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventServiceConfig); i {
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
	file_envoy_config_core_v3_event_service_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EventServiceConfig_GrpcService)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_core_v3_event_service_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_event_service_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_event_service_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_core_v3_event_service_config_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_event_service_config_proto = out.File
	file_envoy_config_core_v3_event_service_config_proto_rawDesc = nil
	file_envoy_config_core_v3_event_service_config_proto_goTypes = nil
	file_envoy_config_core_v3_event_service_config_proto_depIdxs = nil
}
