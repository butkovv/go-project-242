test:
	go mod tidy
	go test -v ./...

install:
	go install

lint:
	golangci-lint run ./...

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
