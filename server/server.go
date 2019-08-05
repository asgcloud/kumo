package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/asgcloud/kumo/storage"
	"github.com/gorilla/mux"
)

const (
	devPort = "3000" // move to env variable
)

// Server is a struct for a web server
type Server struct {
	svc    storage.Service
	router *mux.Router
	stats  *UptimeStats
}

// NewServer creates a server and attaches the db and router
func NewServer(db *storage.Service, router *mux.Router) *Server {
	server := Server{}
	server.svc = *db
	server.router = router
	server.stats = NewUptimeStats()
	return &server
}

// AttachRoutes attaches the routes to the server's router
func (s *Server) AttachRoutes() {
	s.router.HandleFunc("/status", s.HandleStatus())
	s.router.HandleFunc("/servers", s.HandleListServers())
	s.router.HandleFunc("/servers/{serverName}", s.HandleServerByName())
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
		s.stats.Update("requests", "operations")
		log.Printf("Request to status received")
		s.stats.Update("responses")
		data, _ := json.Marshal(s.stats)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		s.stats.Update("operations")
		return
	}
}

// HandleListServers lists all servers in the database
func (s *Server) HandleListServers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.stats.Update("requests", "operations")
		log.Printf("Request to list servers received")
		servers, err := s.svc.FindAllServers(context.Background(), 0, 10)
		if err != nil {
			//TODO better
			panic(err)
		}
		data, _ := json.Marshal(servers)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		s.stats.Update("responses", "operations")
	}
}

// HandleServerByName queries the servers by server name
func (s *Server) HandleServerByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.stats.Update("requests", "operations")
		vars := mux.Vars(r)
		serverName := vars["serverName"]
		log.Printf("Request to find %s", serverName)
		server, err := s.svc.FindServerByName(context.Background(), serverName)
		if err != nil {
			//TODO better
			panic(err)
		}
		data, _ := json.Marshal(server)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		s.stats.Update("responses", "operations")
	}
}
