APP_NAME=server
GOOSE_DBSTRING=root:123456@tcp(127.0.0.1:3306)/goShop
GOOSE_DIR ?= sql/schema

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

# upse/downse/resetse cross-platform
ifeq ($(OS),Windows_NT)
upse:
	@set "GOOSE_DRIVER=mysql" && set "GOOSE_DBSTRING=$(GOOSE_DBSTRING)" && goose -dir=$(GOOSE_DIR) up

downse:
	@set "GOOSE_DRIVER=mysql" && set "GOOSE_DBSTRING=$(GOOSE_DBSTRING)" && goose -dir=$(GOOSE_DIR) down

resetse:
	@set "GOOSE_DRIVER=mysql" && set "GOOSE_DBSTRING=$(GOOSE_DBSTRING)" && goose -dir=$(GOOSE_DIR) reset
else
upse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING="$(GOOSE_DBSTRING)" goose -dir=$(GOOSE_DIR) up

downse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING="$(GOOSE_DBSTRING)" goose -dir=$(GOOSE_DIR) down

resetse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING="$(GOOSE_DBSTRING)" goose -dir=$(GOOSE_DIR) reset
endif

.PHONY: air


