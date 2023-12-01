package app

import (
	"context"
	desc "github.com/fidesy-pay/port-resolver-service/pkg/port-resolver-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdatePort(ctx context.Context, req *desc.UpdatePortRequest) (*desc.UpdatePortResponse, error) {
	err := i.portResolverService.UpdatePort(ctx, req.GetServiceName(), req.GetPort())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "portResolverService.UpdatePort: %v", err)
	}

	return &desc.UpdatePortResponse{}, nil
}
