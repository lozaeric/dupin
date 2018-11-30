package mock

import (
	"errors"
	"strconv"

	"github.com/lozaeric/dupin/domain"
)

type MessageStore struct {
	storage map[string]*domain.Message
	counter int
}

func (s *MessageStore) Message(ID string) (*domain.Message, error) {
	message, found := s.storage[ID]
	if !found {
		return nil, errors.New("not found")
	}
	return message, nil
}

func (s *MessageStore) CreateMessage(message *domain.Message) error {
	message.ID = strconv.Itoa(s.counter)
	s.storage[message.ID] = message
	s.counter++
	return nil
}

func (s *MessageStore) DeleteMessage(ID string) error {
	_, found := s.storage[ID]
	if !found {
		return errors.New("not found")
	}
	delete(s.storage, ID)
	return nil
}

func (s *MessageStore) SearchMessages(kv ...[2]string) ([]*domain.Message, error) {
	return nil, nil
}

func NewMessageStore() (*MessageStore, error) {
	return &MessageStore{
		storage: make(map[string]*domain.Message),
		counter: 1,
	}, nil
}
