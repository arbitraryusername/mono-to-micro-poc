#!/bin/bash

# Create logs directory
mkdir -p logs

# Start yserver with logging
YSERVER_ADDR=":8002" ./yserver > logs/yserver_remote.log 2>&1 &
YSERVER_PID=$!

# Wait a moment for yserver to start
sleep 1

# Start xserver with logging
XSERVER_ADDR=":8001" MODULEY_MODE="remote" MODULEY_URL="http://localhost:8002" ./xserver > logs/xserver_remote.log 2>&1 &
XSERVER_PID=$!

echo "Servers started with logging:"
echo "  yserver (PID $YSERVER_PID) on :8002 -> logs captured in logs/yserver_remote.log"
echo "  xserver (PID $XSERVER_PID) on :8001 -> logs captured in logs/xserver_remote.log"
echo ""
echo "Test with: curl -w '\n' http://localhost:8001/x?input=hello"
echo ""
echo "Press Ctrl+C to stop both servers"

# Wait for Ctrl+C
trap "kill $YSERVER_PID $XSERVER_PID; exit" INT
wait