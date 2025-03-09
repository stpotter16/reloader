run: build
	@./tmp/main

build:
	@go build -o ./tmp/main cmd/app/main.go
