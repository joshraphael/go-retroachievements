package client_test

import (
	"testing"

	"github.com/joshraphael/go-retroachievements/models"
)

func TestGetGame(tt *testing.T) {
	// now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	// require.NoError(tt, err)
	tests := []struct {
		name          string
		id            int
		modifyURL     func(url string) string
		responseCode  int
		responseGame  models.Game
		responseError models.ErrorResponse
		response      func(gameBytes []byte, errorBytes []byte) []byte
		assert        func(t *testing.T, game *models.Game, err error)
	}{}
	_ = tests
}
