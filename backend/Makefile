# This Makefile is for managing the Desserted card game project

# Environment Variables
POSTGRES_USER := root
POSTGRES_PASSWORD := bluecomet
DB_NAME := desserted
TEST_DB_NAME := desserted-test
CONTAINER_NAME := postgres-desserted
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

# Database seeding
seeddb:
	docker exec -it ${CONTAINER_NAME} psql --username=${POSTGRES_USER} --dbname=${DB_NAME} -a -f ${CONTAINER_VOLUME_PATH}/seed.sql

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
	docker exec -it ${CONTAINER_NAME} psql --username=${POSTGRES_USER} --dbname=${TEST_DB_NAME} -a -f ${CONTAINER_VOLUME_PATH}/seed.sql

resetdb_test:
	@echo "Resetting test database..."
	@make dropdb_test || true
	@make createdb_test
	@make citext
	@make migrateup_test
	@make seeddb_test

test:
	go test -v -cover ./...

# Additional tools
db_docs:
	dbdocs build ../doc/database.dbml

db_schema:
	dbml2sql --postgres -o ../doc/schema.sql ../doc/database.dbml

sqlc:
	sqlc generate

server:
	go run main.go

# Define phony targets
.PHONY: network postgres createdb dropdb citext migrateup migrateup1 migratedown migratedown1 new_migration seeddb resetdb createdb_test seeddb_test test db_docs db_schema sqlc server

