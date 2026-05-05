APP_NAME := mustag
BUILD_DIR := build

VERSION := $(shell git describe --tags --always)
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: build clean

build:
	@echo "==> building $(APP_NAME)"
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) .

clean:
	rm -rf $(BUILD_DIR)