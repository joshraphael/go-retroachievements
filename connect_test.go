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

func TestGetCodeNotes(tt *testing.T) {
	tests := []struct {
		name            string
		params          models.GetCodeNotesParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetCodeNotes
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetCodeNotes, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetCodeNotesParameters{
				GameID: 13214,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetCodeNotes, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/dorequest.php?g=13214&r=codenotes2\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetCodeNotesParameters{
				GameID: 13214,
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
			assert: func(t *testing.T, resp *models.GetCodeNotes, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetCodeNotesParameters{
				GameID: 13214,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetCodeNotes{
				Success: true,
				CodeNotes: []models.GetCodeNotesCodeNote{
					{
						User:    "jamiras",
						Address: "0x00000",
						Note:    "test note 1",
					},
					{
						User:    "jamiras",
						Address: "0x00001",
						Note:    "test note 2",
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetCodeNotes, err error) {
				require.NotNil(t, resp)
				require.True(t, resp.Success)
				require.Len(t, resp.CodeNotes, 2)
				require.Equal(t, "jamiras", resp.CodeNotes[0].User)
				require.Equal(t, "0x00000", resp.CodeNotes[0].Address)
				require.Equal(t, "test note 1", resp.CodeNotes[0].Note)
				require.Equal(t, "jamiras", resp.CodeNotes[1].User)
				require.Equal(t, "0x00001", resp.CodeNotes[1].Address)
				require.Equal(t, "test note 2", resp.CodeNotes[1].Note)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/dorequest.php"
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
				ConnectConfig: &retroachievements.ClientConnectConfig{
					ConnectSecret:   "some_other_secret",
					ConnectUsername: "jamiras",
				},
			})
			resp, err := client.GetCodeNotes(test.params)
			test.assert(t, resp, err)
		})
	}
}
