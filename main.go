package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/asgcloud/kumo/db"
	"github.com/asgcloud/kumo/server"
)

// TODO: Delete and refer to real values from environment variables
const (
	host     = "localhost"
	port     = 5432
	user     = "sqladmin"
	password = "sqlpassword"
	dbname   = "kumo"
)

func main() {
	fmt.Println("Hello World")

	postgres := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := db.NewPostgres(postgres)
	if err != nil {
		log.Fatalf("Could not connect to postgres DBMS: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to the database")
	mux := mux.NewRouter()

	server := server.NewServer(db, mux)
	server.AttachRoutes()
	server.Run()
}
