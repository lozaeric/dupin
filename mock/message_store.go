package mock

import (
	"errors"

	"github.com/lozaeric/dupin/domain"
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

func (s *MessageStore) Search(kv ...[2]string) ([]*domain.Message, error) {
	return nil, nil
}

func NewMessageStore() (*MessageStore, error) {
	return &MessageStore{
		storage: make(map[string]*domain.Message),
	}, nil
}
