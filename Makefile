APP_NAME := "pismo"

build:
	@echo "Building the $(APP_NAME) application..."
	go build -o output/pismo .

run:
	@echo "Running the $(APP_NAME) application..."
	./output/pismo

test:
	@echo "Running tests..."
	go test -v ./...

docker-build:
	@echo "Building the Docker image..."
	docker build -t $(APP_NAME) .

docker-run:
	@echo "Running the Docker container..."
	docker run -p 8080:8080 $(APP_NAME)

help:
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Usage:"
	@echo "  make build        Build the Go application"
	@echo "  make run          Run the application"
	@echo "  make docker-build Build the Docker image with Podman"
	@echo "  make docker-run   Run the Docker container with Podman"
	@echo "  make test         Run tests"
	@echo "  make help         Show this help message"