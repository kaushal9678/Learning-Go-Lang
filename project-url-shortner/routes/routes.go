package routes

import (
	"project-url-shortner/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	    // POST route for creating a new short URL
    server.POST("/shorten", handler.CreateShortURL)

    // GET route for redirecting from a short URL to the original URL
    // The `:shortUrl` part is a path parameter that Gin will capture.
    server.GET("/:shortUrl", handler.RedirectToLongURL)
}