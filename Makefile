.PHONY: run build test clean migrate-up migrate-down migrate-create docker-up docker-down swagger

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test -v -cover ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	rm -rf bin/
	rm -f coverage.out

migrate-up:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/academic_db?sslmode=disable" up

migrate-down:
	migrate -path database/migrations -database "postgresql://postgres:postgres@localhost:5432/academic_db?sslmode=disable" down

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir database/migrations -seq $$name

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

swagger:
	swag init -g cmd/api/main.go -o docs/swagger

install-tools:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

deps:
	go mod download
	go mod tidy

help:
	@echo "Available commands:"
	@echo "  make run           - Run the application"
	@echo "  make build         - Build the application"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make clean         - Clean build files"
	@echo "  make migrate-up    - Run database migrations"
	@echo "  make migrate-down  - Rollback database migrations"
	@echo "  make migrate-create- Create new migration"
	@echo "  make docker-up     - Start Docker containers"
	@echo "  make docker-down   - Stop Docker containers"
	@echo "  make swagger       - Generate Swagger docs"
	@echo "  make install-tools - Install required tools"
	@echo "  make deps          - Download dependencies"