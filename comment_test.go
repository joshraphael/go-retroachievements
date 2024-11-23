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

func TestGetComments(tt *testing.T) {
	count := 10
	offset := 12
	submitted, err := time.Parse(time.RFC3339Nano, "2013-04-08T12:37:12.000000Z")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetCommentsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetComments
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetComments, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetCommentsParameters{
				Type: models.GetCommentsUser{
					Username: "jamiras",
				},
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
			assert: func(t *testing.T, resp *models.GetComments, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetComments.php?c=10&i=jamiras&o=12&t=3&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetCommentsParameters{
				Type: models.GetCommentsUser{
					Username: "jamiras",
				},
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
			assert: func(t *testing.T, resp *models.GetComments, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success - game",
			params: models.GetCommentsParameters{
				Type: models.GetCommentsGame{
					GameID: 123,
				},
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetComments{
				Count: 1,
				Total: 1,
				Results: []models.GetCommentsResult{
					{
						User:        "Scott",
						Submitted:   submitted,
						CommentText: "Next I think we should have leaderboards/time trials for each act...",
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetComments, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 1, resp.Count)
				require.Equal(t, 1, resp.Total)
				require.Len(t, resp.Results, 1)
				require.Equal(t, "Scott", resp.Results[0].User)
				require.Equal(t, submitted, resp.Results[0].Submitted)
				require.Equal(t, "Next I think we should have leaderboards/time trials for each act...", resp.Results[0].CommentText)
				require.NoError(t, err)
			},
		},
		{
			name: "success - achievement",
			params: models.GetCommentsParameters{
				Type: models.GetCommentsAchievement{
					AchievementID: 123,
				},
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetComments{
				Count: 1,
				Total: 1,
				Results: []models.GetCommentsResult{
					{
						User:        "Scott",
						Submitted:   submitted,
						CommentText: "Next I think we should have leaderboards/time trials for each act...",
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetComments, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 1, resp.Count)
				require.Equal(t, 1, resp.Total)
				require.Len(t, resp.Results, 1)
				require.Equal(t, "Scott", resp.Results[0].User)
				require.Equal(t, submitted, resp.Results[0].Submitted)
				require.Equal(t, "Next I think we should have leaderboards/time trials for each act...", resp.Results[0].CommentText)
				require.NoError(t, err)
			},
		},
		{
			name: "success - achievement",
			params: models.GetCommentsParameters{
				Type: models.GetCommentsUser{
					Username: "123",
				},
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetComments{
				Count: 1,
				Total: 1,
				Results: []models.GetCommentsResult{
					{
						User:        "Scott",
						Submitted:   submitted,
						CommentText: "Next I think we should have leaderboards/time trials for each act...",
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetComments, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 1, resp.Count)
				require.Equal(t, 1, resp.Total)
				require.Len(t, resp.Results, 1)
				require.Equal(t, "Scott", resp.Results[0].User)
				require.Equal(t, submitted, resp.Results[0].Submitted)
				require.Equal(t, "Next I think we should have leaderboards/time trials for each act...", resp.Results[0].CommentText)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetComments.php"
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
			resp, err := client.GetComments(test.params)
			test.assert(t, resp, err)
		})
	}
}
