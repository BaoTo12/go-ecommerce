APP_NAME=server

run:
	docker compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d && go run ./cmd/$(APP_NAME)
dev:
	go run ./cmd/$(APP_NAME)
kill:
	docker compose kill 
up:
	docker compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d
down:
	docker compose down 

.PHONY: run