.PHONY: help build run test clean lint fmt

help:
	@echo "Available commands:"
	@echo "  make build   - Build the application"
	@echo "  make run     - Run the application"
	@echo "  make test    - Run tests"
	@echo "  make clean   - Clean build artifacts"
	@echo "  make fmt     - Format code"
	@echo "  make lint    - Run linter"
	@echo "  make deps    - Download dependencies"

build:
	go build -o lemonade .

run:
	go run main.go

test:
	go test -v ./...

clean:
	rm -f lemonade
	go clean

fmt:
	go fmt ./...

lint:
	golangci-lint run

deps:
	go mod download
	go mod tidy

dev:
	go run main.go
