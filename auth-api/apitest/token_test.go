package apitest

/*
var (
	token = &domain.Token{
		ApplicationID: mock.GenerateValidID(),
		UserID:        mock.GenerateValidID(),
		Scopes:        []string{"read", "write"},
	}
)

func TestCreateToken(t *testing.T) {
	assert := assert.New(t)
	cases := []*struct {
		expectedStatus int
		dto            *domain.Token
	}{
		{
			http.StatusBadRequest,
			&domain.Token{
				ApplicationID: "invalid",
				UserID:        mock.GenerateValidID(),
			},
		},
		{
			http.StatusBadRequest,
			&domain.Token{
				ApplicationID: mock.GenerateValidID(),
				UserID:        "invalid",
			},
		},
		{
			http.StatusOK,
			token,
		},
	}

	for _, c := range cases {
		r, err := cli.R().SetBody(c.dto).Post("/tokens")
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			t := new(domain.Token)
			err := json.Unmarshal(r.Body(), t)
			assert.Nil(err)
			assert.NotEmpty(t.ID)
			assert.NotEmpty(t.DateCreated)
			assert.NotEmpty(t.ExpirationDate)
			c.dto.ID = t.ID
			c.dto.DateCreated = t.DateCreated
			c.dto.ExpirationDate = t.ExpirationDate
			assert.Equal(c.dto, t)
		}
	}
}

func TestToken(t *testing.T) {
	assert := assert.New(t)

	cases := []*struct {
		expectedStatus int
		ID             string
		expectedToken  *domain.Token
	}{
		{
			http.StatusNotFound,
			"0000000000000000000000000",
			nil,
		},
		{
			http.StatusBadRequest,
			"invalid",
			nil,
		},
		{
			http.StatusOK,
			token.ID,
			token,
		},
	}

	for _, c := range cases {
		r, err := cli.R().Get("/tokens/" + c.ID)
		assert.Nil(err)
		assert.Equal(c.expectedStatus, r.StatusCode())
		if c.expectedStatus == http.StatusOK {
			t := new(domain.Token)
			err := json.Unmarshal(r.Body(), t)
			assert.Nil(err)
			assert.Equal(c.expectedToken, t)
		}
	}
}

func TestExpiratedToken(t *testing.T) {
	assert := assert.New(t)

	time.Sleep(100 * time.Millisecond)
	r, err := cli.R().Get("/tokens/" + token.ID)
	assert.Nil(err)
	assert.Equal(http.StatusNotFound, r.StatusCode())
}
*/
