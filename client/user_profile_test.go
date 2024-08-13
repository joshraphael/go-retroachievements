package client_test

import (
	"testing"

	"github.com/joshraphael/go-retroachievements/models"
)

func TestGetUserProfile(t *testing.T) {
	test := []struct {
		name     string
		response models.Profile
	}{
		{
			name:     "test",
			response: models.Profile{},
		},
	}
	_ = test
}
