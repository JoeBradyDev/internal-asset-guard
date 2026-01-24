package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	// 1. Connection
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGDATABASE"),
	)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("CRITICAL: Unable to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	// 2. Orchestrate (Calling functions from migrate.go and seed.go)
	fmt.Println("Starting database setup...")

	if err := runMigrations(ctx, conn); err != nil {
		log.Fatalf("CRITICAL: Migration failed: %v", err)
	}

	seedCount, _ := strconv.Atoi(os.Getenv("SEED_ASSET_COUNT"))
	if seedCount > 0 {
		if err := seedRealisticAssets(ctx, conn, seedCount); err != nil {
			log.Printf("Warning: Seeding failed: %v", err)
		}
	}

	// 3. Keep Alive
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	port := ":8080"
	fmt.Printf("Asset-service listening on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("CRITICAL: Server failed: %v", err)
	}
}
