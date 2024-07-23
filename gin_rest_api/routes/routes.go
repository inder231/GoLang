package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
		// Setting handler to handle http get request
		server.GET("/events", getEvents)

		// Getting single event
		server.GET("/events/:id", getSingleEvent)
	
		// Create a new event
		server.POST("/events", createEvent)

		// Update the event
		server.PUT("/events/:id", updateEvent)

		// Delete the event
		server.DELETE("/events/:id", deleteEvent)
}