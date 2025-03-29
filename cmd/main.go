package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	theLog := log.New(os.Stdout, "Server: ", log.LstdFlags)

	theServer := server.NewServer(theLog)

	theLog.Printf("Starting server the server on http://localhost:8080/")
	if err := theServer.HttpServer.ListenAndServe(); err != nil {
		theLog.Fatalf("Failed to start server: %v", err)
	}

}
