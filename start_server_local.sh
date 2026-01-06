#!/bin/bash

# AiGen START
# Start xserver in LOCAL mode (single server with embedded moduley)
# Build first with: ./build_local.sh

# Create logs directory
mkdir -p logs

# Check if binary exists
if [ ! -f "./xserver-local" ]; then
    echo "Error: xserver-local not found. Run ./build_local.sh first"
    exit 1
fi

# Start xserver in local mode with logging
XSERVER_ADDR=":8001" ./xserver-local > logs/xserver_local.log 2>&1 &
XSERVER_PID=$!

echo "Single server started in LOCAL mode with logging:"
echo "  xserver-local (PID $XSERVER_PID) on :8001 -> logs/xserver_local.log"
echo ""
echo "Test with: curl -w '\n' http://localhost:8001/x?input=hello"
echo ""
echo "Press Ctrl+C to stop server"

# Wait for Ctrl+C
trap "kill $XSERVER_PID; exit" INT
wait
# AiGen END