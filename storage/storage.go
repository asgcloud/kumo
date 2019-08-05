package storage

import (
	"context"

	"github.com/asgcloud/kumo/schema"
)

// Service is an interface for a server repository
type Service interface {
	Close() error
	InsertServer(ctx context.Context, server schema.Server) error
	FindAllServers(ctx context.Context, skip uint64, take uint64) ([]schema.Server, error)
	FindServerByName(ctx context.Context, name string) (schema.Server, error)
}
