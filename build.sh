#!/bin/bash

APP_NAME="creator-hub"
OUTPUT_DIR="bin"

mkdir -p $OUTPUT_DIR

echo "🔨 Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o $OUTPUT_DIR/${APP_NAME}.exe

echo "🔨 Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_DIR/${APP_NAME}-linux

echo "🔨 Building for macOS (Intel - amd64)..."
GOOS=darwin GOARCH=amd64 go build -o $OUTPUT_DIR/${APP_NAME}-mac-intel

echo "🔨 Building for macOS (ARM - M1/M2)..."
GOOS=darwin GOARCH=arm64 go build -o $OUTPUT_DIR/${APP_NAME}-mac-arm

echo "✅ Done! Binaries are in ./$OUTPUT_DIR"