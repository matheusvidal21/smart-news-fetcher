package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/matheusvidal21/smart-news-fetcher/pkg/logger"
	"log"
)

func main() {
	if err := logger.InitializeLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.CloseLogger()

	server, err := NewServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}
	defer server.DB.Close()

	server.InitializeRoutes()

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
