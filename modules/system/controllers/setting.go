package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/topseom/gin-simple/modules/system/forms"
)

type SettingController struct{}

func (ctrl SettingController) FindAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "findAll setting",
	})
}

func (ctrl SettingController) FindOne(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "findOne",
		"id":      id,
	})
}

func (ctrl SettingController) Create(c *gin.Context) {

	var body forms.SettingForm

	if c.ShouldBindJSON(&body) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		return
	}
	c.JSON(200, gin.H{
		"message": "create",
		"body":    body,
	})
}
func (ctrl SettingController) Remove(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "remove",
		"id":      id,
	})
}
