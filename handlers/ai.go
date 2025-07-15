package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fit-ai-api/models"
	"fit-ai-api/services"

	"github.com/gin-gonic/gin"
)

// AIHandler handles AI-related endpoints
type AIHandler struct {
	firebaseService *services.FirebaseService
	aiService       *services.AIService
}

// NewAIHandler creates a new AI handler instance
func NewAIHandler(firebaseService *services.FirebaseService, aiService *services.AIService) *AIHandler {
	return &AIHandler{
		firebaseService: firebaseService,
		aiService:       aiService,
	}
}

// GenerateWorkoutPlan generates a personalized workout plan for a user
func (h *AIHandler) GenerateWorkoutPlan(c *gin.Context) {
	// Get user ID from URL parameter
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	// Fetch user data from Firestore
	userData, err := h.firebaseService.GetDocumentByID("users", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user data: " + err.Error(),
		})
		return
	}

	// Check if user data was found
	if userData == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Parse user data into our model
	var userDataModel models.UserData

	// Convert map to JSON bytes
	jsonData, err := json.Marshal(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to marshal user data: " + err.Error(),
		})
		return
	}

	// Unmarshal JSON to our model
	err = json.Unmarshal(jsonData, &userDataModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse user data: " + err.Error(),
		})
		return
	}

	// Generate workout plan using AI
	workoutPlan, err := h.aiService.GenerateWorkoutPlan(userDataModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate workout plan: " + err.Error(),
		})
		return
	}

	// Return the generated workout plan
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    workoutPlan,
		"message": "Workout plan generated successfully",
	})
}

// GetWorkoutPlanByID retrieves a specific workout plan by ID
func (h *AIHandler) GetWorkoutPlanByID(c *gin.Context) {
	planID := c.Param("plan_id")
	if planID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Plan ID is required",
		})
		return
	}

	// Convert plan ID to integer
	id, err := strconv.Atoi(planID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid plan ID format",
		})
		return
	}

	// For now, we'll return a mock workout plan
	// In a real implementation, you would fetch this from a database
	mockPlan := &models.WorkoutPlan{
		ID:                   id,
		Name:                 "Push Pull Legs",
		Description:          "Classic 3-day split focusing on push, pull, and leg movements",
		AIFeedbackCycle:      12,
		PlanValidityPeriod:   28,
		SessionsCompleted:    8,
		HasNewPlanSuggestion: true,
		Sessions: []models.WorkoutSession{
			{
				ID:   "session_1",
				Name: "Push Day",
				Note: "Focus on chest, shoulders, and triceps",
				Exercises: []models.Exercise{
					{
						ID:   1,
						Name: "Bench Press",
						Sets: 4,
						Reps: 8,
						Weight: models.WeightInfo{
							Value: 185,
							Unit:  "LB",
						},
						Type: "weight",
					},
					{
						ID:   2,
						Name: "Overhead Press",
						Sets: 3,
						Reps: 10,
						Weight: models.WeightInfo{
							Value: 135,
							Unit:  "LB",
						},
						Type: "weight",
					},
				},
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mockPlan,
	})
}

// UpdateWorkoutPlan updates an existing workout plan
func (h *AIHandler) UpdateWorkoutPlan(c *gin.Context) {
	planID := c.Param("plan_id")
	if planID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Plan ID is required",
		})
		return
	}

	// Parse the request body
	var workoutPlan models.WorkoutPlan
	if err := c.ShouldBindJSON(&workoutPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	// In a real implementation, you would update the plan in the database
	// For now, we'll just return success
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Workout plan updated successfully",
		"data":    workoutPlan,
	})
}

// DeleteWorkoutPlan deletes a workout plan
func (h *AIHandler) DeleteWorkoutPlan(c *gin.Context) {
	planID := c.Param("plan_id")
	if planID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Plan ID is required",
		})
		return
	}

	// In a real implementation, you would delete the plan from the database
	// For now, we'll just return success
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Workout plan deleted successfully",
	})
}

// GetUserWorkoutPlans retrieves all workout plans for a specific user
func (h *AIHandler) GetUserWorkoutPlans(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	// In a real implementation, you would fetch all plans for this user from the database
	// For now, we'll return a mock list
	mockPlans := []models.WorkoutPlan{
		{
			ID:                   1,
			Name:                 "Push Pull Legs",
			Description:          "Classic 3-day split",
			AIFeedbackCycle:      12,
			PlanValidityPeriod:   28,
			SessionsCompleted:    8,
			HasNewPlanSuggestion: true,
		},
		{
			ID:                   2,
			Name:                 "Full Body",
			Description:          "Full body workout for beginners",
			AIFeedbackCycle:      8,
			PlanValidityPeriod:   21,
			SessionsCompleted:    5,
			HasNewPlanSuggestion: false,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mockPlans,
		"count":   len(mockPlans),
	})
}
