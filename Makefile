NAME    := myhttp
VERSION ?= $(shell git describe --tags --abbrev=0)

all: test

build: test
	@echo "+ $@"
	@go build -o "$(NAME)" myhttp.go

lint:
	@echo "+ $@"
	@golint

test: lint
	@echo "+ $@"
	@go test

clean:
	@echo "+ $@"
	@$(RM) -f "$(NAME)"

.PHONY: all build lint test clean
