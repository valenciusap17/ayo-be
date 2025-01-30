tidy:
	@go mod tidy

run:
	@go run cmd/main.go

docker-up:
	docker-compose -f docker-compose.yaml up -d

docker-down:
	docker-compose -f docker-compose.yaml down