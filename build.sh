#!/bin/bash
mkdir -p ./build

# Build for macOS
env GOOS=darwin GOARCH=386 go build -o ./build/main-darwin main.go

# Build for Windows
env GOOS=darwin GOARCH=386 go build -o ./build/main-win.exe main.go

# Build for ARM
env GOOS=linux GOARCH=arm GOARM=7 go build -o ./build/main-arm main.go