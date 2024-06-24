package main

import (
	"fmt"
	"net/http"

	"github.com/avijeetpandey/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

// fetching all the events
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

// creating event
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	fmt.Println(err)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse request",
		})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, event)
}
