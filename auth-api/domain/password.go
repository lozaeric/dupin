package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type SecureInfo struct {
	UserID         string `json:"user_id" validate:"required,len=20,alphanum"`
	HashedPassword []byte `json:"hashed_password"`
}

func (u *SecureInfo) IsCorrect(password string) bool {
	return bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(password)) == nil
}

func NewSecureInfo(userID, password string) (*SecureInfo, bool) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), BcryptCost)
	if err != nil {
		return nil, false
	}
	return &SecureInfo{userID, hash}, true
}

type SecureInfoStore interface {
	SecureInfo(string) (*SecureInfo, error)
	Create(*SecureInfo) error
}
