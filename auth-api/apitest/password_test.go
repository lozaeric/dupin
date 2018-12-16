package apitest

import (
	"net/http"
	"testing"

	"github.com/lozaeric/dupin/toolkit/mock"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	assert := assert.New(t)
	dto := map[string]string{
		"user_id":  mock.GenerateValidID(),
		"password": "holamundo",
	}
	r, err := cli.R().SetBody(dto).
		Post("/passwords")
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())

	r, err = cli.R().SetBody(dto).
		Post("/passwords/validate")
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())

	dto["password"] = "chaumundo"
	r, err = cli.R().SetBody(dto).
		Post("/passwords/validate")
	assert.Nil(err)
	assert.Equal(http.StatusBadRequest, r.StatusCode())
}
