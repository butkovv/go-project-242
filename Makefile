test:
	go mod tidy
	find . -name ".DS_Store" -type f -delete
	go test -v ./...

install:
	go install

lint:
	golangci-lint run ./...

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
