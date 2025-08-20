package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.GET("/events/:id", getEventById)
	server.PUT("/events/:id", updateEventById)
	server.DELETE("/events/:id", deleteEventById)
	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.GET("/users/:email", getUserByEmail)
	server.POST("/signin", signin)
}