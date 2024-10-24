package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/gochat/src/domain"
	"github.com/lukinhas563/gochat/src/model/api/request"
	"github.com/lukinhas563/gochat/src/shared/service/logger"
	"go.uber.org/zap"
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
	logger.Info("Init Register from UserController", zap.String("journey", "Register"))

	var userRequest request.UserRegister
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error to validate user info", err, zap.String("journey", "Register"))

		c.JSON(http.StatusBadRequest, "Fields errors")
		return
	}

	if err := uc.domain.CreateUser(userRequest); err != nil {
		logger.Error("Error to save user info into database", err, zap.String("journey", "Register"))

		c.JSON(http.StatusInternalServerError, "Error to register")
		return
	}

	logger.Info("User registred successfully", zap.String("journey", "Register"))
	c.JSON(http.StatusOK, "Registered successfully")
}

func (uc *userController) Login(c *gin.Context) {
	logger.Info("Init Login from UserController", zap.String("journey", "Login"))

	var userLogin request.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Error to validate user info", err, zap.String("journey", "Login"))

		c.JSON(http.StatusBadRequest, "Fields errors")
		return
	}

	if err := uc.domain.LoginUser(userLogin); err != nil {
		logger.Error("Error to login the user", err, zap.String("journey", "Login"))

		c.JSON(http.StatusBadRequest, "Invalid password")
		return
	}

	logger.Info("User signined successfully", zap.String("journey", "Login"))
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
