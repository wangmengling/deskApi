package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"deskapi/models"
	"net/http"
)

// Create an case
func CaseCreate(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	cases := models.Case{}
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
