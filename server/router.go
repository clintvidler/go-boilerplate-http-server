package server

import (
	"github.com/clintvidler/go-boilerplate-http-server/handlers"
	"github.com/gorilla/mux"
)

func (s *Server) setupRoutes() {
	r := mux.NewRouter().StrictSlash(true)

	defaultHandler := handlers.NewDefaultHandler(s.logger)

	r.HandleFunc("/", defaultHandler.Placeholder)

	s.router = r
}
