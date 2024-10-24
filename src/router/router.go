package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/gochat/src/controller"
	"github.com/lukinhas563/gochat/src/model/database/sqlite"
)

func InitRouter(router *gin.RouterGroup, database sqlite.SqliteDatabase) {
	userController := controller.NewUserController(database)

	router.POST("/register", userController.Register)
	router.GET("/login", userController.Login)
	router.GET("/confirm", userController.Confirm)
	router.GET("/send", userController.Send)
	router.GET("/reset", userController.Reset)
}
