package app

import (
	"net/http"

	"github.com/lozaeric/dupin/toolkit/auth"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/messages-api/controllers"
)

func setRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Group("/messages", auth.Middleware).
		GET("/:id", controllers.Message).
		POST("", controllers.CreateMessage).
		PUT("/:id", controllers.UpdateMessage)
	router.GET("/search/messages", controllers.SearchMessages)
}
