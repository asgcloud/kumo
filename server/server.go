package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/asgcloud/kumo/db"
	"github.com/gorilla/mux"
)

const (
	devPort = "3000" // move to env variable
)

// Server is a struct for a web server
type Server struct {
	db     *db.PostgresRepository
	router *mux.Router
	stats  UptimeStats
}

// NewServer creates a server and attaches the db and router
func NewServer(db *db.PostgresRepository, router *mux.Router) *Server {
	server := Server{}
	server.db = db
	server.router = router
	server.stats = UptimeStats{
		Status:              "OK",
		StartTime:           time.Now().String(),
		RequestsReceived:    0,
		ResponsesProvided:   0,
		OperationsCompleted: 0,
	}
	return &server
}

// AttachRoutes attaches the routes to the server's router
func (s *Server) AttachRoutes() {
	s.router.HandleFunc("/status", s.HandleStatus())
	s.router.HandleFunc("/servers", s.HandleListServers())
}

// Run starts the server
func (s *Server) Run() {
	if err := http.ListenAndServe(":"+devPort, s.router); err != nil {
		log.Fatalf("Server has stopped working: %v", err)
	}
}

// HandleStatus returns statistics on the health and activity of the api
func (s *Server) HandleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.stats.RequestsReceived++
		s.stats.OperationsCompleted++
		log.Printf("Request to status received")
		data, _ := json.Marshal(s.stats)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		s.stats.ResponsesProvided++
		s.stats.OperationsCompleted++
		return
	}
}

// HandleListServers lists all servers in the database
func (s *Server) HandleListServers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.stats.RequestsReceived++
		s.stats.OperationsCompleted++
		log.Printf("Request to list servers received")
		servers, err := s.db.ListServers(context.Background(), 0, 10)
		if err != nil {
			//TODO better
			panic(err)
		}
		data, _ := json.Marshal(servers)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		s.stats.ResponsesProvided++
		s.stats.OperationsCompleted++
		return
	}
}
