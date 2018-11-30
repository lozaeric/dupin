package mock

import (
	"errors"
	"strconv"

	"github.com/lozaeric/dupin/domain"
)

type UserStore struct {
	storage map[string]*domain.User
	counter int
}

func (s *UserStore) User(ID string) (*domain.User, error) {
	user, found := s.storage[ID]
	if !found {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (s *UserStore) Create(user *domain.User) error {
	user.ID = strconv.Itoa(s.counter)
	s.storage[user.ID] = user
	s.counter++
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
		counter: 1,
	}, nil
}
