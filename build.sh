#!/usr/bin/env bash
set -euo pipefail

APP_NAME="cli-snoopy"
OUTPUT_DIR="dist"

rm -rf "${OUTPUT_DIR}"
mkdir -p "${OUTPUT_DIR}"

echo "Building Linux binary"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o "${OUTPUT_DIR}/${APP_NAME}-linux-amd64" \
    .

echo "Building macOS amd64 binary"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o "${OUTPUT_DIR}/${APP_NAME}-darwin-amd64" \
    .

echo "Building macOS arm64 binary"
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o "${OUTPUT_DIR}/${APP_NAME}-darwin-arm64" \
    .

echo "Building Windows binary"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o "${OUTPUT_DIR}/${APP_NAME}-windows-amd64.exe" \
    .

echo "Done"
echo "Binaries are stored in ./${OUTPUT_DIR}/"