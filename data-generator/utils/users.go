package utils

import (
	"sync"
	"time"

	rdata "github.com/Pallinder/go-randomdata"

	"github.com/go-resty/resty"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var usersCli = resty.New().
	SetTimeout(100 * time.Millisecond).
	SetRetryCount(1).
	AddRetryCondition(func(r *resty.Response, err error) bool {
		return err != nil || r.StatusCode() >= 500
	}).
	SetHostURL("http://localhost:8082")

func CreateUsers(n int) []*User {
	users := make(chan *User, n)
	for i := 0; i < n; i++ {
		users <- &User{
			Name:     rdata.FirstName(4),
			LastName: rdata.LastName(),
			Email:    rdata.Email(),
			Password: rdata.Alphanumeric(5),
		}
	}
	close(users)

	parallelProcess(n, func(wg *sync.WaitGroup) {
		for u := range users {
			usersCli.R().SetBody(u).Post("/users")
			wg.Done()
		}
	})

	return
}
