package redis

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/lozaeric/dupin/auth-api/domain"
)

type PasswordStore struct {
	client *redis.Client
}

func (s *PasswordStore) Password(Username string) (*domain.Password, error) {
	info := new(domain.Password)
	b, err := s.client.Get(Username).Bytes()
	if err != nil {
		return nil, err
	}
	return info, json.Unmarshal(b, info)
}

func (s *PasswordStore) Save(pwd *domain.Password) error {
	b, err := json.Marshal(pwd)
	if err != nil {
		return err
	}
	return s.client.Set(pwd.Username, b, 0).Err()
}

func NewPasswordStore() (*PasswordStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: RedisURL,
		DB:   passwordsDatabase,
		// client config
	})

	return &PasswordStore{
		client: client,
	}, client.Ping().Err()
}
