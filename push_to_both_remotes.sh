#!/bin/bash

# Check if a commit message is supplied
if [ -z "$1" ]; then
  echo "Error: Commit message is empty. Usage: ./push_to_both_remotes.sh [commit_message]"
  exit 1
fi

# Check if in a Git repository
if [ ! -d .git ]; then
  echo "Error: This directory has not been initialized with Git."
  exit 1
fi

# Determine the current branch
current_branch=$(git symbolic-ref --short HEAD 2>/dev/null)

if [ -z "$current_branch" ]; then
  echo "Error: Not currently on any branch."
  exit 1
fi

# Add all changes to the staging area
git add .

# Check if there are staged changes
staged_changes=$(git diff --name-only --cached)

if [ -z "$staged_changes" ]; then
  echo "Error: No changes staged for commit."
  exit 1
fi

# Initialize second remote if it doesn't exist
if ! git remote get-url "second-remote" > /dev/null 2>&1; then
  echo "Initializing second remote."
  read -p "Enter the URL for the second remote repository: " second_remote_url
  git remote add second-remote $second_remote_url
fi

# Commit changes using the provided message
git commit -m "$1"

# Push changes to 'origin'
git push origin $current_branch

if [ $? -ne 0 ]; then
  echo "Error: Failed to push to origin."
  exit 1
fi

# Push changes to 'second-remote'
git push second-remote $current_branch

if [ $? -ne 0 ]; then
  echo "Error: Failed to push to second-remote."
  exit 1
fi

# Print message indicating that the push operation is complete
echo "Pushed to both repositories successfully."
