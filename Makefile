# Makefile for DCA Project
# Simplifies common development tasks

SHELL := /bin/bash
.ONESHELL:
.SHELLFLAGS := -o pipefail -c

# Disable built-in rules to avoid unexpected behavior
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables

# Go executable
GO ?= go

# Project directories
CMD_DIR := cmd/dca
PKG_DIR := ./...

# Default target
.PHONY: all
all: help

## --- Development Targets --- ##

# Display this help message
.PHONY: help
help:
	@echo "DCA Development Commands"
	@echo ""
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  run          Run the application (go run)"
	@echo "  build        Build the binary (go build)"
	@echo "  test         Run all tests with verbose output"
	@echo "  test-quiet   Run all tests without verbose output"
	@echo "  test-cover   Generate coverage report"
	@echo "  fmt          Format all Go files"
	@echo "  check        Run fmt, build, and test (CI validation)"
	@echo "  clean        Remove compiled binary and temporary files"
	@echo "  version      Show Go version"
	@echo ""
	@echo "For more information, see the Makefile source."

# Run the application
.PHONY: run
run:
	$(GO) run ./$(CMD_DIR)

# Build the binary
.PHONY: build
build:
	$(GO) build -o dca ./$(CMD_DIR)

# Run tests with verbose output
.PHONY: test
test:
	$(GO) test -v $(PKG_DIR)

# Run tests without verbose output
.PHONY: test-quiet
test-quiet:
	$(GO) test $(PKG_DIR)

# Generate coverage report
.PHONY: test-cover
test-cover:
	$(GO) test -coverprofile=coverage.out $(PKG_DIR) && $(GO) tool cover -func=coverage.out

# Format all Go files
.PHONY: fmt
fmt:
	$(GO) fmt $(PKG_DIR)

# Run fmt, build, and test (CI-friendly validation)
.PHONY: check
check: fmt build test

# Remove compiled binary and temporary files
.PHONY: clean
clean:
	rm -f dca
	rm -f coverage.out

# Show Go version
.PHONY: version
version:
	@echo "Go version: $$($(GO) version)"
