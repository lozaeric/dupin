package domain

import (
	"github.com/lozaeric/dupin/toolkit/utils"
	"github.com/lozaeric/dupin/toolkit/validation"
)

var (
	messageInstance  = new(Message)
	messageUpdatable = map[string]bool{
		"seen": true,
	}
	messageSearchable = map[string]bool{
		"receiver_id": true,
		"sender_id":   true,
		"seen":        true,
	}
)

type Message struct {
	ID          string `json:"id" bson:"id"`
	Text        string `json:"text" bson:"text" validate:"required"`
	SenderID    string `json:"sender_id" bson:"sender_id" validate:"required,len=20,alphanum"`
	ReceiverID  string `json:"receiver_id" bson:"receiver_id" validate:"required,len=20,alphanum"`
	DateCreated string `json:"date_created" bson:"date_created"`
	DateUpdated string `json:"date_updated" bson:"date_updated"`
	Seen        bool   `json:"seen" bson:"seen"`
}

type MessageStore interface {
	Message(string) (*Message, error)
	Create(*Message) error
	Update(*Message) error
	Search(field, value string) ([]*Message, error)
}

func (m *Message) Update(values map[string]interface{}) (err error) {
	err = validation.Update(m, values, messageUpdatable)
	if err != nil {
		return
	}
	m.DateUpdated = utils.Now()
	return
}

func CheckMessageValues(values map[string]interface{}) error {
	return validation.Check(messageInstance, values, messageSearchable)
}
