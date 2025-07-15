package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fit-ai-api/handlers"
	"fit-ai-api/models"
	"fit-ai-api/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db, err := initDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run database migrations
	if err := models.AutoMigrate(db); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Fit AI API is running",
		})
	})

	// Initialize Firebase service
	firebaseService, err := services.NewFirebaseService()
	if err != nil {
		log.Printf("Warning: Firebase service initialization failed: %v", err)
		log.Println("Firestore endpoints will not be available")
		firebaseService = nil
	}

	// Initialize AI service
	aiService := services.NewAIService()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(db)
	var firestoreHandler *handlers.FirestoreHandler
	var aiHandler *handlers.AIHandler
	if firebaseService != nil {
		firestoreHandler = handlers.NewFirestoreHandler(firebaseService)
		aiHandler = handlers.NewAIHandler(firebaseService, aiService)
	}

	// API routes group
	api := r.Group("/api/v1")
	{
		// User endpoints
		api.GET("/users", userHandler.GetUsers)
		api.GET("/users/:id", userHandler.GetUser)
		api.POST("/users", userHandler.CreateUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)

		// Firestore endpoints
		if firestoreHandler != nil {
			api.GET("/firestore/:id", firestoreHandler.GetDocumentByID)
			api.GET("/firestore/collection/:collection/:id", firestoreHandler.GetDocumentByIDWithCollection)
		}

		// AI Workout Plan endpoints
		if aiHandler != nil {
			api.POST("/ai/workout-plan/:user_id", aiHandler.GenerateWorkoutPlan)
			api.GET("/ai/workout-plan/:plan_id", aiHandler.GetWorkoutPlanByID)
			api.PUT("/ai/workout-plan/:plan_id", aiHandler.UpdateWorkoutPlan)
			api.DELETE("/ai/workout-plan/:plan_id", aiHandler.DeleteWorkoutPlan)
			api.GET("/ai/workout-plans/:user_id", aiHandler.GetUserWorkoutPlans)
		}

		// TODO: Add your fitness-related endpoints here
		api.GET("/workouts", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Workouts endpoint - coming soon!",
			})
		})
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Use Docker container connection string
		dsn = "host=localhost user=postgres password=postgres dbname=fit_ai_db port=5432 sslmode=disable TimeZone=UTC"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}
