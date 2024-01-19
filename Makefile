.PHONY: default run build test docs clean

APP_NAME=go_jobs

default: run

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs