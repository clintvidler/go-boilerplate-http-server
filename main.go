package main

import (
	"os"

	"github.com/clintvidler/go-boilerplate-http-server/server"
	"github.com/clintvidler/go-boilerplate-http-server/services"
)

func main() {
	logger := services.NewLogger()

	server := server.NewServer(logger)

	server.Serve(os.Getenv("ADDR"))
}
