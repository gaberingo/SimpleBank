runpostgres:
	docker run --name db_simple_bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rootpass -d postgres

createdb:
	docker exec -it db_simple_bank createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it db_simple_bank dropdb simple_bank

migrateup:
	migrate --path db/migration --database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" --verbode down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: runpostgres createdb dropdb migratedown migrateup sqlc test server

