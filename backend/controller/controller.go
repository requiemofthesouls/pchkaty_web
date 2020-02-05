package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"pchkaty_web/backend/db"
	"strconv"
)

type CreateUserForm struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Email      string `form:"email" json:"email" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	Name       string `form:"name" json:"name"`
	Surname    string `form:"surname" json:"surname"`
	Patronymic string `form:"patronymic" json:"patronymic"`
}

var user db.User
var users []db.User
var record db.Record

// Pong tests that database is working
func Pong(c *gin.Context) {
	err := db.DB.DB().Ping()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ping is not ok. connection with database lost!"})
	}
}

func Main(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

// Returns all users if not `user_id` specified.
// If you pass `user_id` query param it finds user with this user_id in database.
func GetUser(c *gin.Context) {
	qsId := c.Query("id")
	var usersObj *gorm.DB

	if qsId == "" {
		usersObj = db.DB.Find(&users)
		c.JSON(http.StatusOK, usersObj.Value)

	} else {
		userId, _ := strconv.ParseInt(qsId, 10, 64)
		usersObj = db.DB.Find(&user, userId)

		if usersObj.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": usersObj.Error.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{qsId: usersObj.Value})
		}
	}
}

// Creates user in database
func CreateUser(c *gin.Context) {
	var json CreateUserForm
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		userObj := db.User{
			Username:   json.Username,
			Email:      json.Email,
			Password:   json.Password,
			Name:       json.Name,
			Surname:    json.Surname,
			Patronymic: json.Patronymic}

		db.DB.Create(&userObj)
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func GetRecord(c *gin.Context) {
	records := db.DB.Find(&record)
	c.JSON(http.StatusOK, records.Value)
}
