package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	tokenStr := c.GetHeader("x-auth")
	if tokenStr == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "token is invalid",
		})
		return
	}

	claims, err := tokenChecker.Parse(tokenStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := tokenChecker.Token(claims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Set("token", token)
	c.Next()
}

func ParseToken(c *gin.Context) (*Token, error) {
	tk, found := c.Get("token")
	if !found {
		return nil, errors.New("token data doesnt exist")
	}
	return tk.(*Token), nil
}
