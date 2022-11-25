package handlers

import "github.com/gin-gonic/gin"

/*
post method
user sign in function
*/
func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login endpoint",
	})
}

/*
post method
user sign up function
*/
func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register endpoint",
	})
}
