package main

import (
	"greeter/db"
	"greeter/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", checkServer)
	routes.RegisterRoutes(server)
	server.Run(":8081")
}

func checkServer(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is running"})
}
