package routes

import (
	"github.com/avijeetpandey/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)

	// authenticated server group
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.GET("/events/:id", getEventById)

	// user routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
