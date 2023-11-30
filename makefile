.PHONY: default run run-with-docs dev build test docs clean

# Variables
APP_NAME = API_Empregos_Golang

# Tasks
default: run-with-docs

run:
	@echo "Running $(APP_NAME)..."
	@go run main.go
run-with-docs:
	@echo "Running and generating docs$(APP_NAME)..."
	@swag init
	@go run main.go
dev:
	@echo "Running in dev mode $(APP_NAME)..."
	@air main.go --port=8080
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) main.go
test:
	@echo "Testing $(APP_NAME)..."
	@go test ./...
docs:
	@echo "Generating docs $(APP_NAME)..."
	@swag init
clean:
	@echo "Cleaning $(APP_NAME)..."
	@rm -rf $(APP_NAME)
	@rm -rf ./docs
