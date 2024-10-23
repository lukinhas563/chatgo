package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/gochat/src/router"
)

func main() {
	server := gin.Default()

	router.InitRouter(&server.RouterGroup)

	server.Run()
}
