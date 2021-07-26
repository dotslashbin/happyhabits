package core

import (
	"github.com/gin-gonic/gin"
)

func Gin() *gin.Engine {
	framework := gin.Default()

	framework.Use(func(context *gin.Context) {

	})

	InitializeRoutes(framework)

	return framework

}
