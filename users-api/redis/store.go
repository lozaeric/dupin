package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/lozaeric/dupin/users-api/domain"
)

type UserStore struct {
	client         *redis.Client
	eventPublisher *eventPublisher
}

func (s *UserStore) User(ID string) (*domain.User, error) {
	key := usersPrefix + ID
	u := new(domain.User)
	b, err := s.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return u, json.Unmarshal(b, u)
}

func (s *UserStore) Save(u *domain.User) error {
	key := usersPrefix + u.ID
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return s.eventPublisher.SendEventIfNeeded(
		u.ID,
		s.client.Set(key, b, 0).Err,
	)
}

func (s *UserStore) Update(u *domain.User) error {
	key := usersPrefix + u.ID
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return s.eventPublisher.SendEventIfNeeded(
		u.ID,
		s.client.Set(key, b, 0).Err,
	)
}

func NewUserStore() *UserStore {
	client := redis.NewClient(&redis.Options{
		Addr:        redisURL,
		DialTimeout: 50 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	return &UserStore{
		client:         client,
		eventPublisher: newEventPublisher(),
	}
}
