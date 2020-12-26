package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
	"pchkaty_web/controller"
	"pchkaty_web/controller/user"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
)

func initSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}

// SetupRouter setup routing here
func SetupRouter(db *gorm.DB) *gin.Engine {
	initSentry()

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(sentrygin.New(sentrygin.Options{Repanic: true}))

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
