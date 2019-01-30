package oauth

import (
	"encoding/base64"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	oauth2 "gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/generates"

	"gopkg.in/oauth2.v3/utils/uuid"
)

var secret = os.Getenv("SECRET_JWT")

var SigningMethod = jwt.SigningMethodHS256

type TokenGenerate struct {
	jwtgen *generates.JWTAccessGenerate
}

type CustomClaims struct {
	AccessClaims *generates.JWTAccessClaims
	Scope        string `json:"scope"`
}

func (c *CustomClaims) Valid() error {
	return c.AccessClaims.Valid()
}

func (t *TokenGenerate) Token(data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
	jwtClaims := &generates.JWTAccessClaims{
		ClientID:  data.Client.GetID(),
		UserID:    data.UserID,
		ExpiredAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
	}
	claims := &CustomClaims{
		AccessClaims: jwtClaims,
		Scope:        data.TokenInfo.GetScope(),
	}

	token := jwt.NewWithClaims(t.jwtgen.SignedMethod, claims)
	access, err = token.SignedString(t.jwtgen.SignedKey)
	if err != nil {
		return
	}
	if isGenRefresh {
		refresh = base64.URLEncoding.EncodeToString(uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).Bytes())
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}
	return
}

func newTokenGenerate() *TokenGenerate {
	return &TokenGenerate{
		jwtgen: generates.NewJWTAccessGenerate([]byte(secret), SigningMethod),
	}
}
