package routes

import (
	"net/http"
	"rest-api/models"
	"rest-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents ( c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return     // return from the function if there is an error in fetching events from db     }
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent ( c *gin.Context ) {

	token := c.Request.Header.Get("Authorization")
	
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
	}
	
	// Validate token
	err := utils.VerifyToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// New event
	var event models.Event
	// Bind the json data to event variable
	// c.BindJSON(&event)
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	event.ID = 1
	event.UserID = 1

	// Save the event to db
	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return     // return from the function if there is an error in saving event to db     }
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Created new event.", "event": event})
}

func getSingleEvent(c *gin.Context)  {
	// Fetch event from db
    // and return it as JSON

	// Extract param request object and convert int64 type
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}
	if event == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event not found."})
        return
    }

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}

	_, err = models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body!"})
        return
	}

	updatedEvent.ID = id

	err = updatedEvent.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}

	err = event.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})

}