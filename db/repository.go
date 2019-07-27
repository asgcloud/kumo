package db

import (
	"context"

	"github.com/asgcloud/kumo/schema"
)

// Repository is an interface for a server repository
type Repository interface {
	Close()
	InsertServer(ctx context.Context, server schema.Server) error
	ListServers(ctx context.Context, uint64 skip, uint64 take) ([]schema.Server, error)
}
