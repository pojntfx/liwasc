// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: node_and_port_scan.proto

package v1

import (
	context "context"
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

type NodeScanStartMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeScanTimeout int64  `protobuf:"varint,1,opt,name=NodeScanTimeout,proto3" json:"NodeScanTimeout,omitempty"`
	PortScanTimeout int64  `protobuf:"varint,2,opt,name=PortScanTimeout,proto3" json:"PortScanTimeout,omitempty"`
	MACAddress      string `protobuf:"bytes,3,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"` // Scopes the scan to one node. Set to "" to scan all.
}

func (x *NodeScanStartMessage) Reset() {
	*x = NodeScanStartMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeScanStartMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeScanStartMessage) ProtoMessage() {}

func (x *NodeScanStartMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeScanStartMessage.ProtoReflect.Descriptor instead.
func (*NodeScanStartMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{0}
}

func (x *NodeScanStartMessage) GetNodeScanTimeout() int64 {
	if x != nil {
		return x.NodeScanTimeout
	}
	return 0
}

func (x *NodeScanStartMessage) GetPortScanTimeout() int64 {
	if x != nil {
		return x.PortScanTimeout
	}
	return 0
}

func (x *NodeScanStartMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

type NodeScanMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Done      bool   `protobuf:"varint,3,opt,name=Done,proto3" json:"Done,omitempty"`
}

func (x *NodeScanMessage) Reset() {
	*x = NodeScanMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeScanMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeScanMessage) ProtoMessage() {}

func (x *NodeScanMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeScanMessage.ProtoReflect.Descriptor instead.
func (*NodeScanMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{1}
}

func (x *NodeScanMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NodeScanMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NodeScanMessage) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type NodeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt  string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Priority   int64  `protobuf:"varint,3,opt,name=Priority,proto3" json:"Priority,omitempty"`
	MACAddress string `protobuf:"bytes,4,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
	IPAddress  string `protobuf:"bytes,5,opt,name=IPAddress,proto3" json:"IPAddress,omitempty"`
	NodeScanID int64  `protobuf:"varint,6,opt,name=NodeScanID,proto3" json:"NodeScanID,omitempty"`
	PoweredOn  bool   `protobuf:"varint,7,opt,name=PoweredOn,proto3" json:"PoweredOn,omitempty"`
}

func (x *NodeMessage) Reset() {
	*x = NodeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMessage) ProtoMessage() {}

func (x *NodeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMessage.ProtoReflect.Descriptor instead.
func (*NodeMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{2}
}

func (x *NodeMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NodeMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NodeMessage) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *NodeMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

func (x *NodeMessage) GetIPAddress() string {
	if x != nil {
		return x.IPAddress
	}
	return ""
}

func (x *NodeMessage) GetNodeScanID() int64 {
	if x != nil {
		return x.NodeScanID
	}
	return 0
}

func (x *NodeMessage) GetPoweredOn() bool {
	if x != nil {
		return x.PoweredOn
	}
	return false
}

type PortScanMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Done      bool   `protobuf:"varint,3,opt,name=Done,proto3" json:"Done,omitempty"`
	NodeID    int64  `protobuf:"varint,4,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
}

func (x *PortScanMessage) Reset() {
	*x = PortScanMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortScanMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortScanMessage) ProtoMessage() {}

func (x *PortScanMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortScanMessage.ProtoReflect.Descriptor instead.
func (*PortScanMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{3}
}

func (x *PortScanMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PortScanMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PortScanMessage) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *PortScanMessage) GetNodeID() int64 {
	if x != nil {
		return x.NodeID
	}
	return 0
}

type PortMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt         string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Priority          int64  `protobuf:"varint,3,opt,name=Priority,proto3" json:"Priority,omitempty"`
	PortNumber        int64  `protobuf:"varint,4,opt,name=PortNumber,proto3" json:"PortNumber,omitempty"`
	TransportProtocol string `protobuf:"bytes,5,opt,name=TransportProtocol,proto3" json:"TransportProtocol,omitempty"`
	PortScanID        int64  `protobuf:"varint,6,opt,name=PortScanID,proto3" json:"PortScanID,omitempty"`
}

func (x *PortMessage) Reset() {
	*x = PortMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMessage) ProtoMessage() {}

func (x *PortMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMessage.ProtoReflect.Descriptor instead.
func (*PortMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{4}
}

func (x *PortMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PortMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PortMessage) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *PortMessage) GetPortNumber() int64 {
	if x != nil {
		return x.PortNumber
	}
	return 0
}

func (x *PortMessage) GetTransportProtocol() string {
	if x != nil {
		return x.TransportProtocol
	}
	return ""
}

