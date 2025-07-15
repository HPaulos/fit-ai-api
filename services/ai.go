package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"fit-ai-api/models"
)

// AIProvider represents different AI service providers
type AIProvider string

const (
	OpenAI   AIProvider = "OPEN_AI"
	DeepSeek AIProvider = "DEEPSEEK"
)

// AIService handles AI-powered workout plan generation
type AIService struct {
	openaiKey   string
	deepseekKey string
	selectedAI  AIProvider
	client      *http.Client
}

// NewAIService creates a new AI service instance
func NewAIService() *AIService {
	openaiKey := os.Getenv("OPEN_AI_API_KEY")
	deepseekKey := os.Getenv("DEEPSEEK_AI_API_KEY")
	selectedAI := AIProvider(strings.ToUpper(os.Getenv("SELECTED_AI")))

	// Default to OpenAI if not specified
	if selectedAI == "" {
		selectedAI = OpenAI
	}

	return &AIService{
		openaiKey:   openaiKey,
		deepseekKey: deepseekKey,
		selectedAI:  selectedAI,
		client: &http.Client{
			Timeout: 120 * time.Second, // Increased timeout for complex prompts
		},
	}
}

// GenerateWorkoutPlan generates a personalized workout plan based on user data
func (ai *AIService) GenerateWorkoutPlan(userData models.UserData) (*models.WorkoutPlan, error) {
	// Create the prompt for the AI
	prompt := ai.createWorkoutPrompt(userData)

	// Call the AI API based on selected provider
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

	prompt := fmt.Sprintf(WorkoutPlanTemplate,
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

// callAIAPI makes a request to the selected AI API
func (ai *AIService) callAIAPI(prompt string) (string, error) {
	switch ai.selectedAI {
	case OpenAI:
		return ai.callOpenAI(prompt)
	case DeepSeek:
		return ai.callDeepSeek(prompt)
	default:
		return "", fmt.Errorf("unsupported AI provider: %s", ai.selectedAI)
	}
}

// callOpenAI makes a request to OpenAI API
func (ai *AIService) callOpenAI(prompt string) (string, error) {
	// Check if API key is available
	if ai.openaiKey == "" {
		return "", fmt.Errorf("OPEN_AI_API_KEY environment variable is not set")
	}

	// Create OpenAI request
	requestBody := map[string]interface{}{
		"model": "gpt-4",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": WorkoutPlanPrompt,
			},
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"temperature": 0.7,
		"max_tokens":  2000,
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ai.openaiKey)

	// Make the request
	resp, err := ai.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("OpenAI API request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("OpenAI API error: %s - %s", resp.Status, string(body))
	}

	// Parse OpenAI response
	var openaiResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.Unmarshal(body, &openaiResponse)
	if err != nil {
		return "", fmt.Errorf("failed to parse OpenAI response: %w", err)
	}

	// Extract the response content
	if len(openaiResponse.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return openaiResponse.Choices[0].Message.Content, nil
}

// callDeepSeek makes a request to DeepSeek API
func (ai *AIService) callDeepSeek(prompt string) (string, error) {
	// Check if API key is available
	if ai.deepseekKey == "" {
		return "", fmt.Errorf("DEEPSEEK_AI_API_KEY environment variable is not set")
	}

	// Create DeepSeek request with context for better timeout handling
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	requestBody := map[string]interface{}{
		"model": "deepseek-chat",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": WorkoutPlanPrompt,
			},
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"temperature": 0.7,
		"max_tokens":  2000,
		"response_format": map[string]string{
			"type": "json_object",
		},
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.deepseek.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ai.deepseekKey)

	// Make the request
	resp, err := ai.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("DeepSeek API request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body with timeout
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("DeepSeek API error: %s - %s", resp.Status, string(body))
	}

	// Parse DeepSeek response
	var deepseekResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error,omitempty"`
	}

	err = json.Unmarshal(body, &deepseekResponse)
	if err != nil {
		return "", fmt.Errorf("failed to parse DeepSeek response: %w", err)
	}

	// Check for API errors in response
	if deepseekResponse.Error != nil {
		return "", fmt.Errorf("DeepSeek API error: %s", deepseekResponse.Error.Message)
	}

	// Extract the response content
	if len(deepseekResponse.Choices) == 0 {
		return "", fmt.Errorf("no response from DeepSeek")
	}

	return deepseekResponse.Choices[0].Message.Content, nil
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

// GetSelectedAI returns the currently selected AI provider
func (ai *AIService) GetSelectedAI() AIProvider {
	return ai.selectedAI
}
