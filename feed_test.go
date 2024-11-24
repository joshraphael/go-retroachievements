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

func TestGetActiveClaims(tt *testing.T) {
	created, err := time.Parse(time.DateTime, "2024-11-24 13:45:14")
	require.NoError(tt, err)
	doneTime, err := time.Parse(time.DateTime, "2025-02-24 13:45:14")
	require.NoError(tt, err)
	update, err := time.Parse(time.DateTime, "2024-11-24 13:45:14")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetActiveClaimsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetActiveClaims
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp []models.GetActiveClaims, err error)
	}{
		{
			name:   "fail to call endpoint",
			params: models.GetActiveClaimsParameters{},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetActiveClaims, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetActiveClaims.php?y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name:   "error response",
			params: models.GetActiveClaimsParameters{},
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
			assert: func(t *testing.T, resp []models.GetActiveClaims, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name:   "success",
			params: models.GetActiveClaimsParameters{},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetActiveClaims{
				{
					ID:          14667,
					User:        "Tayadaoc",
					GameID:      29805,
					GameTitle:   "Tetras",
					GameIcon:    "/Images/097197.png",
					ConsoleID:   47,
					ConsoleName: "PC-8000/8800",
					ClaimType:   0,
					SetType:     0,
					Status:      0,
					Extension:   0,
					Special:     0,
					Created: models.DateTime{
						Time: created,
					},
					DoneTime: models.DateTime{
						Time: doneTime,
					},
					Updated: models.DateTime{
						Time: update,
					},
					UserIsJrDev: 0,
					MinutesLeft: 132413,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetActiveClaims, err error) {
				require.NotNil(t, resp)
				require.Len(t, resp, 1)
				require.Equal(t, 14667, resp[0].ID)
				require.Equal(t, "Tayadaoc", resp[0].User)
				require.Equal(t, 29805, resp[0].GameID)
				require.Equal(t, "Tetras", resp[0].GameTitle)
				require.Equal(t, "/Images/097197.png", resp[0].GameIcon)
				require.Equal(t, 47, resp[0].ConsoleID)
				require.Equal(t, "PC-8000/8800", resp[0].ConsoleName)
				require.Equal(t, 0, resp[0].ClaimType)
				require.Equal(t, 0, resp[0].SetType)
				require.Equal(t, 0, resp[0].Status)
				require.Equal(t, 0, resp[0].Extension)
				require.Equal(t, 0, resp[0].Special)
				require.Equal(t, created, resp[0].Created.Time)
				require.Equal(t, doneTime, resp[0].DoneTime.Time)
				require.Equal(t, update, resp[0].Updated.Time)
				require.Equal(t, 0, resp[0].UserIsJrDev)
				require.Equal(t, 132413, resp[0].MinutesLeft)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetActiveClaims.php"
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
			resp, err := client.GetActiveClaims(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetClaims(tt *testing.T) {
	created, err := time.Parse(time.DateTime, "2024-11-24 13:45:14")
	require.NoError(tt, err)
	doneTime, err := time.Parse(time.DateTime, "2025-02-24 13:45:14")
	require.NoError(tt, err)
	update, err := time.Parse(time.DateTime, "2024-11-24 13:45:14")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetClaimsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetClaims
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp []models.GetClaims, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetClaimsParameters{
				Kind: &models.GetClaimsParametersKindDropped{},
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetClaims, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetClaims.php?k=2&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetClaimsParameters{
				Kind: &models.GetClaimsParametersKindCompleted{},
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
			assert: func(t *testing.T, resp []models.GetClaims, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetClaimsParameters{
				Kind: &models.GetClaimsParametersKindCompleted{},
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetClaims{
				{
					ID:          14667,
					User:        "Tayadaoc",
					GameID:      29805,
					GameTitle:   "Tetras",
					GameIcon:    "/Images/097197.png",
					ConsoleID:   47,
					ConsoleName: "PC-8000/8800",
					ClaimType:   0,
					SetType:     0,
					Status:      0,
					Extension:   0,
					Special:     0,
					Created: models.DateTime{
						Time: created,
					},
					DoneTime: models.DateTime{
						Time: doneTime,
					},
					Updated: models.DateTime{
						Time: update,
					},
					UserIsJrDev: 0,
					MinutesLeft: 132413,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetClaims, err error) {
				require.NotNil(t, resp)
				require.Len(t, resp, 1)
				require.Equal(t, 14667, resp[0].ID)
				require.Equal(t, "Tayadaoc", resp[0].User)
				require.Equal(t, 29805, resp[0].GameID)
				require.Equal(t, "Tetras", resp[0].GameTitle)
				require.Equal(t, "/Images/097197.png", resp[0].GameIcon)
				require.Equal(t, 47, resp[0].ConsoleID)
				require.Equal(t, "PC-8000/8800", resp[0].ConsoleName)
				require.Equal(t, 0, resp[0].ClaimType)
				require.Equal(t, 0, resp[0].SetType)
				require.Equal(t, 0, resp[0].Status)
				require.Equal(t, 0, resp[0].Extension)
				require.Equal(t, 0, resp[0].Special)
				require.Equal(t, created, resp[0].Created.Time)
				require.Equal(t, doneTime, resp[0].DoneTime.Time)
				require.Equal(t, update, resp[0].Updated.Time)
				require.Equal(t, 0, resp[0].UserIsJrDev)
				require.Equal(t, 132413, resp[0].MinutesLeft)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetClaims.php"
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
			resp, err := client.GetClaims(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetTopTenUsers(tt *testing.T) {
	tests := []struct {
		name            string
		params          models.GetTopTenUsersParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetTopTenUsers
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp []models.GetTopTenUsers, err error)
	}{
		{
			name:   "fail to call endpoint",
			params: models.GetTopTenUsersParameters{},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetTopTenUsers, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetTopTenUsers.php?y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name:   "error response",
			params: models.GetTopTenUsersParameters{},
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
			assert: func(t *testing.T, resp []models.GetTopTenUsers, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name:   "success",
			params: models.GetTopTenUsersParameters{},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetTopTenUsers{
				{
					Username:      "MaxMilyin",
					HarcordPoints: 443235,
					RetroPoints:   1825214,
				},
				{
					Username:      "Sarconius",
					HarcordPoints: 427984,
					RetroPoints:   2913697,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetTopTenUsers, err error) {
				require.NotNil(t, resp)
				require.Len(t, resp, 2)
				require.Equal(t, resp[0].Username, "MaxMilyin")
				require.Equal(t, resp[0].HarcordPoints, 443235)
				require.Equal(t, resp[0].RetroPoints, 1825214)
				require.Equal(t, resp[1].Username, "Sarconius")
				require.Equal(t, resp[1].HarcordPoints, 427984)
				require.Equal(t, resp[1].RetroPoints, 2913697)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetTopTenUsers.php"
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
			resp, err := client.GetTopTenUsers(test.params)
			test.assert(t, resp, err)
		})
	}
}
