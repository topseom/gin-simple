package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/topseom/gin-simple/modules/system/controllers"
)

type LanguageRoute struct{}

func (routes LanguageRoute) GetRoutes(r *gin.Engine) {
	ctrl := new(controllers.LanguageController)
	route := r.Group("/v1/system/languages")
	route.GET("/", ctrl.FindAll)
	route.GET("/:id", ctrl.FindOne)
	route.POST("/", ctrl.Create)
	route.DELETE("/:id", ctrl.Remove)
}
