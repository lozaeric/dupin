package apitest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/lozaeric/dupin/domain"
	"github.com/stretchr/testify/assert"
)

var (
	user     = new(domain.User)
	userJSON = ""
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	invalidUserDTo := &domain.User{
		Name:     "",
		LastName: "a",
		Email:    "b",
	}
	validUserDTO := &domain.User{
		Name:     "eric",
		LastName: "loza",
		Email:    "eric@loza",
	}

	r, err := cli.R().SetBody(invalidUserDTo).Post("/users")
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusBadRequest)

	r, err = cli.R().SetBody(validUserDTO).Post("/users")
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusOK)

	err = json.Unmarshal(r.Body(), &user)
	assert.Nil(err)
	assert.Equal(validUserDTO.Name, user.Name)
	assert.Equal(validUserDTO.LastName, user.LastName)
	assert.Equal(validUserDTO.Email, user.Email)
	userJSON = string(r.Body())
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	r, err := cli.R().Get("/users/0")
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusNotFound)

	r, err = cli.R().Get("/users/" + user.ID)
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusOK)
	assert.Equal(string(r.Body()), userJSON)
}
