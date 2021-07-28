package handlers

import (
	"fmt"
	"net/http"

	"github.com/dotslashbin/happyhabits/api/models"
	"github.com/gin-gonic/gin"
)

func GetLogs(context *gin.Context) {
	fmt.Println("You called GetLogs...")
}

func SubmitFoodLog(context *gin.Context) {
	var food models.Food

	if err := context.ShouldBindJSON(&food); err != nil {
		// TODO: format this error
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": food})
}
