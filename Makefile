# Name of your binary executable
BINARY_NAME=projector

build:
	@go build -o bin/$(BINARY_NAME) -v

run: build
	@./bin/$(BINARY_NAME)

test:
	@go test -v ./...

clean:
	@go clean
	@rm -rf bin/
