package services

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/lozaeric/dupin/messages-api/domain"
)

const usersPrefix = "messages_users_"
const redisURL = "redis:6379"

var usersCache = newCache()

type cache struct {
	client *redis.Client
}

func (c *cache) Get(ID string) (*domain.User, error) {
	key := usersPrefix + ID
	u := new(domain.User)
	b, err := c.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return u, json.Unmarshal(b, u)
}

func (c *cache) Save(u *domain.User) error {
	key := usersPrefix + u.ID
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return c.client.Set(key, b, 0).Err()
}

func (c *cache) Remove(ID string) error {
	key := usersPrefix + ID
	return c.client.Del(key).Err()
}

func newCache() *cache {
	client := redis.NewClient(&redis.Options{
		Addr:        redisURL,
		DialTimeout: 200 * time.Millisecond,
		ReadTimeout: 200 * time.Millisecond,
	})
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	return &cache{
		client: client,
	}
}
