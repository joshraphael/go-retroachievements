package http_test

import (
	"testing"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	r := raHttp.NewRequest("http://localhost")

	r2 := raHttp.NewRequest("http://localhost")

	require.Equal(t, r, r2)
}
