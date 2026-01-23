package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
)

// 1. Embed migrations
//go:embed db/migrations/*.sql
var migrationEmbeddedFS embed.FS

func main() {
	ctx := context.Background()

	// 2. Database Connection String
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

	// 3. Run Migrations
	fmt.Println("Running database migrations...")
	if err := runMigrations(ctx, conn); err != nil {
		log.Fatalf("CRITICAL: Migration failed: %v", err)
	}
	fmt.Println("Migrations successful.")

	// 4. Start HTTP Server (Prevents the 'all goroutines are asleep' deadlock)
	// This keeps the container running and provides a health check for Kubernetes.
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	port := ":8080"
	fmt.Printf("Asset-service listening on %s\n", port)

	// This function blocks forever, keeping the pod alive.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("CRITICAL: Server failed: %v", err)
	}
}

func runMigrations(ctx context.Context, conn *pgx.Conn) error {
	migrator, err := migrate.NewMigrator(ctx, conn, "public.schema_version")
	if err != nil {
		return fmt.Errorf("could not create migrator: %w", err)
	}

	// Zoom into the specific folder in the embedded FS
	migrationRoot, err := fs.Sub(migrationEmbeddedFS, "db/migrations")
	if err != nil {
		return fmt.Errorf("failed to create sub-filesystem: %w", err)
	}

	// Load and Migrate
	err = migrator.LoadMigrations(migrationRoot)
	if err != nil {
		return fmt.Errorf("failed to load migrations: %w", err)
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("migration execution failed: %w", err)
	}

	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		return fmt.Errorf("failed to get current version: %w", err)
	}
	fmt.Printf("Database updated to version: %d\n", ver)

	return nil
}
