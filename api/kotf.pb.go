// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: kotf.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Output  string `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kotf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_kotf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_kotf_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Result) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Result) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type TerraformInitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CloudRegion string `protobuf:"bytes,3,opt,name=cloudRegion,proto3" json:"cloudRegion,omitempty"`
	Hosts       string `protobuf:"bytes,4,opt,name=hosts,proto3" json:"hosts,omitempty"`
	Provider    string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *TerraformInitRequest) Reset() {
	*x = TerraformInitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kotf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformInitRequest) ProtoMessage() {}

func (x *TerraformInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kotf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformInitRequest.ProtoReflect.Descriptor instead.
func (*TerraformInitRequest) Descriptor() ([]byte, []int) {
	return file_kotf_proto_rawDescGZIP(), []int{1}
}

func (x *TerraformInitRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *TerraformInitRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TerraformInitRequest) GetCloudRegion() string {
	if x != nil {
		return x.CloudRegion
	}
	return ""
}

func (x *TerraformInitRequest) GetHosts() string {
	if x != nil {
		return x.Hosts
	}
	return ""
}

func (x *TerraformInitRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type TerraformApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CloudRegion string `protobuf:"bytes,3,opt,name=cloudRegion,proto3" json:"cloudRegion,omitempty"`
}

func (x *TerraformApplyRequest) Reset() {
	*x = TerraformApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kotf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformApplyRequest) ProtoMessage() {}

func (x *TerraformApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kotf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformApplyRequest.ProtoReflect.Descriptor instead.
func (*TerraformApplyRequest) Descriptor() ([]byte, []int) {
	return file_kotf_proto_rawDescGZIP(), []int{2}
}

func (x *TerraformApplyRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *TerraformApplyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TerraformApplyRequest) GetCloudRegion() string {
	if x != nil {
		return x.CloudRegion
	}
	return ""
}

type TerraformDestroyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CloudRegion string `protobuf:"bytes,3,opt,name=cloudRegion,proto3" json:"cloudRegion,omitempty"`
}

func (x *TerraformDestroyRequest) Reset() {
	*x = TerraformDestroyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kotf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformDestroyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformDestroyRequest) ProtoMessage() {}

func (x *TerraformDestroyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kotf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformDestroyRequest.ProtoReflect.Descriptor instead.
func (*TerraformDestroyRequest) Descriptor() ([]byte, []int) {
	return file_kotf_proto_rawDescGZIP(), []int{3}
}

func (x *TerraformDestroyRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *TerraformDestroyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TerraformDestroyRequest) GetCloudRegion() string {
	if x != nil {
		return x.CloudRegion
	}
	return ""
}

var File_kotf_proto protoreflect.FileDescriptor

var file_kotf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6b, 0x6f, 0x74, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x22, 0x4c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0xa0, 0x01, 0x0a, 0x14, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0x6f, 0x0a, 0x15, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x22, 0x71, 0x0a, 0x17, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x32, 0xa7, 0x01, 0x0a, 0x07, 0x4b, 0x6f, 0x74, 0x66, 0x41,
	0x70, 0x69, 0x12, 0x30, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74,
	0x72, 0x6f, 0x79, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kotf_proto_rawDescOnce sync.Once
	file_kotf_proto_rawDescData = file_kotf_proto_rawDesc
)

func file_kotf_proto_rawDescGZIP() []byte {
	file_kotf_proto_rawDescOnce.Do(func() {
		file_kotf_proto_rawDescData = protoimpl.X.CompressGZIP(file_kotf_proto_rawDescData)
	})
	return file_kotf_proto_rawDescData
}

var file_kotf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kotf_proto_goTypes = []interface{}{
	(*Result)(nil),                  // 0: api.Result
	(*TerraformInitRequest)(nil),    // 1: api.TerraformInitRequest
	(*TerraformApplyRequest)(nil),   // 2: api.TerraformApplyRequest
	(*TerraformDestroyRequest)(nil), // 3: api.TerraformDestroyRequest
}
var file_kotf_proto_depIdxs = []int32{
	1, // 0: api.KotfApi.Init:input_type -> api.TerraformInitRequest
	2, // 1: api.KotfApi.Apply:input_type -> api.TerraformApplyRequest
	3, // 2: api.KotfApi.Destroy:input_type -> api.TerraformDestroyRequest
	0, // 3: api.KotfApi.Init:output_type -> api.Result
	0, // 4: api.KotfApi.Apply:output_type -> api.Result
	0, // 5: api.KotfApi.Destroy:output_type -> api.Result
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kotf_proto_init() }
func file_kotf_proto_init() {
	if File_kotf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kotf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_kotf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformInitRequest); i {
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
		file_kotf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformApplyRequest); i {
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
		file_kotf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformDestroyRequest); i {
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
			RawDescriptor: file_kotf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kotf_proto_goTypes,
		DependencyIndexes: file_kotf_proto_depIdxs,
		MessageInfos:      file_kotf_proto_msgTypes,
	}.Build()
	File_kotf_proto = out.File
	file_kotf_proto_rawDesc = nil
	file_kotf_proto_goTypes = nil
	file_kotf_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KotfApiClient is the client API for KotfApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KotfApiClient interface {
	Init(ctx context.Context, in *TerraformInitRequest, opts ...grpc.CallOption) (*Result, error)
	Apply(ctx context.Context, in *TerraformApplyRequest, opts ...grpc.CallOption) (*Result, error)
	Destroy(ctx context.Context, in *TerraformDestroyRequest, opts ...grpc.CallOption) (*Result, error)
}

type kotfApiClient struct {
	cc grpc.ClientConnInterface
}

func NewKotfApiClient(cc grpc.ClientConnInterface) KotfApiClient {
	return &kotfApiClient{cc}
}

func (c *kotfApiClient) Init(ctx context.Context, in *TerraformInitRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/api.KotfApi/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kotfApiClient) Apply(ctx context.Context, in *TerraformApplyRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/api.KotfApi/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kotfApiClient) Destroy(ctx context.Context, in *TerraformDestroyRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/api.KotfApi/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KotfApiServer is the server API for KotfApi service.
type KotfApiServer interface {
	Init(context.Context, *TerraformInitRequest) (*Result, error)
	Apply(context.Context, *TerraformApplyRequest) (*Result, error)
	Destroy(context.Context, *TerraformDestroyRequest) (*Result, error)
}

// UnimplementedKotfApiServer can be embedded to have forward compatible implementations.
type UnimplementedKotfApiServer struct {
}

func (*UnimplementedKotfApiServer) Init(context.Context, *TerraformInitRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedKotfApiServer) Apply(context.Context, *TerraformApplyRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (*UnimplementedKotfApiServer) Destroy(context.Context, *TerraformDestroyRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}

func RegisterKotfApiServer(s *grpc.Server, srv KotfApiServer) {
	s.RegisterService(&_KotfApi_serviceDesc, srv)
}

func _KotfApi_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerraformInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotfApiServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KotfApi/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotfApiServer).Init(ctx, req.(*TerraformInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KotfApi_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerraformApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotfApiServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KotfApi/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotfApiServer).Apply(ctx, req.(*TerraformApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KotfApi_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerraformDestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotfApiServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KotfApi/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotfApiServer).Destroy(ctx, req.(*TerraformDestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KotfApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.KotfApi",
	HandlerType: (*KotfApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _KotfApi_Init_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _KotfApi_Apply_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _KotfApi_Destroy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kotf.proto",
}
