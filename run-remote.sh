#!/bin/bash
set -e

# Kill old process if running
[ -f app.pid ] && kill $(cat app.pid) 2>/dev/null
rm -f app.pid main

# Build ARM64 binary
GOOS=linux GOARCH=arm64 go build -o main cmd/main.go

nohup ./main > app.log 2>&1 &
echo $! > app.pid

# Remove source code only after successful build and start
rm -rf cmd/ config/ internal/ .kamal/ deploy.sh Dockerfile go.mod go.sum .gitignore
