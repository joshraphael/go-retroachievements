package http_test

import (
	"net/http"
	"strconv"
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
		raHttp.UserAgent(),
		raHttp.APIToken("secret_token"),
		raHttp.BearerToken("secret_bearer"),
		raHttp.U("myUsername"),
		raHttp.M(10),
		raHttp.F(int(now.Unix())),
		raHttp.T(int(later.Unix())),
		raHttp.D(now),
		raHttp.I([]string{strconv.Itoa(2837), strconv.Itoa(4535)}),
		raHttp.K([]string{"test1", "test2"}),
		raHttp.G(345),
		raHttp.A(1),
		raHttp.C(20),
		raHttp.O(34),
		raHttp.H(1),
	)

	expected := &raHttp.Request{
		Method: "POST",
		Path:   "/api/v1/some_resource",
		Host:   "http://localhost",
		Headers: map[string]string{
			"Authorization": "Bearer secret_bearer",
			"User-Agent":    "go-retroachievements/v0.0.0",
		},
		Params: map[string]string{
			"y": "secret_token",
			"m": "10",
			"u": "myUsername",
			"f": "1709400423",
			"t": "1709401023",
			"d": "2024-03-02",
			"i": "2837,4535",
			"k": "test1,test2",
			"g": "345",
			"a": "1",
			"c": "20",
			"o": "34",
			"h": "1",
		},
	}

	require.Equal(t, expected, actual)
}
