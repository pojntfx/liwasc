// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: metadata.proto

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

type ScannerMetadataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnets []string `protobuf:"bytes,1,rep,name=Subnets,proto3" json:"Subnets,omitempty"`
	Device  string   `protobuf:"bytes,2,opt,name=Device,proto3" json:"Device,omitempty"`
}

func (x *ScannerMetadataMessage) Reset() {
	*x = ScannerMetadataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScannerMetadataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScannerMetadataMessage) ProtoMessage() {}

func (x *ScannerMetadataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScannerMetadataMessage.ProtoReflect.Descriptor instead.
func (*ScannerMetadataMessage) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *ScannerMetadataMessage) GetSubnets() []string {
	if x != nil {
		return x.Subnets
	}
	return nil
}

func (x *ScannerMetadataMessage) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type NodeMetadataReferenceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MACAddress string `protobuf:"bytes,1,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
}

func (x *NodeMetadataReferenceMessage) Reset() {
	*x = NodeMetadataReferenceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMetadataReferenceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetadataReferenceMessage) ProtoMessage() {}

func (x *NodeMetadataReferenceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetadataReferenceMessage.ProtoReflect.Descriptor instead.
func (*NodeMetadataReferenceMessage) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *NodeMetadataReferenceMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

type NodeMetadataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MACAddress   string `protobuf:"bytes,1,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
	Vendor       string `protobuf:"bytes,3,opt,name=Vendor,proto3" json:"Vendor,omitempty"`
	Registry     string `protobuf:"bytes,4,opt,name=Registry,proto3" json:"Registry,omitempty"`
	Organization string `protobuf:"bytes,5,opt,name=Organization,proto3" json:"Organization,omitempty"`
	Address      string `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	Visible      bool   `protobuf:"varint,7,opt,name=Visible,proto3" json:"Visible,omitempty"`
}

func (x *NodeMetadataMessage) Reset() {
	*x = NodeMetadataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMetadataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetadataMessage) ProtoMessage() {}

func (x *NodeMetadataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetadataMessage.ProtoReflect.Descriptor instead.
func (*NodeMetadataMessage) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *NodeMetadataMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

func (x *NodeMetadataMessage) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *NodeMetadataMessage) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *NodeMetadataMessage) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *NodeMetadataMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *NodeMetadataMessage) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

type PortMetadataReferenceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortNumber        int64  `protobuf:"varint,1,opt,name=PortNumber,proto3" json:"PortNumber,omitempty"`
	TransportProtocol string `protobuf:"bytes,2,opt,name=TransportProtocol,proto3" json:"TransportProtocol,omitempty"`
}

func (x *PortMetadataReferenceMessage) Reset() {
	*x = PortMetadataReferenceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMetadataReferenceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMetadataReferenceMessage) ProtoMessage() {}

func (x *PortMetadataReferenceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMetadataReferenceMessage.ProtoReflect.Descriptor instead.
func (*PortMetadataReferenceMessage) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *PortMetadataReferenceMessage) GetPortNumber() int64 {
	if x != nil {
		return x.PortNumber
	}
	return 0
}

func (x *PortMetadataReferenceMessage) GetTransportProtocol() string {
	if x != nil {
		return x.TransportProtocol
	}
	return ""
}

type PortMetadataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName             string `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	PortNumber              int64  `protobuf:"varint,2,opt,name=PortNumber,proto3" json:"PortNumber,omitempty"`
	TransportProtocol       string `protobuf:"bytes,3,opt,name=TransportProtocol,proto3" json:"TransportProtocol,omitempty"`
	Description             string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Assignee                string `protobuf:"bytes,5,opt,name=Assignee,proto3" json:"Assignee,omitempty"`
	Contact                 string `protobuf:"bytes,6,opt,name=Contact,proto3" json:"Contact,omitempty"`
	RegistrationDate        string `protobuf:"bytes,7,opt,name=RegistrationDate,proto3" json:"RegistrationDate,omitempty"`
	ModificationDate        string `protobuf:"bytes,8,opt,name=ModificationDate,proto3" json:"ModificationDate,omitempty"`
	Reference               string `protobuf:"bytes,9,opt,name=Reference,proto3" json:"Reference,omitempty"`
	ServiceCode             string `protobuf:"bytes,10,opt,name=ServiceCode,proto3" json:"ServiceCode,omitempty"`
	UnauthorizedUseReported string `protobuf:"bytes,11,opt,name=UnauthorizedUseReported,proto3" json:"UnauthorizedUseReported,omitempty"`
	AssignmentNotes         string `protobuf:"bytes,12,opt,name=AssignmentNotes,proto3" json:"AssignmentNotes,omitempty"`
}

