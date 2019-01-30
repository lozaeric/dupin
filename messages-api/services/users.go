package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty"
	"github.com/lozaeric/dupin/messages-api/domain"
)

var usersCli = resty.New().
	SetTimeout(50 * time.Millisecond).
	SetHostURL("http://users:8080")

func User(ID string) (*domain.User, error) {
	r, err := usersCli.R().Get("/users/" + ID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode() != http.StatusOK {
		return nil, errors.New("status code: " + strconv.Itoa(r.StatusCode()))
	}

	user := new(domain.User)
	return user, json.Unmarshal(r.Body(), user)
}