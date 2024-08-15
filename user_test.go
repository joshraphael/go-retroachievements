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
		username        string
		modifyURL       func(url string) string
		responseCode    int
		responseProfile models.Profile
		responseError   models.ErrorResponse
		response        func(profileBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, profile *models.Profile, err error)
	}{
		{
			name:     "fail to call endpoint",
			username: "Test",
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
			response: func(profileBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, profile *models.Profile, err error) {
				require.Nil(t, profile)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserProfile.php?u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name:     "error response",
			username: "Test",
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
			response: func(profileBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, profile *models.Profile, err error) {
				require.Nil(t, profile)
				require.EqualError(t, err, "parsing response object: error responses: [401] Not Authorized")
			},
		},
		{
			name:     "success",
			username: "Test",
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseProfile: models.Profile{
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
			response: func(profileBytes []byte, errorBytes []byte) []byte {
				return profileBytes
			},
			assert: func(t *testing.T, profile *models.Profile, err error) {
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
				profileBytes, err := json.Marshal(test.responseProfile)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(profileBytes, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
			profile, err := client.GetUserProfile(test.username)
			test.assert(t, profile, err)
		})
	}
}

func TestGetUserRecentAchievements(tt *testing.T) {
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(tt, err)
	tests := []struct {
		name                 string
		username             string
		lookbackMinutes      int
		modifyURL            func(url string) string
		responseCode         int
		responseAchievements []models.Achievement
		responseError        models.ErrorResponse
		response             func(achievementsBytes []byte, errorBytes []byte) []byte
		assert               func(t *testing.T, achievements []models.Achievement, err error)
	}{
		{
			name:            "fail to call endpoint",
			username:        "Test",
			lookbackMinutes: 60,
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
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetUserRecentAchievements.php?m=60&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name:            "error response",
			username:        "Test",
			lookbackMinutes: 60,
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
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "parsing response list: error responses: [401] Not Authorized")
			},
		},
		{
			name:            "error response",
			username:        "Test",
			lookbackMinutes: 60,
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseAchievements: []models.Achievement{
				{
					Date: models.DateTime{
						Time: now,
					},
					HardcoreMode:  1,
					AchievementID: 34425,
					Title:         "Beat Level 1",
					Description:   "Finish level 1",
					BadgeName:     "840124",
					Points:        10,
					TrueRatio:     234,
					Type:          "win_condition",
					Author:        "jamiras",
					GameTitle:     "Final Fantasy XXXXIIII",
					GameIcon:      "/Images/056340.png",
					GameID:        34897,
					ConsoleName:   "SNES",
					BadgeURL:      "/Badge/840124.png",
					GameURL:       "/game/34897",
				},
			},
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return achievementsBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
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
				achievementsBytes, err := json.Marshal(test.responseAchievements)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(achievementsBytes, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
			achievements, err := client.GetUserRecentAchievements(test.username, test.lookbackMinutes)
			test.assert(t, achievements, err)
		})
	}
}

func TestGetAchievementsEarnedBetween(tt *testing.T) {
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(tt, err)
	later := now.Add(10 * time.Minute)
	tests := []struct {
		name                 string
		username             string
		fromTime             time.Time
		toTime               time.Time
		modifyURL            func(url string) string
		responseCode         int
		responseAchievements []models.Achievement
		responseError        models.ErrorResponse
		response             func(achievementsBytes []byte, errorBytes []byte) []byte
		assert               func(t *testing.T, achievements []models.Achievement, err error)
	}{
		{
			name:     "fail to call endpoint",
			username: "Test",
			fromTime: now,
			toTime:   later,
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
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetAchievementsEarnedBetween.php?f=1709400423&t=1709401023&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name:     "error response",
			username: "Test",
			fromTime: now,
			toTime:   later,
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
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "parsing response list: error responses: [401] Not Authorized")
			},
		},
		{
			name:     "error response",
			username: "Test",
			fromTime: now,
			toTime:   later,
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseAchievements: []models.Achievement{
				{
					Date: models.DateTime{
						Time: now,
					},
					HardcoreMode:  1,
					AchievementID: 34425,
					Title:         "Beat Level 1",
					Description:   "Finish level 1",
					BadgeName:     "840124",
					Points:        10,
					TrueRatio:     234,
					Type:          "win_condition",
					Author:        "jamiras",
					GameTitle:     "Final Fantasy XXXXIIII",
					GameIcon:      "/Images/056340.png",
					GameID:        34897,
					ConsoleName:   "SNES",
					BadgeURL:      "/Badge/840124.png",
					GameURL:       "/game/34897",
				},
			},
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return achievementsBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
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
				achievementsBytes, err := json.Marshal(test.responseAchievements)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(achievementsBytes, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
			achievements, err := client.GetAchievementsEarnedBetween(test.username, test.fromTime, test.toTime)
			test.assert(t, achievements, err)
		})
	}
}

func TestGetAchievementsEarnedOnDay(tt *testing.T) {
	now, err := time.Parse(time.DateTime, "2024-03-02 17:27:03")
	require.NoError(tt, err)
	tests := []struct {
		name                 string
		username             string
		date                 time.Time
		modifyURL            func(url string) string
		responseCode         int
		responseAchievements []models.Achievement
		responseError        models.ErrorResponse
		response             func(achievementsBytes []byte, errorBytes []byte) []byte
		assert               func(t *testing.T, achievements []models.Achievement, err error)
	}{
		{
			name:     "fail to call endpoint",
			username: "Test",
			date:     now,
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
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetAchievementsEarnedOnDay.php?d=2024-03-02&u=Test&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name:     "error response",
			username: "Test",
			date:     now,
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
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
				require.Nil(t, achievements)
				require.EqualError(t, err, "parsing response list: error responses: [401] Not Authorized")
			},
		},
		{
			name:     "error response",
			username: "Test",
			date:     now,
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseAchievements: []models.Achievement{
				{
					Date: models.DateTime{
						Time: now,
					},
					HardcoreMode:  1,
					AchievementID: 34425,
					Title:         "Beat Level 1",
					Description:   "Finish level 1",
					BadgeName:     "840124",
					Points:        10,
					TrueRatio:     234,
					Type:          "win_condition",
					Author:        "jamiras",
					GameTitle:     "Final Fantasy XXXXIIII",
					GameIcon:      "/Images/056340.png",
					GameID:        34897,
					ConsoleName:   "SNES",
					BadgeURL:      "/Badge/840124.png",
					GameURL:       "/game/34897",
				},
			},
			response: func(achievementsBytes []byte, errorBytes []byte) []byte {
				return achievementsBytes
			},
			assert: func(t *testing.T, achievements []models.Achievement, err error) {
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
				achievementsBytes, err := json.Marshal(test.responseAchievements)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(achievementsBytes, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
			achievements, err := client.GetAchievementsEarnedOnDay(test.username, test.date)
			test.assert(t, achievements, err)
		})
	}
}
