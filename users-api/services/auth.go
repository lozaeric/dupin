package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-resty/resty"
)

var authCli = resty.New().
	SetTimeout(100 * time.Millisecond).
	SetHostURL("http://auth:8080")

func CreatePassword(userID, password string) error {
	dto := map[string]string{
		"user_id":  userID,
		"password": password,
	}
	r, err := authCli.R().SetBody(dto).Post("/passwords")
	if err != nil || r.StatusCode() != http.StatusOK {
		return errors.New("auth-api error")
	}
	return nil
}
