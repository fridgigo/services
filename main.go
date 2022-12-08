package main

import (
	"log"
	"os"

	"github.com/fridgigo/services/handlers"
	"github.com/fridgigo/services/handlers/user"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// PORT
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Creating a new Server with gin
	router := gin.New()
	router.Use(gin.Logger())

	// user struct instance
	user := &user.User{}

	// Routers
	router.GET("/ping", handlers.Ping)
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/api/users/login", user.Login)
		v1.POST("/api/users/register", user.Register)
		v1.GET("/api/users/user-info", user.GetUser)
	}

	router.Run(":" + port)
}
