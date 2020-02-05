package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
	"pchkaty_web/backend/config"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open("postgres", config.PostgresAddr)

	if err != nil {
		panic(fmt.Sprintf("%s | Failed to connect database", config.PostgresAddr))
	} else {
		log.Printf("DB Connection ok!")
	}

	db.AutoMigrate(&User{}, &Record{})

	DB = db

}
