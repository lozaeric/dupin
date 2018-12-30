package users

import (
	"encoding/json"
	"net/http"

	"github.com/lozaeric/dupin/toolkit/apierr"
	"github.com/lozaeric/dupin/toolkit/utils"
	"github.com/lozaeric/dupin/toolkit/validation"
	"github.com/lozaeric/dupin/users-api/domain"
)

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
		return nil, apierr.New(http.StatusInternalServerError, "database error") // 500
	}
	/*
		if err := services.CreatePassword(user.ID, user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, "auth-api error")
		} else {
			c.JSON(http.StatusOK, user)
		}
	*/
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
