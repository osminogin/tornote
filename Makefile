GITCOMMIT = $(shell git rev-parse --short HEAD)
GOLDFLAGS = "-X main.GitCommit=$(GITCOMMIT)"
GOTOOLS = github.com/mattn/goveralls golang.org/x/tools/cmd/cover github.com/jteeuwen/go-bindata/...

default: test

build: default

test:
	@echo "--> Running tests"
	@go tool vet .
	@go test -v ./...

tools:
	@echo "--> Getting tools"
	@$(foreach gotool,$(GOTOOLS),$(shell go get -u $(gotool)))

deps: bindata
	@echo "--> Getting dependencies"
	@go get -t -v ./...

install: deps
	@echo "--> Build and install binary"
	@go install -ldflags $(GOLDFLAGS) ./...

format:
	@echo "--> Running go fmt"
	@go fmt ./...


.PHONY: all deps format tests install bindata
