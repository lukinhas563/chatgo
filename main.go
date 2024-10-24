package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lukinhas563/gochat/src/controller"
	"github.com/lukinhas563/gochat/src/domain"
	"github.com/lukinhas563/gochat/src/model/database/sqlite"
	"github.com/lukinhas563/gochat/src/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	DB_PATH := os.Getenv("DB_PATH")
	if DB_PATH == "" {
		panic("Environment DB_PATH not defined")
	}

	server := gin.Default()

	database := sqlite.NewSqliteDatabase()
	if err := database.Connect(DB_PATH); err != nil {
		panic(err)
	}
	defer database.Close()
	fmt.Println("Connected on database")

	userDomain := domain.NewUserDomain(database)
	userController := controller.NewUserController(userDomain)

	router.InitRouter(&server.RouterGroup, userController)

	server.Run()
}
