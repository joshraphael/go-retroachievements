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

func TestGetGameLeaderboards(tt *testing.T) {
	count := 10
	offset := 10
	tests := []struct {
		name            string
		params          models.GetGameLeaderboardsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetGameLeaderboards
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetGameLeaderboards, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetGameLeaderboardsParameters{
				GameID: 14402,
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGameLeaderboards, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetGameLeaderboards.php?c=10&i=14402&o=10&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetGameLeaderboardsParameters{
				GameID: 14402,
				Count:  &count,
				Offset: &offset,
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
			assert: func(t *testing.T, resp *models.GetGameLeaderboards, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetGameLeaderboardsParameters{
				GameID: 14402,
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetGameLeaderboards{
				Count: 1,
				Total: 2,
				Results: []models.GetGameLeaderboardsResult{
					{
						ID:          114798,
						RankAsc:     true,
						Title:       "Speedrun Monster Max",
						Description: "Complete the game from start to finish as fast as possible without using passwords",
						Format:      "TIME",
						TopEntry: &models.GetGameLeaderboardsTopEntry{
							User:           "joshraphael",
							Score:          2267,
							FormattedScore: "0:37.78",
						},
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGameLeaderboards, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 1, resp.Count)
				require.Equal(t, 2, resp.Total)
				require.Len(t, resp.Results, 1)
				require.Equal(t, 114798, resp.Results[0].ID)
				require.True(t, resp.Results[0].RankAsc)
				require.Equal(t, "Speedrun Monster Max", resp.Results[0].Title)
				require.Equal(t, "Complete the game from start to finish as fast as possible without using passwords", resp.Results[0].Description)
				require.Equal(t, "TIME", resp.Results[0].Format)
				require.NotNil(t, resp.Results[0].TopEntry)
				require.Equal(t, "joshraphael", resp.Results[0].TopEntry.User)
				require.Equal(t, 2267, resp.Results[0].TopEntry.Score)
				require.Equal(t, "0:37.78", resp.Results[0].TopEntry.FormattedScore)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetGameLeaderboards.php"
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetGameLeaderboards(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetLeaderboardEntries(tt *testing.T) {
	count := 10
	offset := 10
	dateSubmitted, err := time.Parse(models.RFC3339NumColonTZFormat, "2024-10-05T18:30:59+00:00")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetLeaderboardEntriesParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetLeaderboardEntries
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetLeaderboardEntries, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetLeaderboardEntriesParameters{
				LeaderboardID: 14402,
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
			assert: func(t *testing.T, resp *models.GetLeaderboardEntries, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetLeaderboardEntries.php?c=10&i=14402&o=10&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetLeaderboardEntriesParameters{
				LeaderboardID: 14402,
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
			assert: func(t *testing.T, resp *models.GetLeaderboardEntries, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetLeaderboardEntriesParameters{
				LeaderboardID: 14402,
				Count:         &count,
				Offset:        &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetLeaderboardEntries{
				Count: 1,
				Total: 2,
				Results: []models.GetLeaderboardEntriesResult{
					{
						User: "ramenoid",
						DateSubmitted: models.RFC3339NumColonTZ{
							Time: dateSubmitted,
						},
						Score:          1908730,
						FormattedScore: "1,908,730",
						Rank:           1,
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetLeaderboardEntries, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 1, resp.Count)
				require.Equal(t, 2, resp.Total)
				require.Len(t, resp.Results, 1)
				require.Equal(t, "ramenoid", resp.Results[0].User)
				require.Equal(t, dateSubmitted, resp.Results[0].DateSubmitted.Time)
				require.Equal(t, 1908730, resp.Results[0].Score)
				require.Equal(t, "1,908,730", resp.Results[0].FormattedScore)
				require.Equal(t, 1, resp.Results[0].Rank)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetLeaderboardEntries.php"
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetLeaderboardEntries(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetUserGameLeaderboards(tt *testing.T) {
	count := 10
	offset := 10
	dateUpdated, err := time.Parse(models.RFC3339NumColonTZFormat, "2024-10-05T18:30:59+00:00")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetUserGameLeaderboardsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUserGameLeaderboards
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetUserGameLeaderboards, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserGameLeaderboardsParameters{
				GameID:   515,
				Username: "test",
				Count:    &count,
				Offset:   &offset,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetUserGameLeaderboards, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserGameLeaderboards.php?c=10&i=515&o=10&u=test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserGameLeaderboardsParameters{
				GameID:   515,
				Username: "test",
				Count:    &count,
				Offset:   &offset,
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
			assert: func(t *testing.T, resp *models.GetUserGameLeaderboards, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserGameLeaderboardsParameters{
				GameID:   515,
				Username: "test",
				Count:    &count,
				Offset:   &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUserGameLeaderboards{
				Count: 1,
				Total: 2,
				Results: []models.GetUserGameLeaderboardsResult{
					{
						ID:          114798,
						RankAsc:     true,
						Title:       "Speedrun Monster Max",
						Description: "Complete the game from start to finish as fast as possible without using passwords",
						Format:      "TIME",
						UserEntry: models.GetUserGameLeaderboardsUserEntry{
							User: "ramenoid",
							DateUpdated: models.RFC3339NumColonTZ{
								Time: dateUpdated,
							},
							Score:          1908730,
							FormattedScore: "1,908,730",
							Rank:           1,
						},
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetUserGameLeaderboards, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 1, resp.Count)
				require.Equal(t, 2, resp.Total)
				require.Len(t, resp.Results, 1)
				require.Equal(t, 114798, resp.Results[0].ID)
				require.True(t, resp.Results[0].RankAsc)
				require.Equal(t, "Speedrun Monster Max", resp.Results[0].Title)
				require.Equal(t, "Complete the game from start to finish as fast as possible without using passwords", resp.Results[0].Description)
				require.Equal(t, "TIME", resp.Results[0].Format)
				require.Equal(t, "ramenoid", resp.Results[0].UserEntry.User)
				require.Equal(t, 1908730, resp.Results[0].UserEntry.Score)
				require.Equal(t, "1,908,730", resp.Results[0].UserEntry.FormattedScore)
				require.Equal(t, 1, resp.Results[0].UserEntry.Rank)
				require.Equal(t, dateUpdated, resp.Results[0].UserEntry.DateUpdated.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserGameLeaderboards.php"
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetUserGameLeaderboards(test.params)
			test.assert(t, resp, err)
		})
	}
}
