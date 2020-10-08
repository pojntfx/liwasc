package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pojntfx/liwasc/pkg/networking"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/validators"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	AUTHORIZATION_METADATA_KEY = "X-Liwasc-Authorization"
)

type MetadataService struct {
	proto.UnimplementedMetadataServiceServer

	subnets            []string
	device             string
	interfaceInspector *networking.InterfaceInspector
	contextValidator   *validators.ContextValidator
}

func NewMetadataService(interfaceInspector *networking.InterfaceInspector, contextValidator *validators.ContextValidator) *MetadataService {
	return &MetadataService{interfaceInspector: interfaceInspector, contextValidator: contextValidator}
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

func (s *MetadataService) GetMetadata(ctx context.Context, _ *empty.Empty) (*proto.MetadataMessage, error) {
	// Authorize
	valid, err := s.contextValidator.Validate(ctx)
	if err != nil || !valid {
		return nil, status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	protoMetadataMessage := &proto.MetadataMessage{
		Subnets: s.subnets,
		Device:  s.device,
	}

	return protoMetadataMessage, nil
}
