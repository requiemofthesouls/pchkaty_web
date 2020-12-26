package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home/home.html", gin.H{})
}

func EditHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/ckeditor.html", gin.H{})
}

func AdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/admin.html", gin.H{})
}
