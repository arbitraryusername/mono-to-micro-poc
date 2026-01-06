#!/bin/bash

# AiGen START
# Start servers in REMOTE mode (two servers with RPC communication)
# Build first with: ./build_remote.sh

# Create logs directory
mkdir -p logs

# Check if binaries exist
if [ ! -f "./xserver-remote" ]; then
    echo "Error: xserver-remote not found. Run ./build_remote.sh first"
    exit 1
fi

if [ ! -f "./yserver" ]; then
    echo "Error: yserver not found. Run ./build_remote.sh first"
    exit 1
fi

# Start yserver with logging
YSERVER_ADDR=":8002" ./yserver > logs/yserver_remote.log 2>&1 &
YSERVER_PID=$!

# Wait a moment for yserver to start
sleep 1

# Start xserver-remote with logging (no MODULEY_MODE needed - determined at build time)
XSERVER_ADDR=":8001" MODULEY_URL="http://localhost:8002" ./xserver-remote > logs/xserver_remote.log 2>&1 &
XSERVER_PID=$!

echo "Servers started in REMOTE mode with logging:"
echo "  yserver (PID $YSERVER_PID) on :8002 -> logs/yserver_remote.log"
echo "  xserver-remote (PID $XSERVER_PID) on :8001 -> logs/xserver_remote.log"
echo ""
echo "Test with: curl -w '\n' http://localhost:8001/x?input=hello"
echo ""
echo "Press Ctrl+C to stop both servers"

# Wait for Ctrl+C
trap "kill $YSERVER_PID $XSERVER_PID; exit" INT
wait
# AiGen END