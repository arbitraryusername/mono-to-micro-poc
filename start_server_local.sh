#!/bin/bash

# Create logs directory
mkdir -p logs

# Start xserver in local mode with logging
XSERVER_ADDR=":8001" ./xserver > logs/xserver_local.log 2>&1 &
XSERVER_PID=$!

echo "Single server started with logging:"
echo "  xserver (PID $XSERVER_PID) on :8001 -> logs/xserver_local.log"
echo ""
echo "Test with: curl -w '\n' http://localhost:8001/x?input=hello"
echo ""
echo "Press Ctrl+C to stop server"

# Wait for Ctrl+C
trap "kill $XSERVER_PID; exit" INT
wait