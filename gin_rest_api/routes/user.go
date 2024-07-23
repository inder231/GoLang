package routes

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func signup (c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body."})
        return
    }

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}