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

func TestGetUserProfile(tt *testing.T) {
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetUserProfileParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUserProfile
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, profile *models.GetUserProfile, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserProfileParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, profile *models.GetUserProfile, err error) {
				require.Nil(t, profile)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserProfile.php?u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserProfileParameters{
				Username: "Test",
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
			assert: func(t *testing.T, profile *models.GetUserProfile, err error) {
				require.Nil(t, profile)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserProfileParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUserProfile{
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
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, profile *models.GetUserProfile, err error) {
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
			profile, err := client.GetUserProfile(test.params)
			test.assert(t, profile, err)
		})
	}
}

func TestGetUserRecentAchievements(tt *testing.T) {
	lookback := 20
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetUserRecentAchievementsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetUserRecentAchievements
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, achievements []models.GetUserRecentAchievements, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserRecentAchievementsParameters{
				Username:        "Test",
				LookbackMinutes: &lookback,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, achievements []models.GetUserRecentAchievements, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserRecentAchievements.php?m=20&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserRecentAchievementsParameters{
				Username: "Test",
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
			assert: func(t *testing.T, achievements []models.GetUserRecentAchievements, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserRecentAchievementsParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetUserRecentAchievements{
				{
					Title:       "Beat Level 1",
					Description: "Finish level 1",
					Points:      10,
					TrueRatio:   234,
					Author:      "jamiras",
					Date: models.DateTime{
						Time: now,
					},
					HardcoreMode:  1,
					AchievementID: 34425,
					BadgeName:     "840124",
					Type:          "win_condition",
					GameTitle:     "Final Fantasy XXXXIIII",
					GameIcon:      "/Images/056340.png",
					GameID:        34897,
					ConsoleName:   "SNES",
					BadgeURL:      "/Badge/840124.png",
					GameURL:       "/game/34897",
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, achievements []models.GetUserRecentAchievements, err error) {
				require.NotNil(t, achievements)
				require.Len(t, achievements, 1)
				require.Equal(t, models.DateTime{
					Time: now,
				}, achievements[0].Date)
				require.Equal(t, 1, achievements[0].HardcoreMode)
				require.Equal(t, 34425, achievements[0].AchievementID)
				require.Equal(t, "Beat Level 1", achievements[0].Title)
				require.Equal(t, "Finish level 1", achievements[0].Description)
				require.Equal(t, "840124", achievements[0].BadgeName)
				require.Equal(t, 10, achievements[0].Points)
				require.Equal(t, 234, achievements[0].TrueRatio)
				require.Equal(t, "win_condition", achievements[0].Type)
				require.Equal(t, "jamiras", achievements[0].Author)
				require.Equal(t, "Final Fantasy XXXXIIII", achievements[0].GameTitle)
				require.Equal(t, "/Images/056340.png", achievements[0].GameIcon)
				require.Equal(t, 34897, achievements[0].GameID)
				require.Equal(t, "SNES", achievements[0].ConsoleName)
				require.Equal(t, "/Badge/840124.png", achievements[0].BadgeURL)
				require.Equal(t, "/game/34897", achievements[0].GameURL)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserRecentAchievements.php"
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
			achievements, err := client.GetUserRecentAchievements(test.params)
			test.assert(t, achievements, err)
		})
	}
}

func TestGetAchievementsEarnedBetween(tt *testing.T) {
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(tt, err)
	later := now.Add(10 * time.Minute)
	tests := []struct {
		name            string
		params          models.GetAchievementsEarnedBetweenParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetAchievementsEarnedBetween
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, achievements []models.GetAchievementsEarnedBetween, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetAchievementsEarnedBetweenParameters{
				Username: "Test",
				From:     now,
				To:       later,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, achievements []models.GetAchievementsEarnedBetween, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetAchievementsEarnedBetween.php?f=1709400423&t=1709401023&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetAchievementsEarnedBetweenParameters{
				Username: "Test",
				From:     now,
				To:       later,
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
			assert: func(t *testing.T, achievements []models.GetAchievementsEarnedBetween, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetAchievementsEarnedBetweenParameters{
				Username: "Test",
				From:     now,
				To:       later,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetAchievementsEarnedBetween{
				{
					Title:       "Beat Level 1",
					Description: "Finish level 1",
					Points:      10,
					TrueRatio:   234,
					Author:      "jamiras",
					Date: models.DateTime{
						Time: now,
					},
					HardcoreMode:  1,
					AchievementID: 34425,
					BadgeName:     "840124",
					Type:          "win_condition",
					GameTitle:     "Final Fantasy XXXXIIII",
					GameIcon:      "/Images/056340.png",
					GameID:        34897,
					ConsoleName:   "SNES",
					BadgeURL:      "/Badge/840124.png",
					GameURL:       "/game/34897",
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, achievements []models.GetAchievementsEarnedBetween, err error) {
				require.NotNil(t, achievements)
				require.Len(t, achievements, 1)
				require.Equal(t, models.DateTime{
					Time: now,
				}, achievements[0].Date)
				require.Equal(t, 1, achievements[0].HardcoreMode)
				require.Equal(t, 34425, achievements[0].AchievementID)
				require.Equal(t, "Beat Level 1", achievements[0].Title)
				require.Equal(t, "Finish level 1", achievements[0].Description)
				require.Equal(t, "840124", achievements[0].BadgeName)
				require.Equal(t, 10, achievements[0].Points)
				require.Equal(t, 234, achievements[0].TrueRatio)
				require.Equal(t, "win_condition", achievements[0].Type)
				require.Equal(t, "jamiras", achievements[0].Author)
				require.Equal(t, "Final Fantasy XXXXIIII", achievements[0].GameTitle)
				require.Equal(t, "/Images/056340.png", achievements[0].GameIcon)
				require.Equal(t, 34897, achievements[0].GameID)
				require.Equal(t, "SNES", achievements[0].ConsoleName)
				require.Equal(t, "/Badge/840124.png", achievements[0].BadgeURL)
				require.Equal(t, "/game/34897", achievements[0].GameURL)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetAchievementsEarnedBetween.php"
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
			achievements, err := client.GetAchievementsEarnedBetween(test.params)
			test.assert(t, achievements, err)
		})
	}
}

func TestGetAchievementsEarnedOnDay(tt *testing.T) {
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetAchievementsEarnedOnDayParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetAchievementsEarnedOnDay
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, achievements []models.GetAchievementsEarnedOnDay, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetAchievementsEarnedOnDayParameters{
				Username: "Test",
				Date:     now,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, achievements []models.GetAchievementsEarnedOnDay, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetAchievementsEarnedOnDay.php?d=2024-03-02&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetAchievementsEarnedOnDayParameters{
				Username: "Test",
				Date:     now,
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
			assert: func(t *testing.T, achievements []models.GetAchievementsEarnedOnDay, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetAchievementsEarnedOnDayParameters{
				Username: "Test",
				Date:     now,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetAchievementsEarnedOnDay{
				{
					Title:       "Beat Level 1",
					Description: "Finish level 1",
					Points:      10,
					TrueRatio:   234,
					Author:      "jamiras",
					Date: models.DateTime{
						Time: now,
					},
					HardcoreMode:  1,
					AchievementID: 34425,
					BadgeName:     "840124",
					Type:          "win_condition",
					GameTitle:     "Final Fantasy XXXXIIII",
					GameIcon:      "/Images/056340.png",
					GameID:        34897,
					ConsoleName:   "SNES",
					BadgeURL:      "/Badge/840124.png",
					GameURL:       "/game/34897",
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, achievements []models.GetAchievementsEarnedOnDay, err error) {
				require.NotNil(t, achievements)
				require.Len(t, achievements, 1)
				require.Equal(t, models.DateTime{
					Time: now,
				}, achievements[0].Date)
				require.Equal(t, 1, achievements[0].HardcoreMode)
				require.Equal(t, 34425, achievements[0].AchievementID)
				require.Equal(t, "Beat Level 1", achievements[0].Title)
				require.Equal(t, "Finish level 1", achievements[0].Description)
				require.Equal(t, "840124", achievements[0].BadgeName)
				require.Equal(t, 10, achievements[0].Points)
				require.Equal(t, 234, achievements[0].TrueRatio)
				require.Equal(t, "win_condition", achievements[0].Type)
				require.Equal(t, "jamiras", achievements[0].Author)
				require.Equal(t, "Final Fantasy XXXXIIII", achievements[0].GameTitle)
				require.Equal(t, "/Images/056340.png", achievements[0].GameIcon)
				require.Equal(t, 34897, achievements[0].GameID)
				require.Equal(t, "SNES", achievements[0].ConsoleName)
				require.Equal(t, "/Badge/840124.png", achievements[0].BadgeURL)
				require.Equal(t, "/game/34897", achievements[0].GameURL)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetAchievementsEarnedOnDay.php"
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
			achievements, err := client.GetAchievementsEarnedOnDay(test.params)
			test.assert(t, achievements, err)
		})
	}
}

func TestGetGameInfoAndUserProgress(tt *testing.T) {
	released, err := time.Parse(models.LongMonthDateFormat, "June 18, 2001")
	require.NoError(tt, err)
	require.NoError(tt, err)
	modified, err := time.Parse(time.DateTime, "2022-10-25 17:00:49")
	require.NoError(tt, err)
	created, err := time.Parse(time.DateTime, "2022-09-28 00:36:26")
	require.NoError(tt, err)
	granularity := "day"
	highestAwardKind := "mastered"
	awarded, err := time.Parse(models.RFC3339NumColonTZFormat, "2024-05-07T08:48:54+00:00")
	require.NoError(tt, err)
	awardMeta := true
	forumTopicId := 16654
	tests := []struct {
		name            string
		params          models.GetGameInfoAndUserProgressParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetGameInfoAndUserProgress
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, gameProgress *models.GetGameInfoAndUserProgress, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetGameInfoAndUserProgressParameters{
				Username:             "Test",
				GameID:               2991,
				IncludeAwardMetadata: &awardMeta,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, gameProgress *models.GetGameInfoAndUserProgress, err error) {
				require.Nil(t, gameProgress)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetGameInfoAndUserProgress.php?a=1&g=2991&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetGameInfoAndUserProgressParameters{
				Username:             "Test",
				GameID:               2991,
				IncludeAwardMetadata: &awardMeta,
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
			assert: func(t *testing.T, gameProgress *models.GetGameInfoAndUserProgress, err error) {
				require.Nil(t, gameProgress)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetGameInfoAndUserProgressParameters{
				Username:             "Test",
				GameID:               2991,
				IncludeAwardMetadata: &awardMeta,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetGameInfoAndUserProgress{
				Title:              "Twisted Metal: Black",
				SortTitle:          "twisted metal black",
				ConsoleID:          21,
				ForumTopicID:       &forumTopicId,
				Flags:              0,
				ImageIcon:          "/Images/057992.png",
				ImageTitle:         "/Images/056152.png",
				ImageIngame:        "/Images/056151.png",
				ImageBoxArt:        "/Images/050832.png",
				Publisher:          "Sony Computer Entertainment",
				Developer:          "Incognito Entertainment",
				Genre:              "Vehicular Combat",
				ID:                 2991,
				IsFinal:            0,
				RichPresencePatch:  "e7a5e12072a6c976a1146756726fdd8c",
				ConsoleName:        "PlayStation 2",
				NumDistinctPlayers: 1287,
				NumAchievements:    93,
				Achievements: map[int]models.GetGameInfoAndUserProgressAchievement{
					252117: {

						Title:              "Zorko Bros. Scrap & Salvage",
						Description:        "Destroy all enemies in Junkyard in Story Mode",
						Points:             5,
						TrueRatio:          5,
						Author:             "TheJediSonic",
						ID:                 252117,
						NumAwarded:         819,
						NumAwardedHardcore: 327,
						DateModified: models.DateTime{
							Time: modified,
						},
						DateCreated: models.DateTime{
							Time: created,
						},
						BadgeName:    "279805",
						DisplayOrder: 0,
						MemAddr:      "3cf81e50c3ff8387e5034b79478d9a04",
						Type:         "progression",
					},
				},
				Released: &models.DateOnly{
					Time: released,
				},
				ReleasedAtGranularity:    &granularity,
				NumAwardedToUser:         1244,
				NumAwardedToUserHardcore: 1234,
				UserCompletion:           "100.00%",
				UserCompletionHardcore:   "95.00%",
				HighestAwardKind:         &highestAwardKind,
				HighestAwardDate: &models.RFC3339NumColonTZ{
					Time: awarded,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, gameProgress *models.GetGameInfoAndUserProgress, err error) {
				require.NotNil(t, gameProgress)
				require.Equal(t, "Twisted Metal: Black", gameProgress.Title)
				require.Equal(t, "twisted metal black", gameProgress.SortTitle)
				require.Equal(t, 21, gameProgress.ConsoleID)
				require.Equal(t, "PlayStation 2", gameProgress.ConsoleName)
				require.Equal(t, 16654, *gameProgress.ForumTopicID)
				require.Equal(t, 0, gameProgress.Flags)
				require.Equal(t, "/Images/057992.png", gameProgress.ImageIcon)
				require.Equal(t, "/Images/056152.png", gameProgress.ImageTitle)
				require.Equal(t, "/Images/056151.png", gameProgress.ImageIngame)
				require.Equal(t, "/Images/050832.png", gameProgress.ImageBoxArt)
				require.Equal(t, "Sony Computer Entertainment", gameProgress.Publisher)
				require.Equal(t, "Incognito Entertainment", gameProgress.Developer)
				require.Equal(t, "Vehicular Combat", gameProgress.Genre)
				require.Equal(t, released, gameProgress.Released.Time)
				require.Equal(t, 2991, gameProgress.ID)
				require.Equal(t, 0, gameProgress.IsFinal)
				require.Equal(t, "e7a5e12072a6c976a1146756726fdd8c", gameProgress.RichPresencePatch)
				require.Equal(t, "PlayStation 2", gameProgress.ConsoleName)
				require.Equal(t, 1287, gameProgress.NumDistinctPlayers)
				require.Equal(t, 93, gameProgress.NumAchievements)
				require.Len(t, gameProgress.Achievements, 1)
				achievement, ok := gameProgress.Achievements[252117]
				require.True(t, ok)
				require.NotNil(t, achievement)
				require.Equal(t, "Zorko Bros. Scrap & Salvage", achievement.Title)
				require.Equal(t, "Destroy all enemies in Junkyard in Story Mode", achievement.Description)
				require.Equal(t, 5, achievement.Points)
				require.Equal(t, 5, achievement.TrueRatio)
				require.Equal(t, "TheJediSonic", achievement.Author)
				require.Equal(t, 252117, achievement.ID)
				require.Equal(t, 819, achievement.NumAwarded)
				require.Equal(t, 327, achievement.NumAwardedHardcore)
				require.Equal(t, modified, achievement.DateModified.Time)
				require.Equal(t, created, achievement.DateCreated.Time)
				require.Equal(t, "279805", achievement.BadgeName)
				require.Equal(t, 0, achievement.DisplayOrder)
				require.Equal(t, "3cf81e50c3ff8387e5034b79478d9a04", achievement.MemAddr)
				require.Equal(t, "progression", achievement.Type)
				require.Equal(t, models.DateOnly{
					Time: released,
				}, *gameProgress.Released)
				require.Equal(t, "day", *gameProgress.ReleasedAtGranularity)
				require.Equal(t, 1244, gameProgress.NumAwardedToUser)
				require.Equal(t, 1234, gameProgress.NumAwardedToUserHardcore)
				require.Equal(t, "100.00%", gameProgress.UserCompletion)
				require.Equal(t, "95.00%", gameProgress.UserCompletionHardcore)
				require.Equal(t, awarded, gameProgress.HighestAwardDate.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetGameInfoAndUserProgress.php"
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
			gameProgress, err := client.GetGameInfoAndUserProgress(test.params)
			test.assert(t, gameProgress, err)
		})
	}
}

func TestGetUserCompletionProgress(tt *testing.T) {
	highestAwardKind := "mastered"
	awarded, err := time.Parse(models.RFC3339NumColonTZFormat, "2024-05-07T08:48:54+00:00")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetUserCompletionProgressParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUserCompletionProgress
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, completionProgress *models.GetUserCompletionProgress, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserCompletionProgressParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, completionProgress *models.GetUserCompletionProgress, err error) {
				require.Nil(t, completionProgress)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserCompletionProgress.php?u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserCompletionProgressParameters{
				Username: "Test",
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
			assert: func(t *testing.T, completionProgress *models.GetUserCompletionProgress, err error) {
				require.Nil(t, completionProgress)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserCompletionProgressParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUserCompletionProgress{
				Count: 19,
				Total: 18,
				Results: []models.CompletionProgress{
					{
						GameID:             14068,
						Title:              "Donkey Kong",
						ImageIcon:          "/Images/044908.png",
						ConsoleID:          44,
						ConsoleName:        "ColecoVision",
						MaxPossible:        19,
						NumAwarded:         13,
						NumAwardedHardcore: 13,
						MostRecentAwardedDate: models.RFC3339NumColonTZ{
							Time: awarded,
						},
						HighestAwardKind: &highestAwardKind,
						HighestAwardDate: &models.RFC3339NumColonTZ{
							Time: awarded,
						},
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, completionProgress *models.GetUserCompletionProgress, err error) {
				require.NotNil(t, completionProgress)
				require.Equal(t, 19, completionProgress.Count)
				require.Equal(t, 18, completionProgress.Total)
				require.Len(t, completionProgress.Results, 1)
				require.Equal(t, 14068, completionProgress.Results[0].GameID)
				require.Equal(t, "Donkey Kong", completionProgress.Results[0].Title)
				require.Equal(t, "/Images/044908.png", completionProgress.Results[0].ImageIcon)
				require.Equal(t, 44, completionProgress.Results[0].ConsoleID)
				require.Equal(t, "ColecoVision", completionProgress.Results[0].ConsoleName)
				require.Equal(t, 19, completionProgress.Results[0].MaxPossible)
				require.Equal(t, 13, completionProgress.Results[0].NumAwarded)
				require.Equal(t, 13, completionProgress.Results[0].NumAwardedHardcore)
				require.Equal(t, awarded, completionProgress.Results[0].MostRecentAwardedDate.Time)
				require.Equal(t, highestAwardKind, *completionProgress.Results[0].HighestAwardKind)
				require.Equal(t, awarded, completionProgress.Results[0].HighestAwardDate.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserCompletionProgress.php"
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
			userCompletionProgress, err := client.GetUserCompletionProgress(test.params)
			test.assert(t, userCompletionProgress, err)
		})
	}
}

func TestGetUserAwards(tt *testing.T) {
	awardedAt, err := time.Parse(models.RFC3339NumColonTZFormat, "2024-05-03T23:24:11+00:00")
	require.NoError(tt, err)
	title := "Pokemon FireRed Version"
	consoleID := 5
	consoleName := "Game Boy Advance"
	flags := 1
	imageIcons := "/Images/074224.png"
	tests := []struct {
		name            string
		params          models.GetUserAwardsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUserAwards
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, userAwards *models.GetUserAwards, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserAwardsParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, userAwards *models.GetUserAwards, err error) {
				require.Nil(t, userAwards)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserAwards.php?u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserAwardsParameters{
				Username: "Test",
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
			assert: func(t *testing.T, userAwards *models.GetUserAwards, err error) {
				require.Nil(t, userAwards)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserAwardsParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUserAwards{
				TotalAwardsCount:          9,
				HiddenAwardsCount:         0,
				MasteryAwardsCount:        4,
				CompletionAwardsCount:     0,
				BeatenHardcoreAwardsCount: 4,
				BeatenSoftcoreAwardsCount: 0,
				EventAwardsCount:          0,
				SiteAwardsCount:           1,
				VisibleUserAwards: []models.Award{
					{
						AwardedAt: models.RFC3339NumColonTZ{
							Time: awardedAt,
						},
						AwardType:      "Game Beaten",
						AwardData:      515,
						AwardDataExtra: 1,
						DisplayOrder:   0,
						Title:          &title,
						ConsoleID:      &consoleID,
						ConsoleName:    &consoleName,
						Flags:          &flags,
						ImageIcon:      &imageIcons,
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, userAwards *models.GetUserAwards, err error) {
				require.NotNil(t, userAwards)
				require.Equal(t, 9, userAwards.TotalAwardsCount)
				require.Equal(t, 0, userAwards.HiddenAwardsCount)
				require.Equal(t, 4, userAwards.MasteryAwardsCount)
				require.Equal(t, 0, userAwards.CompletionAwardsCount)
				require.Equal(t, 4, userAwards.BeatenHardcoreAwardsCount)
				require.Equal(t, 0, userAwards.BeatenSoftcoreAwardsCount)
				require.Equal(t, 0, userAwards.EventAwardsCount)
				require.Equal(t, 1, userAwards.SiteAwardsCount)
				require.Len(t, userAwards.VisibleUserAwards, 1)
				require.Equal(t, awardedAt, userAwards.VisibleUserAwards[0].AwardedAt.Time)
				require.Equal(t, "Game Beaten", userAwards.VisibleUserAwards[0].AwardType)
				require.Equal(t, 515, userAwards.VisibleUserAwards[0].AwardData)
				require.Equal(t, 1, userAwards.VisibleUserAwards[0].AwardDataExtra)
				require.Equal(t, 0, userAwards.VisibleUserAwards[0].DisplayOrder)
				require.Equal(t, title, *userAwards.VisibleUserAwards[0].Title)
				require.Equal(t, consoleID, *userAwards.VisibleUserAwards[0].ConsoleID)
				require.Equal(t, consoleName, *userAwards.VisibleUserAwards[0].ConsoleName)
				require.Equal(t, flags, *userAwards.VisibleUserAwards[0].Flags)
				require.Equal(t, imageIcons, *userAwards.VisibleUserAwards[0].ImageIcon)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserAwards.php"
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
			userAwards, err := client.GetUserAwards(test.params)
			test.assert(t, userAwards, err)
		})
	}
}

func TestGetUserClaims(tt *testing.T) {
	created, err := time.Parse(time.DateTime, "2024-07-21 05:24:58")
	require.NoError(tt, err)
	done, err := time.Parse(time.DateTime, "2024-10-21 05:24:58")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetUserClaimsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetUserClaims
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, userClaims []models.GetUserClaims, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserClaimsParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, userClaims []models.GetUserClaims, err error) {
				require.Nil(t, userClaims)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserClaims.php?u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserClaimsParameters{
				Username: "Test",
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
			assert: func(t *testing.T, userClaims []models.GetUserClaims, err error) {
				require.Nil(t, userClaims)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserClaimsParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetUserClaims{
				{
					ID:          13657,
					User:        "joshraphael",
					GameID:      4111,
					GameTitle:   "Monster Max",
					GameIcon:    "/Images/059373.png",
					ConsoleID:   4,
					ConsoleName: "Game Boy",
					ClaimType:   0,
					SetType:     0,
					Status:      0,
					Extension:   0,
					Special:     0,
					Created: models.DateTime{
						Time: created,
					},
					DoneTime: models.DateTime{
						Time: done,
					},
					Updated: models.DateTime{
						Time: created,
					},
					UserIsJrDev: 1,
					MinutesLeft: 87089,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, userClaims []models.GetUserClaims, err error) {
				require.NotNil(t, userClaims)
				require.Len(t, userClaims, 1)
				require.Equal(t, 13657, userClaims[0].ID)
				require.Equal(t, "joshraphael", userClaims[0].User)
				require.Equal(t, 4111, userClaims[0].GameID)
				require.Equal(t, "Monster Max", userClaims[0].GameTitle)
				require.Equal(t, "/Images/059373.png", userClaims[0].GameIcon)
				require.Equal(t, 4, userClaims[0].ConsoleID)
				require.Equal(t, "Game Boy", userClaims[0].ConsoleName)
				require.Equal(t, 0, userClaims[0].ClaimType)
				require.Equal(t, 0, userClaims[0].SetType)
				require.Equal(t, 0, userClaims[0].Status)
				require.Equal(t, 0, userClaims[0].Extension)
				require.Equal(t, 0, userClaims[0].Special)
				require.Equal(t, created, userClaims[0].Created.Time)
				require.Equal(t, done, userClaims[0].DoneTime.Time)
				require.Equal(t, created, userClaims[0].Updated.Time)
				require.Equal(t, 1, userClaims[0].UserIsJrDev)
				require.Equal(t, 87089, userClaims[0].MinutesLeft)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserClaims.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			userClaims, err := client.GetUserClaims(test.params)
			test.assert(t, userClaims, err)
		})
	}
}

func TestGetUserGameRankAndScore(tt *testing.T) {
	lastAward, err := time.Parse(time.DateTime, "2024-05-07 08:48:54")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetUserGameRankAndScoreParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetUserGameRankAndScore
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, userGameRankScore []models.GetUserGameRankAndScore, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserGameRankAndScoreParameters{
				Username: "Test",
				GameID:   10,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, userGameRankScore []models.GetUserGameRankAndScore, err error) {
				require.Nil(t, userGameRankScore)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserGameRankAndScore.php?g=10&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserGameRankAndScoreParameters{
				Username: "Test",
				GameID:   10,
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
			assert: func(t *testing.T, userGameRankScore []models.GetUserGameRankAndScore, err error) {
				require.Nil(t, userGameRankScore)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserGameRankAndScoreParameters{
				Username: "Test",
				GameID:   10,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetUserGameRankAndScore{
				{
					User:       "joshraphael",
					UserRank:   699,
					TotalScore: 453,
					LastAward: &models.DateTime{
						Time: lastAward,
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, userGameRankScore []models.GetUserGameRankAndScore, err error) {
				require.NotNil(t, userGameRankScore)
				require.Len(t, userGameRankScore, 1)
				require.Equal(t, "joshraphael", userGameRankScore[0].User)
				require.Equal(t, 699, userGameRankScore[0].UserRank)
				require.Equal(t, 453, userGameRankScore[0].TotalScore)
				require.Equal(t, lastAward, userGameRankScore[0].LastAward.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserGameRankAndScore.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			userGameRankScore, err := client.GetUserGameRankAndScore(test.params)
			test.assert(t, userGameRankScore, err)
		})
	}
}

func TestGetUserPoints(tt *testing.T) {
	tests := []struct {
		name            string
		params          models.GetUserPointsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUserPoints
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, points *models.GetUserPoints, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserPointsParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, points *models.GetUserPoints, err error) {
				require.Nil(t, points)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserPoints.php?u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserPointsParameters{
				Username: "Test",
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
			assert: func(t *testing.T, points *models.GetUserPoints, err error) {
				require.Nil(t, points)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserPointsParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUserPoints{
				Points:         230,
				SoftcorePoints: 342,
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, points *models.GetUserPoints, err error) {
				require.NotNil(t, points)
				require.Equal(t, 230, points.Points)
				require.Equal(t, 342, points.SoftcorePoints)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserPoints.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			points, err := client.GetUserPoints(test.params)
			test.assert(t, points, err)
		})
	}
}

func TestGetUserProgress(tt *testing.T) {
	tests := []struct {
		name            string
		params          models.GetUserProgressParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage map[string]models.GetUserProgress
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, progress *map[string]models.GetUserProgress, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserProgressParameters{
				Username: "Test",
				GameIDs:  []int{1, 2, 5352},
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, resp *map[string]models.GetUserProgress, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserProgress.php?i=1%2C2%2C5352&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserProgressParameters{
				Username: "Test",
				GameIDs:  []int{1, 2, 5352},
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
			assert: func(t *testing.T, resp *map[string]models.GetUserProgress, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserProgressParameters{
				Username: "Test",
				GameIDs:  []int{1, 2, 5352},
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: map[string]models.GetUserProgress{
				"1": {
					NumPossibleAchievements: 36,
					PossibleScore:           305,
					NumAchieved:             13,
					ScoreAchieved:           100,
					NumAchievedHardcore:     13,
					ScoreAchievedHardcore:   100,
				},
				"2": {
					NumPossibleAchievements: 56,
					PossibleScore:           600,
					NumAchieved:             0,
					ScoreAchieved:           0,
					NumAchievedHardcore:     0,
					ScoreAchievedHardcore:   0,
				},
				"5352": {
					NumPossibleAchievements: 13,
					PossibleScore:           230,
					NumAchieved:             10,
					ScoreAchieved:           200,
					NumAchievedHardcore:     10,
					ScoreAchievedHardcore:   200,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *map[string]models.GetUserProgress, err error) {
				require.NotNil(t, resp)
				r := *resp
				// first element
				first, ok := r["1"]
				require.True(t, ok)
				require.Equal(t, 36, first.NumPossibleAchievements)
				require.Equal(t, 305, first.PossibleScore)
				require.Equal(t, 13, first.NumAchieved)
				require.Equal(t, 100, first.ScoreAchieved)
				require.Equal(t, 13, first.NumAchievedHardcore)
				require.Equal(t, 100, first.ScoreAchievedHardcore)

				// second element
				second, ok := r["2"]
				require.True(t, ok)
				require.Equal(t, 56, second.NumPossibleAchievements)
				require.Equal(t, 600, second.PossibleScore)
				require.Equal(t, 0, second.NumAchieved)
				require.Equal(t, 0, second.ScoreAchieved)
				require.Equal(t, 0, second.NumAchievedHardcore)
				require.Equal(t, 0, second.ScoreAchievedHardcore)

				// third element
				third, ok := r["5352"]
				require.True(t, ok)
				require.Equal(t, 13, third.NumPossibleAchievements)
				require.Equal(t, 230, third.PossibleScore)
				require.Equal(t, 10, third.NumAchieved)
				require.Equal(t, 200, third.ScoreAchieved)
				require.Equal(t, 10, third.NumAchievedHardcore)
				require.Equal(t, 200, third.ScoreAchievedHardcore)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserProgress.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			resp, err := client.GetUserProgress(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetUserRecentlyPlayedGames(tt *testing.T) {
	count := 10
	offset := 0
	lastPlayed, err := time.Parse(time.DateTime, "2024-05-07 08:48:54")
	require.NoError(tt, err)
	lastPlayed2, err := time.Parse(time.DateTime, "2024-09-19 10:08:09")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetUserRecentlyPlayedGamesParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetUserRecentlyPlayedGames
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp []models.GetUserRecentlyPlayedGames, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserRecentlyPlayedGamesParameters{
				Username: "Test",
				Count:    &count,
				Offset:   &offset,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, resp []models.GetUserRecentlyPlayedGames, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserRecentlyPlayedGames.php?c=10&o=0&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserRecentlyPlayedGamesParameters{
				Username: "Test",
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
			assert: func(t *testing.T, resp []models.GetUserRecentlyPlayedGames, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserRecentlyPlayedGamesParameters{
				Username: "Test",
				Count:    &count,
				Offset:   &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetUserRecentlyPlayedGames{
				{
					NumPossibleAchievements: 36,
					PossibleScore:           305,
					NumAchieved:             13,
					ScoreAchieved:           100,
					NumAchievedHardcore:     13,
					ScoreAchievedHardcore:   100,
					GameID:                  123,
					ConsoleID:               2,
					ConsoleName:             "Game Cube",
					Title:                   "Batman",
					ImageIcon:               "/img/something.png",
					ImageTitle:              "batman image",
					ImageIngame:             "/img/ingame.png",
					ImageBoxArt:             "/img/boxart.png",
					LastPlayed: models.DateTime{
						Time: lastPlayed,
					},
					AchievementsTotal: 16,
				},
				{
					NumPossibleAchievements: 66,
					PossibleScore:           355,
					NumAchieved:             100,
					ScoreAchieved:           1990,
					NumAchievedHardcore:     56,
					ScoreAchievedHardcore:   101,
					GameID:                  234,
					ConsoleID:               6,
					ConsoleName:             "PlayStation",
					Title:                   "Call Of Duty",
					ImageIcon:               "/img/duty.png",
					ImageTitle:              "cod image",
					ImageIngame:             "/img/cod_battlefield.png",
					ImageBoxArt:             "/img/cod_cover.png",
					LastPlayed: models.DateTime{
						Time: lastPlayed2,
					},
					AchievementsTotal: 42,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetUserRecentlyPlayedGames, err error) {
				require.NoError(t, err)
				require.Equal(t, 36, resp[0].NumPossibleAchievements)
				require.Equal(t, 305, resp[0].PossibleScore)
				require.Equal(t, 13, resp[0].NumAchieved)
				require.Equal(t, 100, resp[0].ScoreAchieved)
				require.Equal(t, 13, resp[0].NumAchievedHardcore)
				require.Equal(t, 100, resp[0].ScoreAchievedHardcore)
				require.Equal(t, 123, resp[0].GameID)
				require.Equal(t, 2, resp[0].ConsoleID)
				require.Equal(t, "Game Cube", resp[0].ConsoleName)
				require.Equal(t, "Batman", resp[0].Title)
				require.Equal(t, "/img/something.png", resp[0].ImageIcon)
				require.Equal(t, "batman image", resp[0].ImageTitle)
				require.Equal(t, "/img/ingame.png", resp[0].ImageIngame)
				require.Equal(t, "/img/boxart.png", resp[0].ImageBoxArt)
				require.Equal(t, lastPlayed, resp[0].LastPlayed.Time)
				require.Equal(t, 16, resp[0].AchievementsTotal)

				require.Equal(t, 66, resp[1].NumPossibleAchievements)
				require.Equal(t, 355, resp[1].PossibleScore)
				require.Equal(t, 100, resp[1].NumAchieved)
				require.Equal(t, 1990, resp[1].ScoreAchieved)
				require.Equal(t, 56, resp[1].NumAchievedHardcore)
				require.Equal(t, 101, resp[1].ScoreAchievedHardcore)
				require.Equal(t, 234, resp[1].GameID)
				require.Equal(t, 6, resp[1].ConsoleID)
				require.Equal(t, "PlayStation", resp[1].ConsoleName)
				require.Equal(t, "Call Of Duty", resp[1].Title)
				require.Equal(t, "/img/duty.png", resp[1].ImageIcon)
				require.Equal(t, "cod image", resp[1].ImageTitle)
				require.Equal(t, "/img/cod_battlefield.png", resp[1].ImageIngame)
				require.Equal(t, "/img/cod_cover.png", resp[1].ImageBoxArt)
				require.Equal(t, lastPlayed2, resp[1].LastPlayed.Time)
				require.Equal(t, 42, resp[1].AchievementsTotal)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserRecentlyPlayedGames.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			resp, err := client.GetUserRecentlyPlayedGames(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetUserSummary(tt *testing.T) {
	gameCount := 10
	achievementCount := 5
	rank := 130
	memberSince, err := time.Parse(time.DateTime, "2017-06-18 18:49:00")
	require.NoError(tt, err)
	lastPlayed, err := time.Parse(time.DateTime, "2024-11-17 04:00:35")
	require.NoError(tt, err)
	releaseDate, err := time.Parse(models.LongMonthDateFormat, "September 27, 2011")
	require.NoError(tt, err)
	cheevoType := "progression"
	tests := []struct {
		name            string
		params          models.GetUserSummaryParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUserSummary
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetUserSummary, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserSummaryParameters{
				Username:          "Test",
				GamesCount:        &gameCount,
				AchievementsCount: &achievementCount,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, resp *models.GetUserSummary, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserSummary.php?a=5&g=10&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserSummaryParameters{
				Username:          "Test",
				GamesCount:        &gameCount,
				AchievementsCount: &achievementCount,
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
			assert: func(t *testing.T, resp *models.GetUserSummary, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserSummaryParameters{
				Username:          "Test",
				GamesCount:        &gameCount,
				AchievementsCount: &achievementCount,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUserSummary{
				User:                "Jamiras",
				RichPresenceMsg:     "In Chapter 3: The Paladin Clan • 3h13 • Level 5 • 3 • 4918 Rings • 12/83",
				LastGameID:          9404,
				ContribCount:        179822,
				ContribYield:        1208742,
				TotalPoints:         116649,
				TotalSoftcorePoints: 1350,
				TotalTruePoints:     322472,
				Permissions:         4,
				Untracked:           0,
				ID:                  43495,
				UserWallActive:      0,
				Motto:               "",
				RecentlyPlayedCount: 1,
				UserPic:             "/UserPic/jamiras.png",
				TotalRanked:         78332,
				Status:              "Offline",
				Rank:                &rank,
				MemberSince: models.DateTime{
					Time: memberSince,
				},
				LastActivity: models.GetUserSummaryLastActivity{
					ID:   0,
					User: "jamiras",
				},
				RecentlyPlayed: []models.GetUserSummaryRecentlyPlayed{
					{
						GameID:      9404,
						ConsoleID:   18,
						ConsoleName: "Nintendo DS",
						Title:       "Solatorobo: Red the Hunter",
						ImageIcon:   "/Images/088320.png",
						ImageTitle:  "/Images/073286.png",
						ImageIngame: "/Images/073287.png",
						ImageBoxArt: "/Images/028653.png",
						LastPlayed: models.DateTime{
							Time: lastPlayed,
						},
						AchievementsTotal: 133,
					},
				},
				Awarded: map[string]models.GetUserSummaryAwarded{
					"9404": {
						NumPossibleAchievements: 133,
						PossibleScore:           935,
						NumAchieved:             16,
						ScoreAchieved:           95,
						NumAchievedHardcore:     16,
						ScoreAchievedHardcore:   95,
					},
				},
				RecentAchievements: map[string]map[string]models.GetUserSummaryRecentAchievements{
					"9404": {
						"328833": {
							ID:          328833,
							GameID:      9404,
							GameTitle:   "Solatorobo: Red the Hunter",
							Title:       "Chapter 3: The Paladin Clan",
							Description: "Complete the Chapter 3",
							Points:      10,
							Type:        &cheevoType,
							BadgeName:   "368292",
							IsAwarded:   "1",
							DateAwarded: models.DateTime{
								Time: lastPlayed,
							},
							HardcoreAchieved: 1,
						},
					},
				},
				LastGame: models.GetUserSummaryLastGame{
					ID:           9404,
					Title:        "Solatorobo: Red the Hunter",
					ConsoleID:    18,
					ConsoleName:  "Nintendo DS",
					ForumTopicID: 21569,
					Flags:        0,
					ImageIcon:    "/Images/088320.png",
					ImageTitle:   "/Images/073286.png",
					ImageIngame:  "/Images/073287.png",
					ImageBoxArt:  "/Images/028653.png",
					Publisher:    "XSEED Games",
					Developer:    "CyberConnect2 | CyberConnect",
					Genre:        "Action RPG",
					Released: models.LongMonthDate{
						Time: releaseDate,
					},
					IsFinal: 0,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetUserSummary, err error) {
				require.NoError(t, err)
				require.Equal(t, "Jamiras", resp.User)
				require.Equal(t, "In Chapter 3: The Paladin Clan • 3h13 • Level 5 • 3 • 4918 Rings • 12/83", resp.RichPresenceMsg)
				require.Equal(t, 9404, resp.LastGameID)
				require.Equal(t, 179822, resp.ContribCount)
				require.Equal(t, 1208742, resp.ContribYield)
				require.Equal(t, 116649, resp.TotalPoints)
				require.Equal(t, 1350, resp.TotalSoftcorePoints)
				require.Equal(t, 322472, resp.TotalTruePoints)
				require.Equal(t, 4, resp.Permissions)
				require.Equal(t, 0, resp.Untracked)
				require.Equal(t, 43495, resp.ID)
				require.Equal(t, 0, resp.UserWallActive)
				require.Equal(t, "", resp.Motto)
				require.Equal(t, 1, resp.RecentlyPlayedCount)
				require.Equal(t, "/UserPic/jamiras.png", resp.UserPic)
				require.Equal(t, 78332, resp.TotalRanked)
				require.Equal(t, "Offline", resp.Status)
				require.NotNil(t, resp.Rank)
				require.Equal(t, 130, *resp.Rank)
				require.Equal(t, memberSince, resp.MemberSince.Time)
				require.Equal(t, 0, resp.LastActivity.ID)
				require.Equal(t, "jamiras", resp.LastActivity.User)
				require.Len(t, resp.RecentlyPlayed, 1)
				require.Equal(t, 9404, resp.RecentlyPlayed[0].GameID)
				require.Equal(t, 18, resp.RecentlyPlayed[0].ConsoleID)
				require.Equal(t, "Nintendo DS", resp.RecentlyPlayed[0].ConsoleName)
				require.Equal(t, "Solatorobo: Red the Hunter", resp.RecentlyPlayed[0].Title)
				require.Equal(t, "/Images/088320.png", resp.RecentlyPlayed[0].ImageIcon)
				require.Equal(t, "/Images/073286.png", resp.RecentlyPlayed[0].ImageTitle)
				require.Equal(t, "/Images/073287.png", resp.RecentlyPlayed[0].ImageIngame)
				require.Equal(t, "/Images/028653.png", resp.RecentlyPlayed[0].ImageBoxArt)
				require.Equal(t, 133, resp.RecentlyPlayed[0].AchievementsTotal)
				require.Len(t, resp.Awarded, 1)
				award, ok := resp.Awarded["9404"]
				require.True(t, ok)
				require.Equal(t, 133, award.NumPossibleAchievements)
				require.Equal(t, 935, award.PossibleScore)
				require.Equal(t, 16, award.NumAchieved)
				require.Equal(t, 95, award.ScoreAchieved)
				require.Equal(t, 16, award.NumAchievedHardcore)
				require.Equal(t, 95, award.ScoreAchievedHardcore)
				require.Len(t, resp.RecentAchievements, 1)
				recent, ok := resp.RecentAchievements["9404"]
				require.True(t, ok)
				require.Len(t, recent, 1)
				cheevo, ok := recent["328833"]
				require.True(t, ok)
				require.Equal(t, 328833, cheevo.ID)
				require.Equal(t, 9404, cheevo.GameID)
				require.Equal(t, "Solatorobo: Red the Hunter", cheevo.GameTitle)
				require.Equal(t, "Chapter 3: The Paladin Clan", cheevo.Title)
				require.Equal(t, "Complete the Chapter 3", cheevo.Description)
				require.Equal(t, 10, cheevo.Points)
				require.Equal(t, "368292", cheevo.BadgeName)
				require.Equal(t, "1", cheevo.IsAwarded)
				require.Equal(t, 1, cheevo.HardcoreAchieved)
				require.NotNil(t, cheevo.Type)
				require.Equal(t, cheevoType, *cheevo.Type)
				require.Equal(t, lastPlayed, cheevo.DateAwarded.Time)
				require.Equal(t, 9404, resp.LastGame.ID)
				require.Equal(t, "Solatorobo: Red the Hunter", resp.LastGame.Title)
				require.Equal(t, 18, resp.LastGame.ConsoleID)
				require.Equal(t, "Nintendo DS", resp.LastGame.ConsoleName)
				require.Equal(t, 21569, resp.LastGame.ForumTopicID)
				require.Equal(t, 0, resp.LastGame.Flags)
				require.Equal(t, "/Images/088320.png", resp.LastGame.ImageIcon)
				require.Equal(t, "/Images/073286.png", resp.LastGame.ImageTitle)
				require.Equal(t, "/Images/073287.png", resp.LastGame.ImageIngame)
				require.Equal(t, "/Images/028653.png", resp.LastGame.ImageBoxArt)
				require.Equal(t, "XSEED Games", resp.LastGame.Publisher)
				require.Equal(t, "CyberConnect2 | CyberConnect", resp.LastGame.Developer)
				require.Equal(t, "Action RPG", resp.LastGame.Genre)
				require.Equal(t, releaseDate, resp.LastGame.Released.Time)
				require.Equal(t, 0, resp.LastGame.IsFinal)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserSummary.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			resp, err := client.GetUserSummary(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetUserCompletedGames(tt *testing.T) {
	tests := []struct {
		name            string
		params          models.GetUserCompletedGamesParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetUserCompletedGames
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp []models.GetUserCompletedGames, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserCompletedGamesParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, resp []models.GetUserCompletedGames, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserCompletedGames.php?u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserCompletedGamesParameters{
				Username: "Test",
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
			assert: func(t *testing.T, resp []models.GetUserCompletedGames, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserCompletedGamesParameters{
				Username: "Test",
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetUserCompletedGames{
				{
					GameID:       24941,
					Title:        "Dragon Quest IV: Chapters of the Chosen [Subset - Plentiful Plunder]",
					ImageIcon:    "/Images/075762.png",
					ConsoleID:    18,
					ConsoleName:  "Nintendo DS",
					MaxPossible:  202,
					NumAwarded:   202,
					PctWon:       "1.0000",
					HardcoreMode: "1",
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetUserCompletedGames, err error) {
				require.NoError(t, err)
				require.Len(t, resp, 1)
				require.Equal(t, 24941, resp[0].GameID)
				require.Equal(t, "Dragon Quest IV: Chapters of the Chosen [Subset - Plentiful Plunder]", resp[0].Title)
				require.Equal(t, "/Images/075762.png", resp[0].ImageIcon)
				require.Equal(t, 18, resp[0].ConsoleID)
				require.Equal(t, "Nintendo DS", resp[0].ConsoleName)
				require.Equal(t, 202, resp[0].MaxPossible)
				require.Equal(t, 202, resp[0].NumAwarded)
				require.Equal(t, "1.0000", resp[0].PctWon)
				require.Equal(t, "1", resp[0].HardcoreMode)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserCompletedGames.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			resp, err := client.GetUserCompletedGames(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetUserWantToPlayList(tt *testing.T) {
	count := 10
	offset := 23
	tests := []struct {
		name            string
		params          models.GetUserWantToPlayListParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUserWantToPlayList
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetUserWantToPlayList, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserWantToPlayListParameters{
				Username: "Test",
				Count:    &count,
				Offset:   &offset,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, resp *models.GetUserWantToPlayList, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserWantToPlayList.php?c=10&o=23&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserWantToPlayListParameters{
				Username: "Test",
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
			assert: func(t *testing.T, resp *models.GetUserWantToPlayList, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserWantToPlayListParameters{
				Username: "Test",
				Count:    &count,
				Offset:   &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUserWantToPlayList{
				Count: 1,
				Total: 1,
				Results: []models.GetUserWantToPlayListResult{
					{
						ID:                    189,
						Title:                 "Super Mario Galaxy",
						ConsoleID:             19,
						ConsoleName:           "Wii",
						ImageIcon:             "/Images/079076.png",
						PointsTotal:           0,
						AchievementsPublished: 0,
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetUserWantToPlayList, err error) {
				require.NoError(t, err)
				require.Equal(t, 1, resp.Count)
				require.Equal(t, 1, resp.Total)
				require.Len(t, resp.Results, 1)
				require.Equal(t, 189, resp.Results[0].ID)
				require.Equal(t, "Super Mario Galaxy", resp.Results[0].Title)
				require.Equal(t, 19, resp.Results[0].ConsoleID)
				require.Equal(t, "Wii", resp.Results[0].ConsoleName)
				require.Equal(t, "/Images/079076.png", resp.Results[0].ImageIcon)
				require.Equal(t, 0, resp.Results[0].PointsTotal)
				require.Equal(t, 0, resp.Results[0].AchievementsPublished)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserWantToPlayList.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			resp, err := client.GetUserWantToPlayList(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetUsersIFollow(tt *testing.T) {
	count := 10
	offset := 23
	tests := []struct {
		name            string
		params          models.GetUsersIFollowParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUsersIFollow
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetUsersIFollow, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUsersIFollowParameters{
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, resp *models.GetUsersIFollow, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUsersIFollow.php?c=10&o=23&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUsersIFollowParameters{
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
			assert: func(t *testing.T, resp *models.GetUsersIFollow, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUsersIFollowParameters{
				Count:  &count,
				Offset: &offset,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUsersIFollow{
				Count: 20,
				Total: 120,
				Results: []models.GetUsersIFollowResult{
					{
						User:           "zuliman92",
						Points:         1882,
						PointsSoftcore: 258,
						IsFollowingMe:  true,
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetUsersIFollow, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 20, resp.Count)
				require.Equal(t, 120, resp.Total)
				require.Len(t, resp.Results, 1)
				require.Equal(t, "zuliman92", resp.Results[0].User)
				require.Equal(t, 1882, resp.Results[0].Points)
				require.Equal(t, 258, resp.Results[0].PointsSoftcore)
				require.True(t, resp.Results[0].IsFollowingMe)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUsersIFollow.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			resp, err := client.GetUsersIFollow(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetUserSetRequests(tt *testing.T) {
	all := true
	tests := []struct {
		name            string
		params          models.GetUserSetRequestsParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetUserSetRequests
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetUserSetRequests, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetUserSetRequestsParameters{
				Username: "Test",
				All:      &all,
			},
			modifyURL: func(url string) string {
				return ""
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
			assert: func(t *testing.T, resp *models.GetUserSetRequests, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserSetRequests.php?t=1&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetUserSetRequestsParameters{
				Username: "Test",
				All:      &all,
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
			assert: func(t *testing.T, resp *models.GetUserSetRequests, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetUserSetRequestsParameters{
				Username: "Test",
				All:      &all,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetUserSetRequests{
				TotalRequests: 5,
				PointsForNext: 5000,
				RequestedSets: []models.GetUserSetRequestsRequestedSet{
					{
						GameID:      1,
						Title:       "Sonic the Hedgehog",
						ConsoleID:   1,
						ConsoleName: "Genesis/Mega Drive",
						ImageIcon:   "/Images/085573.png",
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetUserSetRequests, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 5, resp.TotalRequests)
				require.Equal(t, 5000, resp.PointsForNext)
				require.Len(t, resp.RequestedSets, 1)
				require.Equal(t, 1, resp.RequestedSets[0].GameID)
				require.Equal(t, "Sonic the Hedgehog", resp.RequestedSets[0].Title)
				require.Equal(t, 1, resp.RequestedSets[0].ConsoleID)
				require.Equal(t, "Genesis/Mega Drive", resp.RequestedSets[0].ConsoleName)
				require.Equal(t, "/Images/085573.png", resp.RequestedSets[0].ImageIcon)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetUserSetRequests.php"
				if r.URL.Path != expectedPath {
					t.Errorf("Expected to request '%s', got: %s", expectedPath, r.URL.Path)
				}
				w.WriteHeader(test.responseCode)
				responseMessage, err := json.Marshal(test.responseMessage)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(responseMessage, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()
			client := retroachievements.New(test.modifyURL(server.URL), "go-retroachievements/v0.0.0", "some_secret")
			resp, err := client.GetUserSetRequests(test.params)
			test.assert(t, resp, err)
		})
	}
}
