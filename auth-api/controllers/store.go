package controllers

import (
	"github.com/lozaeric/dupin/auth-api/domain"
	"github.com/lozaeric/dupin/auth-api/nanoid"
	"github.com/lozaeric/dupin/auth-api/redis"
)

var (
	tokenStore     domain.TokenStore
	tokenGenerator domain.TokenGenerator
)

func init() {
	var err error
	tokenStore, err = redis.NewTokenStore()
	if err != nil {
		panic(err)
	}
	tokenGenerator = new(nanoid.TokenGenerator)
}
