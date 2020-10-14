package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pojntfx/liwasc/pkg/databases"
	"github.com/pojntfx/liwasc/pkg/networking"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
)

type MetadataNeoService struct {
	proto.UnimplementedMetadataNeoServiceServer

	subnets []string
	device  string

	interfaceInspector              *networking.InterfaceInspector
	mac2vendorDatabase              *databases.MAC2VendorDatabase
	serviceNamesPortNumbersDatabase *databases.ServiceNamesPortNumbersDatabase
}

func NewMetadataNeoService(
	interfaceInspector *networking.InterfaceInspector,
	mac2vendorDatabase *databases.MAC2VendorDatabase,
	serviceNamesPortNumbersDatabase *databases.ServiceNamesPortNumbersDatabase,
) *MetadataNeoService {
	return &MetadataNeoService{
		interfaceInspector:              interfaceInspector,
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
	protoScannerMetadataNeoMessage := &proto.ScannerMetadataNeoMessage{
		Subnets: s.subnets,
		Device:  s.device,
	}

	return protoScannerMetadataNeoMessage, nil
}
