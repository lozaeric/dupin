package users

import (
	"encoding/json"
	"net/http"

	"github.com/lozaeric/dupin/toolkit/apierr"
	"github.com/lozaeric/dupin/toolkit/utils"
	"github.com/lozaeric/dupin/toolkit/validation"
	"github.com/lozaeric/dupin/users-api/domain"
	"github.com/lozaeric/dupin/users-api/redis"
	"github.com/lozaeric/dupin/users-api/services"
)

var userStore = redis.NewUserStore()

func Create(data []byte) (*domain.User, *apierr.ApiError) {
	user := new(domain.User)
	if err := json.Unmarshal(data, user); err != nil {
		return nil, apierr.New(http.StatusBadRequest, "invalid values")
	}
	user.ID = utils.GenerateID()
	user.DateCreated = utils.Now()
	user.DateUpdated = user.DateCreated
	if err := validation.Validate(user); err != nil {
		return nil, apierr.New(http.StatusBadRequest, err.Error())
	}
	if err := userStore.Save(user); err != nil {
		return nil, apierr.New(http.StatusInternalServerError, "database error")
	}
	if err := savePassword(user.ID, data); err != nil {
		return nil, err
	}
	return user, nil
}

func Get(ID string) (*domain.User, *apierr.ApiError) {
	if !validation.IsValidID(ID) {
		return nil, apierr.New(http.StatusBadRequest, "invalid ID")
	}
	user, err := userStore.User(ID)
	if err != nil {
		return nil, apierr.New(http.StatusNotFound, err.Error())
	}
	return user, nil
}

func Update(ID string, data []byte) (*domain.User, *apierr.ApiError) {
	if !validation.IsValidID(ID) {
		return nil, apierr.New(http.StatusBadRequest, "invalid ID")
	}
	user, err := userStore.User(ID)
	if err != nil {
		return nil, apierr.New(http.StatusNotFound, "user not found")
	}

	values := make(map[string]interface{})
	if err := json.Unmarshal(data, &values); err != nil {
		return nil, apierr.New(http.StatusBadRequest, "invalid body")
	}
	if err := user.Update(values); err != nil {
		return nil, apierr.New(http.StatusBadRequest, "invalid values")
	}
	if err := userStore.Update(user); err != nil {
		return nil, apierr.New(http.StatusInternalServerError, "database error")
	}
	return user, nil
}

func Delete(ID string) (*domain.User, *apierr.ApiError) {
	if !validation.IsValidID(ID) {
		return nil, apierr.New(http.StatusBadRequest, "invalid ID")
	}
	user, err := userStore.User(ID)
	if err != nil {
		return nil, apierr.New(http.StatusNotFound, "user not found")
	}
	if user.Deleted {
		return user, apierr.New(http.StatusBadRequest, "user is already deleted")
	}
	user.Deleted = true
	if err := userStore.Update(user); err != nil {
		return nil, apierr.New(http.StatusInternalServerError, "database error")
	}
	return user, nil
}

func savePassword(ID string, data []byte) *apierr.ApiError {
	info := make(map[string]interface{})
	if err := json.Unmarshal(data, &info); err != nil {
		return apierr.New(http.StatusBadRequest, "invalid password")
	}
	password, ok := info["password"].(string)
	if !ok || password == "" {
		return apierr.New(http.StatusBadRequest, "invalid password")
	}
	if err := services.CreatePassword(ID, password); err != nil {
		return apierr.New(http.StatusInternalServerError, "error creating password") // 500
	}
	return nil
}
