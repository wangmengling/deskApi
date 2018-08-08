package routes

import (
	"github.com/gin-gonic/gin"
	"deskapi/controllers"
)

func Routes(router *gin.Engine)  {
	router.GET("/",controllers.Index);

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//router.POST("/case", controllers.CaseCreate)
	//router.POST("/case/list/:pageIndex/:pageSize", controllers.CasesList)
	router.POST("/case/list", controllers.CasesList)
	router.POST("/case/detail",controllers.CasesDetail)
	router.POST("/color/list",controllers.ColorsList)
	router.POST("/style/list",controllers.StylesList)
}