GITCOMMIT = $(shell git rev-parse --short HEAD)
GOLDFLAGS = "-X main.GitCommit=$(GITCOMMIT)"

all: format tests install

deps:
	@echo "--> Getting dependencies"
	@go get ./...

format:
	@echo "--> Running go fmt"
	@go fmt ./...

tests:
	@echo "--> Running go test"
	@go test -v ./...

install: deps
	@echo "--> Build and install binary"
	@go install -ldflags $(GOLDFLAGS) github.com/osminogin/tornote/tornote

.PHONY: all deps format tests install