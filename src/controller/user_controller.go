package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
}

func NewUserController() *userController {
	return &userController{}
}

func (uc *userController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"RESULT": "User Register",
	})
}

func (uc *userController) Login(c *gin.Context) {
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
