package oauth

import (
	"log"

	"github.com/lozaeric/dupin/auth-api/passwords"
	"github.com/lozaeric/dupin/auth-api/redis"
	oredis "gopkg.in/go-oauth2/redis.v3"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
)

var (
	manager *manage.Manager
	srv     *server.Server
)

func init() {
	manager = manage.NewDefaultManager()
	manager.MapTokenStorage(oredis.NewRedisStore(&oredis.Options{
		Addr: redis.RedisURL,
		DB:   redis.TokensDatabase,
	}))
	clientStore, err := redis.NewClientStore()
	if err != nil {
		panic(err)
	}
	manager.MapClientStorage(clientStore)
	// api gateway cli
	clientStore.Save(&models.Client{
		ID:     "123123123", // todo: must be safe, nanoid?
		Secret: "111222333", // todo: must be safe, nanoid?
	})

	srv = server.NewDefaultServer(manager)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(passwords.ValidatePWD)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
}
