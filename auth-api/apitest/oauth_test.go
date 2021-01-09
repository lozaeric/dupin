package apitest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	assert := assert.New(t)
	dto := map[string]string{
		"username": "eric",
		"password": "holamundo",
	}
	r, err := cli.R().SetBody(dto).Post("/passwords")
	assert.Nil(err)
	assert.Equal(http.StatusCreated, r.StatusCode())
}

func TestToken(t *testing.T) {
	assert := assert.New(t)

	cases := []*struct {
		expectedStatus int
		data           map[string]string
	}{
		{
			http.StatusOK,
			map[string]string{
				"grant_type":    "client_credentials",
				"client_id":     "123123123",
				"client_secret": "111222333",
				"scope":         "read",
			},
		},
		{
			http.StatusOK,
			map[string]string{
				"grant_type":    "client_credentials",
				"client_id":     "123123123",
				"client_secret": "111222333",
				"scope":         "read",
				"username":      "eric",
				"password":      "holamundo",
			},
		},
		{
			http.StatusUnauthorized,
			map[string]string{
				"grant_type":    "client_credentials",
				"client_id":     "123123123",
				"client_secret": "000000000",
				"scope":         "read",
			},
		},
	}

	for _, c := range cases {
		r, err := cli.R().SetFormData(c.data).Post("/token")
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if r.StatusCode() == http.StatusOK {
			dto := make(map[string]interface{})
			assert.Nil(json.Unmarshal(r.Body(), &dto))
			token, ok := dto["access_token"].(string)
			assert.True(ok)
			assert.NotEmpty(token)
		}
	}
}
