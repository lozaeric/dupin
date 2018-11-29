package domain

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type UserStore interface {
	User(string) (*User, error)
	CreateUser(*User) error
	DeleteUser(string) error
	Validate(*User) error
}
