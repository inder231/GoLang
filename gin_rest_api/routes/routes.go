package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
		// Setting handler to handle http get request
		server.GET("/events", getEvents)

		// Getting single event
		server.GET("/events/:id", getSingleEvent)
	
		// Group authenticated middlware in one group
		authenticated := server.Group("/")
		// Attach middlware using Use method
		authenticated.Use(middlewares.Authenticate)
		// Create Event
		authenticated.POST("/events", createEvent)
		// Update Event
		authenticated.PUT("/events/:id", updateEvent)
		// Delete Event
		authenticated.DELETE("/events/:id", deleteEvent)

		// // Create a new event
		// server.POST("/events", middlewares.Authenticate, createEvent)

		// // Update the event
		// server.PUT("/events/:id", updateEvent)

		// // Delete the event
		// server.DELETE("/events/:id", deleteEvent)

		// Signup route
		server.POST("/signup", signup)

		// Login route
		server.POST("/login", login)
}