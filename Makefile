.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: build
build:
	go build -o bot cmd/bot/main.go

.PHONY: test
test:
	go test ./...