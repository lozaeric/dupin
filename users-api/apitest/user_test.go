package apitest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Deleted     bool   `json:"deleted"`
}

var (
	user = &UserDTO{
		Name:     "eric",
		LastName: "loza",
		Email:    "eric@lz.com",
		Password: "123",
	}
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	cases := []*struct {
		expectedStatus int
		dto            *UserDTO
	}{
		{
			http.StatusBadRequest,
			&UserDTO{
				Name:     "",
				LastName: "loza",
				Email:    "eric@lz.com",
				Password: "123",
			},
		},
		{
			http.StatusBadRequest,
			&UserDTO{
				Name:     "eric",
				LastName: "",
				Email:    "eric@lz.com",
				Password: "123",
			},
		},
		{
			http.StatusBadRequest,
			&UserDTO{
				Name:     "eric",
				LastName: "loza",
				Email:    "",
				Password: "123",
			},
		},
		{
			http.StatusBadRequest,
			&UserDTO{
				Name:     "eric",
				LastName: "loza",
				Email:    "eric@lz.com",
				Password: "",
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
			u := new(UserDTO)
			err := json.Unmarshal(r.Body(), u)
			assert.Nil(err)
			assert.NotEmpty(u.ID)
			assert.NotEmpty(u.DateCreated)
			assert.NotEmpty(u.DateUpdated)
			assert.Empty(u.Password)
			c.dto.ID = u.ID
			c.dto.DateCreated = u.DateCreated
			c.dto.DateUpdated = u.DateUpdated
			c.dto.Password = ""
			assert.Equal(c.dto, u)
		}
	}
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	cases := []*struct {
		expectedStatus int
		ID             string
		expectedUser   *UserDTO
	}{
		{
			http.StatusNotFound,
			"00000000000000000001",
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
			u := new(UserDTO)
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
		token          string
		expectedUser   *UserDTO
	}{
		{
			http.StatusOK,
			user.ID,
			map[string]string{
				"last_name": "LZ",
			},
			fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, user.ID),
			user,
		},
		{
			http.StatusBadRequest,
			user.ID,
			map[string]string{
				"id": "123",
			},
			fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, user.ID),
			nil,
		},
		{
			http.StatusBadRequest,
			user.ID,
			map[string]string{
				"name": "",
			},
			fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, user.ID),
			nil,
		},
		{
			http.StatusNotFound,
			"00000000000000000002",
			map[string]string{},
			fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, "00000000000000000002"),
			nil,
		},
		{
			http.StatusForbidden,
			"00000000000000000003",
			map[string]string{},
			fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, "00000000000000000002"),
			nil,
		},
	}
	for _, c := range cases {
		r, err := cli.R().SetHeader("x-auth", c.token).SetBody(c.userDTO).Put("/users/" + c.ID)
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			u := new(UserDTO)
			err := json.Unmarshal(r.Body(), u)
			c.expectedUser.DateUpdated = u.DateUpdated
			assert.Nil(err)
			assert.Equal(c.expectedUser, u)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	cases := []*struct {
		expectedStatus int
		ID             string
		token          string
		expectedUser   *UserDTO
	}{
		{
			http.StatusOK,
			user.ID,
			fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, user.ID),
			user,
		},
		{
			http.StatusNotFound,
			"00000000000000000002",
			fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, "00000000000000000002"),
			nil,
		},
		{
			http.StatusBadRequest,
			user.ID,
			fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, user.ID),
			nil,
		},
	}
	for _, c := range cases {
		r, err := cli.R().SetHeader("x-auth", c.token).Post("/users/" + c.ID + "/delete")
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			u := new(UserDTO)
			err := json.Unmarshal(r.Body(), u)
			c.expectedUser.DateUpdated = u.DateUpdated
			assert.Nil(err)
			assert.Equal(c.expectedUser, u)
			assert.True(c.expectedUser.Deleted)
		}
	}
}
