package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/utils"
	"github.com/rs/xid"
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

	message.ID = xid.New().String()
	message.DateCreated = utils.Now()
	return conn.DB(database).C(messagesCollection).Insert(message)
}

func (s *MessageStore) Delete(ID string) error {
	conn := s.session.Copy()
	defer conn.Close()

	return conn.DB(database).C(messagesCollection).Remove(bson.M{"id": ID})
}

func (s *MessageStore) Search(kv ...[2]string) ([]*domain.Message, error) {
	return nil, nil
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
