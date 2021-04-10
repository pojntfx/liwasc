package services

//go:generate sh -c "mkdir -p ../api/proto/v1 && protoc --go_out=paths=source_relative,plugins=grpc:../api/proto/v1 -I=../../api/proto/v1 ../../api/proto/v1/*.proto"

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/pojntfx/liwasc/pkg/api/proto/v1"
	"github.com/pojntfx/liwasc/pkg/networking"
	"github.com/pojntfx/liwasc/pkg/stores"
	"github.com/pojntfx/liwasc/pkg/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	AUTHORIZATION_METADATA_KEY = "X-Liwasc-Authorization"
)

type MetadataService struct {
	api.UnimplementedMetadataServiceServer

	subnets []string
	device  string

	interfaceInspector *networking.InterfaceInspector

	mac2vendorDatabase              *stores.MAC2VendorDatabase
	serviceNamesPortNumbersDatabase *stores.ServiceNamesPortNumbersDatabase

	contextValidator *validators.ContextValidator
}

func NewMetadataService(
	interfaceInspector *networking.InterfaceInspector,

	mac2vendorDatabase *stores.MAC2VendorDatabase,
	serviceNamesPortNumbersDatabase *stores.ServiceNamesPortNumbersDatabase,

	contextValidator *validators.ContextValidator,
) *MetadataService {
	return &MetadataService{
		interfaceInspector: interfaceInspector,

		mac2vendorDatabase:              mac2vendorDatabase,
		serviceNamesPortNumbersDatabase: serviceNamesPortNumbersDatabase,

		contextValidator: contextValidator,
	}
}

func (s *MetadataService) Open() error {
	subnets, err := s.interfaceInspector.GetIPv4Subnets()
	if err != nil {
		return err
	}

	s.subnets = subnets
	s.device = s.interfaceInspector.GetDevice()

	return nil
}

func (s *MetadataService) GetMetadataForScanner(ctx context.Context, _ *empty.Empty) (*api.ScannerMetadataMessage, error) {
	// Authorize
	valid, err := s.contextValidator.Validate(ctx)
	if err != nil || !valid {
		return nil, status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	protoScannerMetadataMessage := &api.ScannerMetadataMessage{
		Subnets: s.subnets,
		Device:  s.device,
	}

	return protoScannerMetadataMessage, nil
}

func (s *MetadataService) GetMetadataForNode(ctx context.Context, nodeMetadataReferenceMessage *api.NodeMetadataReferenceMessage) (*api.NodeMetadataMessage, error) {
	// Authorize
	valid, err := s.contextValidator.Validate(ctx)
	if err != nil || !valid {
		return nil, status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	dbNodeMetadata, err := s.mac2vendorDatabase.GetVendor(nodeMetadataReferenceMessage.GetMACAddress())
	if err != nil {
		log.Printf("could not find node %v in DB: %v\n", nodeMetadataReferenceMessage.GetMACAddress(), err)

		return nil, status.Errorf(codes.NotFound, "could not find node in DB")
	}

	protoNodeMetadataMessage := &api.NodeMetadataMessage{
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

func (s *MetadataService) GetMetadataForPort(ctx context.Context, portMetadataReferenceMessage *api.PortMetadataReferenceMessage) (*api.PortMetadataMessage, error) {
	// Authorize
	valid, err := s.contextValidator.Validate(ctx)
	if err != nil || !valid {
		return nil, status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

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

	protoPortMetadataMessage := &api.PortMetadataMessage{
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
