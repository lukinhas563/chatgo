package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/gochat/src/controller"
)

func InitRouter(router *gin.RouterGroup) {
	userController := controller.NewUserController()

	router.GET("/register", userController.Register)
	router.GET("/login", userController.Login)
	router.GET("/confirm", userController.Confirm)
	router.GET("/send", userController.Send)
	router.GET("/reset", userController.Reset)
}
