package core

import (
	handlers "github.com/dotslashbin/happyhabits/api/handlers/get"
	"github.com/gin-gonic/gin"
)

/**
 * InitializeRoutes method that initializes the routes
 */
func InitializeRoutes(gin *gin.Engine) {
	gin.GET("/test", handlers.GetLogs)
}
