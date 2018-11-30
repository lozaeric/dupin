package apitest

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	r, err := cli.R().SetBody(user).Post("/users")
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusOK)
	assert.Equal(string(r.Body()), userJSON)

	r, err = cli.R().SetBody(invalidUser).Post("/users")
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusBadRequest)
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	r, err := cli.R().Get("/users/404")
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusNotFound)

	r, err = cli.R().Get("/users/1")
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusOK)
	assert.Equal(string(r.Body()), userJSON)
}
