build:
	@go build -o bin/fs

run: build
	@./bin/fs

fmt:
	@go fmt ./...

test:
	@go test ./...
