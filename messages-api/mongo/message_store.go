package mongo

import (
	"errors"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lozaeric/dupin/messages-api/domain"
	"github.com/lozaeric/dupin/toolkit/utils"
)

type MessageStore struct {
	session *mgo.Session
}

func (s *MessageStore) Message(ID string) (*domain.Message, error) {
	conn := s.session.Copy()
	defer conn.Close()

	message := new(domain.Message)
	err := conn.DB(database).C(messagesCollection).Find(bson.M{"id": ID}).One(message)
	return message, err
}

func (s *MessageStore) Create(message *domain.Message) error {
	conn := s.session.Copy()
	defer conn.Close()

	message.ID = utils.GenerateID()
	message.DateCreated = utils.Now()
	message.DateUpdated = message.DateCreated
	return conn.DB(database).C(messagesCollection).Insert(message)
}

func (s *MessageStore) Search(userID, field, value string) ([]*domain.Message, error) {
	if userID == "" {
		return nil, errors.New("invalid user id")
	}
	conn := s.session.Copy()
	defer conn.Close()

	messages := []*domain.Message{}
	query := bson.M{ // it works ?
		"$or": []bson.M{
			{"sender_id": userID},
			{"receiver_id": userID},
		},
	}
	if field != "" && value != "" {
		query[field] = value
	}
	err := conn.DB(database).C(messagesCollection).Find(query).All(&messages)
	return messages, err
}

func (s *MessageStore) Update(message *domain.Message) error {
	conn := s.session.Copy()
	defer conn.Close()

	return conn.DB(database).C(messagesCollection).Update(bson.M{"id": message.ID}, message)
}

func NewMessageStore() (*MessageStore, error) {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	// session config
	return &MessageStore{
		session: session,
	}, nil
}
