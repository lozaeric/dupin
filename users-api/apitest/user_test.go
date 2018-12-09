package apitest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/lozaeric/dupin/toolkit/mock"
	"github.com/lozaeric/dupin/users-api/domain"
	"github.com/stretchr/testify/assert"
)

var (
	user = &domain.User{
		Name:     "eric",
		LastName: "loza",
		Email:    "eric@lz.com",
	}
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	cases := []*struct {
		expectedStatus int
		dto            *domain.User
	}{
		{
			http.StatusBadRequest,
			&domain.User{
				Name:     "",
				LastName: "loza",
				Email:    "eric@lz.com",
			},
		},
		{
			http.StatusBadRequest,
			&domain.User{
				Name:     "eric",
				LastName: "",
				Email:    "eric@lz.com",
			},
		},
		{
			http.StatusBadRequest,
			&domain.User{
				Name:     "eric",
				LastName: "loza",
				Email:    "",
			},
		},
		{
			http.StatusOK,
			user,
		},
	}

	for _, c := range cases {
		r, err := cli.R().SetBody(c.dto).Post("/users")
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			u := new(domain.User)
			err := json.Unmarshal(r.Body(), u)
			assert.Nil(err)
			assert.NotEmpty(u.ID)
			assert.NotEmpty(u.DateCreated)
			assert.NotEmpty(u.DateUpdated)
			c.dto.ID = u.ID
			c.dto.DateCreated = u.DateCreated
			c.dto.DateUpdated = u.DateUpdated
			assert.Equal(c.dto, u)
		}
	}
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	cases := []*struct {
		expectedStatus int
		ID             string
		expectedUser   *domain.User
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
			user.ID,
			user,
		},
	}

	for _, c := range cases {
		r, err := cli.R().Get("/users/" + c.ID)
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			u := new(domain.User)
			err := json.Unmarshal(r.Body(), u)
			assert.Nil(err)
			assert.Equal(c.expectedUser, u)
		}
	}
}
func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)
	user.LastName = "LZ"

	cases := []*struct {
		expectedStatus int
		ID             string
		userDTO        map[string]string
		expectedUser   *domain.User
	}{
		{
			http.StatusOK,
			user.ID,
			map[string]string{
				"last_name": "LZ",
			},
			user,
		},
		{
			http.StatusBadRequest,
			user.ID,
			map[string]string{
				"id": "123",
			},
			nil,
		},
		{
			http.StatusBadRequest,
			user.ID,
			map[string]string{
				"name": "",
			},
			nil,
		},
		{
			http.StatusNotFound,
			mock.GenerateValidID(),
			map[string]string{},
			nil,
		},
	}
	for _, c := range cases {
		r, err := cli.R().SetBody(c.userDTO).Put("/users/" + c.ID)
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			u := new(domain.User)
			err := json.Unmarshal(r.Body(), u)
			c.expectedUser.DateUpdated = u.DateUpdated
			assert.Nil(err)
			assert.Equal(c.expectedUser, u)
		}
	}
}
