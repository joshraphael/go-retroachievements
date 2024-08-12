package http_test

import (
	"net/http"
	"testing"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	actual := raHttp.NewRequest(
		"http://localhost",
		raHttp.Path("/api/v1/some_resource"),
		raHttp.Method(http.MethodPost),
		raHttp.APIToken("secret_token"),
		raHttp.BearerToken("secret_bearer"),
		raHttp.Username("myUsername"),
		raHttp.LookbackMinutes(10),
	)

	expected := &raHttp.Request{
		Method: "POST",
		Path:   "/api/v1/some_resource",
		Host:   "http://localhost",
		Headers: map[string]string{
			"Authorization": "Bearer secret_bearer",
		},
		Params: map[string]string{
			"y": "secret_token",
			"m": "10",
			"u": "myUsername",
		},
	}

	require.Equal(t, expected, actual)
}
