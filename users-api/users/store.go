package users

import (
	"github.com/lozaeric/dupin/users-api/domain"
	"github.com/lozaeric/dupin/users-api/redis"
)

var userStore domain.UserStore

func init() {
	var err error
	userStore, err = redis.NewUserStore()
	if err != nil {
		panic(err)
	}
}
