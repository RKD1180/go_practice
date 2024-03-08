package main

import (
	"greeter/db"
	"greeter/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", checkServer)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8081")
}

func checkServer(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is running"})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBind(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	event.Id = 1
	event.UserId = 1

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}
