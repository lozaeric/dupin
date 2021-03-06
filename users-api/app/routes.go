package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/toolkit/auth"
	"github.com/lozaeric/dupin/users-api/controllers"
)

func setRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/users/:id", controllers.User)
	router.POST("/users", controllers.CreateUser)

	router.Group("/users", auth.Middleware).
		PUT("/:id", controllers.UpdateUser)
}
