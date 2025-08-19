package main

import (
	"net/http"

	"example.com/rest-api-go/models"
	"github.com/gin-gonic/gin"
)
func main(){
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.Run(":8080") //localhost:8080
}
func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
func createEvents(context *gin.Context){
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save event"})
		return
	}
	event.ID = 1;
	event.UserID = 1;

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event":event})
}