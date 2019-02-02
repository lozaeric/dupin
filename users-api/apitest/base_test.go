package apitest

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-resty/resty"
	"github.com/jarcoal/httpmock"
	"github.com/lozaeric/dupin/users-api/app"
	"github.com/stretchr/testify/assert"
)

var cli = resty.New().
	SetTimeout(100 * time.Millisecond).
	SetHostURL("http://localhost:8080")

func TestMain(m *testing.M) {
	go app.Run()
	time.Sleep(3 * time.Second)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setupMocks()

	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	assert := assert.New(t)
	r, err := cli.R().Get("/ping")
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())
}

func setupMocks() {
	httpmock.RegisterResponder("POST", "http://auth:8080/passwords",
		httpmock.NewStringResponder(http.StatusOK, "ok"))
	httpmock.RegisterNoResponder(httpmock.InitialTransport.RoundTrip)
}
