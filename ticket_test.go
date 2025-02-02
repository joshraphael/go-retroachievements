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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetTicketByID(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetMostTicketedGames(tt *testing.T) {
	count := 10
	offset := 12
	tests := []struct {
		name            string
		params          models.GetMostTicketedGamesParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetMostTicketedGames
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetMostTicketedGames, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetMostTicketedGamesParameters{
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
			assert: func(t *testing.T, resp *models.GetMostTicketedGames, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetTicketData.php?c=10&f=1&o=12&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetMostTicketedGamesParameters{
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
			assert: func(t *testing.T, resp *models.GetMostTicketedGames, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetMostTicketedGamesParameters{
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetMostTicketedGames{
				MostReportedGames: []models.GetMostTicketedGamesMostReportedGame{
					{
						GameID:      8301,
						GameTitle:   "Magical Vacation",
						GameIcon:    "/Images/080447.png",
						Console:     "Game Boy Advance",
						OpenTickets: 4,
					},
				},
				URL: "https://retroachievements.org/manage/most-reported-games",
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetMostTicketedGames, err error) {
				require.NotNil(t, resp)
				require.Len(t, resp.MostReportedGames, 1)
				require.Equal(t, 8301, resp.MostReportedGames[0].GameID)
				require.Equal(t, "Magical Vacation", resp.MostReportedGames[0].GameTitle)
				require.Equal(t, "/Images/080447.png", resp.MostReportedGames[0].GameIcon)
				require.Equal(t, "Game Boy Advance", resp.MostReportedGames[0].Console)
				require.Equal(t, 4, resp.MostReportedGames[0].OpenTickets)
				require.Equal(t, "https://retroachievements.org/manage/most-reported-games", resp.URL)
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetMostTicketedGames(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetMostRecentTickets(tt *testing.T) {
	count := 10
	offset := 12
	achievementType := "progression"
	reportedAt, err := time.Parse(time.DateTime, "2014-02-22 23:23:53")
	require.NoError(tt, err)
	hardcore := 1
	resolvedAt, err := time.Parse(time.DateTime, "2014-02-24 22:51:10")
	require.NoError(tt, err)
	resolvedBy := "Scott"
	tests := []struct {
		name            string
		params          models.GetMostRecentTicketsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetMostRecentTickets
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetMostRecentTickets, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetMostRecentTicketsParameters{
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
			assert: func(t *testing.T, resp *models.GetMostRecentTickets, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetTicketData.php?c=10&o=12&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetMostRecentTicketsParameters{
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
			assert: func(t *testing.T, resp *models.GetMostRecentTickets, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetMostRecentTicketsParameters{
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetMostRecentTickets{
				OpenTickets: 1951,
				URL:         "https://retroachievements.org/tickets",
				RecentTickets: []models.GetMostRecentTicketsRecentTicket{
					{
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
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetMostRecentTickets, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 1951, resp.OpenTickets)
				require.Equal(t, "https://retroachievements.org/tickets", resp.URL)
				require.Len(t, resp.RecentTickets, 1)
				require.Equal(t, 1, resp.RecentTickets[0].ID)
				require.Equal(t, 3778, resp.RecentTickets[0].AchievementID)
				require.Equal(t, "Exhibition Match", resp.RecentTickets[0].AchievementTitle)
				require.Equal(t, "Complete Stage 3: Flugelheim Museum", resp.RecentTickets[0].AchievementDesc)
				require.NotNil(t, resp.RecentTickets[0].AchievementType)
				require.Equal(t, achievementType, *resp.RecentTickets[0].AchievementType)
				require.Equal(t, 10, resp.RecentTickets[0].Points)
				require.Equal(t, "04357", resp.RecentTickets[0].BadgeName)
				require.Equal(t, "Batman", resp.RecentTickets[0].AchievementAuthor)
				require.Equal(t, 45, resp.RecentTickets[0].GameID)
				require.Equal(t, "Genesis/Mega Drive", resp.RecentTickets[0].ConsoleName)
				require.Equal(t, "Batman: The Video Game", resp.RecentTickets[0].GameTitle)
				require.Equal(t, "/Images/053393.png", resp.RecentTickets[0].GameIcon)
				require.Equal(t, reportedAt, resp.RecentTickets[0].ReportedAt.Time)
				require.Equal(t, 0, resp.RecentTickets[0].ReportType)
				require.Equal(t, 2, resp.RecentTickets[0].ReportState)
				require.NotNil(t, resp.RecentTickets[0].Hardcore)
				require.Equal(t, hardcore, *resp.RecentTickets[0].Hardcore)
				require.NotNil(t, resp.RecentTickets[0].ResolvedAt)
				require.Equal(t, resolvedAt, resp.RecentTickets[0].ResolvedAt.Time)
				require.NotNil(t, resp.RecentTickets[0].ResolvedBy)
				require.Equal(t, resolvedBy, *resp.RecentTickets[0].ResolvedBy)
				require.Equal(t, "Resolved", resp.RecentTickets[0].ReportStateDescription)
				require.Equal(t, "Invalid ticket type", resp.RecentTickets[0].ReportTypeDescription)
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetMostRecentTickets(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetGameTicketStats(tt *testing.T) {
	unofficial := true
	metadata := true
	achievementType := "progression"
	reportedAt, err := time.Parse(time.DateTime, "2014-02-22 23:23:53")
	require.NoError(tt, err)
	hardcore := 1
	resolvedAt, err := time.Parse(time.DateTime, "2014-02-24 22:51:10")
	require.NoError(tt, err)
	resolvedBy := "Scott"
	tests := []struct {
		name            string
		params          models.GetGameTicketStatsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetGameTicketStats
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetGameTicketStats, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetGameTicketStatsParameters{
				GameID:                1,
				Unofficial:            &unofficial,
				IncludeTicketMetadata: &metadata,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGameTicketStats, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetTicketData.php?d=1&f=5&g=1&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetGameTicketStatsParameters{
				GameID:                1,
				Unofficial:            &unofficial,
				IncludeTicketMetadata: &metadata,
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
			assert: func(t *testing.T, resp *models.GetGameTicketStats, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetGameTicketStatsParameters{
				GameID:                1,
				Unofficial:            &unofficial,
				IncludeTicketMetadata: &metadata,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetGameTicketStats{
				GameID:      20,
				GameTitle:   "Alex Kidd in the Enchanted Castle",
				ConsoleName: "Genesis/Mega Drive",
				OpenTickets: 4,
				URL:         "https://retroachievements.org/game/20/tickets",
				Tickets: []models.GetGameTicketStatsTicket{
					{
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
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGameTicketStats, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 20, resp.GameID)
				require.Equal(t, "Alex Kidd in the Enchanted Castle", resp.GameTitle)
				require.Equal(t, "Genesis/Mega Drive", resp.ConsoleName)
				require.Equal(t, 4, resp.OpenTickets)
				require.Equal(t, "https://retroachievements.org/game/20/tickets", resp.URL)
				require.Len(t, resp.Tickets, 1)
				require.Equal(t, 1, resp.Tickets[0].ID)
				require.Equal(t, 3778, resp.Tickets[0].AchievementID)
				require.Equal(t, "Exhibition Match", resp.Tickets[0].AchievementTitle)
				require.Equal(t, "Complete Stage 3: Flugelheim Museum", resp.Tickets[0].AchievementDesc)
				require.NotNil(t, resp.Tickets[0].AchievementType)
				require.Equal(t, achievementType, *resp.Tickets[0].AchievementType)
				require.Equal(t, 10, resp.Tickets[0].Points)
				require.Equal(t, "04357", resp.Tickets[0].BadgeName)
				require.Equal(t, "Batman", resp.Tickets[0].AchievementAuthor)
				require.Equal(t, 45, resp.Tickets[0].GameID)
				require.Equal(t, "Genesis/Mega Drive", resp.Tickets[0].ConsoleName)
				require.Equal(t, "Batman: The Video Game", resp.Tickets[0].GameTitle)
				require.Equal(t, "/Images/053393.png", resp.Tickets[0].GameIcon)
				require.Equal(t, reportedAt, resp.Tickets[0].ReportedAt.Time)
				require.Equal(t, 0, resp.Tickets[0].ReportType)
				require.Equal(t, 2, resp.Tickets[0].ReportState)
				require.NotNil(t, resp.Tickets[0].Hardcore)
				require.Equal(t, hardcore, *resp.Tickets[0].Hardcore)
				require.NotNil(t, resp.Tickets[0].ResolvedAt)
				require.Equal(t, resolvedAt, resp.Tickets[0].ResolvedAt.Time)
				require.NotNil(t, resp.Tickets[0].ResolvedBy)
				require.Equal(t, resolvedBy, *resp.Tickets[0].ResolvedBy)
				require.Equal(t, "Resolved", resp.Tickets[0].ReportStateDescription)
				require.Equal(t, "Invalid ticket type", resp.Tickets[0].ReportTypeDescription)
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetGameTicketStats(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetDeveloperTicketStats(tt *testing.T) {
	tests := []struct {
		name            string
		params          models.GetDeveloperTicketStatsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetDeveloperTicketStats
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetDeveloperTicketStats, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetDeveloperTicketStatsParameters{
				Username: "jamiras",
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetDeveloperTicketStats, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetTicketData.php?u=jamiras&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetDeveloperTicketStatsParameters{
				Username: "jamiras",
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
			assert: func(t *testing.T, resp *models.GetDeveloperTicketStats, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetDeveloperTicketStatsParameters{
				Username: "jamiras",
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetDeveloperTicketStats{
				User:     "jamiras",
				Open:     0,
				Closed:   46,
				Resolved: 68,
				Total:    114,
				URL:      "https://retroachievements.org/user/jamiras/tickets",
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetDeveloperTicketStats, err error) {
				require.NotNil(t, resp)
				require.Equal(t, "jamiras", resp.User)
				require.Equal(t, 0, resp.Open)
				require.Equal(t, 46, resp.Closed)
				require.Equal(t, 68, resp.Resolved)
				require.Equal(t, 114, resp.Total)
				require.Equal(t, "https://retroachievements.org/user/jamiras/tickets", resp.URL)
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetDeveloperTicketStats(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetAchievementTicketStats(tt *testing.T) {
	achievementType := "progression"
	tests := []struct {
		name            string
		params          models.GetAchievementTicketStatsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetAchievementTicketStats
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetAchievementTicketStats, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetAchievementTicketStatsParameters{
				AchievementID: 284759,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementTicketStats, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetTicketData.php?a=284759&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetAchievementTicketStatsParameters{
				AchievementID: 284759,
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
			assert: func(t *testing.T, resp *models.GetAchievementTicketStats, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetAchievementTicketStatsParameters{
				AchievementID: 284759,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetAchievementTicketStats{
				AchievementID:          284759,
				AchievementTitle:       "The End of the Beginning",
				AchievementDescription: "Receive the package from the King of Baron, and begin your quest to the Mist Cavern.",
				AchievementType:        &achievementType,
				URL:                    "https://retroachievements.org/achievement/284759/tickets",
				OpenTickets:            0,
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementTicketStats, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 284759, resp.AchievementID)
				require.Equal(t, "The End of the Beginning", resp.AchievementTitle)
				require.Equal(t, "Receive the package from the King of Baron, and begin your quest to the Mist Cavern.", resp.AchievementDescription)
				require.NotNil(t, resp.AchievementType)
				require.Equal(t, achievementType, *resp.AchievementType)
				require.Equal(t, "https://retroachievements.org/achievement/284759/tickets", resp.URL)
				require.Equal(t, 0, resp.OpenTickets)
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetAchievementTicketStats(test.params)
			test.assert(t, resp, err)
		})
	}
}
