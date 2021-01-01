package system

import (
	"github.com/gin-gonic/gin"

	"github.com/topseom/gin-simple/modules/system/routes"
)

func LoadRoutes(r *gin.Engine) {
	new(routes.LanguageRoute).GetRoutes(r)
	new(routes.SettingRoute).GetRoutes(r)
}
