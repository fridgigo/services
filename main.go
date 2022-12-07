package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/fridgigo/services/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// PORT
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Creating a new Server with gin
	router := gin.New()
	router.Use(gin.Logger())

	// Routers
	router.GET("/ping", handlers.Ping)
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/api/users/login", handlers.Login)
		v1.POST("/api/users/register", handlers.Register)
		v1.GET("/api/users/user-info", handlers.GetUser)
	}

	router.Run(":" + port)
}
