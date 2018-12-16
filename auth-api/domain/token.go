package domain

import (
	"time"

	nanoid "github.com/matoous/go-nanoid"
)

type Token struct {
	ID             string   `json:"id"`
	ApplicationID  string   `json:"application_id" validate:"required,len=20,alphanum"`
	UserID         string   `json:"user_id" validate:"required,len=20,alphanum"`
	Scopes         []string `json:"scopes"`
	DateCreated    string   `json:"date_created"`
	ExpirationDate string   `json:"expiration_date"`
}

func NewToken(applicationID, userID string) (*Token, error) {
	tokenID, err := nanoid.Generate(TokenAlphabet, TokenLength)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	token := &Token{
		ID:             tokenID,
		ApplicationID:  applicationID,
		UserID:         userID,
		Scopes:         []string{"read", "write"}, //TODO
		DateCreated:    now.Format(time.RFC3339),
		ExpirationDate: now.Add(TokenTTL).Format(time.RFC3339),
	}
	return token, nil
}

type TokenStore interface {
	Token(string) (*Token, error)
	Create(*Token) error
}
