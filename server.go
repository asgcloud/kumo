package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

const (
	serverPort = 8080
)

// Server is a struct for a web server
type Server struct {
	db     *sql.DB
	router mux.Router
}
