package main

import (
	"context"
	"log"

	"github.com/Frely25/Verum/internal/core/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgxpool.New(
		ctx,
		cfg.DatabaseURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("PostgreSQL connected")
}

// PATCH 41.51.125.63:8080/classes/5
