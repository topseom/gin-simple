package config

import (
	"github.com/gin-gonic/gin"

	"github.com/topseom/gin-simple/modules/system"
)

func LoadRoutes(r *gin.Engine) {
	system.LoadRoutes(r)
}
