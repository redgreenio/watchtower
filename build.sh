#!/usr/bin/env sh
set -e
echo "Building Linux binary"
env GOOS=linux GOARCH=arm go build -v -o bin/linux/xwatchtower

echo "Building macOS binary"
env GOOS=darwin GOARCH=amd64 go build -v -o bin/macOS/xwatchtower
