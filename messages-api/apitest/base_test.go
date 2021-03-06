package apitest

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-resty/resty"
	"github.com/jarcoal/httpmock"
	"github.com/lozaeric/dupin/messages-api/app"
	"github.com/lozaeric/dupin/messages-api/domain"
	"github.com/stretchr/testify/assert"
)

var (
	validUsers  = []string{"00000000000000000000", "11111111111111111111", "88888888888888888888", "99999999999999999999"}
	validTokens = map[string]string{
		validUsers[0]: fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, validUsers[0]),
		validUsers[1]: fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, validUsers[1]),
		validUsers[2]: fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, validUsers[2]),
		validUsers[3]: fmt.Sprintf(`{"client_id":"1","user_id":"%s","scope":"read"}`, validUsers[3]),
	}
)
var cli = resty.New().
	SetTimeout(350 * time.Millisecond).
	SetHostURL("http://localhost:8080")

func TestMain(m *testing.M) {
	go app.Run()
	time.Sleep(3 * time.Second)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	setupMocks()

	os.Exit(m.Run())
}

func setupMocks() {
	validUser := &domain.User{
		Name:     "eric",
		LastName: "loza",
		Email:    "lozaeric@gmail.com",
	}

	for _, u := range validUsers {
		validUser.ID = u
		res, _ := httpmock.NewJsonResponder(http.StatusOK, validUser)
		httpmock.RegisterResponder("GET", "http://users:8080/users/"+u, res)
	}
	httpmock.RegisterNoResponder(httpmock.InitialTransport.RoundTrip)
}

func TestPing(t *testing.T) {
	assert := assert.New(t)
	r, err := cli.R().Get("/ping")
	assert.Nil(err)
	assert.Equal(http.StatusOK, r.StatusCode())
}
