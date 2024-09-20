# Include go binaries into path
export PATH := $(GOPATH)/bin:$(PATH)

BUILD=$(shell date +%FT%T)
VERSION= $(shell git rev-parse --short HEAD)
CURRENT_BRANCH_NAME= $(shell git rev-parse --abbrev-ref HEAD)
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"


BIN := $(CURDIR)/bin/
SOURCE_PATH := GOBIN=$(BIN) CURDIR=$(shell pwd) TEST_SOURCE_PATH=$(PWD) CURRENT_BRANCH_NAME=$(CURRENT_BRANCH_NAME)

mod-action-%:
	@echo "Run go mod ${*}...."
	GOBIN=$(BIN) GO111MODULE=on go mod $*
	@echo "Done go mod  ${*}"

mod: mod-action-verify mod-action-tidy mod-action-vendor mod-action-download mod-action-verify ## Download all dependencies

tests: clean-cache-test ## run all tests
	$(SOURCE_PATH) go test ./... -race -v -coverprofile coverage.out
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

clean-cache-test: ## clean cache
	@echo "Cleaning test cache..."
	$(SOURCE_PATH) go clean -testcache

clean: clean-cache-test ## clean cache
	@echo "Cleaning..."
	rm -fr ./vendor
	go clean -i -r -x -cache -testcache -modcache
