#!/bin/bash

echo "Testing ai-done notification..."
echo "Watch your menu bar for the icon to change from 💤 to ✅"
echo ""

# Modify the hook file to trigger the watcher
echo '{
  "name": "AI Done Notification",
  "version": "1.0.0",
  "description": "Notify when AI agent completes",
  "when": {
    "type": "agentStop"
  },
  "then": {
    "type": "runCommand",
    "command": "echo AI task completed"
  }
}' > ~/.kiro/hooks/ai-done-test.json

echo "Hook file updated. You should see:"
echo "  1. Menu bar icon change to ✅"
echo "  2. Hear a 'Glass' sound"
echo "  3. See a notification: 'AI Done: Code generation complete!'"
echo "  4. Icon returns to 💤 after 3 seconds"
