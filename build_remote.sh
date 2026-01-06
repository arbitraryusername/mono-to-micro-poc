#!/bin/bash

# AiGen START
# Build script for remote mode (two servers with RPC communication)
# Make executable: chmod +x build_remote.sh

echo "Building xserver in REMOTE mode..."
go build -tags=remote -ldflags="-s -w" -trimpath -o xserver-remote ./cmd/xserver

if [ $? -eq 0 ]; then
    echo "✓ xserver-remote built successfully"
    ls -lh xserver-remote | awk '{print "  Size:", $5}'
else
    echo "✗ xserver-remote build failed"
    exit 1
fi

echo ""
echo "Building yserver..."
go build -ldflags="-s -w" -trimpath -o yserver ./cmd/yserver

if [ $? -eq 0 ]; then
    echo "✓ yserver built successfully"
    ls -lh yserver | awk '{print "  Size:", $5}'
else
    echo "✗ yserver build failed"
    exit 1
fi
# AiGen END

