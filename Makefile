# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOLINT=golangci-lint
BINARY_NAME=go-ldenweek
BINARY_UNIX=$(BINARY_NAME)

# Build directory
BUILD_DIR=./bin

.PHONY: all build clean test lint fmt tidy install help

all: test lint build

build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/go-ldenweek

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

test:
	$(GOTEST) -v ./...

lint:
	$(GOLINT) run ./...

fmt:
	$(GOCMD) fmt ./...

tidy:
	$(GOMOD) tidy

install:
	$(GOCMD) install ./cmd/go-ldenweek

help:
	@echo "make - Run tests, lint, and build the binary"
	@echo "make build - Build the binary"
	@echo "make clean - Remove the binary and clean Go cache"
	@echo "make test - Run tests"
	@echo "make lint - Run golangci-lint"
	@echo "make fmt - Format Go code"
	@echo "make tidy - Tidy Go modules"
	@echo "make install - Install the binary"
