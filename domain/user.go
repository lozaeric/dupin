package domain

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"date_created"`
}

type UserStore interface {
	User(string) (*User, error)
	CreateUser(*User) error
	DeleteUser(string) error
}
