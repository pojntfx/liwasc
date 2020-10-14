package services

import (
	"context"
	"fmt"
	"log"
	"strconv"

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

func (s *MetadataNeoService) GetMetadataForNode(_ context.Context, nodeMetadataReferenceMessage *proto.NodeMetadataReferenceNeoMessage) (*proto.NodeMetadataNeoMessage, error) {
	dbNodeMetadata, err := s.mac2vendorDatabase.GetVendor(nodeMetadataReferenceMessage.GetMACAddress())
	if err != nil {
		log.Printf("could not find node %v in DB: %v\n", nodeMetadataReferenceMessage.GetMACAddress(), err)

		return nil, status.Errorf(codes.NotFound, "could not find node in DB")
	}

	protoNodeMetadataMessage := &proto.NodeMetadataNeoMessage{
		Address:      dbNodeMetadata.Address.String,
		MACAddress:   nodeMetadataReferenceMessage.GetMACAddress(),
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

func (s *MetadataNeoService) GetMetadataForPort(_ context.Context, portMetadataReferenceMessage *proto.PortMetadataReferenceNeoMessage) (*proto.PortMetadataNeoMessage, error) {
	dbPortMetadata, err := s.serviceNamesPortNumbersDatabase.GetService(int(portMetadataReferenceMessage.GetPortNumber()), portMetadataReferenceMessage.GetTransportProtocol())
	if err != nil || (dbPortMetadata != nil && len(dbPortMetadata) == 0) {
		log.Printf("could not find port %v in DB: %v\n", fmt.Sprintf("%v/%v", portMetadataReferenceMessage.GetPortNumber(), portMetadataReferenceMessage.GetTransportProtocol()), err)

		return nil, status.Errorf(codes.NotFound, "could not find port in DB")
	}

	portNumber, err := strconv.Atoi(dbPortMetadata[0].PortNumber)
	if err != nil {
		log.Printf("could not find valid port number for port %v in DB: %v\n", fmt.Sprintf("%v/%v", portMetadataReferenceMessage.GetPortNumber(), portMetadataReferenceMessage.GetTransportProtocol()), err)

		return nil, status.Errorf(codes.Unknown, "could not find valid port number in DB")
	}

	protoPortMetadataMessage := &proto.PortMetadataNeoMessage{
		Assignee:                dbPortMetadata[0].Assignee,
		AssignmentNotes:         dbPortMetadata[0].AssignmentNotes,
		Contact:                 dbPortMetadata[0].Contact,
		Description:             dbPortMetadata[0].Description,
		ModificationDate:        dbPortMetadata[0].ModificationDate,
		PortNumber:              int64(portNumber),
		Reference:               dbPortMetadata[0].Reference,
		RegistrationDate:        dbPortMetadata[0].RegistrationDate,
		ServiceCode:             dbPortMetadata[0].ServiceCode,
		ServiceName:             dbPortMetadata[0].ServiceName,
		TransportProtocol:       dbPortMetadata[0].TransportProtocol,
		UnauthorizedUseReported: dbPortMetadata[0].UnauthorizedUseReported,
	}

	return protoPortMetadataMessage, nil
}
