# Run migrations to update local DB
migrate:
	tern migrate -c db/migrations/tern.conf -m db/migrations/

# Generate Go code from SQL
generate:
	sqlc generate

# Clean build of the DB layer
db-build: migrate generate

# Run the actual service
run:
	go run main.g
