NAME    := myhttp
VERSION ?= $(shell git describe --tags --abbrev=0)

all: test

build:
	@echo "+ $@"
	@GOOS=windows GOARCH=amd64 go build -o ./bin/$(NAME)_windows_amd64
	@GOOS=linux GOARCH=amd64 go build -o ./bin/$(NAME)_linux_amd64
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/$(NAME)_darwin_amd64

lint:
	@echo "+ $@"
	@golint

test: lint
	@echo "+ $@"
	@go test

clean:
	@echo "+ $@"
	@$(RM) -rf ./bin

.PHONY: all build lint test clean
