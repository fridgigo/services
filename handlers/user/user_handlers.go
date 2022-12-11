package user

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fridgigo/services/helper"
	"github.com/fridgigo/services/internal/driver"

	"github.com/gin-gonic/gin"
)

/*
*********
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
*********
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

	// connect to db, if there is an error, then panic it
	db, err := driver.Connect()
	if err != nil {
		log.Println("Error on connection: ", err)
	}
	// close db connection at the end
	defer db.Close()

	// check for user if user exits in db
	// if user exits in db
	// return the error message back as a response
	user, err := db.Query(fmt.Sprintf("SELECT email FROM public.users where email = '%s' ", u.Email))
	if err != nil {
		log.Println("Error on sql query: ", err)
		return
	}
	for user.Next() {
		var email string
		err = user.Scan(&email)
		if err != nil {
			log.Panic("Error on user finding: ", err)
		}
		c.JSON(http.StatusBadRequest, gin.H{"Message": "This user is exists."})
		return
	}

	// if user is not exists in db
	// complete user object
	// then insert it into db
	u.CreatedAt = time.Now().Unix() // must changed
	u.UpdatedAt = time.Now().Unix() // must improved
	// generate random number
	u.ConfirmationNumber = helper.RandomNumber()

	// INSERT-query executing
	_, err = db.Exec(fmt.Sprintf("INSERT INTO users (email, first_name, last_name, password, confirmed, deleted) VALUES ('%v', '%v', '%v', '%v', '%v', '%v')", u.Email, u.FirstName, u.LastName, u.Password, u.Confirmed, u.Deleted))
	if err != nil {
		log.Println("Error in inserting: ", err)
		return
	}

	// TODO:
	// send an email to user for confirm his/her account
	err = helper.SendMail(u.Email, "Confirm your Account.", u.ConfirmationNumber)
	if err != nil {
		log.Println("Error is sending email: ", err)
	}

	// return back success message
	c.JSON(http.StatusOK, gin.H{"Message": "Your account has been successfully created. Please confirm your email address."})
}

/*
*********
post method
verify users account
**********
*/
func (u *User) ConfirmUser(c *gin.Context) {
	// TODOs:
	// check for user
	// update confirmation status (set true)
	// generate a token (jwt)
}

/*
*********
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
