package controller

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/gochat/src/model/api/request"
	"github.com/lukinhas563/gochat/src/model/database/sqlite"
)

type userController struct {
	database sqlite.SqliteDatabase
}

func NewUserController(database sqlite.SqliteDatabase) *userController {
	return &userController{
		database: database,
	}
}

func (uc *userController) Register(c *gin.Context) {
	var userRequest request.UserRegister
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Fields errors")
		return
	}

	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userRequest.Password))
	userRequest.Password = hex.EncodeToString(hash.Sum(nil))

	if err := uc.database.InsertUser(userRequest); err != nil {
		c.JSON(http.StatusInternalServerError, "Error to register the user")
		return
	}

	c.JSON(http.StatusOK, "Registered successfully")
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
