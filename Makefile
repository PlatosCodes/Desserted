DB_URL=postgresql://root:bluecomet@localhost:5432/desserted?sslmode=disable

network:
	docker network create desserted-network
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=bluecomet -d postgres:15-alpine
	
createdb: 
	docker exec -it postgres createdb --username=root --owner=root desserted

dropdb:
	docker exec -it postgres dropdb --username=root desserted

citext:
	docker exec -it postgres psql --username=root desserted -c "CREATE EXTENSION IF NOT EXISTS citext;"

migrateup:
	migrate -path backend/db/migration -database "${DB_URL}" -verbose up

migrateup1:
	migrate -path backend/db/migration -database "${DB_URL}" -verbose up 1

migratedown: 
	migrate -path backend/db/migration -database "${DB_URL}" -verbose down

migratedown1: 
	migrate -path backend/db/migration -database "${DB_URL}" -verbose down 1

new_migration:
	migrate create -ext sql -dir backend/db/migration -seq $(name)

db_docs:
	dbdocs build doc/database.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/database.dbml

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb citext migrateup migrateup1 migratedown migratedown1 new_migration db_docs db_schema sqlc server
