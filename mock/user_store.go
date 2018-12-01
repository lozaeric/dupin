package mock

import (
	"errors"

	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/utils"
)

type UserStore struct {
	storage map[string]*domain.User
}

func (s *UserStore) User(ID string) (*domain.User, error) {
	user, found := s.storage[ID]
	if !found {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (s *UserStore) Create(user *domain.User) error {
	user.ID = GenerateValidID()
	user.DateCreated = utils.Now()
	s.storage[user.ID] = user
	return nil
}

func (s *UserStore) Delete(ID string) error {
	_, found := s.storage[ID]
	if !found {
		return errors.New("not found")
	}
	delete(s.storage, ID)
	return nil
}

func NewUserStore() (*UserStore, error) {
	return &UserStore{
		storage: make(map[string]*domain.User),
	}, nil
}
