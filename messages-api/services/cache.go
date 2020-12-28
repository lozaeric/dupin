package services

import (
	"errors"

	lru "github.com/hashicorp/golang-lru"
	"github.com/lozaeric/dupin/messages-api/domain"
)

const maxCacheSize = 10000

var localCache, _ = lru.New(maxCacheSize)

func getFromCache(ID string) (*domain.User, error) {
	value, found := localCache.Get(ID)
	if !found {
		return nil, errors.New("user doesnt exist in cache")
	}
	return value.(*domain.User), nil
}

func saveToCache(user *domain.User) {
	localCache.Add(user.ID, user)
}

func removefromCache(ID string) bool {
	return localCache.Remove(ID)
}
