package nanoid

import (
	"github.com/lozaeric/dupin/auth-api/domain"
	nano "github.com/matoous/go-nanoid"
)

type TokenGenerator struct{}

func (*TokenGenerator) Generate(applicationID, userID string) (*domain.Token, error) {
	tokenID, err := nano.Generate(domain.TokenAlphabet, domain.TokenLength)
	if err != nil {
		return nil, err
	}

	token := new(domain.Token)
	token.ID = tokenID
	token.ApplicationID = applicationID
	token.UserID = userID
	token.Scopes = []string{"read", "write"} //TODO
	token.CalculateDates()
	return token, nil
}
