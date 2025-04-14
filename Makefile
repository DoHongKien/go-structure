APP_NAME := server

all:
	@echo "Hello World!"
	@echo "This is a Makefile for the ${APP_NAME} application."

dev:
	go run ./cmd/${APP_NAME}

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

docker_up:
	docker-compose -f environment/docker-compose-dev.yml up

docker_down:
	docker-compose -f environment/docker-compose-dev.yml down

wire:
	cd ./internal/wire && wire
.PHONY: all dev swag docker_up docker_down wire