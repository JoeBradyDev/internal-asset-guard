package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
)

//go:embed db/migrations/*.sql
var migrationEmbeddedFS embed.FS

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
