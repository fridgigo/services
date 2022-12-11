package handlers

import (
	"github.com/gin-gonic/gin"
)

/*
check server whether is reachable
*/
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
