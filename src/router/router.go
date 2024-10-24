package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/gochat/src/controller"
)

func InitRouter(router *gin.RouterGroup, userController controller.UserController) {
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.GET("/confirm", userController.Confirm)
	router.GET("/send", userController.Send)
	router.GET("/reset", userController.Reset)
}
