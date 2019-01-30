package redis

import (
	"encoding/json"

	"github.com/go-redis/redis"
	oauth2 "gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/models"
)

type ClientStore struct {
	client *redis.Client
}

func (s *ClientStore) GetByID(ID string) (oauth2.ClientInfo, error) {
	client := new(models.Client)
	b, err := s.client.Get(ID).Bytes()
	if err != nil {
		return nil, err
	}
	return client, json.Unmarshal(b, &client)
}

func (s *ClientStore) Save(client *models.Client) error {
	b, err := json.Marshal(client)
	if err != nil {
		return err
	}
	return s.client.Set(client.ID, b, 0).Err()
}

func NewClientStore() (*ClientStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: RedisURL,
		DB:   clientsDatabase,
		// client config
	})

	return &ClientStore{
		client: client,
	}, client.Ping().Err()
}