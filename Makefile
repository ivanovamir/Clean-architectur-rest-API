.SILENT:

include .env

dep:
	go mod download

vet:
	go vet -json
	go vet -json ./cmd
	go vet -json ./internal
	go vet -json ./pkg/dto
	go vet -json ./pkg/handler
	go vet -json ./pkg/models
	go vet -json ./pkg/repository
	go vet -json ./pkg/service

run:
	go run cmd/main.go

build:
	go build cmd/main.go

clean:
	go clean