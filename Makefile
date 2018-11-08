.PHONY: build test deps clean

all: build
	@./linkedlist

deps:
	@go get ./...

build:
	@go build .

test:
	@go test -v -cover -coverprofile=coverage.txt -covermode=atomic .

clean:
	@rm -rf linkedlist
