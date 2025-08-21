package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api-go/models"
	"example.com/rest-api-go/utils"
	"github.com/gin-gonic/gin"
)
func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
func createEvents(context *gin.Context){
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}
	userId,err := utils.VerifyToken(token); if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save event"})
		return
	}
	//event.ID = 1;
	event.UserID = userId;

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event":event})
}

func getEventById(context * gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64);
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch Event"})
		return
	}
	context.JSON(http.StatusOK, event)
}
func updateEventById(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10,64);
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventById(eventId);
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch Event"})
		return
	}
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	// Bind the request body to the event model
	// This will update the event with the new values provided in the request body
	// The ID will remain the same as the one in the URL
	// and the event will be updated in the database
	// The UserID will also remain the same as the one in the event model
	// This is because we are updating an existing event, not creating a new one
	// The UserID is not provided in the request body, so it will remain the same
	// as the one in the event model
	// This is a PUT request, so we expect the entire event to be updated
	// If you want to update only specific fields, you can use PATCH instead of PUT
	// and only provide the fields you want to update	
	var updatedEvent models.Event
	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEvent.ID = eventId
	
	if err := updatedEvent.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": event})
}
func deleteEventById(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10,64);
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventById(eventId);
	fmt.Println("Event:", event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch Event"})
		return
	}
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	if err := event.DeleteEvent(); err != nil {
		fmt.Println("Error deleting event:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
