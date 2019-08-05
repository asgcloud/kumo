package postgres

import (
	"context"
	"errors"

	// This loads the postgres drivers.
	pg "github.com/go-pg/pg"
	_ "github.com/lib/pq"

	"github.com/asgcloud/kumo/schema"
	"github.com/asgcloud/kumo/storage"
)

type postgres struct{ db *pg.DB }

// New connects to a new session of postgres
func New(host, port, user, password, dbName string) (storage.Service, error) {

	opts := &pg.Options{
		User:     user,
		Password: password,
		Addr:     host + ":" + port,
		Database: dbName,
	}

	db := pg.Connect(opts)
	if db == nil {
		return nil, errors.New("failed to connect to database")
	}

	return &postgres{db}, nil
}

// Close closes the database connection
func (p *postgres) Close() error {
	err := p.db.Close()
	if err != nil {
		return err
	}
	return nil
}

// InsertServer adds a new server to the database
func (p *postgres) InsertServer(ctx context.Context, server schema.Server) error {
	return p.db.Insert(server)
}

// FindAllServers returns all servers
func (p *postgres) FindAllServers(ctx context.Context, skip uint64, take uint64) ([]schema.Server, error) {

	var servers []schema.Server

	err := p.db.Model(&servers).Select()
	if err != nil {
		return nil, err
	}

	return servers, nil
}

// FindServerByName queries the database for an individual server by its name
func (p *postgres) FindServerByName(ctx context.Context, name string) (schema.Server, error) {
	server := &schema.Server{}
	err := p.db.Model(server).Where("server_name = ?", name).Select()
	if err != nil {
		return schema.Server{}, err
	}
	return *server, nil
}
