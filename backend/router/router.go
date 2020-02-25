package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"pchkaty_web/backend/controller"
	"pchkaty_web/backend/controller/user"
)

// SetupRouter setup routing here
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	// Users
	userAPI := user.InitUserAPI(db)
	router.GET("/users", userAPI.FindAll)
	router.GET("/users/:id", userAPI.FindByID)
	router.POST("/users", userAPI.Create)
	router.PUT("/users/:id", userAPI.Update)
	router.DELETE("/users/:id", userAPI.Delete)

	router.GET("/", controller.Main)
	router.GET("/ping", controller.Pong)

	return router
}
