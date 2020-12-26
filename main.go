package main

import (
	"pchkaty_web/config"
	"pchkaty_web/db"
	"pchkaty_web/router"
)

func init() {
	config.LoadEnv()
}

func main() {
	DB := db.InitDB()
	defer DB.Close()

	r := router.SetupRouter(DB)
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	_ = r.Run(":8080")

}
