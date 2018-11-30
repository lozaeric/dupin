package apitest

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-resty/resty"
	"github.com/lozaeric/dupin/app"
	"github.com/lozaeric/dupin/domain"
	"github.com/stretchr/testify/assert"
)

const userJSON = `{"id":"1","name":"eric","last_name":"loza","email":"eric@loza"}`

var user = &domain.User{
	Name:     "eric",
	LastName: "loza",
	Email:    "eric@loza",
}

var invalidUser = &domain.User{
	Name:     "",
	LastName: "loza",
	Email:    "eric@loza",
}

var cli = resty.New().
	SetTimeout(15 * time.Millisecond).
	SetHostURL("http://localhost:8080")

func TestMain(m *testing.M) {
	go app.Run()
	time.Sleep(2 * time.Second)
	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	assert := assert.New(t)
	r, err := cli.R().Get("/ping")
	assert.Nil(err)
	assert.Equal(r.StatusCode(), http.StatusOK)
}
