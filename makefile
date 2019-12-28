# BINARY_NAME defaults to the name of the repository
BINARY_NAME := parking_service
BUILD_INFO_FLAGS := -X main.BuildTime=$(shell date -u '+%Y-%m-%d_%H:%M:%S') -X main.BuildCommitHash=$(shell git rev-parse HEAD)
GOBIN := $(GOPATH)/bin
LIST_NO_VENDOR := $(go list ./... | grep -v /vendor/)
OSX_BUILD_FLAGS := -s
STATIC_BUILD_FLAGS := -linkmode external -extldflags -static -w

# `make` -- run in wercker (golang image uses Debian)
default: check fmt deps test linux

# `make dev` / `make osx` -- run when doing local development (on OSX)
dev: osx
osx: check fmt deps test.osx build

.PHONY: build
build:
	# Build project
	go build -ldflags "$(BUILD_INFO_FLAGS) $(OSX_BUILD_FLAGS)" -a -o $(BINARY_NAME) .

.PHONY: linux
linux:
	# Build project for linux
	env GOOS=linux GOARCH=amd64 go build -ldflags "$(BUILD_INFO_FLAGS)" -a -o $(BINARY_NAME).linux .

# `make docker` -- build Docker image with linux bits
.PHONY: docker
docker: deps linux

.PHONY: clean
clean:
	go clean -i
	rm -rf ./vendor/*/
	rm -f $(BINARY_NAME)

deps:
	# Install or update govend
	go get -u github.com/govend/govend
	# Fetch vendored dependencies
	$(GOBIN)/govend -v

.PHONY: fmt
fmt:
	# Format all Go source files (excluding vendored packages)
	go fmt $(LIST_NO_VENDOR)

.PHONY: test
test:
	# Run all tests, with coverage (excluding vendored packages)
	go test -a -v -cover $(LIST_NO_VENDOR)
