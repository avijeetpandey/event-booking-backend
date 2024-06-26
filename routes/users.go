package routes

import (
	"net/http"

	"github.com/avijeetpandey/event-booking/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to parse request",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to create user",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"data":    user,
	})
}
