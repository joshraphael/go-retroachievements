package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/joshraphael/go-retroachievements/client"
	"github.com/joshraphael/go-retroachievements/models"
	"github.com/stretchr/testify/require"
)

func TestGetUserProfile(tt *testing.T) {
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		username        string
		responseCode    int
		responseProfile models.Profile
		responseError   models.ErrorResponse
		response        func(profileBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, profile *models.Profile, err error)
	}{
		{
			name:         "error response",
			username:     "Test",
			responseCode: http.StatusUnauthorized,
			responseError: models.ErrorResponse{
				Message: "test",
				Errors: []models.ErrorDetail{
					{
						Status: http.StatusUnauthorized,
						Code:   "unauthorized",
						Title:  "Not Authorized",
					},
				},
			},
			response: func(profileBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, profile *models.Profile, err error) {
				require.Nil(t, profile)
				require.EqualError(t, err, "parsing response object: error responses: [401] Not Authorized")
			},
		},
		{
			name:         "success",
			username:     "Test",
			responseCode: http.StatusOK,
			responseProfile: models.Profile{
				User:    "xXxSnip3rxXx",
				UserPic: "/some/resource.png",
				MemberSince: models.DateTime{
					Time: now,
				},
				RichPresenceMsg:     "Playing Super Mario 64",
				LastGameID:          5436,
				ContribCount:        10,
				ContribYield:        1,
				TotalPoints:         1000,
				TotalSoftcorePoints: 234,
				TotalTruePoints:     512,
				Permissions:         1,
				Untracked:           0,
				ID:                  445526,
				UserWallActive:      true,
				Motto:               "Playing games",
			},
			response: func(profileBytes []byte, errorBytes []byte) []byte {
				return profileBytes
			},
			assert: func(t *testing.T, profile *models.Profile, err error) {
				require.NoError(t, err)
				require.NotNil(t, profile)
				require.Equal(t, "xXxSnip3rxXx", profile.User)
				require.Equal(t, "/some/resource.png", profile.UserPic)
				require.Equal(t, models.DateTime{
					Time: now,
				}, profile.MemberSince)
				require.Equal(t, "Playing Super Mario 64", profile.RichPresenceMsg)
				require.Equal(t, 5436, profile.LastGameID)
				require.Equal(t, 10, profile.ContribCount)
				require.Equal(t, 1, profile.ContribYield)
				require.Equal(t, 1000, profile.TotalPoints)
				require.Equal(t, 234, profile.TotalSoftcorePoints)
				require.Equal(t, 512, profile.TotalTruePoints)
				require.Equal(t, 1, profile.Permissions)
				require.Equal(t, 0, profile.Untracked)
				require.Equal(t, 445526, profile.ID)
				require.True(t, profile.UserWallActive)
				require.Equal(t, "Playing games", profile.Motto)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserProfile.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				profileBytes, err := json.Marshal(test.responseProfile)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				w.Write(test.response(profileBytes, errBytes))
			}))
			defer server.Close()

			client := client.New(server.URL, "some_secret")
			profile, err := client.GetUserProfile(test.username)
			test.assert(t, profile, err)
		})
	}
}
