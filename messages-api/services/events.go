package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const eventChannel = "users_events"

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
		DialTimeout: 200 * time.Millisecond,
		ReadTimeout: 200 * time.Millisecond,
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
			err := usersCache.Remove(userID)

			if err != nil {
				fmt.Println("[USER-CACHE] " + userID + " has changed and will be removed from cache. Err: " + err.Error())
			} else {
				fmt.Println("[USER-CACHE] " + userID + " has changed and will be removed from cache.")
			}
		}
	}()
}
