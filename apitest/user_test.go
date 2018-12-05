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
			c.dto.ID = u.ID
			c.dto.DateCreated = u.DateCreated
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
