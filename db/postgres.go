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

// NewPostgres connects to a new session of postgres
func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

// Close closes the database connection
func (r *PostgresRepository) Close() {
	r.db.Close()
}

// InsertServer adds a new server to the database
func (r *PostgresRepository) InsertServer(ctx context.Context, server schema.Server) error {
	_, err := r.db.Query("INSERT INTO servers (server_id, server_name) VALUES ($1, $2)", server.ID, server.Name)
	if err != nil {
		return err
	}
	return nil
}

// ListServers lists all servers
func (r *PostgresRepository) ListServers(ctx context.Context, skip uint64, take uint64) ([]schema.Server, error) {
	servers := []schema.Server{}
	rows, err := r.db.Query("SELECT * FROM servers ORDER BY server_id DESC OFFSET $1 LIMIT $2", skip, take)
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
