package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"deskapi/models"
)

func ColorsList(c *gin.Context)  {
	db := c.MustGet("db").(*mgo.Database)
	//查询所有数据
	colors := []models.Colors{}
	db.C(models.CollectionColors).Find(nil).All(&colors)

	count,_ := db.C(models.CollectionColors).Count()
	c.JSON(200, gin.H{
		"code":1,

		"data": gin.H{
			"list":colors,
			"pageTotal":count,
			"count":count,
		},
	})
}