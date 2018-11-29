package mongo

import (
	"errors"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lozaeric/dupin/domain"
	"github.com/rs/xid"
)

type UserStore struct {
	session *mgo.Session
}

func (s *UserStore) User(ID string) (*domain.User, error) {
	conn := s.session.Copy()
	defer conn.Close()

	user := new(domain.User)
	err := conn.DB(usersCollection).C(usersCollection).Find(bson.M{"id": ID}).One(user)
	return user, err
}

func (s *UserStore) CreateUser(user *domain.User) error {
	conn := s.session.Copy()
	defer conn.Close()

	user.ID = xid.New().String()
	return conn.DB(usersCollection).C(usersCollection).Insert(user)
}

func (s *UserStore) DeleteUser(ID string) error {
	conn := s.session.Copy()
	defer conn.Close()

	return conn.DB(usersCollection).C(usersCollection).Remove(bson.M{"id": ID})
}

func (s *UserStore) Validate(user *domain.User) error {
	if user.Name == "" {
		return errors.New("name is empty")
	}
	if user.LastName == "" {
		return errors.New("last name is empty")
	}
	if user.Email == "" {
		return errors.New("email is empty")
	}
	return nil
}

func NewUserStore() (*UserStore, error) {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	// session config
	return &UserStore{
		session: session,
	}, nil
}
