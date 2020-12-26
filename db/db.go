package db

import (
	_ "github.com/jackc/pgx/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"pchkaty_web/config"
)

func InitDB() *gorm.DB {

	db, err := gorm.Open("postgres", config.PostgresAddr)

	if err != nil {
		log.Panicf("%s | Failed to connect database | %s", config.PostgresAddr, err)
	} else {
		log.Info("DB Connection ok!")
	}

	db.AutoMigrate(&User{}, &Record{})

	return db
}
