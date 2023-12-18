# This Makefile is for managing the Desserted card game project

# Environment Variables
POSTGRES_USER := root
POSTGRES_PASSWORD := bluecomet
DB_NAME := desserted
TEST_DB_NAME := desserted-test
CONTAINER_NAME := postgres
CONTAINER_VOLUME_PATH := /data
LOCAL_DB_PATH := /Users/alexmerola/Developer/desserted/backend/db
DB_URL := postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable
TEST_DB_URL := postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${TEST_DB_NAME}?sslmode=disable


# Docker PostgreSQL container setup
network:
	docker network create desserted-network

postgres:
	docker run --name ${CONTAINER_NAME} -p 5432:5432 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d -v ${LOCAL_DB_PATH}:${CONTAINER_VOLUME_PATH} postgres:15-alpine

createdb: 
	docker exec -it ${CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${DB_NAME}

dropdb:
	docker exec -it ${CONTAINER_NAME} dropdb --username=${POSTGRES_USER} ${DB_NAME}

citext:
	docker exec -it ${CONTAINER_NAME} psql --username=${POSTGRES_USER} ${DB_NAME} -c "CREATE EXTENSION IF NOT EXISTS citext;"

# Database migrations
migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrateup1:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1

migratedown: 
	migrate -path db/migration -database "${DB_URL}" -verbose down

migratedown1: 
	migrate -path db/migration -database "${DB_URL}" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

# Database seeding -- CHANGE TO seed.sql before deploying
seeddb:
	docker exec -it ${CONTAINER_NAME} psql --username=${POSTGRES_USER} --dbname=${DB_NAME} -a -f ${CONTAINER_VOLUME_PATH}/seed.sql

# Change the seed file before deploying
resetdb:
	@echo "Resetting database..."
	@make dropdb || true
	@make createdb
	@make citext
	@make migrateup
	@make seeddb

# Create a test database
createdb_test:
	docker exec -it ${CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${TEST_DB_NAME}

migrateup_test:
	migrate -path db/migration -database "${TEST_DB_URL}" -verbose up

# Drop the test database
dropdb_test:
	docker exec -it ${CONTAINER_NAME} dropdb --username=${POSTGRES_USER} ${TEST_DB_NAME}

# Seed the test database
seeddb_test:
    PGPASSWORD=${POSTGRES_PASSWORD} psql -h localhost -U ${POSTGRES_USER} -d ${TEST_DB_NAME} -a -f ${LOCAL_DB_PATH}/seed.sql

resetdb_test:
	@echo "Resetting test database..."
	@make dropdb_test || true
	@make createdb_test
	@make citext
	@make migrateup_test
	@make seeddb_test

sqlc:
	sqlc generate

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=desserted \
		proto/*.proto	

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/PlatosCodes/desserted/backend/db/sqlc Store

server:
	go run main.go

test:
	go test -v -cover ./...

# Additional tools
db_docs:
	dbdocs build doc/database.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/database.dbml

# Define phony targets
.PHONY: network postgres createdb dropdb citext migrateup migrateup1 migratedown migratedown1 new_migration seeddb resetdb createdb_test seeddb_test sqlc proto server test db_docs db_schema 

