package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/messages-api/controllers"
	"github.com/lozaeric/dupin/toolkit/auth"
)

func setRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/messages", controllers.CreateMessage)
	router.PUT("/messages/:id", controllers.UpdateMessage)
	router.GET("/search/messages", controllers.SearchMessages)
	// auth middleware
	router.Group("/messages", auth.Middleware).
		GET("/:id", controllers.Message)
}
