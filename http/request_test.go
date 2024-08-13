package http_test

import (
	"net/http"
	"testing"
	"time"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(t, err)
	later := now.Add(10 * time.Minute)
	actual := raHttp.NewRequest(
		"http://localhost",
		raHttp.Path("/api/v1/some_resource"),
		raHttp.Method(http.MethodPost),
		raHttp.APIToken("secret_token"),
		raHttp.BearerToken("secret_bearer"),
		raHttp.Username("myUsername"),
		raHttp.LookbackMinutes(10),
		raHttp.FromTime(now),
		raHttp.ToTime(later),
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
			"f": "1709400423",
			"t": "1709401023",
		},
	}

	require.Equal(t, expected, actual)
}
