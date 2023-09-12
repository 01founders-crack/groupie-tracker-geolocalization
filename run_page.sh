#!/bin/bash

go run backend/main.go &

sleep 1s

# Check the operating system
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    open http://localhost:3000/
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    xdg-open http://localhost:3000/
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" ]]; then
    # Windows (Cygwin, Git Bash, or Windows Subsystem for Linux)
    start http://localhost:3000/
else
    # Unsupported OS
    echo "Unsupported operating system: $OSTYPE"
fi