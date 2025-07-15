#!/bin/bash

# Test script for AI Workout Plan API with OpenAI Integration
# Make sure the server is running on localhost:8080
# Make sure AI_API_KEY is set in your environment

echo "🧪 Testing AI Workout Plan API with OpenAI"
echo "=========================================="

# Check if AI_API_KEY is set
if [ -z "$AI_API_KEY" ]; then
    echo "⚠️  Warning: AI_API_KEY environment variable is not set"
    echo "   The API will return an error if OpenAI integration is enabled"
fi

# Test 1: Generate workout plan for user
echo ""
echo "1. Generating personalized workout plan for user i05zVUkMmkabNryrIdD4vwnBPkO2..."
echo "   (This will call OpenAI GPT-4 to generate a real workout plan)"
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
echo "🤖 OpenAI Integration Features:"
echo "   • Real AI-powered workout plan generation"
echo "   • Personalized based on user profile from Firestore"
echo "   • Considers fitness level, goals, and available equipment"
echo "   • Generates appropriate exercises, sets, reps, and weights"
echo "   • Uses GPT-4 for high-quality fitness recommendations"
echo ""
echo "📋 Available endpoints:"
echo "   POST   /api/v1/ai/workout-plan/:user_id     - Generate workout plan (OpenAI)"
echo "   GET    /api/v1/ai/workout-plan/:plan_id     - Get specific plan"
echo "   PUT    /api/v1/ai/workout-plan/:plan_id     - Update plan"
echo "   DELETE /api/v1/ai/workout-plan/:plan_id     - Delete plan"
echo "   GET    /api/v1/ai/workout-plans/:user_id    - Get all user plans"
echo ""
echo "🔧 Environment Setup:"
echo "   AI_API_KEY=your-openai-api-key-here" 