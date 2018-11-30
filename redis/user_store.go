package redis

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/lozaeric/dupin/domain"
	"github.com/rs/xid"
)

type UserStore struct {
	client *redis.Client
}

func (s *UserStore) User(ID string) (*domain.User, error) {
	key := usersPrefix + ID
	user := new(domain.User)
	b, err := s.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return user, json.Unmarshal(b, user)
}

func (s *UserStore) CreateUser(user *domain.User) error {
	user.ID = xid.New().String()
	key := usersPrefix + user.ID
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return s.client.Set(key, b, 0).Err()
}

func (s *UserStore) DeleteUser(ID string) error {
	key := usersPrefix + ID
	return s.client.Del(key).Err()
}

func NewUserStore() (*UserStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		// client config
	})

	return &UserStore{
		client: client,
	}, nil
}
