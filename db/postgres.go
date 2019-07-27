package db

import (
	"context"
	"database/sql"

	"github.com/asgcloud/kumo/schema"
)

// PostgresRepository is a structure for a prostgres sql DBMS
type PostgresRepository struct {
	db *sql.DB
}

// Close closes the database connection
func (r *PostgresRepository) Close() {
	r.db.Close()
}

// InsertServer adds a new server to the database
func (r *PostgresRepository) InsertServer(ctx context.Context, server schema.Server) error {
	_, err := r.db.Query("INSERT INTO servers (id, name) VALUES ($1, $2)", server.ID, server.Name)
	if err != nil {
		return err
	}
	return nil
}

// ListServers lists all servers
func (r *PostgresRepository) ListServers(ctx context.Context, uint64 skip, uint64 take) ([]schema.Server, error) {
	servers := []schema.Server{}
	rows, err := r.db.Query("SELECT * FROM servers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		server := schema.Server{}
		err := rows.Scan(&server.ID, &server.Name)
		if err != nil {
			return nil, err
		}
		servers = append(servers, server)
	}
	return servers, nil
}
