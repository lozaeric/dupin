package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type eventPublisher struct {
	client *redis.Client
}

func (p *eventPublisher) SendEventIfNeeded(ID string, save func() error) error {
	if err := save(); err != nil {
		return err
	}
	fmt.Println("USER " + ID + " has changed and an event will be published")
	return p.client.Publish(eventChannel, ID).Err()
}

func newEventPublisher() *eventPublisher {
	client := redis.NewClient(&redis.Options{
		Addr:        redisURL,
		DialTimeout: 200 * time.Millisecond,
		ReadTimeout: 200 * time.Millisecond,
	})
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	return &eventPublisher{
		client: client,
	}
}
