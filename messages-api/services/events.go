package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const eventChannel = "users_events"
const redisURL = "redis:6379"

var subscriber *eventSubscriber

type eventSubscriber struct {
	client *redis.Client
}

func (s *eventSubscriber) EventChannel() <-chan *redis.Message {
	sub := s.client.Subscribe(eventChannel)
	return sub.Channel()
}

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:        redisURL,
		DialTimeout: 50 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}
	subscriber = &eventSubscriber{
		client: client,
	}

	go func() {
		for m := range subscriber.EventChannel() {
			userID := m.Payload
			removefromCache(userID)
			fmt.Println("USER " + userID + " has changed and will be removed from cache")
		}
	}()
}
