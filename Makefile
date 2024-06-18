# Variables
APP_NAME=epoch-conv
VERSION := $(shell cat VERSION)
EXISTING_TAG := $(shell git tag -l "$(VERSION)")

# Build all platforms
build-all: pre-build-check fmt imports staticcheck vet
	@echo "Building for all platforms..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=${VERSION}'" -o $(APP_NAME)_linux_amd64
	GOOS=linux GOARCH=arm64 go build -ldflags="-X 'main.Version=${VERSION}'" -o $(APP_NAME)_linux_arm64
	GOOS=darwin GOARCH=arm64 go build -ldflags="-X 'main.Version=${VERSION}'" -o $(APP_NAME)_macos_apple
	GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'main.Version=${VERSION}'" -o $(APP_NAME)_macos_intel
	GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.Version=${VERSION}'" -o $(APP_NAME)_windows.exe

.PHONY: fmt imports staticcheck vet build-all clean

# Code quality checks
fmt:
	@echo "Running gofmt..."
	@gofmt -d -e -s . 2>&1 | read; if [ $$? == 0 ]; then echo "Code is not formatted, please run 'gofmt -w .'" && exit 1; fi

imports:
	@echo "Running goimports..."
	@goimports -l . 2>&1 | read; if [ $$? == 0 ]; then echo "Imports are not properly organized, please run 'goimports -w .'" && exit 1; fi

staticcheck:
	@echo "Running staticcheck..."
	@staticcheck ./... || (echo "Staticcheck identified problems" && exit 1)

vet:
	@echo "Running go vet..."
	@go vet ./... || (echo "Go vet identified problems" && exit 1)

# Pre-build check to ensure version tag does not already exist
pre-build-check:
	@if [ "$(EXISTING_TAG)" = "$(VERSION)" ]; then \
		echo "Error: Tag $(VERSION) already exists.  Please update the VERSION file."; \
		exit 1; \
	fi

# Clean Up
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)_*
