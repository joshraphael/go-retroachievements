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

func TestGetTicketByID(tt *testing.T) {
	achievementType := "progression"
	reportedAt, err := time.Parse(time.DateTime, "2014-02-22 23:23:53")
	require.NoError(tt, err)
	hardcore := 1
	resolvedAt, err := time.Parse(time.DateTime, "2014-02-24 22:51:10")
	require.NoError(tt, err)
	resolvedBy := "Scott"
	tests := []struct {
		name            string
		params          models.GetTicketByIDParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetTicketByID
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetTicketByID, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetTicketByIDParameters{
				TicketID: 1,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetTicketByID, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetTicketData.php?i=1&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetTicketByIDParameters{
				TicketID: 1,
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
			assert: func(t *testing.T, resp *models.GetTicketByID, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetTicketByIDParameters{
				TicketID: 1,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetTicketByID{
				ID:                1,
				AchievementID:     3778,
				AchievementTitle:  "Exhibition Match",
				AchievementDesc:   "Complete Stage 3: Flugelheim Museum",
				AchievementType:   &achievementType,
				Points:            10,
				BadgeName:         "04357",
				AchievementAuthor: "Batman",
				GameID:            45,
				ConsoleName:       "Genesis/Mega Drive",
				GameTitle:         "Batman: The Video Game",
				GameIcon:          "/Images/053393.png",
				ReportedAt: models.DateTime{
					Time: reportedAt,
				},
				ReportType:  0,
				ReportState: 2,
				Hardcore:    &hardcore,
				ReportNotes: "This achievement didn't trigger for some reason?",
				ReportedBy:  "qwe",
				ResolvedAt: &models.DateTime{
					Time: resolvedAt,
				},
				ResolvedBy:             &resolvedBy,
				ReportStateDescription: "Resolved",
				ReportTypeDescription:  "Invalid ticket type",
				URL:                    "https://retroachievements.org/ticket/1",
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetTicketByID, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 1, resp.ID)
				require.Equal(t, 3778, resp.AchievementID)
				require.Equal(t, "Exhibition Match", resp.AchievementTitle)
				require.Equal(t, "Complete Stage 3: Flugelheim Museum", resp.AchievementDesc)
				require.NotNil(t, resp.AchievementType)
				require.Equal(t, achievementType, *resp.AchievementType)
				require.Equal(t, 10, resp.Points)
				require.Equal(t, "04357", resp.BadgeName)
				require.Equal(t, "Batman", resp.AchievementAuthor)
				require.Equal(t, 45, resp.GameID)
				require.Equal(t, "Genesis/Mega Drive", resp.ConsoleName)
				require.Equal(t, "Batman: The Video Game", resp.GameTitle)
				require.Equal(t, "/Images/053393.png", resp.GameIcon)
				require.Equal(t, reportedAt, resp.ReportedAt.Time)
				require.Equal(t, 0, resp.ReportType)
				require.Equal(t, 2, resp.ReportState)
				require.NotNil(t, resp.ResolvedBy)
				require.Equal(t, resolvedBy, *resp.ResolvedBy)
				require.Equal(t, "Resolved", resp.ReportStateDescription)
				require.Equal(t, "Invalid ticket type", resp.ReportTypeDescription)
				require.Equal(t, "https://retroachievements.org/ticket/1", resp.URL)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetTicketData.php"
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
			resp, err := client.GetTicketByID(test.params)
			test.assert(t, resp, err)
		})
	}
}
