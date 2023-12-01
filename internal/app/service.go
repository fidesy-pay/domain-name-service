package app

import (
	"context"
	desc "github.com/fidesy-pay/port-resolver-service/pkg/port-resolver-service"
)

type (
	PortResolverService interface {
		GetPort(ctx context.Context, serviceName string) (int64, error)
		UpdatePort(ctx context.Context, serviceName string, port int64) error
	}
)
type Implementation struct {
	desc.UnimplementedPortResolverServiceServer

	portResolverService PortResolverService
}

func New(portResolverService PortResolverService) *Implementation {
	return &Implementation{
		portResolverService: portResolverService,
	}
}
