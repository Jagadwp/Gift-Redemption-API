.PHONY: help test test-coverage run migrate generate clean tidy build

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

test: ## Run all tests
	go test -v ./internal/service/...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./internal/service/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

generate: ## Generate swagger docs
	swag init -g cmd/server/main.go -o docs
	@echo "Swagger docs generated"

run: generate ## Generate swagger docs and run the application
	go run ./cmd/server

migrate: ## Run migrations
	go run ./cmd/server

clean: ## Clean build artifacts
	rm -f coverage.out coverage.html
	rm -rf tmp/

tidy: ## Tidy go modules
	go mod tidy

build: generate ## Generate swagger docs and build the application
	go build -o bin/gift-redemption ./cmd/server