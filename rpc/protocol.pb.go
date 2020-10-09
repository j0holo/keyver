// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: rpc/protocol.proto

package rpc

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

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_rpc_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_rpc_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_rpc_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_rpc_protocol_proto protoreflect.FileDescriptor

var file_rpc_protocol_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x34,
	0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x6a, 0x0a, 0x02, 0x44, 0x42, 0x12, 0x31,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x30, 0x68, 0x6f, 0x6c, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x76, 0x65, 0x72,
	0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_protocol_proto_rawDescOnce sync.Once
	file_rpc_protocol_proto_rawDescData = file_rpc_protocol_proto_rawDesc
)

func file_rpc_protocol_proto_rawDescGZIP() []byte {
	file_rpc_protocol_proto_rawDescOnce.Do(func() {
		file_rpc_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_protocol_proto_rawDescData)
	})
	return file_rpc_protocol_proto_rawDescData
}

var file_rpc_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_protocol_proto_goTypes = []interface{}{
	(*SetRequest)(nil), // 0: protocol.SetRequest
	(*GetRequest)(nil), // 1: protocol.GetRequest
	(*Response)(nil),   // 2: protocol.Response
}
var file_rpc_protocol_proto_depIdxs = []int32{
	1, // 0: protocol.DB.Get:input_type -> protocol.GetRequest
	0, // 1: protocol.DB.Set:input_type -> protocol.SetRequest
	2, // 2: protocol.DB.Get:output_type -> protocol.Response
	2, // 3: protocol.DB.Set:output_type -> protocol.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_protocol_proto_init() }
func file_rpc_protocol_proto_init() {
	if File_rpc_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_rpc_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_rpc_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_rpc_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_protocol_proto_goTypes,
		DependencyIndexes: file_rpc_protocol_proto_depIdxs,
		MessageInfos:      file_rpc_protocol_proto_msgTypes,
	}.Build()
	File_rpc_protocol_proto = out.File
	file_rpc_protocol_proto_rawDesc = nil
	file_rpc_protocol_proto_goTypes = nil
	file_rpc_protocol_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DBClient is the client API for DB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DBClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Response, error)
}

type dBClient struct {
	cc grpc.ClientConnInterface
}

func NewDBClient(cc grpc.ClientConnInterface) DBClient {
	return &dBClient{cc}
}

func (c *dBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.DB/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protocol.DB/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServer is the server API for DB service.
type DBServer interface {
	Get(context.Context, *GetRequest) (*Response, error)
	Set(context.Context, *SetRequest) (*Response, error)
}

// UnimplementedDBServer can be embedded to have forward compatible implementations.
type UnimplementedDBServer struct {
}

func (*UnimplementedDBServer) Get(context.Context, *GetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDBServer) Set(context.Context, *SetRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}

func RegisterDBServer(s *grpc.Server, srv DBServer) {
	s.RegisterService(&_DB_serviceDesc, srv)
}

func _DB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.DB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.DB/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.DB",
	HandlerType: (*DBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DB_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _DB_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/protocol.proto",
}
