package controller

import (
	"golangservices/entity"
	"golangservices/request"
	"golangservices/response"
	"golangservices/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.AuthResponse{
			Message: "Invalid request",
			Status:  false,
		})
		return
	}

	token, err := service.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusOK, response.AuthResponse{
			Message: err.Error(),
			Status:  false,
		})
		return
	}

	c.JSON(http.StatusCreated, response.AuthResponse{
		JWT:     token,
		Message: "Registration was successful",
		Status:  true,
	})
}

func Login(c *gin.Context) {
	var loginRequest request.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.AuthResponse{
			Message: "Invalid request",
			Status:  false,
		})
		return
	}

	token, err := service.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.AuthResponse{
			Message: err.Error(),
			Status:  false,
		})
		return
	}

	c.JSON(http.StatusOK, response.AuthResponse{
		JWT:     token,
		Message: "Login successful",
		Status:  true,
	})
}