func (x *PortMetadataMessage) Reset() {
	*x = PortMetadataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMetadataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMetadataMessage) ProtoMessage() {}

func (x *PortMetadataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMetadataMessage.ProtoReflect.Descriptor instead.
func (*PortMetadataMessage) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{4}
}

func (x *PortMetadataMessage) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *PortMetadataMessage) GetPortNumber() int64 {
	if x != nil {
		return x.PortNumber
	}
	return 0
}

func (x *PortMetadataMessage) GetTransportProtocol() string {
	if x != nil {
		return x.TransportProtocol
	}
	return ""
}

func (x *PortMetadataMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PortMetadataMessage) GetAssignee() string {
	if x != nil {
		return x.Assignee
	}
	return ""
}

func (x *PortMetadataMessage) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *PortMetadataMessage) GetRegistrationDate() string {
	if x != nil {
		return x.RegistrationDate
	}
	return ""
}

func (x *PortMetadataMessage) GetModificationDate() string {
	if x != nil {
		return x.ModificationDate
	}
	return ""
}

func (x *PortMetadataMessage) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *PortMetadataMessage) GetServiceCode() string {
	if x != nil {
		return x.ServiceCode
	}
	return ""
}

func (x *PortMetadataMessage) GetUnauthorizedUseReported() string {
	if x != nil {
		return x.UnauthorizedUseReported
	}
	return ""
}

func (x *PortMetadataMessage) GetAssignmentNotes() string {
	if x != nil {
		return x.AssignmentNotes
	}
	return ""
}

var File_metadata_proto protoreflect.FileDescriptor

var file_metadata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e,
	0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x16, 0x53, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x1c, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x6c, 0x0a, 0x1c, 0x50, 0x6f, 0x72,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x72,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50,
	0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xd9, 0x03, 0x0a, 0x13, 0x50, 0x6f, 0x72, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x38, 0x0a, 0x17, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x17, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x32, 0xf8, 0x02, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c,
	0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x7f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2f, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c,
	0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x7f, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2f, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65,
	0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x25,
	0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6a,
	0x6e, 0x74, 0x66, 0x78, 0x2f, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_proto_rawDescOnce sync.Once
	file_metadata_proto_rawDescData = file_metadata_proto_rawDesc
)

func file_metadata_proto_rawDescGZIP() []byte {
	file_metadata_proto_rawDescOnce.Do(func() {
		file_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_proto_rawDescData)
	})
	return file_metadata_proto_rawDescData
}

