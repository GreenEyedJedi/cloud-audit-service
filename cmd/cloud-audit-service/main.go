package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GreenEyedJedi/cloud-audit-service/internal/config"
	"github.com/GreenEyedJedi/cloud-audit-service/internal/handlers"
)

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handlers.HealthzHandler)

	address := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on %s...", address)

	err = http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
