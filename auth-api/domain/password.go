package domain

type Password struct {
	Username       string `json:"username" validate:"required,alphanum"`
	HashedPassword []byte `json:"hashed_password"`
}

type PasswordStore interface {
	Password(string) (*Password, error)
	Save(*Password) error
}
