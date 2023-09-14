#!/bin/bash

# Check if remote URL is supplied
if [ -z "$1" ]; then
  echo "Error: Remote URL is empty. Usage: ./setup_multiple_remotes.sh [second_remote_url]"
  exit 1
fi

# Step 1: List existing remotes for validation
echo "Existing remotes:"
git remote -v

# Step 2: Add a new remote repository
# Use the provided URL as the second remote repository's URL
git remote add second-remote $1

# Step 3: Confirm that the new remote has been added
echo "Updated remotes:"
git remote -v

# Print a message indicating that the setup is complete
echo "Remote setup complete."
