package controllers

import (
	"os"

	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/mock"
	"github.com/lozaeric/dupin/mongo"
	"github.com/lozaeric/dupin/redis"
)

var (
	messageStore domain.MessageStore
	userStore    domain.UserStore
)

func init() {
	if os.Getenv("ENV") == "production" {
		var err error
		userStore, err = redis.NewUserStore()
		if err != nil {
			panic(err)
		}
		messageStore, err = mongo.NewMessageStore()
		if err != nil {
			panic(err)
		}
	} else {
		userStore, _ = mock.NewUserStore()
		messageStore, _ = mock.NewMessageStore()
	}
}
