package domain

import (
	"github.com/lozaeric/dupin/domain/validation"
	"github.com/lozaeric/dupin/utils"
)

var messageUpdatable = map[string]bool{
	"seen": true,
}

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
