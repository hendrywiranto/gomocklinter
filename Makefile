.PHONY: test

test:
	@go test -race -v ./...

build:
	@go build -o gomocklinter main.go
