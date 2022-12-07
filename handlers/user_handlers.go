package handlers

import (
	"net/http"

	"github.com/fridgigo/services/models"
	"github.com/gin-gonic/gin"
)

/*
post method
user sign in function
*/
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// connect to db and check for user exists
	// if user exits and password is correct
	// then generate a json web token (jwt)
	// return it back as a response
	// if user not exists in db or password is wrong
	// then send back error message as a response
}

/*
post method
user sign up function
*/
func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// connect to db and check for user exists
	// if user exits in db
	// return the error message back as a response
	// if user not exists in db
	// then insert in into db
}

/*
get method
get user infos
*/
func GetUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// connect to db and check for user exists
	// if user exits
	// then response the user datas
	// if user not exist
	// then return back the error message as a response
}
