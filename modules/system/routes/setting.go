package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/topseom/gin-simple/modules/system/controllers"
)

type SettingRoute struct{}

func (routes SettingRoute) GetRoutes(r *gin.Engine) {
	ctrl := new(controllers.SettingController)
	route := r.Group("/v1/system/settings")
	route.GET("/", ctrl.FindAll)
	route.GET("/:id", ctrl.FindOne)
	route.POST("/", ctrl.Create)
	route.DELETE("/:id", ctrl.Remove)
}
