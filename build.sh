#!/bin/bash

# make this executable: chmod +x build.sh

# Release build for xserver (optimized, stripped)
go build -ldflags="-s -w" -trimpath -o xserver ./cmd/xserver

# Release build for yserver (optimized, stripped)
go build -ldflags="-s -w" -trimpath -o yserver ./cmd/yserver