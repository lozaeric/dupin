package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/controllers"
)

func setRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong ping")
	})
	router.GET("/users/:id", controllers.User)
	router.POST("/users", controllers.CreateUser)
	router.GET("/messages/:id", controllers.Message)
	router.POST("/messages", controllers.CreateMessage)
	router.GET("/search/messages", controllers.SearchMessages)
}
