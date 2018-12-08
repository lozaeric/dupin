package controllers

import (
	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/mongo"
)

var messageStore domain.MessageStore

func init() {
	var err error
	messageStore, err = mongo.NewMessageStore()
	if err != nil {
		panic(err)
	}
}
