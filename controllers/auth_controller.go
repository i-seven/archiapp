package controllers

import (
	"backendAf/models"
	"backendAf/services"
	"backendAf/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signupPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name"`
}

type loginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context) {
	var p signupPayload
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "INVALID_PAYLOAD", err.Error())
		return
	}

	user := models.User{
		Email:    p.Email,
		Password: p.Password,
		Name:     p.Name,
	}
	u, err := services.UserService.SignUp(user)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "SIGNUP_FAILED", err.Error())
		return
	}
	c.JSON(http.StatusCreated, u)
}

func Login(c *gin.Context) {
	var p loginPayload
	if err := c.ShouldBindJSON(&p); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "INVALID_PAYLOAD", err.Error())
		return
	}

	token, user, err := services.UserService.Login(p.Email, p.Password)
	if err != nil {
		utils.JSONError(c, http.StatusUnauthorized, "AUTH_FAILED", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func Me(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, user)
}
