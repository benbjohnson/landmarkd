package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

type Server struct {
	router     *mux.Router
	httpServer *http.Server
}

// Creates a new server.
func New() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

// Starts the server.
func (s *Server) ListenAndServe() error {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", DefaultPort),
		Handler: s.router,
	}
	s.router.HandleFunc("/track", s.trackHandler).Methods("GET")
	s.router.HandleFunc("/track.gif", s.trackHandler).Methods("GET")

	log.Printf("Running at http://localhost%s/", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Tracks a single event to a project.
func (s *Server) trackHandler(w http.ResponseWriter, r *http.Request) error {
	// TODO: Parse query parameters.
	// TODO: Lookup project by API Key.
	// TODO: Send event against project.
	return nil
}
