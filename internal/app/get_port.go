package app

import (
	"context"
	"errors"
	port_resolver_service "github.com/fidesy-pay/port-resolver-service/internal/pkg/port-resolver-service"
	desc "github.com/fidesy-pay/port-resolver-service/pkg/port-resolver-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetPort(ctx context.Context, req *desc.GetPortRequest) (*desc.GetPortResponse, error) {
	port, err := i.portResolverService.GetPort(ctx, req.GetServiceName())
	if err != nil {
		if errors.Is(err, port_resolver_service.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "portResolverService.GetPort: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "portResolverService.GetPort: %v", err)
	}

	return &desc.GetPortResponse{
		Port: port,
	}, nil
}
