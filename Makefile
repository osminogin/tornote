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
	@go install -ldflags $(GOLDFLAGS) github.com/osminogin/tornote/tornote

bindata: deps
	@echo "--> Generate bindata"
	@go-bindata -pkg tornote templates/... \
		public/vendor/jquery/dist/jquery.min.js \
		public/vendor/bootstrap/dist/css/bootstrap.min.css \
		public/vendor/bootstrap/dist/js/bootstrap.min.js \
		public/main.js \
		public/styles.css


.PHONY: all deps format tests install bindata