func (x *PortMessage) GetPortScanID() int64 {
	if x != nil {
		return x.PortScanID
	}
	return 0
}

var File_node_and_port_scan_proto protoreflect.FileDescriptor

var file_node_and_port_scan_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e,
	0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63,
	0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x53, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x6b, 0x0a, 0x0f, 0x50,
	0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0xc5, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x72,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x44,
	0x32, 0xaf, 0x04, 0x0a, 0x16, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74,
	0x53, 0x63, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x30, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c,
	0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66,
	0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69,
	0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x6a, 0x0a, 0x10, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2b,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66,
	0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69,
	0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x6e, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x73, 0x12, 0x27,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66,
	0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f,
	0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69,
	0x77, 0x61, 0x73, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x6a, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78,
	0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f,
	0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69,
	0x77, 0x61, 0x73, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x6f, 0x6a, 0x6e, 0x74, 0x66, 0x78, 0x2f, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_and_port_scan_proto_rawDescOnce sync.Once
	file_node_and_port_scan_proto_rawDescData = file_node_and_port_scan_proto_rawDesc
)

func file_node_and_port_scan_proto_rawDescGZIP() []byte {
	file_node_and_port_scan_proto_rawDescOnce.Do(func() {
		file_node_and_port_scan_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_and_port_scan_proto_rawDescData)
	})
	return file_node_and_port_scan_proto_rawDescData
}

