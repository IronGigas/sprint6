package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "Server: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Printf("Starting server the server on http://localhost:8080/")
	if err := srv.HttpServer.ListenAndServe(); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}

}
