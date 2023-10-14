# Makefile

.PHONY: build run test docker-build docker-run docker-compose-up docker-compose-down

# Remove the previous built binary to make sure that latest changes should be reflected in the build
# You can keep this part to clean up any local artifacts
clean:
	@rm -rf go-bank
# Build the Go application
build: clean
	@go build -o go-bank

# Run the Go application using Docker Compose
run: build docker-compose-up

# Run tests
test:
	@go test -v ./...

# Build the Docker image for the Go application
docker-build:
	@docker-compose build

# Run the Docker containers using Docker Compose
docker-run:
	@docker-compose up -d

# Run Docker Compose (build and run)
docker-compose-up: docker-build docker-run

# Stop and remove the Docker containers
docker-compose-down:
	@docker-compose down
