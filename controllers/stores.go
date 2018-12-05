package controllers

import (
	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/mongo"
	"github.com/lozaeric/dupin/redis"
)

var (
	messageStore domain.MessageStore
	userStore    domain.UserStore
)

func init() {
	var err error
	userStore, err = redis.NewUserStore()
	if err != nil {
		panic(err)
	}
	messageStore, err = mongo.NewMessageStore()
	if err != nil {
		panic(err)
	}
}
