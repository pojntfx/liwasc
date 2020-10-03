// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: node_wake.proto

package proto

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

type NodeWakeTriggerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MACAddress string `protobuf:"bytes,1,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
	Timeout    int64  `protobuf:"varint,2,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
}

func (x *NodeWakeTriggerMessage) Reset() {
	*x = NodeWakeTriggerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_wake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeWakeTriggerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeWakeTriggerMessage) ProtoMessage() {}

func (x *NodeWakeTriggerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_wake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeWakeTriggerMessage.ProtoReflect.Descriptor instead.
func (*NodeWakeTriggerMessage) Descriptor() ([]byte, []int) {
	return file_node_wake_proto_rawDescGZIP(), []int{0}
}

func (x *NodeWakeTriggerMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

func (x *NodeWakeTriggerMessage) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type NodeWakeReferenceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MACAddress string `protobuf:"bytes,1,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
	NodeWakeID int64  `protobuf:"varint,2,opt,name=NodeWakeID,proto3" json:"NodeWakeID,omitempty"`
}

func (x *NodeWakeReferenceMessage) Reset() {
	*x = NodeWakeReferenceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_wake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeWakeReferenceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeWakeReferenceMessage) ProtoMessage() {}

func (x *NodeWakeReferenceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_wake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeWakeReferenceMessage.ProtoReflect.Descriptor instead.
func (*NodeWakeReferenceMessage) Descriptor() ([]byte, []int) {
	return file_node_wake_proto_rawDescGZIP(), []int{1}
}

func (x *NodeWakeReferenceMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

func (x *NodeWakeReferenceMessage) GetNodeWakeID() int64 {
	if x != nil {
		return x.NodeWakeID
	}
	return 0
}

type LucidNodeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoweredOn  bool   `protobuf:"varint,1,opt,name=PoweredOn,proto3" json:"PoweredOn,omitempty"`
	MACAddress string `protobuf:"bytes,2,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
}

func (x *LucidNodeMessage) Reset() {
	*x = LucidNodeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_wake_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LucidNodeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LucidNodeMessage) ProtoMessage() {}

func (x *LucidNodeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_wake_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LucidNodeMessage.ProtoReflect.Descriptor instead.
func (*LucidNodeMessage) Descriptor() ([]byte, []int) {
	return file_node_wake_proto_rawDescGZIP(), []int{2}
}

func (x *LucidNodeMessage) GetPoweredOn() bool {
	if x != nil {
		return x.PoweredOn
	}
	return false
}

func (x *LucidNodeMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

var File_node_wake_proto protoreflect.FileDescriptor

var file_node_wake_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x77, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x22, 0x52, 0x0a,
	0x16, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x22, 0x5a, 0x0a, 0x18, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x49, 0x44, 0x22, 0x50, 0x0a,
	0x10, 0x4c, 0x75, 0x63, 0x69, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32,
	0x8d, 0x02, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0f, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a,
	0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77,
	0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78,
	0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x7d, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4e,
	0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x55, 0x70, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e,
	0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x61, 0x6b, 0x65, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e,
	0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4c, 0x75, 0x63,
	0x69, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f,
	0x6a, 0x6e, 0x74, 0x66, 0x78, 0x2f, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_wake_proto_rawDescOnce sync.Once
	file_node_wake_proto_rawDescData = file_node_wake_proto_rawDesc
)

func file_node_wake_proto_rawDescGZIP() []byte {
	file_node_wake_proto_rawDescOnce.Do(func() {
		file_node_wake_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_wake_proto_rawDescData)
	})
	return file_node_wake_proto_rawDescData
}

var file_node_wake_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_node_wake_proto_goTypes = []interface{}{
	(*NodeWakeTriggerMessage)(nil),   // 0: com.pojtinger.felicitas.liwasc.NodeWakeTriggerMessage
	(*NodeWakeReferenceMessage)(nil), // 1: com.pojtinger.felicitas.liwasc.NodeWakeReferenceMessage
	(*LucidNodeMessage)(nil),         // 2: com.pojtinger.felicitas.liwasc.LucidNodeMessage
}
var file_node_wake_proto_depIdxs = []int32{
	0, // 0: com.pojtinger.felicitas.liwasc.NodeWakeService.TriggerNodeWake:input_type -> com.pojtinger.felicitas.liwasc.NodeWakeTriggerMessage
	1, // 1: com.pojtinger.felicitas.liwasc.NodeWakeService.SubscribeToNodeWakeUp:input_type -> com.pojtinger.felicitas.liwasc.NodeWakeReferenceMessage
	1, // 2: com.pojtinger.felicitas.liwasc.NodeWakeService.TriggerNodeWake:output_type -> com.pojtinger.felicitas.liwasc.NodeWakeReferenceMessage
	2, // 3: com.pojtinger.felicitas.liwasc.NodeWakeService.SubscribeToNodeWakeUp:output_type -> com.pojtinger.felicitas.liwasc.LucidNodeMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_wake_proto_init() }
func file_node_wake_proto_init() {
	if File_node_wake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_wake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeWakeTriggerMessage); i {
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
		file_node_wake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeWakeReferenceMessage); i {
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
		file_node_wake_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LucidNodeMessage); i {
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
			RawDescriptor: file_node_wake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_wake_proto_goTypes,
		DependencyIndexes: file_node_wake_proto_depIdxs,
		MessageInfos:      file_node_wake_proto_msgTypes,
	}.Build()
	File_node_wake_proto = out.File
	file_node_wake_proto_rawDesc = nil
	file_node_wake_proto_goTypes = nil
	file_node_wake_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeWakeServiceClient is the client API for NodeWakeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeWakeServiceClient interface {
	TriggerNodeWake(ctx context.Context, in *NodeWakeTriggerMessage, opts ...grpc.CallOption) (*NodeWakeReferenceMessage, error)
	// Use -1 as NodeWakeID and a valid MAC address to return the latest wake
	// state for the MAC Address
	// Use a valid NodeWakeID and "" as a MAC address to
	// return the matching wake state
	SubscribeToNodeWakeUp(ctx context.Context, in *NodeWakeReferenceMessage, opts ...grpc.CallOption) (NodeWakeService_SubscribeToNodeWakeUpClient, error)
}

type nodeWakeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeWakeServiceClient(cc grpc.ClientConnInterface) NodeWakeServiceClient {
	return &nodeWakeServiceClient{cc}
}

func (c *nodeWakeServiceClient) TriggerNodeWake(ctx context.Context, in *NodeWakeTriggerMessage, opts ...grpc.CallOption) (*NodeWakeReferenceMessage, error) {
	out := new(NodeWakeReferenceMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.NodeWakeService/TriggerNodeWake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeWakeServiceClient) SubscribeToNodeWakeUp(ctx context.Context, in *NodeWakeReferenceMessage, opts ...grpc.CallOption) (NodeWakeService_SubscribeToNodeWakeUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeWakeService_serviceDesc.Streams[0], "/com.pojtinger.felicitas.liwasc.NodeWakeService/SubscribeToNodeWakeUp", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeWakeServiceSubscribeToNodeWakeUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeWakeService_SubscribeToNodeWakeUpClient interface {
	Recv() (*LucidNodeMessage, error)
	grpc.ClientStream
}

type nodeWakeServiceSubscribeToNodeWakeUpClient struct {
	grpc.ClientStream
}

func (x *nodeWakeServiceSubscribeToNodeWakeUpClient) Recv() (*LucidNodeMessage, error) {
	m := new(LucidNodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeWakeServiceServer is the server API for NodeWakeService service.
type NodeWakeServiceServer interface {
	TriggerNodeWake(context.Context, *NodeWakeTriggerMessage) (*NodeWakeReferenceMessage, error)
	// Use -1 as NodeWakeID and a valid MAC address to return the latest wake
	// state for the MAC Address
	// Use a valid NodeWakeID and "" as a MAC address to
	// return the matching wake state
	SubscribeToNodeWakeUp(*NodeWakeReferenceMessage, NodeWakeService_SubscribeToNodeWakeUpServer) error
}

// UnimplementedNodeWakeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeWakeServiceServer struct {
}

func (*UnimplementedNodeWakeServiceServer) TriggerNodeWake(context.Context, *NodeWakeTriggerMessage) (*NodeWakeReferenceMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerNodeWake not implemented")
}
func (*UnimplementedNodeWakeServiceServer) SubscribeToNodeWakeUp(*NodeWakeReferenceMessage, NodeWakeService_SubscribeToNodeWakeUpServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToNodeWakeUp not implemented")
}

func RegisterNodeWakeServiceServer(s *grpc.Server, srv NodeWakeServiceServer) {
	s.RegisterService(&_NodeWakeService_serviceDesc, srv)
}

func _NodeWakeService_TriggerNodeWake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeWakeTriggerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeWakeServiceServer).TriggerNodeWake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.NodeWakeService/TriggerNodeWake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeWakeServiceServer).TriggerNodeWake(ctx, req.(*NodeWakeTriggerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeWakeService_SubscribeToNodeWakeUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeWakeReferenceMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeWakeServiceServer).SubscribeToNodeWakeUp(m, &nodeWakeServiceSubscribeToNodeWakeUpServer{stream})
}

type NodeWakeService_SubscribeToNodeWakeUpServer interface {
	Send(*LucidNodeMessage) error
	grpc.ServerStream
}

type nodeWakeServiceSubscribeToNodeWakeUpServer struct {
	grpc.ServerStream
}

func (x *nodeWakeServiceSubscribeToNodeWakeUpServer) Send(m *LucidNodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeWakeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.pojtinger.felicitas.liwasc.NodeWakeService",
	HandlerType: (*NodeWakeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TriggerNodeWake",
			Handler:    _NodeWakeService_TriggerNodeWake_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToNodeWakeUp",
			Handler:       _NodeWakeService_SubscribeToNodeWakeUp_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node_wake.proto",
}
