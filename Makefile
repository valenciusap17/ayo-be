DB_USER=$(shell jq -r '.common.postgres.user' config.jsonc)
DB_PASSWORD=$(shell jq -r '.common.postgres.password' config.jsonc)
DB_PORT=$(shell jq -r '.common.postgres.port' config.jsonc)
DB_HOST=$(shell jq -r '.common.postgres.host' config.jsonc)
DB_NAME=$(shell jq -r '.common.postgres.name' config.jsonc)
DB_CNS="user=$(DB_USER) password=$(DB_PASSWORD) port=$(DB_PORT) host=$(DB_HOST) dbname=$(DB_NAME)"

tidy:
	@go mod tidy

run:
	@go run cmd/main.go

docker-up:
	docker-compose -f docker-compose.yaml up -d

docker-down:
	docker-compose -f docker-compose.yaml down

create-migration:
	@goose create $(name) sql

migrate-up:
	@goose -dir migrations postgres $(DB_CNS) up

migrate-up-by-one:
	@goose -dir migrations postgres $(DB_CNS) up-by-one

migrate-down:
	@goose -dir migrations postgres $(DB_CNS) down