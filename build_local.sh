#!/bin/bash

# AiGen START
# Build script for local mode (single server with embedded moduley)
# Make executable: chmod +x build_local.sh

echo "Building xserver in LOCAL mode..."
go build -tags=local -ldflags="-s -w" -trimpath -o xserver-local ./cmd/xserver

if [ $? -eq 0 ]; then
    echo "✓ xserver-local built successfully"
    ls -lh xserver-local | awk '{print "  Size:", $5}'
else
    echo "✗ Build failed"
    exit 1
fi
# AiGen END

