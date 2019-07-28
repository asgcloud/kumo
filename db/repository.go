package db

import (
	"context"

	"github.com/asgcloud/kumo/schema"
)

// Repository is an interface for a server repository
type Repository interface {
	Close()
	InsertServer(ctx context.Context, server schema.Server) error
	ListServers(ctx context.Context, skip uint64, take uint64) ([]schema.Server, error)
}

var impl Repository

// SetRepository sets the repository
func SetRepository(repository Repository) {
	impl = repository
}

// Close closes the implementation
func Close() {
	impl.Close()
}

// InsertServer inserts a server into the database
func InsertServer(ctx context.Context, server schema.Server) error {
	return impl.InsertServer(ctx, server)
}

// ListServers lists all servers in the repository
func ListServers(ctx context.Context, skip uint64, take uint64) ([]schema.Server, error) {
	return impl.ListServers(ctx, skip, take)
}
