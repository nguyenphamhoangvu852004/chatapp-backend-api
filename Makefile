# Makefile for Go Project

APP_NAME := chatapp-backend-api

.PHONY: build run test clean

build:
	go build -o bin/$(APP_NAME) ./cmd

dev:
	go run ./cmd/server/main.go

run: build
	./bin/$(APP_NAME)

test:
	go test ./...

clean:
	rm -rf bin/