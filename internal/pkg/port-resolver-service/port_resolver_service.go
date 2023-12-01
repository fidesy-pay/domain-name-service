package port_resolver_service

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrNotFound = errors.New("record with provided serviceName not found")
)

type Storage interface {
	Get(ctx context.Context, key string, dst interface{}) (bool, error)
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetPort(ctx context.Context, serviceName string) (int64, error) {
	var port int64
	found, err := s.storage.Get(ctx, serviceName, &port)
	if err != nil {
		return 0, fmt.Errorf("storage.Get: %w", err)
	}

	if !found {
		return 0, ErrNotFound
	}

	return port, nil
}

func (s *Service) UpdatePort(ctx context.Context, serviceName string, port int64) error {
	err := s.storage.Set(ctx, serviceName, port, 0)
	if err != nil {
		return fmt.Errorf("storage.Set: %w", err)
	}

	return nil
}