var file_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_metadata_proto_goTypes = []interface{}{
	(*ScannerMetadataMessage)(nil),       // 0: com.pojtinger.felicitas.liwasc.ScannerMetadataMessage
	(*NodeMetadataReferenceMessage)(nil), // 1: com.pojtinger.felicitas.liwasc.NodeMetadataReferenceMessage
	(*NodeMetadataMessage)(nil),          // 2: com.pojtinger.felicitas.liwasc.NodeMetadataMessage
	(*PortMetadataReferenceMessage)(nil), // 3: com.pojtinger.felicitas.liwasc.PortMetadataReferenceMessage
	(*PortMetadataMessage)(nil),          // 4: com.pojtinger.felicitas.liwasc.PortMetadataMessage
	(*empty.Empty)(nil),                  // 5: google.protobuf.Empty
}
var file_metadata_proto_depIdxs = []int32{
	5, // 0: com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForScanner:input_type -> google.protobuf.Empty
	1, // 1: com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForNode:input_type -> com.pojtinger.felicitas.liwasc.NodeMetadataReferenceMessage
	3, // 2: com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForPort:input_type -> com.pojtinger.felicitas.liwasc.PortMetadataReferenceMessage
	0, // 3: com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForScanner:output_type -> com.pojtinger.felicitas.liwasc.ScannerMetadataMessage
	2, // 4: com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForNode:output_type -> com.pojtinger.felicitas.liwasc.NodeMetadataMessage
	4, // 5: com.pojtinger.felicitas.liwasc.MetadataService.GetMetadataForPort:output_type -> com.pojtinger.felicitas.liwasc.PortMetadataMessage
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_metadata_proto_init() }
func file_metadata_proto_init() {
	if File_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScannerMetadataMessage); i {
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
		file_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMetadataReferenceMessage); i {
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
		file_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMetadataMessage); i {
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
		file_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMetadataReferenceMessage); i {
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
		file_metadata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMetadataMessage); i {
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
			RawDescriptor: file_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metadata_proto_goTypes,
		DependencyIndexes: file_metadata_proto_depIdxs,
		MessageInfos:      file_metadata_proto_msgTypes,
	}.Build()
	File_metadata_proto = out.File
	file_metadata_proto_rawDesc = nil
	file_metadata_proto_goTypes = nil
	file_metadata_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataServiceClient interface {
	GetMetadataForScanner(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ScannerMetadataMessage, error)
	GetMetadataForNode(ctx context.Context, in *NodeMetadataReferenceMessage, opts ...grpc.CallOption) (*NodeMetadataMessage, error)
	GetMetadataForPort(ctx context.Context, in *PortMetadataReferenceMessage, opts ...grpc.CallOption) (*PortMetadataMessage, error)
}

type metadataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataServiceClient(cc grpc.ClientConnInterface) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) GetMetadataForScanner(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ScannerMetadataMessage, error) {
	out := new(ScannerMetadataMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.MetadataService/GetMetadataForScanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetMetadataForNode(ctx context.Context, in *NodeMetadataReferenceMessage, opts ...grpc.CallOption) (*NodeMetadataMessage, error) {
	out := new(NodeMetadataMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.MetadataService/GetMetadataForNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetMetadataForPort(ctx context.Context, in *PortMetadataReferenceMessage, opts ...grpc.CallOption) (*PortMetadataMessage, error) {
	out := new(PortMetadataMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.MetadataService/GetMetadataForPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServiceServer is the server API for MetadataService service.
type MetadataServiceServer interface {
	GetMetadataForScanner(context.Context, *empty.Empty) (*ScannerMetadataMessage, error)
	GetMetadataForNode(context.Context, *NodeMetadataReferenceMessage) (*NodeMetadataMessage, error)
	GetMetadataForPort(context.Context, *PortMetadataReferenceMessage) (*PortMetadataMessage, error)
}

// UnimplementedMetadataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetadataServiceServer struct {
}

func (*UnimplementedMetadataServiceServer) GetMetadataForScanner(context.Context, *empty.Empty) (*ScannerMetadataMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadataForScanner not implemented")
}
func (*UnimplementedMetadataServiceServer) GetMetadataForNode(context.Context, *NodeMetadataReferenceMessage) (*NodeMetadataMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadataForNode not implemented")
}
func (*UnimplementedMetadataServiceServer) GetMetadataForPort(context.Context, *PortMetadataReferenceMessage) (*PortMetadataMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadataForPort not implemented")
}

func RegisterMetadataServiceServer(s *grpc.Server, srv MetadataServiceServer) {
	s.RegisterService(&_MetadataService_serviceDesc, srv)
}

func _MetadataService_GetMetadataForScanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetMetadataForScanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.MetadataService/GetMetadataForScanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetMetadataForScanner(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetMetadataForNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeMetadataReferenceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetMetadataForNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.MetadataService/GetMetadataForNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetMetadataForNode(ctx, req.(*NodeMetadataReferenceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetMetadataForPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortMetadataReferenceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetMetadataForPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.MetadataService/GetMetadataForPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetMetadataForPort(ctx, req.(*PortMetadataReferenceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetadataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.pojtinger.felicitas.liwasc.MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetadataForScanner",
			Handler:    _MetadataService_GetMetadataForScanner_Handler,
		},
		{
			MethodName: "GetMetadataForNode",
			Handler:    _MetadataService_GetMetadataForNode_Handler,
		},
		{
			MethodName: "GetMetadataForPort",
			Handler:    _MetadataService_GetMetadataForPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata.proto",
}
