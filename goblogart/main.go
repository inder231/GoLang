package main

import (
	"goblogart/controllers"
	"goblogart/inits"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()

	/* MongoDB connection commented. */
	// if err := inits.MongoDBInit(); err != nil {
	// 	log.Fatal("Could not connect to MongoDB")
	// }

	inits.MySqlDBInit()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}