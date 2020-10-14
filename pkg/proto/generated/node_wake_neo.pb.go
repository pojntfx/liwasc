// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: node_wake_neo.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type NodeWakeStartNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeWakeTimeout int64 `protobuf:"varint,1,opt,name=NodeWakeTimeout,proto3" json:"NodeWakeTimeout,omitempty"`
	MACAddress      int64 `protobuf:"varint,2,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
}

func (x *NodeWakeStartNeoMessage) Reset() {
	*x = NodeWakeStartNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_wake_neo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeWakeStartNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeWakeStartNeoMessage) ProtoMessage() {}

func (x *NodeWakeStartNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_wake_neo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeWakeStartNeoMessage.ProtoReflect.Descriptor instead.
func (*NodeWakeStartNeoMessage) Descriptor() ([]byte, []int) {
	return file_node_wake_neo_proto_rawDescGZIP(), []int{0}
}

func (x *NodeWakeStartNeoMessage) GetNodeWakeTimeout() int64 {
	if x != nil {
		return x.NodeWakeTimeout
	}
	return 0
}

func (x *NodeWakeStartNeoMessage) GetMACAddress() int64 {
	if x != nil {
		return x.MACAddress
	}
	return 0
}

type NodeWakeNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt  string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Done       bool   `protobuf:"varint,3,opt,name=Done,proto3" json:"Done,omitempty"`
	MACAddress int64  `protobuf:"varint,4,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
	PoweredOne bool   `protobuf:"varint,5,opt,name=PoweredOne,proto3" json:"PoweredOne,omitempty"`
}

func (x *NodeWakeNeoMessage) Reset() {
	*x = NodeWakeNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_wake_neo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeWakeNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeWakeNeoMessage) ProtoMessage() {}

func (x *NodeWakeNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_wake_neo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeWakeNeoMessage.ProtoReflect.Descriptor instead.
func (*NodeWakeNeoMessage) Descriptor() ([]byte, []int) {
	return file_node_wake_neo_proto_rawDescGZIP(), []int{1}
}

func (x *NodeWakeNeoMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NodeWakeNeoMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NodeWakeNeoMessage) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *NodeWakeNeoMessage) GetMACAddress() int64 {
	if x != nil {
		return x.MACAddress
	}
	return 0
}

func (x *NodeWakeNeoMessage) GetPoweredOne() bool {
	if x != nil {
		return x.PoweredOne
	}
	return false
}

var File_node_wake_neo_proto protoreflect.FileDescriptor

var file_node_wake_neo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x77, 0x61, 0x6b, 0x65, 0x5f, 0x6e, 0x65, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73,
	0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x63, 0x0a, 0x17, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x41, 0x43,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65,
	0x57, 0x61, 0x6b, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x65,
	0x32, 0xf8, 0x01, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x4e, 0x65, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c,
	0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61,
	0x6b, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e,
	0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x64, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61,
	0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x4e,
	0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6a, 0x6e, 0x74, 0x66,
	0x78, 0x2f, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_wake_neo_proto_rawDescOnce sync.Once
	file_node_wake_neo_proto_rawDescData = file_node_wake_neo_proto_rawDesc
)

func file_node_wake_neo_proto_rawDescGZIP() []byte {
	file_node_wake_neo_proto_rawDescOnce.Do(func() {
		file_node_wake_neo_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_wake_neo_proto_rawDescData)
	})
	return file_node_wake_neo_proto_rawDescData
}

