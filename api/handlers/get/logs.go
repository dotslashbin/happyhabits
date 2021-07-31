package handlers

import (
	"github.com/dotslashbin/happyhabits/api/services/food"
	"net/http"

	"github.com/dotslashbin/happyhabits/api/models"
	"github.com/gin-gonic/gin"
)

func GetLogs(context *gin.Context) {
	//fmt.Println("You called GetLogs...")
}

func SubmitFoodLog(context *gin.Context) {
	var food models.Food

	if err := context.ShouldBindJSON(&food); err != nil {
		// TODO: format this error
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := srvfoodlog.FoodLogger{}

	service.CreateLog()

	context.JSON(http.StatusOK, gin.H{"data": food})
}