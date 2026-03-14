package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Boyeep/flippy-backend/internal/config"
	"github.com/Boyeep/flippy-backend/internal/repository"
	"github.com/Boyeep/flippy-backend/internal/server"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	db, err := repository.NewPostgresPool(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer db.Close()

	if err := repository.RunInitialSchema(ctx, db); err != nil {
		log.Fatalf("failed to run initial schema: %v", err)
	}

	srv := &http.Server{
		Addr:              cfg.Address(),
		Handler:           server.NewRouter(cfg, db),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("starting flippy-backend on %s", cfg.Address())
	log.Printf("database url configured for %s", cfg.DB.Name)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
