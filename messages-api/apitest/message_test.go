package apitest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/lozaeric/dupin/toolkit/mock"

	"github.com/lozaeric/dupin/messages-api/domain"
	"github.com/stretchr/testify/assert"
)

var (
	validUserID = "11111111111111111111"
	message     = &domain.Message{
		Text:       "hello!!",
		ReceiverID: "99999999999999999999",
	}
	otherMessage = &domain.Message{
		Text:       "good bye!",
		ReceiverID: "88888888888888888888",
	}
)

func TestCreateMessage(t *testing.T) {
	assert := assert.New(t)
	cases := []*struct {
		expectedStatus int
		userID         string
		dto            *domain.Message
	}{
		{
			http.StatusBadRequest,
			validUserID,
			&domain.Message{
				Text:       "chauuu",
				ReceiverID: "",
			},
		},
		{
			http.StatusBadRequest,
			validUserID,
			&domain.Message{
				Text:       "",
				ReceiverID: "00000000000000000001",
			},
		},
		{
			http.StatusOK,
			validUserID,
			message,
		},
		{
			http.StatusOK,
			validUserID,
			otherMessage,
		},
	}

	for _, c := range cases {
		token := validTokens[c.userID]
		r, err := cli.R().SetBody(c.dto).
			SetHeader("x-auth", token).
			Post("/messages")
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
			c.dto.SenderID = m.SenderID
			assert.Equal(c.dto, m)
		}
	}
}

func TestMessage(t *testing.T) {
	assert := assert.New(t)

	cases := []*struct {
		expectedStatus  int
		userID          string
		expectedMessage *domain.Message
	}{
		{
			http.StatusNotFound,
			validUserID,
			&domain.Message{ID: mock.GenerateValidID()},
		},
		{
			http.StatusBadRequest,
			validUserID,
			&domain.Message{ID: "invalid"},
		},
		{
			http.StatusOK,
			validUserID,
			message,
		},
		{
			http.StatusOK,
			validUserID,
			otherMessage,
		},
	}

	for _, c := range cases {
		token := validTokens[c.userID]
		r, err := cli.R().SetBody(c.expectedMessage).
			SetHeader("x-auth", token).
			Get("/messages/" + c.expectedMessage.ID)
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
		userID, params   string
		expectedMessages []*domain.Message
	}{
		{
			http.StatusNotFound,
			validUserID,
			"field=receiver_id&value=00000000000000000000",
			nil,
		},
		{
			http.StatusOK,
			validUserID,
			"",
			[]*domain.Message{message, otherMessage},
		},
		{
			http.StatusOK,
			validUserID,
			"field=receiver_id&value=99999999999999999999",
			[]*domain.Message{message},
		},
	}

	for _, c := range cases {
		token := validTokens[c.userID]
		r, err := cli.R().SetHeader("x-auth", token).
			Get("/search/messages?" + c.params)
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
		userID          string
		messageDTO      map[string]interface{}
		expectedMessage *domain.Message
	}{
		{
			http.StatusBadRequest,
			validUserID,
			map[string]interface{}{
				"receiver_id": "123",
			},
			message,
		},
		{
			http.StatusOK,
			validUserID,
			map[string]interface{}{
				"seen": true,
			},
			message,
		},
	}

	for _, c := range cases {
		token := validTokens[c.userID]
		r, err := cli.R().SetBody(c.messageDTO).
			SetHeader("x-auth", token).
			Put("/messages/" + c.expectedMessage.ID)
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
