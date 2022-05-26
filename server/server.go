package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/clintvidler/go-boilerplate-http-server/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	logger *services.Logger
	router *mux.Router
}

func NewServer(l *services.Logger) *Server {
	return &Server{logger: l}
}

func (s *Server) Serve(addr string) {
	s.setupRoutes()

	cors := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	server := &http.Server{
		Handler:      cors(services.LogRequestResponse(s.router, *s.logger)),
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		log.Println("Starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm/interupt and gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until a signal is received
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting a time for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
