package core

import (
	"github.com/gin-gonic/gin"
)

func Gin() *gin.Engine {
	framework := gin.Default()

	// Firebase
	firebaseAuth := SetupFirebase()

	framework.Use(func(context *gin.Context) {
		context.Set("firebaseAuth", firebaseAuth)
	})

	InitializeRoutes(framework)

	return framework

}
