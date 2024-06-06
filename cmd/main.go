package main

import (
	"golangservices/config"
	"golangservices/controller"
	"golangservices/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	r := gin.Default()
	r.Use(middleware.Logger())

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}

	protected := r.Group("/api/v1")
	protected.Use(middleware.JWTAuth())
	{
		// Define protected routes here
	}

	r.Run(":8001")
}
