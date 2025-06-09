# Makefile for Go Project

APP_NAME := chatapp-backend-api

.PHONY: build run test clean

build:
	go build -o bin/$(APP_NAME) ./cmd/server

run: 
	go run ./cmd/server/

test:
	go test ./...

clean:
	rm -rf bin/