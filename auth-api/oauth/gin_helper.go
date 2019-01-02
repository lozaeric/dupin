package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleAuthorizeRequest the authorization request handling
func HandleAuthorizeRequest(c *gin.Context) {
	err := srv.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

// HandleTokenRequest token request handling
func HandleTokenRequest(c *gin.Context) {
	err := srv.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

// HandleTokenVerify Verify the access token of the middleware
func HandleTokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		ti, err := srv.ValidationBearerToken(c.Request)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("token", ti)
		c.Next()
	}
}
