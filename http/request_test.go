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
		raHttp.UserAgent("go-retroachievements/v0.0.0"),
		raHttp.BearerToken("secret_bearer"),
		raHttp.A(1),
		raHttp.C(20),
		raHttp.D(now.UTC().Format(time.DateOnly)),
		raHttp.F(int(now.Unix())),
		raHttp.G(345),
		raHttp.H(1),
		raHttp.I([]string{strconv.Itoa(2837), strconv.Itoa(4535)}),
		raHttp.K([]string{"test1", "test2"}),
		raHttp.M(10),
		raHttp.O(34),
		raHttp.R("codenotes2"),
		raHttp.T(strconv.Itoa(int(later.Unix()))),
		raHttp.U("myUsername"),
		raHttp.Y("secret_token"),
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
			"a": "1",
			"c": "20",
			"d": "2024-03-02",
			"f": "1709400423",
			"g": "345",
			"h": "1",
			"i": "2837,4535",
			"k": "test1,test2",
			"m": "10",
			"o": "34",
			"r": "codenotes2",
			"t": "1709401023",
			"u": "myUsername",
			"y": "secret_token",
		},
	}

	require.Equal(t, expected, actual)
}
