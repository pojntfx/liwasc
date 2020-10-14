package services

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pojntfx/liwasc/pkg/databases"
	"github.com/pojntfx/liwasc/pkg/networking"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MetadataNeoService struct {
	proto.UnimplementedMetadataNeoServiceServer

	subnets []string
	device  string

	interfaceInspector *networking.InterfaceInspector

	mac2vendorDatabase              *databases.MAC2VendorDatabase
	serviceNamesPortNumbersDatabase *databases.ServiceNamesPortNumbersDatabase
}

func NewMetadataNeoService(
	interfaceInspector *networking.InterfaceInspector,

	mac2vendorDatabase *databases.MAC2VendorDatabase,
	serviceNamesPortNumbersDatabase *databases.ServiceNamesPortNumbersDatabase,
) *MetadataNeoService {
	return &MetadataNeoService{
		interfaceInspector: interfaceInspector,

		mac2vendorDatabase:              mac2vendorDatabase,
		serviceNamesPortNumbersDatabase: serviceNamesPortNumbersDatabase,
	}
}

func (s *MetadataNeoService) Open() error {
	subnets, err := s.interfaceInspector.GetIPv4Subnets()
	if err != nil {
		return err
	}

	s.subnets = subnets
	s.device = s.interfaceInspector.GetDevice()

	return nil
}

func (s *MetadataNeoService) GetMetadataForScanner(context.Context, *empty.Empty) (*proto.ScannerMetadataNeoMessage, error) {
	protoScannerMetadataMessage := &proto.ScannerMetadataNeoMessage{
		Subnets: s.subnets,
		Device:  s.device,
	}

	return protoScannerMetadataMessage, nil
}

func (s *MetadataNeoService) GetMetadataForNode(_ context.Context, nodeReferenceMessage *proto.NodeMetadataReferenceNeoMessage) (*proto.NodeMetadataNeoMessage, error) {
	dbNodeMetadata, err := s.mac2vendorDatabase.GetVendor(nodeReferenceMessage.GetMACAddress())
	if err != nil {
		log.Printf("could not find node %v in DB: %v\n", nodeReferenceMessage.GetMACAddress(), err)

		return nil, status.Errorf(codes.NotFound, "could not find node in DB")
	}

	protoNodeMetadataMessage := &proto.NodeMetadataNeoMessage{
		Address:      dbNodeMetadata.Address.String,
		MACAddress:   nodeReferenceMessage.GetMACAddress(),
		Organization: dbNodeMetadata.Organization.String,
		Registry:     dbNodeMetadata.Registry,
		Vendor:       dbNodeMetadata.Vendor.String,
		Visible: func() bool {
			if dbNodeMetadata.Visibility == 1 {
				return true
			}

			return false
		}(),
	}

	return protoNodeMetadataMessage, nil
}
