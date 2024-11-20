package routes

import (
	"net/http"
	"strconv"

	"traunseenet.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("user_id")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse event id", "error": err.Error()})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event", "error": err.Error()})
		return
	}

	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event, try again later", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event!"})
}

func unregisterFromEvent(context *gin.Context) {
	userID := context.GetInt64("user_id")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse event id", "error": err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not unregister user from event, try again later", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Unregistered from event!"})
}