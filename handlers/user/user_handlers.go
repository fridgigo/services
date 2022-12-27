package user

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fridgigo/services/helper"
	"github.com/fridgigo/services/internal/driver"
	"github.com/fridgigo/services/middleware"

	"github.com/gin-gonic/gin"
)

/*
*********
post method
user sign in function
*/
func (u *User) Login(c *gin.Context) {

	// check for request body
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check for User struct object
	if u.Email == "" && u.Password == "" {
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
	user, err := db.Query(fmt.Sprintf("SELECT email, password, confirmed FROM public.users where email = '%s' ", u.Email))
	if err != nil {
		log.Println("Error on sql query: ", err)
		return
	}
	var email, password string
	var confirmed bool

	for user.Next() {
		err = user.Scan(&email, &password, &confirmed)
		if err != nil {
			log.Panic("Error on user finding: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"Message": "Something went wrong. Please try again."})
			return
		}
	}
	// check if user not confirmed, then return back error message
	if confirmed == false {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "This user has not confirmed their account."})
		return
	}

	// check for password hash
	pass := helper.CheckPasswordHash(u.Password, password)
	if pass != true {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "The password you have entered is incorrect."})
		return
	}

	// TODO:
	jwt, err := middleware.GenerateJWT(email)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"Message": "The user has successfully logged in.", "Token": jwt})
	return
}

/*
*********
post method
user sign up function
*/
func (u *User) Register(c *gin.Context) {

	// check for request body
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
	// hash the password
	u.Password, err = helper.HashPassword(u.Password)
	if err != nil {
		log.Println("there is an error password hashing:", err)
	}

	// INSERT-query executing
	_, err = db.Exec(fmt.Sprintf("INSERT INTO users (email, first_name, last_name, password, confirmed, deleted, confirmation_number) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v')", u.Email, u.FirstName, u.LastName, u.Password, u.Confirmed, u.Deleted, u.ConfirmationNumber))
	if err != nil {
		log.Println("Error in inserting: ", err)
		return
	}

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
	// check for request body
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check for User struct object
	if u.Email == "" && u.ConfirmationNumber == 0 {
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
	user, err := db.Query(fmt.Sprintf("SELECT email, confirmation_number FROM public.users where email = '%s' ", u.Email))
	if err != nil {
		log.Println("Error on sql query: ", err)
		return
	}

	var email string
	var confirmationNumber int

	for user.Next() {
		err = user.Scan(&email, &confirmationNumber)
		if err != nil {
			log.Panic("Error on user finding: ", err)
			return
		}
	}

	// check is confirmation number is correct
	if u.ConfirmationNumber != confirmationNumber {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Something went wrong. Please try again."})
		return
	}

	// confirm user, set user confirmed as true
	// Prepare the update statement
	stmt, err := db.Prepare(fmt.Sprintf("UPDATE users SET confirmed = 'true' WHERE email = '%s'", u.Email))
	if err != nil {
		log.Println("Error (update): ", err)
	}
	defer stmt.Close()

	// Execute the update statement
	_, err = stmt.Exec()
	if err != nil {
		log.Println("Error (execute update): ", err)
	}

	// return back
	c.JSON(http.StatusOK, gin.H{"Message": "User was successfully confirmed."})
	return
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
