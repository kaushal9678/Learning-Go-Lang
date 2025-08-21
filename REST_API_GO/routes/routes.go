package routes

import (
	"example.com/rest-api-go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	//1st approach simple to apply  this middleware.Authenticate middleware of token validation
	//server.POST("/events",middleware.Authenticate, createEvents)

	//2nd Approach when have multiple routes, in this Group we are 
	// using Use method 
	// to apply middleware, it will now execute all group
	// related routes with authentication token
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEventById)
	authenticated.DELETE("/events/:id", deleteEventById)
	
	server.GET("/events/:id", getEventById)
	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.GET("/users/:email", getUserByEmail)
	server.POST("/signin", signin)
}