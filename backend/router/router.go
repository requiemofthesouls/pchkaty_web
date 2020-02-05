package router

import (
	"github.com/gin-gonic/gin"
	"pchkaty_web/backend/controller"
)

// SetupRouter setup routing here
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// routes
	router.GET("/", controller.Main)
	router.GET("/ping", controller.Pong)

	router.GET("/user", controller.GetUser)
	router.POST("/user", controller.CreateUser)

	router.GET("/records", controller.GetRecord)
	return router
}
