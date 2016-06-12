GITCOMMIT = $(shell git rev-parse --short HEAD)
GOLDFLAGS = "-X main.GitCommit=$(GITCOMMIT)"
GOTOOLS = github.com/mattn/goveralls golang.org/x/tools/cmd/cover github.com/jteeuwen/go-bindata/...

default: tests

deps:
	@echo "--> Getting dependencies"
	@$(foreach gotool,$(GOTOOLS),$(shell go get -u $(gotool)))

format:
	@echo "--> Running go fmt"
	@go fmt ./...

tests:
	@echo "--> Running go test"
	@go test -v ./...

install: bindata format
	@echo "--> Build and install binary"
	@go get ./...
	@go install -ldflags $(GOLDFLAGS) ./...

bindata: deps
	@echo "--> Generate bindata"
	@go-bindata -pkg tornote templates/... \
		public/vendor/jquery/dist/jquery.slim.min.js \
		public/vendor/sjcl/sjcl.js \
		public/main.js \
		public/styles.css


.PHONY: all deps format tests install bindata