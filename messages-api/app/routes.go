package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/messages-api/controllers"
)

func setRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong ping")
	})
	router.GET("/messages/:id", controllers.Message)
	router.POST("/messages", controllers.CreateMessage)
	router.PUT("/messages/:id", controllers.UpdateMessage)
	router.GET("/search/messages", controllers.SearchMessages)
}
