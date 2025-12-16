#!/bin/bash

# make this executable: chmod +x start_servers.sh

# Start yserver
YSERVER_ADDR=":8002" ./yserver &
YSERVER_PID=$!

# Wait a moment for yserver to start
sleep 1

# Start xserver
XSERVER_ADDR=":8001" MODULEY_MODE="remote" MODULEY_URL="http://localhost:8002" ./xserver &
XSERVER_PID=$!

echo "Servers started:"
echo "  yserver (PID $YSERVER_PID) on :8002"
echo "  xserver (PID $XSERVER_PID) on :8001"
echo ""
echo "Test with: curl -w '\n' http://localhost:8001/x?input=hello"
echo ""
echo "Press Ctrl+C to stop both servers"

# Wait for Ctrl+C
trap "kill $YSERVER_PID $XSERVER_PID; exit" INT
wait