package main

import (
	"deskapi/config"
	"deskapi/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Connect()
}

func main() {
	router := gin.Default()
	router.Static("static", "./static")
	routes.Routes(router)

	router.Run("3002") // listen and serve on 0.0.0.0:8080
}
