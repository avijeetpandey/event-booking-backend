package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/avijeetpandey/event-booking/db"
	"github.com/avijeetpandey/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("events/:id", getEventById)

	server.Run(":8080")
}

// fetching all the events
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not parse request",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "successfully done",
		"data":    events,
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

	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events try again later",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "event created successfully",
		"data":    event,
	})
}

// fetching an event
func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "event fetched successfully",
		"data":    event,
	})

}
