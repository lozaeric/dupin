package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Run() {
	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router = gin.Default()
	setRoutes()
	router.Run()
}
