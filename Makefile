build:
	@go build -o mugen-api

run: build
	@./mugen-api

test:
	@go test ./... -v
