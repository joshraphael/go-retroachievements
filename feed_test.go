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

func TestGetRecentGameAwards(tt *testing.T) {
	count := 10
	offset := 10
	startingDate, err := time.Parse(models.RFC3339NumColonTZFormat, "2024-10-05T18:30:59+00:00")
	require.NoError(tt, err)
	awardDate, err := time.Parse(models.RFC3339NumColonTZFormat, "2024-11-23T13:39:21+00:00")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetRecentGameAwardsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetRecentGameAwards
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetRecentGameAwards, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetRecentGameAwardsParameters{
				StartingDate: &startingDate,
				Count:        &count,
				Offset:       &offset,
				IncludePartialAwards: &models.GetRecentGameAwardsParametersPartialAwards{
					BeatenSoftcore: true,
					BeatenHardcore: true,
					Completed:      true,
					Mastered:       true,
				},
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetRecentGameAwards, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetRecentGameAwards.php?c=10&d=2024-10-05&k=beaten-softcore%2Cbeaten-hardcore%2Ccompleted%2Cmastered&o=10&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetRecentGameAwardsParameters{
				StartingDate: &startingDate,
				Count:        &count,
				Offset:       &offset,
				IncludePartialAwards: &models.GetRecentGameAwardsParametersPartialAwards{
					BeatenSoftcore: true,
					BeatenHardcore: true,
					Completed:      true,
					Mastered:       true,
				},
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
			assert: func(t *testing.T, resp *models.GetRecentGameAwards, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetRecentGameAwardsParameters{
				StartingDate: &startingDate,
				Count:        &count,
				Offset:       &offset,
				IncludePartialAwards: &models.GetRecentGameAwardsParametersPartialAwards{
					BeatenSoftcore: true,
					BeatenHardcore: true,
					Completed:      true,
					Mastered:       true,
				},
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetRecentGameAwards{
				Total: 2,
				Count: 1,
				Results: []models.GetRecentGameAwardsResult{
					{
						User:      "spoony",
						AwardKind: "mastered",
						AwardDate: models.RFC3339NumColonTZ{
							Time: awardDate,
						},
						GameID:      7317,
						GameTitle:   "~Hack~ Pokémon Brown Version",
						ConsoleID:   4,
						ConsoleName: "Game Boy",
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetRecentGameAwards, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 2, resp.Total)
				require.Equal(t, 1, resp.Count)
				require.Len(t, resp.Results, 1)
				require.Equal(t, "spoony", resp.Results[0].User)
				require.Equal(t, "mastered", resp.Results[0].AwardKind)
				require.Equal(t, awardDate, resp.Results[0].AwardDate.Time)
				require.Equal(t, 7317, resp.Results[0].GameID)
				require.Equal(t, "~Hack~ Pokémon Brown Version", resp.Results[0].GameTitle)
				require.Equal(t, 4, resp.Results[0].ConsoleID)
				require.Equal(t, "Game Boy", resp.Results[0].ConsoleName)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetRecentGameAwards.php"
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
			resp, err := client.GetRecentGameAwards(test.params)
			test.assert(t, resp, err)
		})
	}
}
