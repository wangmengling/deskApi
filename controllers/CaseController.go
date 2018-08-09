package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"deskapi/models"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2/bson"

	"strconv"
)

// Create an case
func CaseCreate(c *gin.Context) {
	fmt.Print("dsds")
	db := c.MustGet("db").(*mgo.Database)
	cases := models.Cases{}
	err := c.Bind(&cases)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionArticle).Insert(cases)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/cases")
}


func CasesList(c *gin.Context)  {

	db := c.MustGet("db").(*mgo.Database)
	//查询所有数据
	cases := []models.Cases{}
	pageSize := c.PostForm("pageSize")
	pageIndex := c.PostForm("pageIndex")
	size, _ := strconv.Atoi(pageSize)
	index, _ := strconv.Atoi(pageIndex)
	query := bson.M{}
	color := c.PostForm("color")
	style := c.PostForm("style")
	fmt.Print(color)
	if len(color) > 0 {
		query["color"] = color
	}

	if len(style) > 0 {
		query["style"] = style
	}

	fmt.Print(query)
	db.C(models.CollectionCases).Find(&query).Limit(size).Skip(index).All(&cases)

	count,_ := db.C(models.CollectionCases).Count()
	c.JSON(200, gin.H{
		"code":1,

		"data": gin.H{
			"list":cases,
			"pageTotal":count,
			"count":count,
			},
	})
}

func CasesDetail(c *gin.Context)  {
	db := c.MustGet("db").(*mgo.Database)
	//查询所有数据
	cases := models.Cases{}
	oID := bson.ObjectIdHex(c.PostForm("_id"))

	db.C(models.CollectionCases).FindId(oID).One(&cases)
	fmt.Print(cases)
	c.JSON(200, gin.H{
		"code":1,

		"data": cases,
	})
}