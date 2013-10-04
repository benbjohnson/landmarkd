package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
func (s *Server) ListenAndServe(port int) error {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.router,
	}
	s.router.HandleFunc("/track", s.trackHandler).Methods("GET")
	s.router.HandleFunc("/track.gif", s.trackHandler).Methods("GET")

	log.Printf("Running at http://localhost%s/", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Tracks a single event to a project.
func (s *Server) trackHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse query parameters.
	// TODO: Lookup project by API Key.
	// TODO: Send event against project.
}
