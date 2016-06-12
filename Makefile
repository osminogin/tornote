GITCOMMIT = $(shell git rev-parse --short HEAD)
GOLDFLAGS = "-X main.GitCommit=$(GITCOMMIT)"
GOTOOLS = github.com/mattn/goveralls golang.org/x/tools/cmd/cover github.com/jteeuwen/go-bindata/...

default: tests

tests: format
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

bindata: tools
	@echo "--> Generate bindata"
	@go-bindata -pkg tornote templates/... \
		public/vendor/jquery/dist/jquery.slim.min.js \
		public/vendor/sjcl/sjcl.js \
		public/main.js \
		public/styles.css

format:
	@echo "--> Running go fmt"
	@go fmt ./...


.PHONY: all deps format tests install bindata