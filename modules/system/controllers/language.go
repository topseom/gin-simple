package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/topseom/gin-simple/modules/system/forms"
)

type LanguageController struct{}

func (ctrl LanguageController) FindAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "findAll language",
	})
}

func (ctrl LanguageController) FindOne(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "findOne",
		"id":      id,
	})
}

func (ctrl LanguageController) Create(c *gin.Context) {

	var body forms.LanguageForm

	if c.ShouldBindJSON(&body) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		return
	}
	c.JSON(200, gin.H{
		"message": "create",
		"body":    body,
	})
}
func (ctrl LanguageController) Remove(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "remove",
		"id":      id,
	})
}
