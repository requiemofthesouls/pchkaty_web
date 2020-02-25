package main

import (
	"pchkaty_web/backend/config"
	"pchkaty_web/backend/db"
	"pchkaty_web/backend/router"
)

func main() {
	config.LoadEnv()

	DB := db.InitDB()
	defer DB.Close()

	r := router.SetupRouter(DB)
	_ = r.Run(":8080")

}
