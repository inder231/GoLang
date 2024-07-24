package main

import (
	"goblogart/inits"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()

	if err := inits.DBInit(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
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