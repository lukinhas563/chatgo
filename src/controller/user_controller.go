package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/gochat/src/domain"
	"github.com/lukinhas563/gochat/src/model/api/request"
)

type UserController interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Confirm(*gin.Context)
	Send(*gin.Context)
	Reset(*gin.Context)
}

type userController struct {
	domain domain.UserDomain
}

func NewUserController(domain domain.UserDomain) *userController {
	return &userController{
		domain: domain,
	}
}

func (uc *userController) Register(c *gin.Context) {
	var userRequest request.UserRegister
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Fields errors")
		return
	}

	if err := uc.domain.CreateUser(userRequest); err != nil {
		c.JSON(http.StatusInternalServerError, "Error to register")
		return
	}

	c.JSON(http.StatusOK, "Registered successfully")
}

func (uc *userController) Login(c *gin.Context) {
	var userLogin request.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, "Fields errors")
		return
	}

	if err := uc.domain.LoginUser(userLogin); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid password")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"RESULT": "User Login",
	})
}

func (uc *userController) Confirm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"RESULT": "User Confirm",
	})
}

func (uc *userController) Send(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"RESULT": "User Send reset",
	})
}

func (uc *userController) Reset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"RESULT": "User reset password",
	})
}
