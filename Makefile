build:
	@go build -o bin/app

run:build
	@./bin/app

.PHONY: db-up db-down

# Start the PostgreSQL container
db-up:
    docker-compose up -d db

# Stop and remove the PostgreSQL container
db-down:
    docker-compose down