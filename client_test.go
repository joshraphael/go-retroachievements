package retroachievements_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/joshraphael/go-retroachievements"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	actual := retroachievements.New(
		retroachievements.RetroAchievementHost,
		"newUserAgent",
		"some_secret",
		retroachievements.HttpClient(&http.Client{
			Transport: http.DefaultTransport,
			Timeout:   5 * time.Minute,
		}),
	)

	expected := &retroachievements.Client{
		Host:   retroachievements.RetroAchievementHost,
		Secret: "some_secret",
		HttpClient: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   5 * time.Minute,
		},
		UserAgent: "newUserAgent",
	}

	require.Equal(t, expected, actual)
}

func TestNewClient(t *testing.T) {
	actual := retroachievements.NewClient("some_secret")

	expected := &retroachievements.Client{
		Host:   retroachievements.RetroAchievementHost,
		Secret: "some_secret",
		HttpClient: &http.Client{
			Transport: http.DefaultTransport,
		},
		UserAgent: actual.UserAgent,
	}

	require.Equal(t, expected, actual)
}
