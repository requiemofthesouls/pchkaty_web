package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Pong tests that api is working
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func Main(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}
