package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
post method
user sign in function
*/
func (u *User) Login(c *gin.Context) {

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"Message": "Success"})

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
func (u *User) Register(c *gin.Context) {

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check for User struct object
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Password == "" || u.Password != u.RepeatPassword {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Something went wrong. Please try again."})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"Message": "Success"})
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
func (u *User) GetUser(c *gin.Context) {

	if err := c.ShouldBindJSON(&u); err != nil {
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
