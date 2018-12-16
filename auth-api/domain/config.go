package domain

import (
	"os"
	"time"

	"github.com/go-playground/validator"
)

var validate = validator.New()
var TokenTTL = time.Hour

const (
	TokenLength     = 25
	TokenAlphabet   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	TokenValidation = "required,len=25,alphanum" //TODO: utils and mock
	BcryptCost      = 7
)

func IsValidID(ID string) bool {
	return validate.Var(ID, TokenValidation) == nil
}

func init() {
	if os.Getenv("ENV") != "production" {
		TokenTTL = 80 * time.Millisecond
	}
}
