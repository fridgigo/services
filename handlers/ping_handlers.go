package handlers

import (
	"github.com/fridgigo/services/internal/driver"
	"github.com/gin-gonic/gin"
)

/*
check server whether is reachable
*/
func Ping(c *gin.Context) {
	driver.Connect()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
