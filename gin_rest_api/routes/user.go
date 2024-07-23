package routes

import (
	"net/http"
	"rest-api/models"
	"rest-api/utils"

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

func login(c *gin.Context) {
	// Implement login logic here
    var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body."})
        return
	}
	// Validate password will validate password and attach user's id with user struct instance
	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate."})
        return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}