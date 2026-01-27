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
  "asset-service/internal/service/pb"
)

func main() {
	ctx := context.Background()

	// 1. DATABASE CONNECTION
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
	fmt.Println("Checking database status and running migrations...")
	if err := runMigrations(ctx, conn); err != nil {
		log.Fatalf("CRITICAL: Migration failed: %v", err)
	}

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
	queries := db.New(conn)
	assetServer := &service.AssetServer{
		Queries: queries,
	}

	// 4. START gRPC SERVER
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("CRITICAL: Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAssetServiceServer(grpcServer, assetServer)

	// Enable reflection for grpcurl and debugging
	reflection.Register(grpcServer)

	fmt.Printf("✅ Asset Service is live on port %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("CRITICAL: Server exited with error: %v", err)
	}
}
