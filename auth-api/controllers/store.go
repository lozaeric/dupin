package controllers

import (
	"github.com/lozaeric/dupin/auth-api/domain"
	"github.com/lozaeric/dupin/auth-api/redis"
)

var (
	tokenStore      domain.TokenStore
	secureInfoStore domain.SecureInfoStore
)

func init() {
	var err error
	tokenStore, err = redis.NewTokenStore()
	if err != nil {
		panic(err)
	}
	secureInfoStore, err = redis.NewPasswordStore()
	if err != nil {
		panic(err)
	}
}
