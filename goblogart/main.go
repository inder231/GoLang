package main

import (
	"goblogart/inits"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
}

func main() {
	
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run()
}