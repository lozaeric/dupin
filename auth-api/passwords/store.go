package passwords

import (
	"github.com/lozaeric/dupin/auth-api/domain"
	"github.com/lozaeric/dupin/auth-api/redis"
)

var passwordStore domain.PasswordStore

func init() {
	var err error
	passwordStore, err = redis.NewPasswordStore()
	if err != nil {
		panic(err)
	}
}
