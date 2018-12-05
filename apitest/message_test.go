package apitest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/mock"
	"github.com/stretchr/testify/assert"
)

var (
	message     = new(domain.Message)
	messageJSON = ""
)

func TestCreateMessage(t *testing.T) {
	assert := assert.New(t)
	invalidMessageDTO := &domain.Message{
		Text:       "",
		SenderID:   "",
		ReceiverID: "2",
	}
	validMessageDTO := &domain.Message{
		Text:       "holaaa",
		SenderID:   "12345678901234567404",
		ReceiverID: "12345678901234567890",
	}

	r, err := cli.R().SetBody(invalidMessageDTO).Post("/messages")
	assert.Nil(err)
	assert.Equal(http.StatusBadRequest, r.StatusCode())

	r, err = cli.R().SetBody(validMessageDTO).Post("/messages")
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())

	err = json.Unmarshal(r.Body(), message)
	assert.Nil(err)
	assert.Equal(validMessageDTO.Text, message.Text)
	assert.Equal(validMessageDTO.SenderID, message.SenderID)
	assert.Equal(validMessageDTO.ReceiverID, message.ReceiverID)
	assert.NotEmpty(message.DateCreated)
	messageJSON = string(r.Body())
}

func TestMessage(t *testing.T) {
	assert := assert.New(t)

	r, err := cli.R().Get("/messages/" + mock.GenerateValidID())
	assert.Nil(err)
	assert.Equal(http.StatusNotFound, r.StatusCode())

	r, err = cli.R().Get("/messages/" + message.ID)
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())
	assert.Equal(messageJSON, string(r.Body()))
}

/*
func TestSearchMessage(t *testing.T) {
	assert := assert.New(t)
	validMessageDTO := &domain.Message{
		Text:       "chauu",
		SenderID:   "12345678901234567404",
		ReceiverID: "12345678901234567890",
	}

	r, err := cli.R().SetBody(validMessageDTO).Post("/messages")
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())

	r, err = cli.R().Get("/search/messages?receiver_id=12345678901234567890")
	var messages []*domain.Message
	err = json.Unmarshal(r.Body(), messages)
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())
	assert.Equal(2, len(messages))
}
*/
