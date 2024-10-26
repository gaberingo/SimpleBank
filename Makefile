# Create and Run Postgres using Docker
runpostgres:
	docker run --name db_simple_bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rootpass -d postgres

# Create Database in db_simple_bank
createdb:
	docker exec -it db_simple_bank createdb --username=root --owner=root simple_bank

# Drop database
dropdb:
	docker exec -it db_simple_bank dropdb simple_bank

# Migrate all migration
migrateup:
	migrate --path db/migration --database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" --verbose down

migrateup_last:
	migrate --path db/migration --database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" --verbose up 1

migratedown_last:
	migrate --path db/migration --database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" --verbose down 1

# Migrate the last one migration
migrateup_last:
	migrate --path db/migration --database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" --verbose up 1

migratedown_last:
	migrate --path db/migration --database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" --verbose down 1

# Generate sqlc
sqlc:
	sqlc generate

# Do all Go test
test:
	go test -v -cover ./...

# Run Server
server:
	go run main.go

# Create Mock
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/gaberingo/SimpleBank/db/sqlc Store

.PHONY: runpostgres createdb dropdb migratedown migrateup migratedown_last migrateup_last sqlc test server mock

