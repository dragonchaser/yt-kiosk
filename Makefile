
APP_NAME = "yt-kiosk"

.PHONY: run
run:
	@echo "Running..."
	@go run cmd/$(APP_NAME)/main.go

.PHONY: all
all: build

.PHONY: all-arch-build
all-arch-build: build build-arm64 build-arm

.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

.PHONY: build-arm64
build-arm64:
	@echo "Building for arm64..."
	@GOOS=linux GOARCH=arm64 go build -o bin/$(APP_NAME)_arm64 cmd/$(APP_NAME)/main.go

build-arm:
	@echo "Building for arm..."
	@GOOS=linux GOARCH=arm go build -o bin/$(APP_NAME)_arm cmd/$(APP_NAME)/main.go