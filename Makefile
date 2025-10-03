APP_NAME=gophkeeper
CLIENT_BIN=bin/client
SERVER_BIN=bin/server

VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE?=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS=-X github.com/gobackev/gophkeeper/internal/buildinfo.Version=$(VERSION) \
        -X github.com/gobackev/gophkeeper/internal/buildinfo.Commit=$(COMMIT) \
        -X github.com/gobackev/gophkeeper/internal/buildinfo.Date=$(DATE)

# Платформы для кросс-компиляции
PLATFORMS = linux/amd64 linux/arm64 windows/amd64 darwin/amd64 darwin/arm64

.PHONY: all client server clean cross-compile

all: client server

client:
	go build -ldflags "$(LDFLAGS)" -o $(CLIENT_BIN) ./cmd/client

server:
	go build -ldflags "$(LDFLAGS)" -o $(SERVER_BIN) ./cmd/server

# Кросс-компиляция для всех платформ
cross-compile:
	@echo "Building for all platforms..."
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -d'/' -f1); \
		ARCH=$$(echo $$platform | cut -d'/' -f2); \
		OUTPUT_DIR=bin/$$OS-$$ARCH; \
		mkdir -p $$OUTPUT_DIR; \
		echo "Building for $$OS/$$ARCH..."; \
		if [ "$$OS" = "windows" ]; then \
			GOOS=$$OS GOARCH=$$ARCH go build -ldflags "$(LDFLAGS)" -o $$OUTPUT_DIR/client.exe ./cmd/client; \
			GOOS=$$OS GOARCH=$$ARCH go build -ldflags "$(LDFLAGS)" -o $$OUTPUT_DIR/server.exe ./cmd/server; \
		else \
			GOOS=$$OS GOARCH=$$ARCH go build -ldflags "$(LDFLAGS)" -o $$OUTPUT_DIR/client ./cmd/client; \
			GOOS=$$OS GOARCH=$$ARCH go build -ldflags "$(LDFLAGS)" -o $$OUTPUT_DIR/server ./cmd/server; \
		fi; \
	done
	@echo "Cross-compilation completed!"

clean:
	rm -rf bin

