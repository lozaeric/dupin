package passwords

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lozaeric/dupin/toolkit/apierr"
	"golang.org/x/crypto/bcrypt"

	"github.com/lozaeric/dupin/auth-api/domain"
	"github.com/lozaeric/dupin/auth-api/redis"
)

var passwordStore = redis.NewPasswordStore()

const BcryptCost = 7

func Create(data []byte) *apierr.ApiError {
	dto := make(map[string]string)
	if err := json.Unmarshal(data, &dto); err != nil {
		return apierr.New(http.StatusBadRequest, "invalid values")
	}
	username, password := dto["username"], dto["password"]
	if username == "" || password == "" { // password need validations and user must exist
		return apierr.New(http.StatusBadRequest, "invalid values")
	}

	hash, err := bcryptHash(password)
	if err != nil {
		return apierr.New(http.StatusInternalServerError, err.Error())
	}
	pwd := &domain.Password{
		Username:       username,
		HashedPassword: hash,
	}
	if err := passwordStore.Save(pwd); err != nil {
		return apierr.New(http.StatusInternalServerError, "database error")
	}
	return nil
}

func ValidatePWD(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("invalid values")
	}
	pwd, err := passwordStore.Password(username)
	if err != nil {
		return "", errors.New("db error, " + err.Error()) // not found?
	} else if !IsCorrect(pwd.HashedPassword, password) {
		return "", errors.New("incorrect values")
	}
	return username, nil
}

func IsCorrect(hashedPassword []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil
}

func bcryptHash(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCost)
	if err != nil {
		return nil, errors.New("password generator error")
	}
	return hash, nil
}
