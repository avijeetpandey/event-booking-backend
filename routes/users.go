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

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to parse request",
		})
		return
	}

	// validating the credentials of the user

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials please try again with valid ones",
		})
		return
	}

	// logic to generate auth token

	// response
	context.JSON(http.StatusOK, gin.H{
		"message": "logged in successfully",
	})
}
