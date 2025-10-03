APP_NAME=gophkeeper
CLIENT_BIN=bin/client
SERVER_BIN=bin/server

VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
DATE?=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS=-X github.com/gobackev/gophkeeper/internal/buildinfo.Version=$(VERSION) \
        -X github.com/gobackev/gophkeeper/internal/buildinfo.Commit=$(COMMIT) \
        -X github.com/gobackev/gophkeeper/internal/buildinfo.Date=$(DATE)

.PHONY: all client server clean

all: client server

client:
	go build -ldflags "$(LDFLAGS)" -o $(CLIENT_BIN) ./cmd/client

server:
	go build -ldflags "$(LDFLAGS)" -o $(SERVER_BIN) ./cmd/server

clean:
	rm -rf bin

