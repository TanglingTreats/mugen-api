build:
	@go build -o mugen-api

run: build
	@./mugen-api --env ".env.local"

test:
	@go test ./... -v
