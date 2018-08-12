package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"deskapi/models"
)

func StylesList(c *gin.Context)  {
	db := c.MustGet("db").(*mgo.Database)
	//查询所有数据
	styles := []models.Styles{}
	db.C(models.CollectionStyles).Find(nil).All(&styles)

	count,_ := db.C(models.CollectionStyles).Count()
	c.JSON(200, gin.H{
		"code":1,

		"data": gin.H{
			"list":styles,
			"pageTotal":count,
			"count":count,
		},
	})
}