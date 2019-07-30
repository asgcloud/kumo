package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/asgcloud/kumo/schema"
)

// PostgresRepository is a structure for a prostgres sql DBMS
type PostgresRepository struct {
	DB *sql.DB
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
	r.DB.Close()
}

// InsertServer adds a new server to the database
func (r *PostgresRepository) InsertServer(ctx context.Context, server schema.Server) error {
	_, err := r.DB.Query("INSERT INTO servers (server_id, server_name) VALUES ($1, $2)", server.ServerID, server.ServerName)
	if err != nil {
		return err
	}
	return nil
}

// ListServers lists all servers
func (r *PostgresRepository) ListServers(ctx context.Context, skip uint64, take uint64) ([]schema.Server, error) {
	servers := []schema.Server{}
	//rows, err := r.DB.Query("SELECT * FROM servers ORDER BY server_id DESC OFFSET $2 LIMIT $3", skip, take)
	rows, err := r.DB.Query("SELECT * FROM servers ORDER BY server_name")
	if err != nil {
		fmt.Println("ERROR!!!!!")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		server := schema.Server{}
		err := rows.Scan(&server.ServerID, &server.ProjectID, &server.ServerName, &server.CPU, &server.RAM, &server.Storage, &server.Status, &server.State, &server.Tenancy, &server.Host)
		if err != nil {
			return nil, err
		}
		servers = append(servers, server)
	}
	return servers, nil
}
