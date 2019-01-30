package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = os.Getenv("SECRET_JWT")

type Token struct {
	ClientID    string `json:"client_id"`
	UserID      string `json:"user_id"`
	Scope       string `json:"scope"`
	ExpiratedAt int64  `json:"expiration_date"`
}

func Middleware(c *gin.Context) {
	tokenStr := c.GetHeader("x-auth")
	if tokenStr == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "token is invalid",
		})
		return
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "token is invalid",
		})
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tk := &Token{
			ClientID:    claims["client_id"].(string),
			UserID:      claims["application_id"].(string),
			ExpiratedAt: claims["expired_at"].(int64),
			Scope:       claims["scope"].(string),
		}
		c.Set("token", tk)
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": "token is invalid",
	})
}

func ParseToken(c *gin.Context) (*Token, error) {
	tk, found := c.Get("token")
	if !found {
		return nil, errors.New("token data doesnt exist")
	}
	return tk.(*Token), nil
}
