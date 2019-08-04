package main

import (
	"flag"
	"log"

	"github.com/gorilla/mux"

	"github.com/asgcloud/kumo/config"
	"github.com/asgcloud/kumo/server"
	"github.com/asgcloud/kumo/storage/postgres"
)

func main() {
	configPath := flag.String("config", "./config/config.json", "path of the config file")

	flag.Parse()

	// Read config
	config, err := config.FromFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	svc, err := postgres.New(
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.DB)
	if err != nil {
		log.Fatalf("Could not connect to postgres DBMS: %v", err)
	}
	defer svc.Close()

	log.Println("Successfully connected to the database")
	mux := mux.NewRouter()

	server := server.NewServer(&svc, mux)
	server.AttachRoutes()
	server.Run()
}
