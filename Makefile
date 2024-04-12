build:
	@go build -o mugen-api

run: build
	@./mugen-api
