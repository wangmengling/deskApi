package main

import (
	"github.com/gin-gonic/gin"
	"deskApi/routes"
)

func main() {
	router := gin.Default()
	router.Static("static", "./static")
	routes.Routes(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
