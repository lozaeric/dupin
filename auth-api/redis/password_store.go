package redis

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/lozaeric/dupin/auth-api/domain"
)

type PasswordStore struct {
	client *redis.Client
}

func (s *PasswordStore) SecureInfo(UserID string) (*domain.SecureInfo, error) {
	key := passwordsPrefix + UserID
	info := new(domain.SecureInfo)
	b, err := s.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return info, json.Unmarshal(b, info)
}

func (s *PasswordStore) Create(info *domain.SecureInfo) error {
	key := passwordsPrefix + info.UserID
	b, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return s.client.Set(key, b, 0).Err()
}

func NewPasswordStore() (*PasswordStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		// client config
	})

	return &PasswordStore{
		client: client,
	}, nil
}