var file_node_wake_neo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_wake_neo_proto_goTypes = []interface{}{
	(*NodeWakeStartNeoMessage)(nil), // 0: com.pojtinger.felicitas.liwasc.neo.NodeWakeStartNeoMessage
	(*NodeWakeNeoMessage)(nil),      // 1: com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoMessage
	(*empty.Empty)(nil),             // 2: google.protobuf.Empty
}
var file_node_wake_neo_proto_depIdxs = []int32{
	0, // 0: com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoService.StartNodeWake:input_type -> com.pojtinger.felicitas.liwasc.neo.NodeWakeStartNeoMessage
	2, // 1: com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoService.SubscribeToNodeWakes:input_type -> google.protobuf.Empty
	1, // 2: com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoService.StartNodeWake:output_type -> com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoMessage
	1, // 3: com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoService.SubscribeToNodeWakes:output_type -> com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_wake_neo_proto_init() }
func file_node_wake_neo_proto_init() {
	if File_node_wake_neo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_wake_neo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeWakeStartNeoMessage); i {
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
		file_node_wake_neo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeWakeNeoMessage); i {
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
			RawDescriptor: file_node_wake_neo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_wake_neo_proto_goTypes,
		DependencyIndexes: file_node_wake_neo_proto_depIdxs,
		MessageInfos:      file_node_wake_neo_proto_msgTypes,
	}.Build()
	File_node_wake_neo_proto = out.File
	file_node_wake_neo_proto_rawDesc = nil
	file_node_wake_neo_proto_goTypes = nil
	file_node_wake_neo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeWakeNeoServiceClient is the client API for NodeWakeNeoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeWakeNeoServiceClient interface {
	StartNodeWake(ctx context.Context, in *NodeWakeStartNeoMessage, opts ...grpc.CallOption) (*NodeWakeNeoMessage, error)
	SubscribeToNodeWakes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (NodeWakeNeoService_SubscribeToNodeWakesClient, error)
}

type nodeWakeNeoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeWakeNeoServiceClient(cc grpc.ClientConnInterface) NodeWakeNeoServiceClient {
	return &nodeWakeNeoServiceClient{cc}
}

func (c *nodeWakeNeoServiceClient) StartNodeWake(ctx context.Context, in *NodeWakeStartNeoMessage, opts ...grpc.CallOption) (*NodeWakeNeoMessage, error) {
	out := new(NodeWakeNeoMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoService/StartNodeWake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeWakeNeoServiceClient) SubscribeToNodeWakes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (NodeWakeNeoService_SubscribeToNodeWakesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeWakeNeoService_serviceDesc.Streams[0], "/com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoService/SubscribeToNodeWakes", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeWakeNeoServiceSubscribeToNodeWakesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeWakeNeoService_SubscribeToNodeWakesClient interface {
	Recv() (*NodeWakeNeoMessage, error)
	grpc.ClientStream
}

type nodeWakeNeoServiceSubscribeToNodeWakesClient struct {
	grpc.ClientStream
}

func (x *nodeWakeNeoServiceSubscribeToNodeWakesClient) Recv() (*NodeWakeNeoMessage, error) {
	m := new(NodeWakeNeoMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeWakeNeoServiceServer is the server API for NodeWakeNeoService service.
type NodeWakeNeoServiceServer interface {
	StartNodeWake(context.Context, *NodeWakeStartNeoMessage) (*NodeWakeNeoMessage, error)
	SubscribeToNodeWakes(*empty.Empty, NodeWakeNeoService_SubscribeToNodeWakesServer) error
}

// UnimplementedNodeWakeNeoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeWakeNeoServiceServer struct {
}

func (*UnimplementedNodeWakeNeoServiceServer) StartNodeWake(context.Context, *NodeWakeStartNeoMessage) (*NodeWakeNeoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNodeWake not implemented")
}
func (*UnimplementedNodeWakeNeoServiceServer) SubscribeToNodeWakes(*empty.Empty, NodeWakeNeoService_SubscribeToNodeWakesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToNodeWakes not implemented")
}

func RegisterNodeWakeNeoServiceServer(s *grpc.Server, srv NodeWakeNeoServiceServer) {
	s.RegisterService(&_NodeWakeNeoService_serviceDesc, srv)
}

func _NodeWakeNeoService_StartNodeWake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeWakeStartNeoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeWakeNeoServiceServer).StartNodeWake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoService/StartNodeWake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeWakeNeoServiceServer).StartNodeWake(ctx, req.(*NodeWakeStartNeoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeWakeNeoService_SubscribeToNodeWakes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeWakeNeoServiceServer).SubscribeToNodeWakes(m, &nodeWakeNeoServiceSubscribeToNodeWakesServer{stream})
}

type NodeWakeNeoService_SubscribeToNodeWakesServer interface {
	Send(*NodeWakeNeoMessage) error
	grpc.ServerStream
}

type nodeWakeNeoServiceSubscribeToNodeWakesServer struct {
	grpc.ServerStream
}

func (x *nodeWakeNeoServiceSubscribeToNodeWakesServer) Send(m *NodeWakeNeoMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeWakeNeoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.pojtinger.felicitas.liwasc.neo.NodeWakeNeoService",
	HandlerType: (*NodeWakeNeoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartNodeWake",
			Handler:    _NodeWakeNeoService_StartNodeWake_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToNodeWakes",
			Handler:       _NodeWakeNeoService_SubscribeToNodeWakes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node_wake_neo.proto",
}
