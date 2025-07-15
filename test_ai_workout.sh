#!/bin/bash

# Test script for AI Workout Plan API with OpenAI Integration
# Make sure the server is running on localhost:8080
# Make sure AI_API_KEY is set in your environment

echo "🧪 Testing AI Workout Plan API with Multi-AI Support"
echo "=================================================="

# Check environment variables
echo "Current AI Configuration:"
echo "  SELECTED_AI: ${SELECTED_AI:-OPEN_AI (default)}"
echo "  OPEN_AI_API_KEY: ${OPEN_AI_API_KEY:+SET}"
echo "  DEEPSEEK_AI_API_KEY: ${DEEPSEEK_AI_API_KEY:+SET}"
echo ""

# Check if required API keys are set
if [ "$SELECTED_AI" = "OPEN_AI" ] && [ -z "$OPEN_AI_API_KEY" ]; then
    echo "⚠️  Warning: OPEN_AI_API_KEY not set but SELECTED_AI=OPEN_AI"
fi

if [ "$SELECTED_AI" = "DEEPSEEK" ] && [ -z "$DEEPSEEK_AI_API_KEY" ]; then
    echo "⚠️  Warning: DEEPSEEK_AI_API_KEY not set but SELECTED_AI=DEEPSEEK"
fi

# Test 1: Generate workout plan for user
echo ""
echo "1. Generating personalized workout plan for user i05zVUkMmkabNryrIdD4vwnBPkO2..."
echo "   (This will call ${SELECTED_AI:-OPEN_AI} to generate a real workout plan)"
curl -X POST http://localhost:8080/api/v1/ai/workout-plan/i05zVUkMmkabNryrIdD4vwnBPkO2 \
  -H "Content-Type: application/json" \
  -s | jq '.'

# Test 2: Get specific workout plan
echo ""
echo "2. Getting workout plan with ID 1..."
curl http://localhost:8080/api/v1/ai/workout-plan/1 \
  -H "Content-Type: application/json" \
  -s | jq '.'

# Test 3: Get all workout plans for user
echo ""
echo "3. Getting all workout plans for user i05zVUkMmkabNryrIdD4vwnBPkO2..."
curl http://localhost:8080/api/v1/ai/workout-plans/i05zVUkMmkabNryrIdD4vwnBPkO2 \
  -H "Content-Type: application/json" \
  -s | jq '.'

# Test 4: Update workout plan
echo ""
echo "4. Updating workout plan with ID 1..."
curl -X PUT http://localhost:8080/api/v1/ai/workout-plan/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Push Pull Legs",
    "description": "Updated description with more focus on strength",
    "aiFeedbackCycle": 14,
    "planValidityPeriod": 30
  }' \
  -s | jq '.'

# Test 5: Health check
echo ""
echo "5. Health check..."
curl http://localhost:8080/health \
  -H "Content-Type: application/json" \
  -s | jq '.'

echo ""
echo "✅ AI Workout Plan API tests completed!"
echo ""
echo "🤖 Multi-AI Integration Features:"
echo "   • Support for both OpenAI GPT-4 and DeepSeek AI"
echo "   • Configurable AI provider via environment variable"
echo "   • Personalized based on user profile from Firestore"
echo "   • Considers fitness level, goals, and available equipment"
echo "   • Generates appropriate exercises, sets, reps, and weights"
echo "   • High-quality fitness recommendations from advanced language models"
echo ""
echo "📋 Available endpoints:"
echo "   POST   /api/v1/ai/workout-plan/:user_id     - Generate workout plan (Configurable AI)"
echo "   GET    /api/v1/ai/workout-plan/:plan_id     - Get specific plan"
echo "   PUT    /api/v1/ai/workout-plan/:plan_id     - Update plan"
echo "   DELETE /api/v1/ai/workout-plan/:plan_id     - Delete plan"
echo "   GET    /api/v1/ai/workout-plans/:user_id    - Get all user plans"
echo ""
echo "🔧 Environment Setup:"
echo "   OPEN_AI_API_KEY=your-openai-api-key-here"
echo "   DEEPSEEK_AI_API_KEY=your-deepseek-api-key-here"
echo "   SELECTED_AI=OPEN_AI or DEEPSEEK" 