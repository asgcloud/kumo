package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/asgcloud/kumo/db"
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

	fmt.Println("Successfully connected to the database")

	servers, err := db.ListServers(context.Background(), 0, 10)
	if err != nil {
		log.Fatalf("Could not query the server list: %v", err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
