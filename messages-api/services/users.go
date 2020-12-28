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
	SetRetryCount(1).
	AddRetryCondition(func(r *resty.Response) (bool, error) {
		return r == nil || r.Error() != nil || r.StatusCode() >= 500, nil
	}).
	SetHostURL("http://users:8080")

func User(ID string) (*domain.User, error) {
	if user, err := getFromCache(ID); err == nil {
		return user, nil
	}

	r, err := usersCli.R().Get("/users/" + ID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode() != http.StatusOK {
		return nil, errors.New("status code: " + strconv.Itoa(r.StatusCode()))
	}

	user := new(domain.User)
	err = json.Unmarshal(r.Body(), user)
	if err == nil {
		saveToCache(user)
	}
	return user, err
}
