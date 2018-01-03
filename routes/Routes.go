package routes

import (
	"github.com/gin-gonic/gin"
	"deskApi/controllers"
)

func Routes(router *gin.Engine)  {
	router.GET("/",controllers.Index);

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}