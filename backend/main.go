package main

import (
	"pchkaty_web/backend/config"
	"pchkaty_web/backend/db"
	"pchkaty_web/backend/router"
)

func init() {
	config.LoadEnv()
	db.Connect()
}

func main() {
	r := router.SetupRouter()
	_ = r.Run(":8080")

}
