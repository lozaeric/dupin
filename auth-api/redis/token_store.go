package redis

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/lozaeric/dupin/auth-api/domain"
)

type TokenStore struct {
	client *redis.Client
}

func (s *TokenStore) Token(ID string) (*domain.Token, error) {
	key := tokensPrefix + ID
	token := new(domain.Token)
	b, err := s.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return token, json.Unmarshal(b, token)
}

func (s *TokenStore) Create(token *domain.Token) error {
	key := tokensPrefix + token.ID
	b, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return s.client.Set(key, b, domain.TokenTTL).Err()
}

func NewTokenStore() (*TokenStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		// client config
	})

	return &TokenStore{
		client: client,
	}, nil
}
