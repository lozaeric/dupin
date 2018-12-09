package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/auth-api/controllers"
)

func setRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//router.GET("/application/:id", controllers.Application)
	//router.POST("/application", controllers.CreateApplication)
	router.GET("/tokens/:id", controllers.Token)
	router.POST("/tokens", controllers.CreateToken)
}
