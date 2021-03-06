package db

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(100)"`
	Email      string `gorm:"type:varchar(100);unique_index"`
	Password   string `gorm:"type:varchar(50)"`
	Name       string `gorm:"type:varchar(50)"`
	Surname    string `gorm:"type:varchar(50)"`
	Patronymic string `gorm:"type:varchar(50)"`
	Superuser  bool   `gorm:"type:bool"`
}

type Record struct {
	gorm.Model
	User    User
	UserID  uint
	Content string
}

type HTMLContent struct {
	gorm.Model
	Content string `gorm:"type:text"`
}
