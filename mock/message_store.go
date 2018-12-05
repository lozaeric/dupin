package mock

import (
	"errors"

	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/utils"
)

type MessageStore struct {
	storage map[string]*domain.Message
}

func (s *MessageStore) Message(ID string) (*domain.Message, error) {
	message, found := s.storage[ID]
	if !found {
		return nil, errors.New("not found")
	}
	return message, nil
}

func (s *MessageStore) Create(message *domain.Message) error {
	message.ID = GenerateValidID()
	message.DateCreated = utils.Now()
	s.storage[message.ID] = message
	return nil
}

func (s *MessageStore) Delete(ID string) error {
	_, found := s.storage[ID]
	if !found {
		return errors.New("not found")
	}
	delete(s.storage, ID)
	return nil
}

func (s *MessageStore) Search(field, value string) ([]*domain.Message, error) {
	return nil, nil
}

func NewMessageStore() (*MessageStore, error) {
	return &MessageStore{
		storage: make(map[string]*domain.Message),
	}, nil
}
