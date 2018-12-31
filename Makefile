.PHONY: build test deps clean

all: build
	@./linkedlist

deps:
	@go get ./...

build:
	@go build -o linkedlist ./cmd/linkedlist/main.go

test:
	@go test -v -cover -coverprofile=coverage.txt -covermode=atomic .

clean:
	@git clean -f -d -X
