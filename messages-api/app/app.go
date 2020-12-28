package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/toolkit/utils"
)

var router *gin.Engine

func Run() {
	if utils.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router = gin.Default()
	setRoutes()
	router.Run()
}
