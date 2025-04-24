package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appConfig "github.com/GreenEyedJedi/cloud-audit-service/internal/config"

	"github.com/GreenEyedJedi/cloud-audit-service/internal/handlers"
)

func main() {
	// Load app config from JSON
	cfg, err := appConfig.Load("config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Load AWS SDK config
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(cfg.AWSRegion))
	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
	}

	// Get AWS client and handler
	s3Client := s3.NewFromConfig(awsCfg)
	s3Handler := handlers.S3Handler{Client: s3Client}

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handlers.HealthzHandler)
	mux.HandleFunc("/s3", s3Handler.ListBuckets)

	// Start server
	address := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on %s...", address)

	err = http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
