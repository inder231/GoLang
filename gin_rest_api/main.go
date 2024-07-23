package main

import (
	"rest-api/db"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	/* Initialize the DB */
	db.InitDB()

	/* Create Instance of a server */
	server := gin.Default()

	/* register routes */
	routes.RegisterRoutes(server)

	// Run server on required port 
	server.Run(":8080") // localhost:8080

}
