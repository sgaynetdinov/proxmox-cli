.PHONY: build build-all clean test help install

# Binary name
BINARY_NAME=proxmox-cli

# Version info
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME_UTC = $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
LDFLAGS = -ldflags="-s -w -X main.version=$(VERSION) -X main.buildTimeUTC=$(BUILD_TIME_UTC) -X main.gitCommit=$(GIT_COMMIT)"

# Default build
build:
	go build $(LDFLAGS) -o $(BINARY_NAME) ./cmd/proxmox-cli

# Build for all platforms
build-all: clean
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 ./cmd/proxmox-cli
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 ./cmd/proxmox-cli
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe ./cmd/proxmox-cli
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 ./cmd/proxmox-cli
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 ./cmd/proxmox-cli

# Create distribution directory
dist:
	mkdir -p dist

# Build with distribution directory
build-dist: dist build-all

# Install to local system
install: build
	sudo mv $(BINARY_NAME) /usr/local/bin/

# Run tests
test:
	go test -v ./...

# Clean built binaries
clean:
	rm -f $(BINARY_NAME)*
	rm -rf dist/

# Tidy dependencies
tidy:
	go mod tidy
	go mod vendor


# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build binary for current platform"
	@echo "  build-all      - Build binaries for all platforms"
	@echo "  build-dist     - Build all binaries in dist/ directory"
	@echo "  install        - Install binary to /usr/local/bin/"
	@echo "  test           - Run tests"
	@echo "  clean          - Remove built binaries"
	@echo "  tidy           - Tidy and vendor dependencies"
	@echo "  fmt            - Format code"
	@echo "  vet            - Vet code"
	@echo "  help           - Show this help"