var file_node_and_port_scan_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_node_and_port_scan_proto_goTypes = []interface{}{
	(*NodeScanStartMessage)(nil), // 0: com.pojtinger.felicitas.liwasc.NodeScanStartMessage
	(*NodeScanMessage)(nil),      // 1: com.pojtinger.felicitas.liwasc.NodeScanMessage
	(*NodeMessage)(nil),          // 2: com.pojtinger.felicitas.liwasc.NodeMessage
	(*PortScanMessage)(nil),      // 3: com.pojtinger.felicitas.liwasc.PortScanMessage
	(*PortMessage)(nil),          // 4: com.pojtinger.felicitas.liwasc.PortMessage
	(*empty.Empty)(nil),          // 5: google.protobuf.Empty
}
var file_node_and_port_scan_proto_depIdxs = []int32{
	0, // 0: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.StartNodeScan:input_type -> com.pojtinger.felicitas.liwasc.NodeScanStartMessage
	5, // 1: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToNodeScans:input_type -> google.protobuf.Empty
	1, // 2: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToNodes:input_type -> com.pojtinger.felicitas.liwasc.NodeScanMessage
	2, // 3: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToPortScans:input_type -> com.pojtinger.felicitas.liwasc.NodeMessage
	3, // 4: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToPorts:input_type -> com.pojtinger.felicitas.liwasc.PortScanMessage
	1, // 5: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.StartNodeScan:output_type -> com.pojtinger.felicitas.liwasc.NodeScanMessage
	1, // 6: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToNodeScans:output_type -> com.pojtinger.felicitas.liwasc.NodeScanMessage
	2, // 7: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToNodes:output_type -> com.pojtinger.felicitas.liwasc.NodeMessage
	3, // 8: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToPortScans:output_type -> com.pojtinger.felicitas.liwasc.PortScanMessage
	4, // 9: com.pojtinger.felicitas.liwasc.NodeAndPortScanService.SubscribeToPorts:output_type -> com.pojtinger.felicitas.liwasc.PortMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_and_port_scan_proto_init() }
func file_node_and_port_scan_proto_init() {
	if File_node_and_port_scan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_and_port_scan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeScanStartMessage); i {
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
		file_node_and_port_scan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeScanMessage); i {
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
		file_node_and_port_scan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMessage); i {
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
		file_node_and_port_scan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortScanMessage); i {
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
		file_node_and_port_scan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMessage); i {
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
			RawDescriptor: file_node_and_port_scan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_and_port_scan_proto_goTypes,
		DependencyIndexes: file_node_and_port_scan_proto_depIdxs,
		MessageInfos:      file_node_and_port_scan_proto_msgTypes,
	}.Build()
	File_node_and_port_scan_proto = out.File
	file_node_and_port_scan_proto_rawDesc = nil
	file_node_and_port_scan_proto_goTypes = nil
	file_node_and_port_scan_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeAndPortScanServiceClient is the client API for NodeAndPortScanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeAndPortScanServiceClient interface {
	StartNodeScan(ctx context.Context, in *NodeScanStartMessage, opts ...grpc.CallOption) (*NodeScanMessage, error)
	SubscribeToNodeScans(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (NodeAndPortScanService_SubscribeToNodeScansClient, error)
	SubscribeToNodes(ctx context.Context, in *NodeScanMessage, opts ...grpc.CallOption) (NodeAndPortScanService_SubscribeToNodesClient, error)
	SubscribeToPortScans(ctx context.Context, in *NodeMessage, opts ...grpc.CallOption) (NodeAndPortScanService_SubscribeToPortScansClient, error)
	SubscribeToPorts(ctx context.Context, in *PortScanMessage, opts ...grpc.CallOption) (NodeAndPortScanService_SubscribeToPortsClient, error)
}

type nodeAndPortScanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeAndPortScanServiceClient(cc grpc.ClientConnInterface) NodeAndPortScanServiceClient {
	return &nodeAndPortScanServiceClient{cc}
}

func (c *nodeAndPortScanServiceClient) StartNodeScan(ctx context.Context, in *NodeScanStartMessage, opts ...grpc.CallOption) (*NodeScanMessage, error) {
	out := new(NodeScanMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.NodeAndPortScanService/StartNodeScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAndPortScanServiceClient) SubscribeToNodeScans(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (NodeAndPortScanService_SubscribeToNodeScansClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAndPortScanService_serviceDesc.Streams[0], "/com.pojtinger.felicitas.liwasc.NodeAndPortScanService/SubscribeToNodeScans", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAndPortScanServiceSubscribeToNodeScansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeAndPortScanService_SubscribeToNodeScansClient interface {
	Recv() (*NodeScanMessage, error)
	grpc.ClientStream
}

type nodeAndPortScanServiceSubscribeToNodeScansClient struct {
	grpc.ClientStream
}

func (x *nodeAndPortScanServiceSubscribeToNodeScansClient) Recv() (*NodeScanMessage, error) {
	m := new(NodeScanMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAndPortScanServiceClient) SubscribeToNodes(ctx context.Context, in *NodeScanMessage, opts ...grpc.CallOption) (NodeAndPortScanService_SubscribeToNodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAndPortScanService_serviceDesc.Streams[1], "/com.pojtinger.felicitas.liwasc.NodeAndPortScanService/SubscribeToNodes", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAndPortScanServiceSubscribeToNodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeAndPortScanService_SubscribeToNodesClient interface {
	Recv() (*NodeMessage, error)
	grpc.ClientStream
}

type nodeAndPortScanServiceSubscribeToNodesClient struct {
	grpc.ClientStream
}

func (x *nodeAndPortScanServiceSubscribeToNodesClient) Recv() (*NodeMessage, error) {
	m := new(NodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAndPortScanServiceClient) SubscribeToPortScans(ctx context.Context, in *NodeMessage, opts ...grpc.CallOption) (NodeAndPortScanService_SubscribeToPortScansClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAndPortScanService_serviceDesc.Streams[2], "/com.pojtinger.felicitas.liwasc.NodeAndPortScanService/SubscribeToPortScans", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAndPortScanServiceSubscribeToPortScansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeAndPortScanService_SubscribeToPortScansClient interface {
	Recv() (*PortScanMessage, error)
	grpc.ClientStream
}

type nodeAndPortScanServiceSubscribeToPortScansClient struct {
	grpc.ClientStream
}

func (x *nodeAndPortScanServiceSubscribeToPortScansClient) Recv() (*PortScanMessage, error) {
	m := new(PortScanMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAndPortScanServiceClient) SubscribeToPorts(ctx context.Context, in *PortScanMessage, opts ...grpc.CallOption) (NodeAndPortScanService_SubscribeToPortsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAndPortScanService_serviceDesc.Streams[3], "/com.pojtinger.felicitas.liwasc.NodeAndPortScanService/SubscribeToPorts", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAndPortScanServiceSubscribeToPortsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeAndPortScanService_SubscribeToPortsClient interface {
	Recv() (*PortMessage, error)
	grpc.ClientStream
}

type nodeAndPortScanServiceSubscribeToPortsClient struct {
	grpc.ClientStream
}

func (x *nodeAndPortScanServiceSubscribeToPortsClient) Recv() (*PortMessage, error) {
	m := new(PortMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeAndPortScanServiceServer is the server API for NodeAndPortScanService service.
type NodeAndPortScanServiceServer interface {
	StartNodeScan(context.Context, *NodeScanStartMessage) (*NodeScanMessage, error)
	SubscribeToNodeScans(*empty.Empty, NodeAndPortScanService_SubscribeToNodeScansServer) error
	SubscribeToNodes(*NodeScanMessage, NodeAndPortScanService_SubscribeToNodesServer) error
	SubscribeToPortScans(*NodeMessage, NodeAndPortScanService_SubscribeToPortScansServer) error
	SubscribeToPorts(*PortScanMessage, NodeAndPortScanService_SubscribeToPortsServer) error
}

// UnimplementedNodeAndPortScanServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeAndPortScanServiceServer struct {
}

func (*UnimplementedNodeAndPortScanServiceServer) StartNodeScan(context.Context, *NodeScanStartMessage) (*NodeScanMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNodeScan not implemented")
}
func (*UnimplementedNodeAndPortScanServiceServer) SubscribeToNodeScans(*empty.Empty, NodeAndPortScanService_SubscribeToNodeScansServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToNodeScans not implemented")
}
func (*UnimplementedNodeAndPortScanServiceServer) SubscribeToNodes(*NodeScanMessage, NodeAndPortScanService_SubscribeToNodesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToNodes not implemented")
}
func (*UnimplementedNodeAndPortScanServiceServer) SubscribeToPortScans(*NodeMessage, NodeAndPortScanService_SubscribeToPortScansServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToPortScans not implemented")
}
func (*UnimplementedNodeAndPortScanServiceServer) SubscribeToPorts(*PortScanMessage, NodeAndPortScanService_SubscribeToPortsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToPorts not implemented")
}

func RegisterNodeAndPortScanServiceServer(s *grpc.Server, srv NodeAndPortScanServiceServer) {
	s.RegisterService(&_NodeAndPortScanService_serviceDesc, srv)
}

func _NodeAndPortScanService_StartNodeScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeScanStartMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAndPortScanServiceServer).StartNodeScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.NodeAndPortScanService/StartNodeScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAndPortScanServiceServer).StartNodeScan(ctx, req.(*NodeScanStartMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAndPortScanService_SubscribeToNodeScans_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeAndPortScanServiceServer).SubscribeToNodeScans(m, &nodeAndPortScanServiceSubscribeToNodeScansServer{stream})
}

type NodeAndPortScanService_SubscribeToNodeScansServer interface {
	Send(*NodeScanMessage) error
	grpc.ServerStream
}

type nodeAndPortScanServiceSubscribeToNodeScansServer struct {
	grpc.ServerStream
}

func (x *nodeAndPortScanServiceSubscribeToNodeScansServer) Send(m *NodeScanMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeAndPortScanService_SubscribeToNodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeScanMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeAndPortScanServiceServer).SubscribeToNodes(m, &nodeAndPortScanServiceSubscribeToNodesServer{stream})
}

type NodeAndPortScanService_SubscribeToNodesServer interface {
	Send(*NodeMessage) error
	grpc.ServerStream
}

type nodeAndPortScanServiceSubscribeToNodesServer struct {
	grpc.ServerStream
}

func (x *nodeAndPortScanServiceSubscribeToNodesServer) Send(m *NodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeAndPortScanService_SubscribeToPortScans_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeAndPortScanServiceServer).SubscribeToPortScans(m, &nodeAndPortScanServiceSubscribeToPortScansServer{stream})
}

type NodeAndPortScanService_SubscribeToPortScansServer interface {
	Send(*PortScanMessage) error
	grpc.ServerStream
}

type nodeAndPortScanServiceSubscribeToPortScansServer struct {
	grpc.ServerStream
}

func (x *nodeAndPortScanServiceSubscribeToPortScansServer) Send(m *PortScanMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeAndPortScanService_SubscribeToPorts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PortScanMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeAndPortScanServiceServer).SubscribeToPorts(m, &nodeAndPortScanServiceSubscribeToPortsServer{stream})
}

type NodeAndPortScanService_SubscribeToPortsServer interface {
	Send(*PortMessage) error
	grpc.ServerStream
}

type nodeAndPortScanServiceSubscribeToPortsServer struct {
	grpc.ServerStream
}

func (x *nodeAndPortScanServiceSubscribeToPortsServer) Send(m *PortMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeAndPortScanService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.pojtinger.felicitas.liwasc.NodeAndPortScanService",
	HandlerType: (*NodeAndPortScanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartNodeScan",
			Handler:    _NodeAndPortScanService_StartNodeScan_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToNodeScans",
			Handler:       _NodeAndPortScanService_SubscribeToNodeScans_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToNodes",
			Handler:       _NodeAndPortScanService_SubscribeToNodes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToPortScans",
			Handler:       _NodeAndPortScanService_SubscribeToPortScans_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToPorts",
			Handler:       _NodeAndPortScanService_SubscribeToPorts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node_and_port_scan.proto",
}
