package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetLogs(context *gin.Context) {
	fmt.Println("You called GetLogs...")
}
