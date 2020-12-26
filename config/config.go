package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var DbUser string
var DbPassword string
var DbName string
var DbHost string
var DbPort string
var PostgresAddr string
var DbSslMode string

func LoadEnv() {
	_ = godotenv.Load()

	DbUser = os.Getenv("DBUSER")
	DbPassword = os.Getenv("DBPASSWORD")
	DbName = os.Getenv("DBNAME")
	DbHost = os.Getenv("DBHOST")
	DbPort = os.Getenv("DBPORT")
	DbSslMode = os.Getenv("DBSSLMODE")

	PostgresAddr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DbHost, DbPort, DbUser, DbPassword, DbName, DbSslMode)

}
