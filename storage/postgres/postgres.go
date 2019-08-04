package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/asgcloud/kumo/schema"
	"github.com/asgcloud/kumo/storage"
)

type postgres struct{ db *sql.DB }

// New connects to a new session of postgres
func New(host, port, user, password, dbName string) (storage.Service, error) {

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgres{db}, nil
}

// Close closes the database connection
func (p *postgres) Close() {
	p.db.Close()
}

// InsertServer adds a new server to the database
func (p *postgres) InsertServer(ctx context.Context, server schema.Server) error {
	_, err := p.db.Query("INSERT INTO servers (server_id, server_name) VALUES ($1, $2)", server.ServerID, server.ServerName)
	if err != nil {
		return err
	}
	return nil
}

// ListServers lists all servers
func (p *postgres) ListServers(ctx context.Context, skip uint64, take uint64) ([]schema.Server, error) {
	servers := []schema.Server{}
	//rows, err := p.db.Query("SELECT * FROM servers ORDER BY server_id DESC OFFSET $2 LIMIT $3", skip, take)
	rows, err := p.db.Query("SELECT * FROM servers ORDER BY server_name")
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
