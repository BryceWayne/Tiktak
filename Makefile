# Makefile for the tiktak project

# Project variables
BINARY_NAME=tiktak
BUILD_DIR=build

# Go related variables
GOFILES=$(wildcard *.go)

# Target Operating Systems and Architectures
WINDOWS=windows
LINUX=linux
DARWIN=darwin

# Makefile rules
.PHONY: all windows_build linux_build darwin_build clean

all: windows_build linux_build darwin_build

windows_build:
	@echo Building for Windows...
	@set GOOS=windows&& set GOARCH=amd64&& go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(GOFILES)
	@echo Build complete: $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe

linux_build:
	@echo Building for Linux...
	@set GOOS=linux&& set GOARCH=amd64&& go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(GOFILES)
	@echo Build complete: $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64

darwin_build:
	@echo Building for Darwin (macOS)...
	@set GOOS=darwin&& set GOARCH=amd64&& go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(GOFILES)
	@echo Build complete: $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64

clean:
	@echo Cleaning up...
	@del /Q $(BUILD_DIR)\*
	@echo Clean complete.
