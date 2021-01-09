package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-resty/resty"
)

var authCli = resty.New().
	SetTimeout(100 * time.Millisecond).
	SetRetryCount(1).
	AddRetryCondition(func(r *resty.Response) (bool, error) {
		return r == nil || r.Error() != nil || r.StatusCode() >= 500, nil
	}).
	SetHostURL("http://auth:8080")

func CreatePassword(userID, password string) error {
	dto := map[string]string{
		"username": userID,
		"password": password,
	}
	r, err := authCli.R().SetBody(dto).Post("/passwords")
	if err != nil || r.StatusCode() != http.StatusCreated {
		return errors.New("auth-api error")
	}
	return nil
}
