package services

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirebaseService struct {
	client *firestore.Client
}

func NewFirebaseService() (*FirebaseService, error) {
	ctx := context.Background()

	// Check if we have a service account key file
	serviceAccountKey := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	var opt option.ClientOption

	if serviceAccountKey != "" {
		// Use service account key file
		opt = option.WithCredentialsFile(serviceAccountKey)
	} else {
		// Use default credentials (for local development)
		opt = option.WithCredentialsFile("serviceAccountKey.json")
	}

	// Get project ID from environment or use default
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		projectID = "your-project-id" // You'll need to set this
		log.Println("Warning: GOOGLE_CLOUD_PROJECT not set, using default project ID")
	}

	// Initialize Firestore client directly
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Printf("Error initializing Firestore client: %v", err)
		return nil, err
	}

	log.Println("Firestore client initialized successfully")
	return &FirebaseService{
		client: client,
	}, nil
}

// GetDocumentByID retrieves a document from Firestore by ID
func (fs *FirebaseService) GetDocumentByID(collection, docID string) (map[string]interface{}, error) {
	ctx := context.Background()

	doc, err := fs.client.Collection(collection).Doc(docID).Get(ctx)
	if err != nil {
		return nil, err
	}

	return doc.Data(), nil
}

// GetDocumentByIDFromCollection retrieves a document from Firestore by ID with custom collection
func (fs *FirebaseService) GetDocumentByIDFromCollection(collection, docID string) (map[string]interface{}, error) {
	ctx := context.Background()

	doc, err := fs.client.Collection(collection).Doc(docID).Get(ctx)
	if err != nil {
		return nil, err
	}

	return doc.Data(), nil
}

// Close closes the Firestore client
func (fs *FirebaseService) Close() error {
	return fs.client.Close()
}
