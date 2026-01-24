package main

import (
	"context"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5"
)

func seedRealisticAssets(ctx context.Context, conn *pgx.Conn, targetCount int) error {
	var currentCount int
	// 1. Only seed if empty
	err := conn.QueryRow(ctx, "SELECT COUNT(*) FROM asset").Scan(&currentCount)
	if err != nil {
		return err
	}
	if currentCount > 0 {
		fmt.Printf("Database already has %d assets. Skipping faker seed.\n", currentCount)
		return nil
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// 2. Pre-fetch Class IDs from migrations to avoid NULL subqueries
	var deviceClassID, networkClassID, softwareClassID int
	err = tx.QueryRow(ctx, "SELECT id FROM cis_asset_class WHERE name = 'Devices'").Scan(&deviceClassID)
	if err != nil { return fmt.Errorf("missing 'Devices' class: %w", err) }

	err = tx.QueryRow(ctx, "SELECT id FROM cis_asset_class WHERE name = 'Network'").Scan(&networkClassID)
	if err != nil { return fmt.Errorf("missing 'Network' class: %w", err) }

	err = tx.QueryRow(ctx, "SELECT id FROM cis_asset_class WHERE name = 'Software'").Scan(&softwareClassID)
	if err != nil { return fmt.Errorf("missing 'Software' class: %w", err) }

	fmt.Printf("Seeding %d assets... (IDs: Dev:%d, Net:%d, Sw:%d)\n", targetCount, deviceClassID, networkClassID, softwareClassID)

	var createdAssetIDs []int
	for i := 0; i < targetCount; i++ {
		var assetName string
		var classID int
		var isSoftware bool
		roll := gofakeit.Number(1, 100)

		// Determine Class based on distribution
		if roll <= 70 { // 70% Devices (Laptops/Servers)
			classID = deviceClassID
			assetName = fmt.Sprintf("DEV-%s-%d", gofakeit.LastName(), gofakeit.Number(1000, 9999))
		} else if roll <= 90 { // 20% Software
			classID = softwareClassID
			assetName = gofakeit.AppName()
			isSoftware = true
		} else { // 10% Network
			classID = networkClassID
			assetName = fmt.Sprintf("NET-%s-SW-%d", gofakeit.City(), i)
		}

		// 3. Insert Base Asset
		var assetID int
		err := tx.QueryRow(ctx, `
			INSERT INTO asset (name, asset_class_id, criticality_id)
			VALUES ($1, $2, (SELECT id FROM cis_asset_criticality ORDER BY RANDOM() LIMIT 1))
			RETURNING id`, assetName, classID).Scan(&assetID)
		if err != nil {
			return fmt.Errorf("failed asset insert at index %d: %w", i, err)
		}
		createdAssetIDs = append(createdAssetIDs, assetID)

		// 4. Insert Details
		if isSoftware {
			_, err = tx.Exec(ctx, `
				INSERT INTO software_detail (asset_id, name, vendor, version)
				VALUES ($1, $2, $3, $4)`,
				assetID, assetName, gofakeit.Company(), gofakeit.AppVersion())
		} else {
			_, err = tx.Exec(ctx, `
				INSERT INTO device_detail (asset_id, hostname, device_type_id, ip_address, mac_address, os_name)
				VALUES ($1, $2,
					(SELECT id FROM device_type WHERE asset_class_id = $3 LIMIT 1),
					$4, $5, $6)`,
				assetID, assetName, classID, gofakeit.IPv4Address(), gofakeit.MacAddress(),
				gofakeit.RandomString([]string{"Windows 11", "Ubuntu 22.04", "macOS Sonoma"}))
		}
		if err != nil {
			return fmt.Errorf("failed details for %s (ClassID: %d): %w", assetName, classID, err)
		}
	}

	// 5. Seed random issues for a quarter of the assets
	issueCount := targetCount / 4
	for i := 0; i < issueCount; i++ {
		randomAssetID := createdAssetIDs[gofakeit.Number(0, len(createdAssetIDs)-1)]
		_, err = tx.Exec(ctx, `
			INSERT INTO asset_issue (asset_id, issue_type_id, status_id, issue_source_id, external_issue_id, description)
			VALUES (
				$1,
				(SELECT id FROM issue_type ORDER BY RANDOM() LIMIT 1),
				(SELECT id FROM cis_issue_status ORDER BY RANDOM() LIMIT 1),
				(SELECT id FROM issue_source ORDER BY RANDOM() LIMIT 1),
				$2, $3)`,
			randomAssetID, fmt.Sprintf("VULN-%d", gofakeit.Number(100000, 999999)), gofakeit.Sentence(6))
		if err != nil {
			log.Printf("Warning: Failed to seed an issue: %v", err)
		}
	}

	return tx.Commit(ctx)
}
