package apitest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/lozaeric/dupin/mock"

	"github.com/lozaeric/dupin/domain"
	"github.com/stretchr/testify/assert"
)

var (
	message = &domain.Message{
		Text:       "holaaa",
		SenderID:   "11111111111111111111",
		ReceiverID: "99999999999999999999",
	}
	otherMessage = &domain.Message{
		Text:       "holaaa",
		SenderID:   "11111111111111111111",
		ReceiverID: "88888888888888888888",
	}
)

func TestCreateMessage(t *testing.T) {
	assert := assert.New(t)
	cases := []*struct {
		expectedStatus int
		dto            *domain.Message
	}{
		{
			http.StatusBadRequest,
			&domain.Message{
				Text:       "chauuu",
				SenderID:   "",
				ReceiverID: "00000000000000000000",
			},
		},
		{
			http.StatusBadRequest,
			&domain.Message{
				Text:       "chauuu",
				SenderID:   "00000000000000000000",
				ReceiverID: "",
			},
		},
		{
			http.StatusBadRequest,
			&domain.Message{
				Text:       "",
				SenderID:   "00000000000000000000",
				ReceiverID: "00000000000000000001",
			},
		},
		{
			http.StatusOK,
			message,
		},
		{
			http.StatusOK,
			otherMessage,
		},
	}

	for _, c := range cases {
		r, err := cli.R().SetBody(c.dto).Post("/messages")
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			m := new(domain.Message)
			err := json.Unmarshal(r.Body(), m)
			assert.Nil(err)
			assert.NotEmpty(m.ID)
			assert.NotEmpty(m.DateCreated)
			assert.NotEmpty(m.DateUpdated)
			c.dto.ID = m.ID
			c.dto.DateCreated = m.DateCreated
			c.dto.DateUpdated = m.DateUpdated
			assert.Equal(c.dto, m)
		}
	}
}

func TestMessage(t *testing.T) {
	assert := assert.New(t)

	cases := []*struct {
		expectedStatus  int
		ID              string
		expectedMessage *domain.Message
	}{
		{
			http.StatusNotFound,
			mock.GenerateValidID(),
			nil,
		},
		{
			http.StatusBadRequest,
			"invalid",
			nil,
		},
		{
			http.StatusOK,
			message.ID,
			message,
		},
		{
			http.StatusOK,
			otherMessage.ID,
			otherMessage,
		},
	}

	for _, c := range cases {
		r, err := cli.R().Get("/messages/" + c.ID)
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			m := new(domain.Message)
			err := json.Unmarshal(r.Body(), m)
			assert.Nil(err)
			assert.Equal(c.expectedMessage, m)
		}
	}
}

func TestSearchMessage(t *testing.T) {
	assert := assert.New(t)

	cases := []*struct {
		expectedStatus   int
		params           string
		expectedMessages []*domain.Message
	}{
		{
			http.StatusNotFound,
			"field=receiver_id&id=00000000000000000000",
			nil,
		},
		{
			http.StatusOK,
			"field=sender_id&id=11111111111111111111",
			[]*domain.Message{message, otherMessage},
		},
		{
			http.StatusOK,
			"field=receiver_id&id=99999999999999999999",
			[]*domain.Message{message},
		},
	}

	for _, c := range cases {
		r, err := cli.R().Get("/search/messages?" + c.params)
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			messages := []*domain.Message{}
			err := json.Unmarshal(r.Body(), &messages)
			assert.Nil(err)
			assert.Equal(c.expectedMessages, messages)
		}
	}
}
func TestUpdateMessage(t *testing.T) {
	assert := assert.New(t)
	message.Seen = true

	cases := []*struct {
		expectedStatus  int
		ID              string
		messageDTO      map[string]interface{}
		expectedMessage *domain.Message
	}{
		{
			http.StatusOK,
			message.ID,
			map[string]interface{}{
				"seen": true,
			},
			message,
		},
		{
			http.StatusBadRequest,
			message.ID,
			map[string]interface{}{
				"id": "123",
			},
			nil,
		},
		{
			http.StatusNotFound,
			mock.GenerateValidID(),
			map[string]interface{}{},
			nil,
		},
	}

	for _, c := range cases {
		r, err := cli.R().SetBody(c.messageDTO).Put("/messages/" + c.ID)
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			m := new(domain.Message)
			err := json.Unmarshal(r.Body(), m)
			c.expectedMessage.DateUpdated = m.DateUpdated
			assert.Nil(err)
			assert.Equal(c.expectedMessage, m)
		}
	}
}
