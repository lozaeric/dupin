package controllers

import (
	"github.com/lozaeric/dupin/messages-api/domain"
	"github.com/lozaeric/dupin/messages-api/mongo"
)

var messageStore domain.MessageStore

func init() {
	var err error
	messageStore, err = mongo.NewMessageStore()
	if err != nil {
		panic(err)
	}
}
