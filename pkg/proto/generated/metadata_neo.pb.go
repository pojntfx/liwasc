// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: metadata_neo.proto

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

type ScannerMetadataNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnets []string `protobuf:"bytes,1,rep,name=Subnets,proto3" json:"Subnets,omitempty"`
	Device  string   `protobuf:"bytes,2,opt,name=Device,proto3" json:"Device,omitempty"`
}

func (x *ScannerMetadataNeoMessage) Reset() {
	*x = ScannerMetadataNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_neo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScannerMetadataNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScannerMetadataNeoMessage) ProtoMessage() {}

func (x *ScannerMetadataNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_neo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScannerMetadataNeoMessage.ProtoReflect.Descriptor instead.
func (*ScannerMetadataNeoMessage) Descriptor() ([]byte, []int) {
	return file_metadata_neo_proto_rawDescGZIP(), []int{0}
}

func (x *ScannerMetadataNeoMessage) GetSubnets() []string {
	if x != nil {
		return x.Subnets
	}
	return nil
}

func (x *ScannerMetadataNeoMessage) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type NodeMetadataReferenceNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MACAddress string `protobuf:"bytes,1,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
}

func (x *NodeMetadataReferenceNeoMessage) Reset() {
	*x = NodeMetadataReferenceNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_neo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMetadataReferenceNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetadataReferenceNeoMessage) ProtoMessage() {}

func (x *NodeMetadataReferenceNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_neo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetadataReferenceNeoMessage.ProtoReflect.Descriptor instead.
func (*NodeMetadataReferenceNeoMessage) Descriptor() ([]byte, []int) {
	return file_metadata_neo_proto_rawDescGZIP(), []int{1}
}

func (x *NodeMetadataReferenceNeoMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

type NodeMetadataNeoMessage struct {
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

func (x *NodeMetadataNeoMessage) Reset() {
	*x = NodeMetadataNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_neo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMetadataNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetadataNeoMessage) ProtoMessage() {}

func (x *NodeMetadataNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_neo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetadataNeoMessage.ProtoReflect.Descriptor instead.
func (*NodeMetadataNeoMessage) Descriptor() ([]byte, []int) {
	return file_metadata_neo_proto_rawDescGZIP(), []int{2}
}

func (x *NodeMetadataNeoMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

func (x *NodeMetadataNeoMessage) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *NodeMetadataNeoMessage) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *NodeMetadataNeoMessage) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *NodeMetadataNeoMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *NodeMetadataNeoMessage) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

type PortMetadataReferenceNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortNumber        string `protobuf:"bytes,1,opt,name=PortNumber,proto3" json:"PortNumber,omitempty"`
	TransportProtocol string `protobuf:"bytes,2,opt,name=TransportProtocol,proto3" json:"TransportProtocol,omitempty"`
}

func (x *PortMetadataReferenceNeoMessage) Reset() {
	*x = PortMetadataReferenceNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_neo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMetadataReferenceNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMetadataReferenceNeoMessage) ProtoMessage() {}

func (x *PortMetadataReferenceNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_neo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMetadataReferenceNeoMessage.ProtoReflect.Descriptor instead.
func (*PortMetadataReferenceNeoMessage) Descriptor() ([]byte, []int) {
	return file_metadata_neo_proto_rawDescGZIP(), []int{3}
}

func (x *PortMetadataReferenceNeoMessage) GetPortNumber() string {
	if x != nil {
		return x.PortNumber
	}
	return ""
}

func (x *PortMetadataReferenceNeoMessage) GetTransportProtocol() string {
	if x != nil {
		return x.TransportProtocol
	}
	return ""
}

type PortMetadataNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName             string `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	PortNumber              string `protobuf:"bytes,2,opt,name=PortNumber,proto3" json:"PortNumber,omitempty"`
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

func (x *PortMetadataNeoMessage) Reset() {
	*x = PortMetadataNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_neo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMetadataNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMetadataNeoMessage) ProtoMessage() {}

func (x *PortMetadataNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_neo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMetadataNeoMessage.ProtoReflect.Descriptor instead.
func (*PortMetadataNeoMessage) Descriptor() ([]byte, []int) {
	return file_metadata_neo_proto_rawDescGZIP(), []int{4}
}

func (x *PortMetadataNeoMessage) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetPortNumber() string {
	if x != nil {
		return x.PortNumber
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetTransportProtocol() string {
	if x != nil {
		return x.TransportProtocol
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetAssignee() string {
	if x != nil {
		return x.Assignee
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetRegistrationDate() string {
	if x != nil {
		return x.RegistrationDate
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetModificationDate() string {
	if x != nil {
		return x.ModificationDate
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetServiceCode() string {
	if x != nil {
		return x.ServiceCode
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetUnauthorizedUseReported() string {
	if x != nil {
		return x.UnauthorizedUseReported
	}
	return ""
}

func (x *PortMetadataNeoMessage) GetAssignmentNotes() string {
	if x != nil {
		return x.AssignmentNotes
	}
	return ""
}

var File_metadata_neo_proto protoreflect.FileDescriptor

var file_metadata_neo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x65, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63,
	0x2e, 0x6e, 0x65, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4d, 0x0a, 0x19, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x41, 0x0a, 0x1f, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x16, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x6f, 0x0a, 0x1f, 0x50, 0x6f,
	0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xdc, 0x03, 0x0a, 0x16,
	0x50, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x65, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6f,
	0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x12, 0x28, 0x0a, 0x0f, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x32, 0xa0, 0x03, 0x0a, 0x12, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x46, 0x6f, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e,
	0x6e, 0x65, 0x6f, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x8d, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73,
	0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61,
	0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x8d, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73,
	0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61,
	0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x25, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6a, 0x6e,
	0x74, 0x66, 0x78, 0x2f, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_neo_proto_rawDescOnce sync.Once
	file_metadata_neo_proto_rawDescData = file_metadata_neo_proto_rawDesc
)

func file_metadata_neo_proto_rawDescGZIP() []byte {
	file_metadata_neo_proto_rawDescOnce.Do(func() {
		file_metadata_neo_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_neo_proto_rawDescData)
	})
	return file_metadata_neo_proto_rawDescData
}

var file_metadata_neo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_metadata_neo_proto_goTypes = []interface{}{
	(*ScannerMetadataNeoMessage)(nil),       // 0: com.pojtinger.felicitas.liwasc.neo.ScannerMetadataNeoMessage
	(*NodeMetadataReferenceNeoMessage)(nil), // 1: com.pojtinger.felicitas.liwasc.neo.NodeMetadataReferenceNeoMessage
	(*NodeMetadataNeoMessage)(nil),          // 2: com.pojtinger.felicitas.liwasc.neo.NodeMetadataNeoMessage
	(*PortMetadataReferenceNeoMessage)(nil), // 3: com.pojtinger.felicitas.liwasc.neo.PortMetadataReferenceNeoMessage
	(*PortMetadataNeoMessage)(nil),          // 4: com.pojtinger.felicitas.liwasc.neo.PortMetadataNeoMessage
	(*empty.Empty)(nil),                     // 5: google.protobuf.Empty
}
var file_metadata_neo_proto_depIdxs = []int32{
	5, // 0: com.pojtinger.felicitas.liwasc.neo.MetadataNeoService.GetMetadataForScanner:input_type -> google.protobuf.Empty
	1, // 1: com.pojtinger.felicitas.liwasc.neo.MetadataNeoService.GetMetadataForNode:input_type -> com.pojtinger.felicitas.liwasc.neo.NodeMetadataReferenceNeoMessage
	3, // 2: com.pojtinger.felicitas.liwasc.neo.MetadataNeoService.GetMetadataForPort:input_type -> com.pojtinger.felicitas.liwasc.neo.PortMetadataReferenceNeoMessage
	0, // 3: com.pojtinger.felicitas.liwasc.neo.MetadataNeoService.GetMetadataForScanner:output_type -> com.pojtinger.felicitas.liwasc.neo.ScannerMetadataNeoMessage
	2, // 4: com.pojtinger.felicitas.liwasc.neo.MetadataNeoService.GetMetadataForNode:output_type -> com.pojtinger.felicitas.liwasc.neo.NodeMetadataNeoMessage
	4, // 5: com.pojtinger.felicitas.liwasc.neo.MetadataNeoService.GetMetadataForPort:output_type -> com.pojtinger.felicitas.liwasc.neo.PortMetadataNeoMessage
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_metadata_neo_proto_init() }
func file_metadata_neo_proto_init() {
	if File_metadata_neo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_neo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScannerMetadataNeoMessage); i {
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
		file_metadata_neo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMetadataReferenceNeoMessage); i {
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
		file_metadata_neo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMetadataNeoMessage); i {
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
		file_metadata_neo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMetadataReferenceNeoMessage); i {
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
		file_metadata_neo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMetadataNeoMessage); i {
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
			RawDescriptor: file_metadata_neo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metadata_neo_proto_goTypes,
		DependencyIndexes: file_metadata_neo_proto_depIdxs,
		MessageInfos:      file_metadata_neo_proto_msgTypes,
	}.Build()
	File_metadata_neo_proto = out.File
	file_metadata_neo_proto_rawDesc = nil
	file_metadata_neo_proto_goTypes = nil
	file_metadata_neo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetadataNeoServiceClient is the client API for MetadataNeoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataNeoServiceClient interface {
	GetMetadataForScanner(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ScannerMetadataNeoMessage, error)
	GetMetadataForNode(ctx context.Context, in *NodeMetadataReferenceNeoMessage, opts ...grpc.CallOption) (*NodeMetadataNeoMessage, error)
	GetMetadataForPort(ctx context.Context, in *PortMetadataReferenceNeoMessage, opts ...grpc.CallOption) (*PortMetadataNeoMessage, error)
}

type metadataNeoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataNeoServiceClient(cc grpc.ClientConnInterface) MetadataNeoServiceClient {
	return &metadataNeoServiceClient{cc}
}

func (c *metadataNeoServiceClient) GetMetadataForScanner(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ScannerMetadataNeoMessage, error) {
	out := new(ScannerMetadataNeoMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.neo.MetadataNeoService/GetMetadataForScanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataNeoServiceClient) GetMetadataForNode(ctx context.Context, in *NodeMetadataReferenceNeoMessage, opts ...grpc.CallOption) (*NodeMetadataNeoMessage, error) {
	out := new(NodeMetadataNeoMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.neo.MetadataNeoService/GetMetadataForNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataNeoServiceClient) GetMetadataForPort(ctx context.Context, in *PortMetadataReferenceNeoMessage, opts ...grpc.CallOption) (*PortMetadataNeoMessage, error) {
	out := new(PortMetadataNeoMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felicitas.liwasc.neo.MetadataNeoService/GetMetadataForPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataNeoServiceServer is the server API for MetadataNeoService service.
type MetadataNeoServiceServer interface {
	GetMetadataForScanner(context.Context, *empty.Empty) (*ScannerMetadataNeoMessage, error)
	GetMetadataForNode(context.Context, *NodeMetadataReferenceNeoMessage) (*NodeMetadataNeoMessage, error)
	GetMetadataForPort(context.Context, *PortMetadataReferenceNeoMessage) (*PortMetadataNeoMessage, error)
}

// UnimplementedMetadataNeoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetadataNeoServiceServer struct {
}

func (*UnimplementedMetadataNeoServiceServer) GetMetadataForScanner(context.Context, *empty.Empty) (*ScannerMetadataNeoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadataForScanner not implemented")
}
func (*UnimplementedMetadataNeoServiceServer) GetMetadataForNode(context.Context, *NodeMetadataReferenceNeoMessage) (*NodeMetadataNeoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadataForNode not implemented")
}
func (*UnimplementedMetadataNeoServiceServer) GetMetadataForPort(context.Context, *PortMetadataReferenceNeoMessage) (*PortMetadataNeoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadataForPort not implemented")
}

func RegisterMetadataNeoServiceServer(s *grpc.Server, srv MetadataNeoServiceServer) {
	s.RegisterService(&_MetadataNeoService_serviceDesc, srv)
}

func _MetadataNeoService_GetMetadataForScanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataNeoServiceServer).GetMetadataForScanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.neo.MetadataNeoService/GetMetadataForScanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataNeoServiceServer).GetMetadataForScanner(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataNeoService_GetMetadataForNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeMetadataReferenceNeoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataNeoServiceServer).GetMetadataForNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.neo.MetadataNeoService/GetMetadataForNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataNeoServiceServer).GetMetadataForNode(ctx, req.(*NodeMetadataReferenceNeoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataNeoService_GetMetadataForPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortMetadataReferenceNeoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataNeoServiceServer).GetMetadataForPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felicitas.liwasc.neo.MetadataNeoService/GetMetadataForPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataNeoServiceServer).GetMetadataForPort(ctx, req.(*PortMetadataReferenceNeoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetadataNeoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.pojtinger.felicitas.liwasc.neo.MetadataNeoService",
	HandlerType: (*MetadataNeoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetadataForScanner",
			Handler:    _MetadataNeoService_GetMetadataForScanner_Handler,
		},
		{
			MethodName: "GetMetadataForNode",
			Handler:    _MetadataNeoService_GetMetadataForNode_Handler,
		},
		{
			MethodName: "GetMetadataForPort",
			Handler:    _MetadataNeoService_GetMetadataForPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata_neo.proto",
}