package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Run() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	setRoutes()
	router.Run()
}
