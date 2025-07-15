.PHONY: help build run stop clean logs db-up db-down db-reset

# Default target
help: ## Show this help message
	@echo "Fit AI API - Available Commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Database commands
db-up: ## Start PostgreSQL database with Docker Compose
	docker-compose up -d postgres
	@echo "PostgreSQL is starting... Wait a few seconds for it to be ready"
	@echo "Database will be available at: localhost:5432"
	@echo "pgAdmin will be available at: http://localhost:8081"

db-down: ## Stop PostgreSQL database
	docker-compose down

db-reset: ## Reset database (delete all data and restart)
	docker-compose down -v
	docker-compose up -d postgres
	@echo "Database reset complete!"

db-logs: ## View database logs
	docker-compose logs -f postgres

# Application commands
run: ## Run the Go application
	go run main.go

build: ## Build the Go application
	go build -o fit-ai-api main.go

test: ## Run tests
	go test ./...

deps: ## Install/update dependencies
	go mod tidy
	go mod download

# Development commands
dev: db-up ## Start development environment (database + app)
	@echo "Starting development environment..."
	@echo "Database: localhost:5432"
	@echo "pgAdmin: http://localhost:8081 (admin@fitai.com / admin)"
	@echo "API: http://localhost:8080"
	@echo ""
	@echo "Starting API server..."
	go run main.go

clean: ## Clean up build artifacts
	rm -f fit-ai-api
	go clean

# Docker commands
docker-build: ## Build Docker images
	docker-compose build

docker-up: ## Start all services
	docker-compose up -d

docker-down: ## Stop all services
	docker-compose down

docker-logs: ## View all logs
	docker-compose logs -f

# Database connection test
test-db: ## Test database connection
	@echo "Testing database connection..."
	@docker-compose exec postgres psql -U postgres -d fit_ai_db -c "SELECT version();"

# Quick setup for new developers
setup: ## Complete setup for new developers
	@echo "Setting up Fit AI API development environment..."
	@echo "1. Installing dependencies..."
	@make deps
	@echo "2. Starting database..."
	@make db-up
	@echo "3. Waiting for database to be ready..."
	@sleep 5
	@echo "4. Testing database connection..."
	@make test-db
	@echo ""
	@echo "Setup complete! 🎉"
	@echo "Run 'make dev' to start the API server"
	@echo "Database: localhost:5432"
	@echo "pgAdmin: http://localhost:8081"
	@echo "API: http://localhost:8080" 