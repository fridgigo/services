package main

import (
	"github.com/fridgigo/sercies/handlers"
	"github.com/gin-gonic/gin"
)

// routes
func Routes() *gin.Engine {
	r := gin.Default()

	// api version nr 1
	v1 := r.Group("v1")
	{
		v1.GET("api/ping", handlers.Ping)
		v1.POST("/api/users/login", handlers.Login)
		v1.POST("/api/users/register", handlers.Register)
		v1.GET("/api/users/user-info", handlers.GetUser)
	}

	return r
}
