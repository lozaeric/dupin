package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty"
)

var authCli = resty.New().
	SetTimeout(50 * time.Millisecond).
	SetHostURL("http://auth:8080")

type Token struct {
	ID             string   `json:"id"`
	ApplicationID  string   `json:"application_id"`
	UserID         string   `json:"user_id"`
	Scopes         []string `json:"scopes"`
	DateCreated    string   `json:"date_created"`
	ExpirationDate string   `json:"expiration_date"`
}

func (t *Token) HasScope(scope string) bool {
	for _, s := range t.Scopes {
		if s == scope {
			return true
		}
	}
	return false
}

func Middleware(c *gin.Context) {
	tokenID := c.GetHeader("x-auth")
	if tokenID == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "token is invalid",
		})
		return
	}

	r, err := authCli.R().Get("/tokens/" + tokenID)
	if r.StatusCode() == http.StatusNotFound || r.StatusCode() == http.StatusBadRequest {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "token is invalid",
		})
		return
	}
	if err != nil || r.StatusCode() != http.StatusOK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "auth-api error",
		})
		return
	}
	token := new(Token)
	json.Unmarshal(r.Body(), token)
	c.Set("token", token)
	c.Next()
}

func ParseToken(c *gin.Context) (*Token, error) {
	tk, found := c.Get("token")
	if !found {
		return nil, errors.New("token data doesnt exist")
	}
	token, ok := tk.(*Token)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	return token, nil
}