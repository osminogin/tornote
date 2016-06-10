GITCOMMIT = $(shell git rev-parse --short HEAD)
GOLDFLAGS = "-X main.GitCommit=$(GITCOMMIT)"
GOTOOLS = "github.com/jteeuwen/go-bindata/..."

default: tests

deps:
	@echo "--> Getting dependencies"
	@go get -u $(GOTOOLS)

format:
	@echo "--> Running go fmt"
	@go fmt ./...

tests:
	@echo "--> Running go test"
	@go test -v ./...

install: bindata
	@echo "--> Build and install binary"
	@go get ./...
	@go install -ldflags $(GOLDFLAGS) ./...

bindata: deps
	@echo "--> Generate bindata"
	@go-bindata -pkg tornote templates/... \
		public/vendor/sjcl/sjcl.js \
		public/main.js \
		public/styles.css


.PHONY: all deps format tests install bindata