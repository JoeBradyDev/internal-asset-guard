package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"asset-service/internal/db"
	"asset-service/internal/service"
	"asset-service/proto"
)

func main() {
	ctx := context.Background()

	// 1. DATABASE CONNECTION
	// We use standard environment variables. Ensure these are set in your shell or docker-compose.
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("CRITICAL: Unable to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	// 2. ORCHESTRATE DATABASE SETUP
	// This ensures the schema is ready before we attempt to serve requests.
	fmt.Println("Checking database status and running migrations...")
	if err := runMigrations(ctx, conn); err != nil {
		log.Fatalf("CRITICAL: Migration failed: %v", err)
	}

	// Optional: Seed the database if SEED_ASSET_COUNT is provided
	seedCountStr := os.Getenv("SEED_ASSET_COUNT")
	if seedCountStr != "" {
		seedCount, _ := strconv.Atoi(seedCountStr)
		if seedCount > 0 {
			fmt.Printf("Seeding %d realistic assets...\n", seedCount)
			if err := seedRealisticAssets(ctx, conn, seedCount); err != nil {
				log.Printf("Warning: Seeding failed: %v", err)
			}
		}
	}

	// 3. INITIALIZE SQLC & SERVICE LAYER
	// Create the SQLC repository (Queries) and inject it into our gRPC service implementation.
	queries := db.New(conn)
	assetServer := &service.AssetServer{
		Queries: queries,
	}

	// 4. START gRPC SERVER
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("CRITICAL: Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	// Register the logic implemented across your service files (assets.go, issues.go, etc.)
	proto.RegisterAssetServiceServer(grpcServer, assetServer)

	// Enable reflection to allow for easier debugging/testing
	reflection.Register(grpcServer)

	fmt.Printf("✅ Asset Service is live on port %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("CRITICAL: Server exited with error: %v", err)
	}
}

// NOTE: Ensure your migration and seeding logic are defined in this package
// or imported if they live in internal/db/orchestrator.go.
// These functions are called above.
func runMigrations(ctx context.Context, conn *pgx.Conn) error {
	// Implementation for applying .sql files to the DB
	return nil
}

func seedRealisticAssets(ctx context.Context, conn *pgx.Conn, count int) error {
	// Implementation for initial data population
	return nil
}
