#!/bin/bash

# AiGen START
# Build script for all binaries (local and remote modes)
# Make executable: chmod +x build_all.sh

echo "=== Building all binaries ==="
echo ""

echo "1. Building xserver-local (local mode)..."
go build -tags=local -ldflags="-s -w" -trimpath -o xserver-local ./cmd/xserver
if [ $? -eq 0 ]; then
    echo "   ✓ xserver-local built successfully"
    ls -lh xserver-local | awk '{print "     Size:", $5}'
else
    echo "   ✗ Build failed"
    exit 1
fi

echo ""
echo "2. Building xserver-remote (remote mode)..."
go build -tags=remote -ldflags="-s -w" -trimpath -o xserver-remote ./cmd/xserver
if [ $? -eq 0 ]; then
    echo "   ✓ xserver-remote built successfully"
    ls -lh xserver-remote | awk '{print "     Size:", $5}'
else
    echo "   ✗ Build failed"
    exit 1
fi

echo ""
echo "3. Building yserver..."
go build -ldflags="-s -w" -trimpath -o yserver ./cmd/yserver
if [ $? -eq 0 ]; then
    echo "   ✓ yserver built successfully"
    ls -lh yserver | awk '{print "     Size:", $5}'
else
    echo "   ✗ Build failed"
    exit 1
fi

echo ""
echo "=== All builds complete ==="
# AiGen END

