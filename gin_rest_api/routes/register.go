package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)


func registerForEvent (c *gin.Context) {

	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
        return
	}
	
	event, err := models.GetEventById(eventId)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}

	err = event.Register(userId)
	fmt.Println("err registering---", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event."})
        return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
	
}

func cancelRegistration (c *gin.Context) {
	
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id."})
        return
	}
	
	event, err := models.GetEventById(eventId)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event."})
		return
	}
	
	err = event.CancelRegistration(userId)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel registration for event."})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cancelled registration!"})

}