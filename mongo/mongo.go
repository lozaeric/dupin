package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lozaeric/dupin/domain"
)

type Store struct {
	session *mgo.Session
}

func (s *Store) User(ID int) (*domain.User, error) {
	conn := s.session.Copy()
	defer conn.Close()

	user := new(domain.User)
	err := conn.DB(usersCollection).C(usersCollection).Find(bson.M{"id": ID}).One(user)
	return user, err
}

func (s *Store) CreateUser(user *domain.User) error {
	conn := s.session.Copy()
	defer conn.Close()

	return conn.DB(usersCollection).C(usersCollection).Insert(user)
}

func (s *Store) DeleteUser(ID int) error {
	conn := s.session.Copy()
	defer conn.Close()

	return conn.DB(usersCollection).C(usersCollection).Remove(bson.M{"id": ID})
}

func NewStore() (*Store, error) {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	// session config
	return &Store{
		session: session,
	}, nil
}
