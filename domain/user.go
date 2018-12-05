package domain

import (
	"github.com/lozaeric/dupin/domain/validation"
	"github.com/lozaeric/dupin/utils"
)

var userUpdatable = map[string]bool{
	"name":      true,
	"last_name": true,
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

type UserStore interface {
	User(string) (*User, error)
	Create(*User) error
	Update(*User) error
}

func (u *User) Update(values map[string]interface{}) (err error) {
	err = validation.Update(u, values, userUpdatable)
	if err != nil {
		return
	}
	u.DateUpdated = utils.Now()
	return
}
