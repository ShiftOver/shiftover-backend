prepare:
	@echo "Preparing required libraries..."
	brew install pre-commit
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
	go install golang.org/x/tools/cmd/goimports@latest

	@echo "Initializing environment for git..."
	pre-commit install -t pre-commit
	pre-commit install -t pre-push
	pre-commit install -t commit-msg

	@echo "Installing npm packages..."
	npm install
	npm install -g migrate-mongo

rehooks:
	@echo "Clear git hooks cache..."
	pre-commit clean

	@echo "Reinitializing environment for git..."
	pre-commit install -t pre-commit
	pre-commit install -t pre-push
	pre-commit install -t commit-msg

dev.up:
	@echo "Starting local development..."
	docker compose -f docker-compose.dev.yaml up

dev.down:
	@echo "Removing local development container..."
	docker-compose down

swagger:
	@echo "Generating swagger docs..."
	swag init

migrate.up:
	@echo "Running migrations..."
	migrate-mongo up

migrate.down:
	@echo "Rolling back migrations..."
	migrate-mongo down

.PHONY: prepare rehooks dev.up dev.down swagger migrate.up migrate.down
