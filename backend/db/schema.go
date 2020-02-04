package db

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"type:varchar(100);unique_index"`
}
