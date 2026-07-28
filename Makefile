run:
	go run cmd/main.go

tidy:
	go mod tidy

fmt:
	gofmt -w .

compose-up:
	docker compose -f docker-compose.yaml up -d

compose-down:
	docker compose -f docker-compose.yaml down
