package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type UserStore interface {
	User(int) (*User, error)
	CreateUser(*User) error
	DeleteUser(int) error
}
