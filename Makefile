GITCOMMIT = $(shell git rev-parse --short HEAD)
GOLDFLAGS = "-X main.GitCommit=$(GITCOMMIT)"
GOTOOLS = "github.com/jteeuwen/go-bindata/..."

default: tests

deps:
	@echo "--> Getting dependencies"
	@go get ./...
	@go get -u $(GOTOOLS)

format:
	@echo "--> Running go fmt"
	@go fmt ./...

tests:
	@echo "--> Running go test"
	@go test -v ./...

install: deps bindata
	@echo "--> Build and install binary"
	@go install -ldflags $(GOLDFLAGS) github.com/osminogin/tornote/tornote

bindata:
	@echo "--> Generate bindata"
	@go-bindata -pkg tornote templates/...

.PHONY: all deps format tests install bindata