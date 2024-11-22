package retroachievements_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
	"github.com/stretchr/testify/require"
)

func TestGetConsoleIDs(tt *testing.T) {
	active := true
	gameSystems := true
	tests := []struct {
		name            string
		params          models.GetConsoleIDsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetConsoleIDs
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp []models.GetConsoleIDs, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetConsoleIDsParameters{
				OnlyActive:      &active,
				OnlyGameSystems: &gameSystems,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetConsoleIDs, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetConsoleIDs.php?a=1&g=1&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetConsoleIDsParameters{
				OnlyActive:      &active,
				OnlyGameSystems: &gameSystems,
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
			assert: func(t *testing.T, resp []models.GetConsoleIDs, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response list: error responses: [401] Not Authorized")
			},
		},
		{
			name: "success",
			params: models.GetConsoleIDsParameters{
				OnlyActive:      &active,
				OnlyGameSystems: &gameSystems,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetConsoleIDs{
				{
					ID:           1,
					Name:         "Genesis/Mega Drive",
					IconURL:      "https://static.retroachievements.org/assets/images/system/md.png",
					Active:       true,
					IsGameSystem: true,
				},
				{
					ID:           2,
					Name:         "Nintendo 64",
					IconURL:      "https://static.retroachievements.org/assets/images/system/n64.png",
					Active:       true,
					IsGameSystem: true,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetConsoleIDs, err error) {
				require.NotNil(t, resp)
				require.Len(t, resp, 2)
				require.Equal(t, 1, resp[0].ID)
				require.Equal(t, "Genesis/Mega Drive", resp[0].Name)
				require.Equal(t, "https://static.retroachievements.org/assets/images/system/md.png", resp[0].IconURL)
				require.True(t, resp[0].Active)
				require.True(t, resp[0].IsGameSystem)
				require.Equal(t, 2, resp[1].ID)
				require.Equal(t, "Nintendo 64", resp[1].Name)
				require.Equal(t, "https://static.retroachievements.org/assets/images/system/n64.png", resp[1].IconURL)
				require.True(t, resp[1].Active)
				require.True(t, resp[1].IsGameSystem)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetConsoleIDs.php"
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
			resp, err := client.GetConsoleIDs(test.params)
			test.assert(t, resp, err)
		})
	}
}
