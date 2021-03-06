package apitest

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-resty/resty"
	"github.com/lozaeric/dupin/auth-api/app"
	"github.com/stretchr/testify/assert"
)

var cli = resty.New().
	SetTimeout(200 * time.Millisecond).
	SetHostURL("http://localhost:8080")

func TestMain(m *testing.M) {
	go app.Run()
	time.Sleep(3 * time.Second)
	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	assert := assert.New(t)
	r, err := cli.R().Get("/ping")
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())
}
