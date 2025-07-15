package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"fit-ai-api/models"

	"github.com/sashabaranov/go-openai"
)

// AIService handles AI-powered workout plan generation
type AIService struct {
	apiKey string
	client *openai.Client
}

// NewAIService creates a new AI service instance
func NewAIService() *AIService {
	apiKey := os.Getenv("AI_API_KEY")
	client := openai.NewClient(apiKey)

	return &AIService{
		apiKey: apiKey,
		client: client,
	}
}

// GenerateWorkoutPlan generates a personalized workout plan based on user data
func (ai *AIService) GenerateWorkoutPlan(userData models.UserData) (*models.WorkoutPlan, error) {
	// Create the prompt for the AI
	prompt := ai.createWorkoutPrompt(userData)

	// Call the AI API
	response, err := ai.callAIAPI(prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to call AI API: %w", err)
	}

	// Parse the AI response into a workout plan
	workoutPlan, err := ai.parseAIResponse(response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	return workoutPlan, nil
}

// createWorkoutPrompt creates a detailed prompt for the AI based on user data
func (ai *AIService) createWorkoutPrompt(userData models.UserData) string {
	user := userData.Data

	prompt := fmt.Sprintf(`Generate a personalized workout plan for the following user:

User Profile:
- Name: %s
- Age: %s (calculated from date of birth)
- Gender: %s
- Fitness Level: %s
- Activity Level: %s
- Height: %.1f %s
- Weight: %.1f %s
- Goals: %v
- Available Equipment: %v
- Location: %s
- Units: %s

Current Stats:
- Total Workouts: %d
- Current Streak: %d
- Longest Streak: %d
- Total Time: %d minutes
- Total Volume: %d

Please generate a complete workout plan in the following JSON format:

{
  "id": 1,
  "name": "Plan Name",
  "description": "Plan description",
  "createdAt": "2024-01-15T00:00:00.000Z",
  "aiFeedbackCycle": 12,
  "planValidityPeriod": 28,
  "sessionsCompleted": 0,
  "planStartDate": "2024-01-15T00:00:00.000Z",
  "hasNewPlanSuggestion": false,
  "sessions": [
    {
      "name": "Session Name",
      "note": "Session notes",
      "exercises": [
        {
          "id": 1,
          "name": "Exercise Name",
          "sets": 3,
          "reps": 10,
          "weight": {"value": 100, "unit": "LB"},
          "type": "weight"
        }
      ]
    }
  ]
}

Guidelines:
1. Create 3-4 sessions per week based on fitness level
2. Use available equipment: %v
3. Focus on user goals: %v
4. Adjust difficulty based on fitness level (%s)
5. Use appropriate weight units (%s)
6. Include a mix of compound and isolation exercises
7. Progressive overload principles
8. Rest days between muscle groups
9. Warm-up and cool-down considerations

Return only the JSON object, no additional text.`,
		user.FullName,
		user.DateOfBirth,
		user.Gender,
		user.FitnessLevel,
		user.ActivityLevel,
		user.Height.Value,
		user.Height.Unit,
		user.Weight.Value,
		user.Weight.Unit,
		user.Goals,
		user.Equipment,
		user.Location,
		user.Preferences.Units,
		user.Stats.TotalWorkouts,
		user.Stats.CurrentStreak,
		user.Stats.LongestStreak,
		user.Stats.TotalTime,
		user.Stats.TotalVolume,
		user.Equipment,
		user.Goals,
		user.FitnessLevel,
		user.Preferences.Units)

	return prompt
}

// callAIAPI makes a request to OpenAI API
func (ai *AIService) callAIAPI(prompt string) (string, error) {
	// Check if API key is available
	if ai.apiKey == "" {
		return "", fmt.Errorf("AI_API_KEY environment variable is not set")
	}

	// Create OpenAI request
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a fitness expert and personal trainer. Generate workout plans in JSON format only. Always return valid JSON that matches the exact structure requested.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Temperature: 0.7,
		MaxTokens:   2000,
	}

	// Call OpenAI API with context
	ctx := context.Background()
	resp, err := ai.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}

	// Extract the response content
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

// parseAIResponse parses the AI response into a WorkoutPlan struct
func (ai *AIService) parseAIResponse(response string) (*models.WorkoutPlan, error) {
	var workoutPlan models.WorkoutPlan

	err := json.Unmarshal([]byte(response), &workoutPlan)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal AI response: %w", err)
	}

	// Set default timestamps if not provided
	if workoutPlan.CreatedAt.IsZero() {
		workoutPlan.CreatedAt = time.Now()
	}
	if workoutPlan.PlanStartDate.IsZero() {
		workoutPlan.PlanStartDate = time.Now()
	}

	return &workoutPlan, nil
}
