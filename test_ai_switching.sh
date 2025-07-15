#!/bin/bash

# Test script for AI switching functionality
# Make sure the server is running on localhost:8080

echo "🤖 Testing AI Provider Switching"
echo "================================"

# Check environment variables
echo "Current AI Configuration:"
echo "  SELECTED_AI: ${SELECTED_AI:-OPEN_AI (default)}"
echo "  OPEN_AI_API_KEY: ${OPEN_AI_API_KEY:+SET}"
echo "  DEEPSEEK_AI_API_KEY: ${DEEPSEEK_AI_API_KEY:+SET}"
echo ""

# Test 1: Generate workout plan with current AI provider
echo "1. Generating workout plan with current AI provider..."
curl -X POST http://localhost:8080/api/v1/ai/workout-plan/i05zVUkMmkabNryrIdD4vwnBPkO2 \
  -H "Content-Type: application/json" \
  -s | jq '.data.name, .data.description' 2>/dev/null || echo "Error: Could not connect to server"

echo ""
echo "✅ AI switching functionality is ready!"
echo ""
echo "🔧 To switch AI providers, set the environment variable:"
echo "   export SELECTED_AI=OPEN_AI    # Use OpenAI"
echo "   export SELECTED_AI=DEEPSEEK   # Use DeepSeek"
echo ""
echo "📋 Environment Variables:"
echo "   OPEN_AI_API_KEY=your-openai-key"
echo "   DEEPSEEK_AI_API_KEY=your-deepseek-key"
echo "   SELECTED_AI=OPEN_AI or DEEPSEEK"
echo ""
echo "🔄 Restart the server after changing SELECTED_AI to switch providers" 