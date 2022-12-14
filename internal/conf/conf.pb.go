// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: internal/conf/conf.proto

package conf

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	common "gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
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

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *common.Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Jwks     *common.Jwks   `protobuf:"bytes,2,opt,name=jwks,proto3" json:"jwks,omitempty"`
	Trace    *common.Trace  `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
	Sentry   *common.Sentry `protobuf:"bytes,4,opt,name=sentry,proto3" json:"sentry,omitempty"`
	Data     *Data          `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Business *Business      `protobuf:"bytes,6,opt,name=business,proto3" json:"business,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *common.Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetJwks() *common.Jwks {
	if x != nil {
		return x.Jwks
	}
	return nil
}

func (x *Bootstrap) GetTrace() *common.Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Bootstrap) GetSentry() *common.Sentry {
	if x != nil {
		return x.Sentry
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetBusiness() *Business {
	if x != nil {
		return x.Business
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relational *common.Relational `protobuf:"bytes,1,opt,name=relational,proto3" json:"relational,omitempty"`
	Redis      *common.Redis      `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetRelational() *common.Relational {
	if x != nil {
		return x.Relational
	}
	return nil
}

func (x *Data) GetRedis() *common.Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublishDuration *duration.Duration `protobuf:"bytes,1,opt,name=publish_duration,json=publishDuration,proto3" json:"publish_duration,omitempty"`
	ShortenLimit    int32              `protobuf:"varint,2,opt,name=shorten_limit,json=shortenLimit,proto3" json:"shorten_limit,omitempty"`
}

func (x *Business) Reset() {
	*x = Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business) ProtoMessage() {}

func (x *Business) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Business.ProtoReflect.Descriptor instead.
func (*Business) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Business) GetPublishDuration() *duration.Duration {
	if x != nil {
		return x.PublishDuration
	}
	return nil
}

func (x *Business) GetShortenLimit() int32 {
	if x != nil {
		return x.ShortenLimit
	}
	return 0
}

var File_internal_conf_conf_proto protoreflect.FileDescriptor

var file_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a,
	0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4a, 0x77, 0x6b, 0x73, 0x52, 0x04,
	0x6a, 0x77, 0x6b, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x5f, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x32, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x22, 0x6c, 0x0a, 0x08, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x63, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x2d, 0x70, 0x73, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_conf_conf_proto_rawDescOnce sync.Once
	file_internal_conf_conf_proto_rawDescData = file_internal_conf_conf_proto_rawDesc
)

func file_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_conf_conf_proto_rawDescData)
	})
	return file_internal_conf_conf_proto_rawDescData
}

var file_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),         // 0: kratos.api.Bootstrap
	(*Data)(nil),              // 1: kratos.api.Data
	(*Business)(nil),          // 2: kratos.api.Business
	(*common.Server)(nil),     // 3: common.Server
	(*common.Jwks)(nil),       // 4: common.Jwks
	(*common.Trace)(nil),      // 5: common.Trace
	(*common.Sentry)(nil),     // 6: common.Sentry
	(*common.Relational)(nil), // 7: common.Relational
	(*common.Redis)(nil),      // 8: common.Redis
	(*duration.Duration)(nil), // 9: common.Duration
}
var file_internal_conf_conf_proto_depIdxs = []int32{
	3, // 0: kratos.api.Bootstrap.server:type_name -> common.Server
	4, // 1: kratos.api.Bootstrap.jwks:type_name -> common.Jwks
	5, // 2: kratos.api.Bootstrap.trace:type_name -> common.Trace
	6, // 3: kratos.api.Bootstrap.sentry:type_name -> common.Sentry
	1, // 4: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	2, // 5: kratos.api.Bootstrap.business:type_name -> kratos.api.Business
	7, // 6: kratos.api.Data.relational:type_name -> common.Relational
	8, // 7: kratos.api.Data.redis:type_name -> common.Redis
	9, // 8: kratos.api.Business.publish_duration:type_name -> common.Duration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_internal_conf_conf_proto_init() }
func file_internal_conf_conf_proto_init() {
	if File_internal_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_internal_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_internal_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Business); i {
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
			RawDescriptor: file_internal_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_internal_conf_conf_proto_depIdxs,
		MessageInfos:      file_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_internal_conf_conf_proto = out.File
	file_internal_conf_conf_proto_rawDesc = nil
	file_internal_conf_conf_proto_goTypes = nil
	file_internal_conf_conf_proto_depIdxs = nil
}
