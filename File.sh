#!/bin/bash

# Load maintenance flag path from config.json
CONFIG_FILE="path/to/config.json"  # Update with the correct path to your config.json

# Use jq (a lightweight JSON processor) to parse the config.json
MAINTENANCE_FLAG=$(jq -r '.maintenance_flag_path' $CONFIG_FILE)

# If jq is not installed, you can install it via your package manager, e.g., `apt-get install jq` or `brew install jq`

if [ "$1" == "on" ]; then
    echo "Enabling maintenance mode..."
    touch $MAINTENANCE_FLAG
    echo "Maintenance mode enabled."
elif [ "$1" == "off" ]; then
    echo "Disabling maintenance mode..."
    rm -f $MAINTENANCE_FLAG
    echo "Maintenance mode disabled."
else
    echo "Usage: $0 {on|off}"
    exit 1
fi
