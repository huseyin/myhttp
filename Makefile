NAME    := myhttp
VERSION ?= $(shell git describe --tags --abbrev=0)

all: test

build:
	@echo "+ $@"
	@GOOS=windows GOARCH=amd64 go build -o ./bin/$(NAME)_$(VERSION)_windows_amd64
	@GOOS=linux GOARCH=amd64 go build -o ./bin/$(NAME)_$(VERSION)_linux_amd64
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/$(NAME)_$(VERSION)_darwin_amd64

deps:
	@echo "+ $@"
	@go mod tidy -v

test: deps
	@echo "+ $@"
	@go test -race ./...

clean:
	@echo "+ $@"
	@$(RM) -rf ./bin

.PHONY: all build deps test clean
