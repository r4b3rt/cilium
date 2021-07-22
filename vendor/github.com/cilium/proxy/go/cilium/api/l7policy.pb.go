// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: cilium/api/l7policy.proto

package cilium

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type L7Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the unix domain socket for the cilium access log.
	AccessLogPath string `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath,proto3" json:"access_log_path,omitempty"`
	// Cilium endpoint security policy to enforce.
	PolicyName string `protobuf:"bytes,2,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// HTTP response body message for 403 status code.
	// If empty, "Access denied" will be used.
	Denied_403Body string `protobuf:"bytes,3,opt,name=denied_403_body,json=denied403Body,proto3" json:"denied_403_body,omitempty"`
	// 'true' if the filter is on ingress listener, 'false' for egress listener.
	// Value from the listener filter will be used if not specified here.
	IsIngress *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=is_ingress,json=isIngress,proto3" json:"is_ingress,omitempty"`
}

func (x *L7Policy) Reset() {
	*x = L7Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_l7policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L7Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L7Policy) ProtoMessage() {}

func (x *L7Policy) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_l7policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L7Policy.ProtoReflect.Descriptor instead.
func (*L7Policy) Descriptor() ([]byte, []int) {
	return file_cilium_api_l7policy_proto_rawDescGZIP(), []int{0}
}

func (x *L7Policy) GetAccessLogPath() string {
	if x != nil {
		return x.AccessLogPath
	}
	return ""
}

func (x *L7Policy) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *L7Policy) GetDenied_403Body() string {
	if x != nil {
		return x.Denied_403Body
	}
	return ""
}

func (x *L7Policy) GetIsIngress() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsIngress
	}
	return nil
}

var File_cilium_api_l7policy_proto protoreflect.FileDescriptor

var file_cilium_api_l7policy_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x37, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69, 0x6c,
	0x69, 0x75, 0x6d, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x4c, 0x37, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x26, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x6e,
	0x69, 0x65, 0x64, 0x5f, 0x34, 0x30, 0x33, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x34, 0x30, 0x33, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x39, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x69, 0x73, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cilium_api_l7policy_proto_rawDescOnce sync.Once
	file_cilium_api_l7policy_proto_rawDescData = file_cilium_api_l7policy_proto_rawDesc
)

func file_cilium_api_l7policy_proto_rawDescGZIP() []byte {
	file_cilium_api_l7policy_proto_rawDescOnce.Do(func() {
		file_cilium_api_l7policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_cilium_api_l7policy_proto_rawDescData)
	})
	return file_cilium_api_l7policy_proto_rawDescData
}

var file_cilium_api_l7policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cilium_api_l7policy_proto_goTypes = []interface{}{
	(*L7Policy)(nil),             // 0: cilium.L7Policy
	(*wrapperspb.BoolValue)(nil), // 1: google.protobuf.BoolValue
}
var file_cilium_api_l7policy_proto_depIdxs = []int32{
	1, // 0: cilium.L7Policy.is_ingress:type_name -> google.protobuf.BoolValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cilium_api_l7policy_proto_init() }
func file_cilium_api_l7policy_proto_init() {
	if File_cilium_api_l7policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cilium_api_l7policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L7Policy); i {
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
			RawDescriptor: file_cilium_api_l7policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cilium_api_l7policy_proto_goTypes,
		DependencyIndexes: file_cilium_api_l7policy_proto_depIdxs,
		MessageInfos:      file_cilium_api_l7policy_proto_msgTypes,
	}.Build()
	File_cilium_api_l7policy_proto = out.File
	file_cilium_api_l7policy_proto_rawDesc = nil
	file_cilium_api_l7policy_proto_goTypes = nil
	file_cilium_api_l7policy_proto_depIdxs = nil
}
