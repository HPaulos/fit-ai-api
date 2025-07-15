package models

import "time"

// WorkoutPlan represents the complete workout plan structure
type WorkoutPlan struct {
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	CreatedAt            time.Time        `json:"createdAt"`
	AIFeedbackCycle      int              `json:"aiFeedbackCycle"`
	PlanValidityPeriod   int              `json:"planValidityPeriod"`
	SessionsCompleted    int              `json:"sessionsCompleted"`
	PlanStartDate        time.Time        `json:"planStartDate"`
	HasNewPlanSuggestion bool             `json:"hasNewPlanSuggestion"`
	SuggestedPlan        *SuggestedPlan   `json:"suggestedPlan,omitempty"`
	Sessions             []WorkoutSession `json:"sessions"`
}

// SuggestedPlan represents a suggested workout plan
type SuggestedPlan struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Reason       string           `json:"reason"`
	Improvements []string         `json:"improvements"`
	Comparison   PlanComparison   `json:"comparison"`
	Sessions     []WorkoutSession `json:"sessions"`
}

// PlanComparison compares current and suggested plans
type PlanComparison struct {
	Current   PlanDetails `json:"current"`
	Suggested PlanDetails `json:"suggested"`
}

// PlanDetails contains plan comparison details
type PlanDetails struct {
	Sessions   int    `json:"sessions"`
	Duration   string `json:"duration"`
	Difficulty string `json:"difficulty"`
	Focus      string `json:"focus"`
}

// WorkoutSession represents a single workout session
type WorkoutSession struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Note      string     `json:"note"`
	Exercises []Exercise `json:"exercises"`
}

// Exercise represents a single exercise in a workout
type Exercise struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Sets   int        `json:"sets"`
	Reps   int        `json:"reps"`
	Weight WeightInfo `json:"weight"`
	Type   string     `json:"type"`
}

// WeightInfo represents weight information for an exercise
type WeightInfo struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// UserData represents the user data structure from Firestore
type UserData struct {
	Collection string        `json:"collection"`
	Data       FirestoreUser `json:"data"`
	DocumentID string        `json:"document_id"`
	Success    bool          `json:"success"`
}

// FirestoreUser represents the user information from Firestore
type FirestoreUser struct {
	ActivityLevel string          `json:"activityLevel"`
	CreatedAt     string          `json:"createdAt"`
	DateOfBirth   string          `json:"dateOfBirth"`
	DisplayName   string          `json:"displayName"`
	Equipment     []string        `json:"equipment"`
	FitnessLevel  string          `json:"fitnessLevel"`
	FullName      string          `json:"fullName"`
	Gender        string          `json:"gender"`
	Goals         []string        `json:"goals"`
	Height        Measurement     `json:"height"`
	Location      string          `json:"location"`
	Preferences   UserPreferences `json:"preferences"`
	Stats         UserStats       `json:"stats"`
	UID           string          `json:"uid"`
	UpdatedAt     string          `json:"updatedAt"`
	Weight        Measurement     `json:"weight"`
}

// Measurement represents height or weight measurement
type Measurement struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

// UserPreferences represents user preferences
type UserPreferences struct {
	Notifications NotificationSettings `json:"notifications"`
	Privacy       PrivacySettings      `json:"privacy"`
	Units         string               `json:"units"`
}

// NotificationSettings represents notification preferences
type NotificationSettings struct {
	Achievements bool `json:"achievements"`
	Nutrition    bool `json:"nutrition"`
	Reminders    bool `json:"reminders"`
	Workout      bool `json:"workout"`
}

// PrivacySettings represents privacy preferences
type PrivacySettings struct {
	Profile  string `json:"profile"`
	Progress string `json:"progress"`
	Workouts string `json:"workouts"`
}

// UserStats represents user statistics
type UserStats struct {
	CurrentStreak int `json:"currentStreak"`
	LongestStreak int `json:"longestStreak"`
	TotalTime     int `json:"totalTime"`
	TotalVolume   int `json:"totalVolume"`
	TotalWorkouts int `json:"totalWorkouts"`
}
