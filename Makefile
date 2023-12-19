dev:
	@~/go/bin/air -c .air.toml

run: build
	@./bin/server

build:
	@go build -o bin/server cmd/api/main.go