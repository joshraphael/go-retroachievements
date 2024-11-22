package retroachievements_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
	"github.com/stretchr/testify/require"
)

func TestGetAchievementUnlocks(tt *testing.T) {
	count := 10
	offset := 10
	achievementType := "progression"
	dateCreated, err := time.Parse(time.DateTime, "2012-11-02 00:03:12")
	require.NoError(tt, err)
	dateModified, err := time.Parse(time.DateTime, "2023-09-30 02:00:49")
	require.NoError(tt, err)
	dateAwarded, err := time.Parse(time.RFC3339Nano, "2024-11-22T17:25:17.000000Z")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetAchievementUnlocksParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetAchievementUnlocks
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetAchievementUnlocks, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetAchievementUnlocksParameters{
				AchievementID: 14402,
				Count:         &count,
				Offset:        &offset,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementUnlocks, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetAchievementUnlocks.php?a=14402&c=10&o=10&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetAchievementUnlocksParameters{
				AchievementID: 14402,
				Count:         &count,
				Offset:        &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
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
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementUnlocks, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error responses: [401] Not Authorized")
			},
		},
		{
			name: "success",
			params: models.GetAchievementUnlocksParameters{
				AchievementID: 14402,
				Count:         &count,
				Offset:        &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetAchievementUnlocks{
				UnlocksCount:         15906,
				UnlocksHardcoreCount: 8223,
				TotalPlayers:         38104,
				Achievement: models.GetAchievementUnlocksAchievement{
					ID:          1,
					Title:       "Ring Collector",
					Description: "Collect 100 rings",
					Points:      5,
					TrueRatio:   7,
					Author:      "Scott",
					DateCreated: models.DateTime{
						Time: dateCreated,
					},
					DateModified: models.DateTime{
						Time: dateModified,
					},
					Type: &achievementType,
				},
				Console: models.GetAchievementUnlocksConsole{
					ID:    1,
					Title: "Genesis/Mega Drive",
				},
				Game: models.GetAchievementUnlocksGame{
					ID:    1,
					Title: "Sonic the Hedgehog",
				},
				Unlocks: []models.GetAchievementUnlocksUnlock{
					{
						User:             "redjedia",
						RAPoints:         524,
						RASoftcorePoints: 1615,
						DateAwarded:      dateAwarded,
						HardcoreMode:     0,
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementUnlocks, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 15906, resp.UnlocksCount)
				require.Equal(t, 8223, resp.UnlocksHardcoreCount)
				require.Equal(t, 38104, resp.TotalPlayers)
				require.Equal(t, 1, resp.Achievement.ID)
				require.Equal(t, "Ring Collector", resp.Achievement.Title)
				require.Equal(t, "Collect 100 rings", resp.Achievement.Description)
				require.Equal(t, 5, resp.Achievement.Points)
				require.Equal(t, 7, resp.Achievement.TrueRatio)
				require.Equal(t, "Scott", resp.Achievement.Author)
				require.Equal(t, dateCreated, resp.Achievement.DateCreated.Time)
				require.Equal(t, dateModified, resp.Achievement.DateModified.Time)
				require.NotNil(t, resp.Achievement.Type)
				require.Equal(t, achievementType, *resp.Achievement.Type)
				require.Equal(t, 1, resp.Console.ID)
				require.Equal(t, "Genesis/Mega Drive", resp.Console.Title)
				require.Equal(t, 1, resp.Game.ID)
				require.Equal(t, "Sonic the Hedgehog", resp.Game.Title)
				require.Len(t, resp.Unlocks, 1)
				require.Equal(t, "redjedia", resp.Unlocks[0].User)
				require.Equal(t, 524, resp.Unlocks[0].RAPoints)
				require.Equal(t, 1615, resp.Unlocks[0].RASoftcorePoints)
				require.Equal(t, dateAwarded, resp.Unlocks[0].DateAwarded)
				require.Equal(t, 0, resp.Unlocks[0].HardcoreMode)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetAchievementUnlocks.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				messageBytes, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(messageBytes, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			resp, err := client.GetAchievementUnlocks(test.params)
			test.assert(t, resp, err)
		})
	}
}
