package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret = os.Getenv("SECRET_JWT")

var tokenChecker TokenChecker

type Token struct {
	ClientID string `json:"client_id"`
	UserID   string `json:"user_id"`
	Scope    string `json:"scope"`
	//	ExpiratedAt int64  `json:"expiration_date"`
}

type TokenChecker interface {
	Parse(string) (jwt.Claims, error)
	Token(jwt.Claims) (*Token, error)
}

type JWTTokenChecker struct {
	secret string
}

func (t *JWTTokenChecker) Parse(tokenStr string) (jwt.Claims, error) {
	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return t.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, errors.New("token expired")
	}
	return jwtToken.Claims, nil
}

func (t *JWTTokenChecker) Token(claims jwt.Claims) (*Token, error) {
	if claims, ok := claims.(jwt.MapClaims); ok {
		return &Token{
			ClientID: claims["client_id"].(string),
			UserID:   claims["application_id"].(string),
			//ExpiratedAt: claims["expired_at"].(int64),
			Scope: claims["scope"].(string),
		}, nil
	}
	return nil, errors.New("invalid token")
}

type MockTokenChecker struct{}

func (t *MockTokenChecker) Parse(tokenStr string) (jwt.Claims, error) {
	claims := make(jwt.MapClaims)
	err := json.Unmarshal([]byte(tokenStr), &claims)
	return claims, err
}

func (t *MockTokenChecker) Token(claims jwt.Claims) (*Token, error) {
	if claims, ok := claims.(jwt.MapClaims); ok {
		return &Token{
			ClientID: claims["client_id"].(string),
			UserID:   claims["application_id"].(string),
			//ExpiratedAt: claims["expired_at"].(int64),
			Scope: claims["scope"].(string),
		}, nil
	}
	return nil, errors.New("invalid token")
}

func init() {
	if os.Getenv("ENV") == "production" {
		tokenChecker = &JWTTokenChecker{
			secret: secret,
		}
	} else {
		tokenChecker = new(MockTokenChecker)
	}
}
