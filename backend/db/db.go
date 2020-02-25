package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
	"pchkaty_web/backend/config"
)

func InitDB() *gorm.DB {

	db, err := gorm.Open("postgres", config.PostgresAddr)

	if err != nil {
		panic(fmt.Sprintf("%s | Failed to connect database", config.PostgresAddr))
	} else {
		log.Printf("DB Connection ok!")
	}

	db.AutoMigrate(&User{}, &Record{})

	return db
}
