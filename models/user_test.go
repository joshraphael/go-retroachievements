package models_test

import (
	"encoding/json"
	"testing"

	"github.com/joshraphael/go-retroachievements/models"
	"github.com/stretchr/testify/require"
)

func TestGetUserSummaryRecentAchievementsUnmarshalJSON(tt *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		assert func(t *testing.T, s models.GetUserSummaryRecentAchievements, err error)
	}{
		{
			name:  "array substituted for object",
			input: []byte(`["test"]`),
			assert: func(t *testing.T, s models.GetUserSummaryRecentAchievements, err error) {
				require.NoError(t, err)
				require.Equal(t, models.GetUserSummaryRecentAchievements{}, s)
			},
		},
		{
			name:  "error parsing",
			input: []byte(`"?>?>>L:"`),
			assert: func(t *testing.T, s models.GetUserSummaryRecentAchievements, err error) {
				require.EqualError(t, err, "json: cannot unmarshal string into Go value of type map[string]map[string]models.GetUserSummaryRecentAchievement")
				require.Equal(t, models.GetUserSummaryRecentAchievements{}, s)
			},
		},
		{
			name:  "success",
			input: []byte(`{"test": {"test": {"ID": 132}}}`),
			assert: func(t *testing.T, s models.GetUserSummaryRecentAchievements, err error) {
				require.NoError(t, err)
				require.Equal(t, models.GetUserSummaryRecentAchievements{
					"test": {
						"test": {
							ID: 132,
						},
					},
				}, s)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			s := models.GetUserSummaryRecentAchievements{}
			err := json.Unmarshal(test.input, &s)
			test.assert(t, s, err)
		})
	}
}

func TestGetUserSummaryLastGameUnmarshalJSON(tt *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		assert func(t *testing.T, s models.GetUserSummaryLastGame, err error)
	}{
		{
			name:  "array substituted for object",
			input: []byte(`["test"]`),
			assert: func(t *testing.T, s models.GetUserSummaryLastGame, err error) {
				require.NoError(t, err)
				require.Equal(t, models.GetUserSummaryLastGame{}, s)
			},
		},
		{
			name:  "error parsing",
			input: []byte(`"?>?>>L:"`),
			assert: func(t *testing.T, s models.GetUserSummaryLastGame, err error) {
				require.EqualError(t, err, "json: cannot unmarshal string into Go value of type models.internalGetUserSummaryLastGame")
				require.Equal(t, models.GetUserSummaryLastGame{}, s)
			},
		},
		{
			name:  "success",
			input: []byte(`{"ID": 132}`),
			assert: func(t *testing.T, s models.GetUserSummaryLastGame, err error) {
				require.NoError(t, err)
				require.Equal(t, models.GetUserSummaryLastGame{
					ID: 132,
				}, s)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			s := models.GetUserSummaryLastGame{}
			err := json.Unmarshal(test.input, &s)
			test.assert(t, s, err)
		})
	}
}
