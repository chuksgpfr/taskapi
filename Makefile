dev:
	go run cmd/main.go -mode=debug
start:
	go run cmd/main.go -mode=release
build:
	go build -o backend cmd/main.go

.PHONY: dev start