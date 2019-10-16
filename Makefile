PACKAGES = $(shell go list ./... | grep -v '/vendor/')
VERSION = $(shell git rev-parse --short HEAD)
COMMIT = $(shell git log -1 --format='%H')


export GO111MODULE = on

BUILD_TAGS = netgo

all: build install

build: go.sum
	GO111MODULE=on go build  $(BUILD_FLAGS)  -o build/gaiad ./cmd/gaiad/
	GO111MODULE=on go build   $(BUILD_FLAGS) -o build/gaiacli ./cmd/gaiacli/

install: go.sum
	go install -tags "$(BUILD_FLAGS)" ./cmd/gaiacli
	go install -tags "$(BUILD_FLAGS)" ./cmd/gaiad

test: 
	@go test -cover $(PACKAGES)
	
go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		go mod verify

.PHONY: all build install go.sum test