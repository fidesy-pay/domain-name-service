package main

import (
	"context"
	"fmt"
	"github.com/fidesy-pay/domain-name-service/internal/app"
	"github.com/fidesy-pay/domain-name-service/internal/config"
	"github.com/fidesy-pay/domain-name-service/internal/pkg/domain-name-service"
	"github.com/fidesy-pay/domain-name-service/internal/pkg/redis"
	desc "github.com/fidesy-pay/domain-name-service/pkg/domain-name-service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	grpcPort string
)

func main() {
	grpcPort = os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		log.Fatalf("GRPC_PORT ENV is required")
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	defer cancel()

	err := config.Init()
	if err != nil {
		log.Fatalf("config.Init: %v", err)
	}

	storage, err := redis.New(ctx)
	if err != nil {
		log.Fatalf("redis.New: %v", err)
	}

	domainNameService := domain_name_service.New(storage)

	impl := app.New(domainNameService)

	grpcServer := grpc.NewServer()
	grpcServer.RegisterService(&desc.DomainNameService_ServiceDesc, impl)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	log.Printf("grpcServer is running at %s port", grpcPort)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve: %v", err)
	}
}
