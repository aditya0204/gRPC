// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.5.1
// source: numbers.proto

package numberspb

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

type NumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"varint,1,opt,name=Num1,proto3" json:"Num1,omitempty"`
	Num2 int32 `protobuf:"varint,2,opt,name=Num2,proto3" json:"Num2,omitempty"`
}

func (x *NumRequest) Reset() {
	*x = NumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numbers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumRequest) ProtoMessage() {}

func (x *NumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_numbers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumRequest.ProtoReflect.Descriptor instead.
func (*NumRequest) Descriptor() ([]byte, []int) {
	return file_numbers_proto_rawDescGZIP(), []int{0}
}

func (x *NumRequest) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *NumRequest) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type NumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=Sum,proto3" json:"Sum,omitempty"`
}

func (x *NumResponse) Reset() {
	*x = NumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numbers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumResponse) ProtoMessage() {}

func (x *NumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_numbers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumResponse.ProtoReflect.Descriptor instead.
func (*NumResponse) Descriptor() ([]byte, []int) {
	return file_numbers_proto_rawDescGZIP(), []int{1}
}

func (x *NumResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type PrimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *PrimeRequest) Reset() {
	*x = PrimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numbers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeRequest) ProtoMessage() {}

func (x *PrimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_numbers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeRequest.ProtoReflect.Descriptor instead.
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return file_numbers_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type PrimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *PrimeResponse) Reset() {
	*x = PrimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_numbers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeResponse) ProtoMessage() {}

func (x *PrimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_numbers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeResponse.ProtoReflect.Descriptor instead.
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return file_numbers_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeResponse) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_numbers_proto protoreflect.FileDescriptor

var file_numbers_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75,
	0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4e, 0x75, 0x6d, 0x32, 0x22, 0x1f,
	0x0a, 0x0b, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53, 0x75, 0x6d, 0x22,
	0x20, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4e, 0x75,
	0x6d, 0x22, 0x21, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x4e, 0x75, 0x6d, 0x32, 0x80, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x13,
	0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x4e, 0x75,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_numbers_proto_rawDescOnce sync.Once
	file_numbers_proto_rawDescData = file_numbers_proto_rawDesc
)

func file_numbers_proto_rawDescGZIP() []byte {
	file_numbers_proto_rawDescOnce.Do(func() {
		file_numbers_proto_rawDescData = protoimpl.X.CompressGZIP(file_numbers_proto_rawDescData)
	})
	return file_numbers_proto_rawDescData
}

var file_numbers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_numbers_proto_goTypes = []interface{}{
	(*NumRequest)(nil),    // 0: numbers.NumRequest
	(*NumResponse)(nil),   // 1: numbers.NumResponse
	(*PrimeRequest)(nil),  // 2: numbers.PrimeRequest
	(*PrimeResponse)(nil), // 3: numbers.PrimeResponse
}
var file_numbers_proto_depIdxs = []int32{
	0, // 0: numbers.SumService.numbers:input_type -> numbers.NumRequest
	2, // 1: numbers.SumService.prime:input_type -> numbers.PrimeRequest
	1, // 2: numbers.SumService.numbers:output_type -> numbers.NumResponse
	3, // 3: numbers.SumService.prime:output_type -> numbers.PrimeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_numbers_proto_init() }
func file_numbers_proto_init() {
	if File_numbers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_numbers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumRequest); i {
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
		file_numbers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumResponse); i {
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
		file_numbers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeRequest); i {
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
		file_numbers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeResponse); i {
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
			RawDescriptor: file_numbers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_numbers_proto_goTypes,
		DependencyIndexes: file_numbers_proto_depIdxs,
		MessageInfos:      file_numbers_proto_msgTypes,
	}.Build()
	File_numbers_proto = out.File
	file_numbers_proto_rawDesc = nil
	file_numbers_proto_goTypes = nil
	file_numbers_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	Numbers(ctx context.Context, in *NumRequest, opts ...grpc.CallOption) (*NumResponse, error)
	Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (SumService_PrimeClient, error)
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Numbers(ctx context.Context, in *NumRequest, opts ...grpc.CallOption) (*NumResponse, error) {
	out := new(NumResponse)
	err := c.cc.Invoke(ctx, "/numbers.SumService/numbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (SumService_PrimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[0], "/numbers.SumService/prime", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServicePrimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_PrimeClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type sumServicePrimeClient struct {
	grpc.ClientStream
}

func (x *sumServicePrimeClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	Numbers(context.Context, *NumRequest) (*NumResponse, error)
	Prime(*PrimeRequest, SumService_PrimeServer) error
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Numbers(context.Context, *NumRequest) (*NumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Numbers not implemented")
}
func (*UnimplementedSumServiceServer) Prime(*PrimeRequest, SumService_PrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method Prime not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Numbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Numbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/numbers.SumService/Numbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Numbers(ctx, req.(*NumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Prime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).Prime(m, &sumServicePrimeServer{stream})
}

type SumService_PrimeServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type sumServicePrimeServer struct {
	grpc.ServerStream
}

func (x *sumServicePrimeServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "numbers.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "numbers",
			Handler:    _SumService_Numbers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "prime",
			Handler:       _SumService_Prime_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "numbers.proto",
}