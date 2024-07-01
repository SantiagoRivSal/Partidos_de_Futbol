package main

import (
	"backend/app"
	"backend/db"
	"log"
	"net/http"
	"os"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // Puerto por defecto si no está configurado
	}

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
