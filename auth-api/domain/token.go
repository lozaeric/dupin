package domain

import "time"

type Token struct {
	ID             string   `json:"id"`
	ApplicationID  string   `json:"application_id" validate:"required,len=20,alphanum"`
	UserID         string   `json:"user_id" validate:"required,len=20,alphanum"`
	Scopes         []string `json:"scopes"`
	DateCreated    string   `json:"date_created"`
	ExpirationDate string   `json:"expiration_date"`
}

func (t *Token) CalculateDates() {
	now := time.Now()
	t.DateCreated = now.Format(time.RFC3339)
	t.ExpirationDate = now.Add(TokenTTL).Format(time.RFC3339)
}

type TokenStore interface {
	Token(string) (*Token, error)
	Create(*Token) error
}

type TokenGenerator interface {
	Generate(string, string) (*Token, error)
}
