package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/auth-api/controllers"
	"github.com/lozaeric/dupin/auth-api/oauth"
)

func setRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/token", oauth.HandleTokenRequest)
	router.POST("/authorize", oauth.HandleAuthorizeRequest)
	router.POST("/passwords", controllers.CreatePassword)
}
