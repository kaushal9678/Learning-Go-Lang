package handler

import (
	"math/rand"
	"net/http"
	"project-url-shortner/models"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = NewSeededRand()// NewSeededRand creates a new rand.Rand seeded with the current time.
func NewSeededRand() *rand.Rand {
	// Seed the random number generator with the current time
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GenerateShortURL generates a random short URL of the specified length.
func GenerateShortURL(length int) string {
	// Create a byte slice to hold the random characters
	b := make([]byte, length)
	// Fill the byte slice with random characters from the charset
	for i := range b {
		// Pick a random character from the charset	
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	// Convert the byte slice to a string and return it
	return string(b)
}
func CreateShortURL(context *gin.Context) {
	var request models.ShortenRequest
	if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		shortURL := GenerateShortURL(7) // Generate a random short URL of length 7
		err := models.InsertURL(shortURL, request.URL) // Insert the URL mapping into the database
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to create short URL"})
			return
		}
		response := models.ShortenResponse{ShortURL: shortURL}
		context.JSON(http.StatusOK, response)
	}	
func RedirectToLongURL(context *gin.Context){
	shortURL := context.Param("shortUrl")
	longURL, err := models.GetLongURL(shortURL)
	if err != nil {
		context.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Short URL not found"})
		return
	}
	context.Redirect(http.StatusMovedPermanently, longURL)
}
