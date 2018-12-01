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
	user     = new(domain.User)
	userJSON = ""
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	invalidUserDTO := &domain.User{
		Name:     "",
		LastName: "a",
		Email:    "b",
	}
	validUserDTO := &domain.User{
		Name:     "eric",
		LastName: "loza",
		Email:    "eric@lz.com",
	}

	r, err := cli.R().SetBody(invalidUserDTO).Post("/users")
	assert.Nil(err)
	assert.Equal(http.StatusBadRequest, r.StatusCode())

	r, err = cli.R().SetBody(validUserDTO).Post("/users")
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())

	err = json.Unmarshal(r.Body(), user)
	assert.Nil(err)
	assert.Equal(validUserDTO.Name, user.Name)
	assert.Equal(validUserDTO.LastName, user.LastName)
	assert.Equal(validUserDTO.Email, user.Email)
	assert.NotEmpty(user.DateCreated)
	userJSON = string(r.Body())
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	r, err := cli.R().Get("/users/" + mock.GenerateValidID())
	assert.Nil(err)
	assert.Equal(http.StatusNotFound, r.StatusCode())

	r, err = cli.R().Get("/users/" + user.ID)
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())
	assert.Equal(userJSON, string(r.Body()))
}
