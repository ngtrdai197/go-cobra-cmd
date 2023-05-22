include .env

# Variables for binding arguments from cmd
MIGRATION_FILE_NAME?=default_migration_file_name

.PHONY: worker client producer consumer migration migrateup migrateup1 migratedown migratedown1 sqlc public-api

public-api:
	go run main.go public-api-cmd

# Delivery email client
client:
	go run main.go client-cmd

# Delivery email worker
worker:
	go run main.go worker-cmd

producer:
	go run main.go kafka-producer-cmd

consumer:
	go run main.go kafka-consumer-cmd

migration:
	migrate create --ext sql --dir db/migrations -seq $(MIGRATION_FILE_NAME)

migrateup:
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose up

migrateup1:
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose up 1

migratedown:
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose down

migratedown1:
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate