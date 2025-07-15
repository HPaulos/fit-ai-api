package handlers

import (
	"net/http"

	"fit-ai-api/services"

	"github.com/gin-gonic/gin"
)

type FirestoreHandler struct {
	firebaseService *services.FirebaseService
}

func NewFirestoreHandler(firebaseService *services.FirebaseService) *FirestoreHandler {
	return &FirestoreHandler{
		firebaseService: firebaseService,
	}
}

// GetDocumentByID retrieves a document from Firestore by ID
func (h *FirestoreHandler) GetDocumentByID(c *gin.Context) {
	docID := c.Param("id")
	if docID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Document ID is required",
		})
		return
	}

	// Default collection name, can be overridden by query parameter
	collection := c.DefaultQuery("collection", "users")

	data, err := h.firebaseService.GetDocumentByIDFromCollection(collection, docID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Document not found",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"data":        data,
		"document_id": docID,
		"collection":  collection,
	})
}

// GetDocumentByIDWithCollection retrieves a document from a specific collection
func (h *FirestoreHandler) GetDocumentByIDWithCollection(c *gin.Context) {
	collection := c.Param("collection")
	docID := c.Param("id")

	if collection == "" || docID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Collection name and document ID are required",
		})
		return
	}

	data, err := h.firebaseService.GetDocumentByIDFromCollection(collection, docID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Document not found",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"data":        data,
		"document_id": docID,
		"collection":  collection,
	})
}
