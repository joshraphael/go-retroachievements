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

func TestGetGame(tt *testing.T) {
	forumTopicId := 16654
	flags := 0
	released, err := time.Parse(time.DateOnly, "2001-06-18")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetGameParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetGame
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetGame, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetGameParameters{
				GameID: 2991,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetGame{
				Title:        "Twisted Metal: Black",
				ConsoleID:    21,
				ForumTopicID: &forumTopicId,
				Flags:        &flags,
				ImageIcon:    "/Images/057992.png",
				ImageTitle:   "/Images/056152.png",
				ImageIngame:  "/Images/056151.png",
				ImageBoxArt:  "/Images/050832.png",
				Publisher:    "Sony Computer Entertainment",
				Developer:    "Incognito Entertainment",
				Genre:        "Vehicular Combat",
				Released: &models.DateOnly{
					Time: released,
				},
				GameTitle:   "Twisted Metal: Black",
				ConsoleName: "Playstation 2",
				Console:     "Playstation 2",
				GameIcon:    "/Images/057992.png",
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGame, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetGame.php?i=2991&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetGameParameters{
				GameID: 2991,
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
			assert: func(t *testing.T, resp *models.GetGame, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetGameParameters{
				GameID: 2991,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetGame{
				Title:        "Twisted Metal: Black",
				ConsoleID:    21,
				ForumTopicID: &forumTopicId,
				Flags:        &flags,
				ImageIcon:    "/Images/057992.png",
				ImageTitle:   "/Images/056152.png",
				ImageIngame:  "/Images/056151.png",
				ImageBoxArt:  "/Images/050832.png",
				Publisher:    "Sony Computer Entertainment",
				Developer:    "Incognito Entertainment",
				Genre:        "Vehicular Combat",
				Released: &models.DateOnly{
					Time: released,
				},
				GameTitle:   "Twisted Metal: Black",
				ConsoleName: "Playstation 2",
				Console:     "Playstation 2",
				GameIcon:    "/Images/057992.png",
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGame, err error) {
				require.NotNil(t, resp)
				require.Equal(t, resp.Title, "Twisted Metal: Black")
				require.Equal(t, "Twisted Metal: Black", resp.GameTitle)
				require.Equal(t, 21, resp.ConsoleID)
				require.Equal(t, "Playstation 2", resp.ConsoleName)
				require.Equal(t, "Playstation 2", resp.Console)
				require.Equal(t, 16654, *resp.ForumTopicID)
				require.Equal(t, 0, *resp.Flags)
				require.Equal(t, "/Images/057992.png", resp.GameIcon)
				require.Equal(t, "/Images/057992.png", resp.ImageIcon)
				require.Equal(t, "/Images/056152.png", resp.ImageTitle)
				require.Equal(t, "/Images/056151.png", resp.ImageIngame)
				require.Equal(t, "/Images/050832.png", resp.ImageBoxArt)
				require.Equal(t, "Sony Computer Entertainment", resp.Publisher)
				require.Equal(t, "Incognito Entertainment", resp.Developer)
				require.Equal(t, "Vehicular Combat", resp.Genre)
				require.Equal(t, released, resp.Released.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetGame.php"
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
			client := retroachievements.New(retroachievements.ClientConfig{
				Host:      test.modifyURL(server.URL),
				UserAgent: "go-retroachievements/v0.0.0",
				APISecret: "some_secret",
			})
			resp, err := client.GetGame(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetGameExtended(tt *testing.T) {
	forumTopicId := 16654
	flags := 0
	released, err := time.Parse(time.DateOnly, "2001-06-18")
	require.NoError(tt, err)
	updated, err := time.Parse(time.RFC3339Nano, "2024-08-15T11:46:06.000000Z")
	require.NoError(tt, err)
	modified, err := time.Parse(time.DateTime, "2022-10-25 17:00:49")
	require.NoError(tt, err)
	created, err := time.Parse(time.DateTime, "2022-09-28 00:36:26")
	require.NoError(tt, err)
	unofficial := true
	tests := []struct {
		name            string
		params          models.GetGameExtentedParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetGameExtented
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetGameExtented, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetGameExtentedParameters{
				GameID:     2991,
				Unofficial: &unofficial,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetGameExtented{
				Title:        "Twisted Metal: Black",
				ConsoleID:    21,
				ForumTopicID: &forumTopicId,
				Flags:        &flags,
				ImageIcon:    "/Images/057992.png",
				ImageTitle:   "/Images/056152.png",
				ImageIngame:  "/Images/056151.png",
				ImageBoxArt:  "/Images/050832.png",
				Publisher:    "Sony Computer Entertainment",
				Developer:    "Incognito Entertainment",
				Genre:        "Vehicular Combat",
				Released: &models.DateOnly{
					Time: released,
				},
				ID:                 2991,
				IsFinal:            0,
				RichPresencePatch:  "e7a5e12072a6c976a1146756726fdd8c",
				Updated:            &updated,
				ConsoleName:        "PlayStation 2",
				NumDistinctPlayers: 1287,
				NumAchievements:    93,
				Achievements: map[int]models.GetGameExtentedAchievement{
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
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGameExtented, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetGameExtended.php?f=5&i=2991&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetGameExtentedParameters{
				GameID: 2991,
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
			assert: func(t *testing.T, resp *models.GetGameExtented, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetGameExtentedParameters{
				GameID: 2991,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetGameExtented{
				Title:        "Twisted Metal: Black",
				ConsoleID:    21,
				ForumTopicID: &forumTopicId,
				Flags:        &flags,
				ImageIcon:    "/Images/057992.png",
				ImageTitle:   "/Images/056152.png",
				ImageIngame:  "/Images/056151.png",
				ImageBoxArt:  "/Images/050832.png",
				Publisher:    "Sony Computer Entertainment",
				Developer:    "Incognito Entertainment",
				Genre:        "Vehicular Combat",
				Released: &models.DateOnly{
					Time: released,
				},
				ID:                 2991,
				IsFinal:            0,
				RichPresencePatch:  "e7a5e12072a6c976a1146756726fdd8c",
				Updated:            &updated,
				ConsoleName:        "PlayStation 2",
				NumDistinctPlayers: 1287,
				NumAchievements:    93,
				Achievements: map[int]models.GetGameExtentedAchievement{
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
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGameExtented, err error) {
				require.NotNil(t, resp)
				require.Equal(t, "Twisted Metal: Black", resp.Title)
				require.Equal(t, 21, resp.ConsoleID)
				require.Equal(t, "PlayStation 2", resp.ConsoleName)
				require.Equal(t, 16654, *resp.ForumTopicID)
				require.Equal(t, 0, *resp.Flags)
				require.Equal(t, "/Images/057992.png", resp.ImageIcon)
				require.Equal(t, "/Images/056152.png", resp.ImageTitle)
				require.Equal(t, "/Images/056151.png", resp.ImageIngame)
				require.Equal(t, "/Images/050832.png", resp.ImageBoxArt)
				require.Equal(t, "Sony Computer Entertainment", resp.Publisher)
				require.Equal(t, "Incognito Entertainment", resp.Developer)
				require.Equal(t, "Vehicular Combat", resp.Genre)
				require.Equal(t, released, resp.Released.Time)
				require.Equal(t, 2991, resp.ID)
				require.Equal(t, 0, resp.IsFinal)
				require.Equal(t, "e7a5e12072a6c976a1146756726fdd8c", resp.RichPresencePatch)
				require.Equal(t, updated, *resp.Updated)
				require.Equal(t, "PlayStation 2", resp.ConsoleName)
				require.Equal(t, 1287, resp.NumDistinctPlayers)
				require.Equal(t, 93, resp.NumAchievements)
				require.Len(t, resp.Achievements, 1)
				achievement, ok := resp.Achievements[252117]
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
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetGameExtended.php"
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
			resp, err := client.GetGameExtended(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetGameHashes(tt *testing.T) {
	patchUrl := "https://github.com/RetroAchievements/RAPatches/raw/main/MD/Translation/Russian/1-Sonic1-Russian.zip"
	tests := []struct {
		name            string
		params          models.GetGameHashesParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetGameHashes
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetGameHashes, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetGameHashesParameters{
				GameID: 2991,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGameHashes, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetGameHashes.php?i=2991&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetGameHashesParameters{
				GameID: 2991,
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
			assert: func(t *testing.T, resp *models.GetGameHashes, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetGameHashesParameters{
				GameID: 1,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetGameHashes{
				Results: []models.GetGameHashesResult{
					{
						Name: "Sonic The Hedgehog (USA, Europe) (Ru) (NewGame).md",
						MD5:  "1b1d9ac862c387367e904036114c4825",
						Labels: []string{
							"nointro",
							"rapatches",
						},
						PatchUrl: &patchUrl,
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetGameHashes, err error) {
				require.NotNil(t, resp)
				require.Len(t, resp.Results, 1)
				require.Equal(t, "Sonic The Hedgehog (USA, Europe) (Ru) (NewGame).md", resp.Results[0].Name)
				require.Equal(t, "1b1d9ac862c387367e904036114c4825", resp.Results[0].MD5)
				require.Len(t, resp.Results[0].Labels, 2)
				require.Equal(t, resp.Results[0].Labels[0], "nointro")
				require.Equal(t, resp.Results[0].Labels[1], "rapatches")
				require.NotNil(t, resp.Results[0].PatchUrl)
				require.Equal(t, "https://github.com/RetroAchievements/RAPatches/raw/main/MD/Translation/Russian/1-Sonic1-Russian.zip", *resp.Results[0].PatchUrl)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetGameHashes.php"
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
			resp, err := client.GetGameHashes(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetAchievementCount(tt *testing.T) {
	tests := []struct {
		name            string
		params          models.GetAchievementCountParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetAchievementCount
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetAchievementCount, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetAchievementCountParameters{
				GameID: 14402,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementCount, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetAchievementCount.php?i=14402&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetAchievementCountParameters{
				GameID: 14402,
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
			assert: func(t *testing.T, resp *models.GetAchievementCount, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetAchievementCountParameters{
				GameID: 14402,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetAchievementCount{
				GameID: 14402,
				AchievementIDs: []int{
					79434,
					79435,
					79436,
					79437,
					79438,
					79439,
					79440,
					79441,
					79442,
					79443,
					79444,
					79445,
					325413,
					325414,
					325415,
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementCount, err error) {
				require.NotNil(t, resp)
				require.Equal(t, 14402, resp.GameID)
				require.Len(t, resp.AchievementIDs, 15)
				require.Equal(t, 79434, resp.AchievementIDs[0])
				require.Equal(t, 79435, resp.AchievementIDs[1])
				require.Equal(t, 79436, resp.AchievementIDs[2])
				require.Equal(t, 79437, resp.AchievementIDs[3])
				require.Equal(t, 79438, resp.AchievementIDs[4])
				require.Equal(t, 79439, resp.AchievementIDs[5])
				require.Equal(t, 79440, resp.AchievementIDs[6])
				require.Equal(t, 79441, resp.AchievementIDs[7])
				require.Equal(t, 79442, resp.AchievementIDs[8])
				require.Equal(t, 79443, resp.AchievementIDs[9])
				require.Equal(t, 79444, resp.AchievementIDs[10])
				require.Equal(t, 79445, resp.AchievementIDs[11])
				require.Equal(t, 325413, resp.AchievementIDs[12])
				require.Equal(t, 325414, resp.AchievementIDs[13])
				require.Equal(t, 325415, resp.AchievementIDs[14])
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetAchievementCount.php"
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
			resp, err := client.GetAchievementCount(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetAchievementDistribution(tt *testing.T) {
	hardcore := true
	unofficial := true
	official := false
	tests := []struct {
		name            string
		params          models.GetAchievementDistributionParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetAchievementDistribution
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp *models.GetAchievementDistribution, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetAchievementDistributionParameters{
				GameID:     14402,
				Hardcore:   &hardcore,
				Unofficial: &unofficial,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementDistribution, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetAchievementDistribution.php?f=5&h=1&i=14402&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetAchievementDistributionParameters{
				GameID:     14402,
				Hardcore:   &hardcore,
				Unofficial: &unofficial,
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
			assert: func(t *testing.T, resp *models.GetAchievementDistribution, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response object: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetAchievementDistributionParameters{
				GameID:     14402,
				Hardcore:   &hardcore,
				Unofficial: &official,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: models.GetAchievementDistribution{
				"1":  105,
				"2":  28,
				"3":  33,
				"4":  30,
				"5":  20,
				"6":  15,
				"7":  4,
				"8":  29,
				"9":  8,
				"10": 4,
				"11": 1,
				"12": 0,
				"13": 0,
				"14": 0,
				"15": 3,
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp *models.GetAchievementDistribution, err error) {
				require.NotNil(t, resp)
				r := *resp
				val, ok := r["1"]
				require.True(t, ok)
				require.Equal(t, 105, val)
				val, ok = r["2"]
				require.True(t, ok)
				require.Equal(t, 28, val)
				val, ok = r["3"]
				require.True(t, ok)
				require.Equal(t, 33, val)
				val, ok = r["4"]
				require.True(t, ok)
				require.Equal(t, 30, val)
				val, ok = r["5"]
				require.True(t, ok)
				require.Equal(t, 20, val)
				val, ok = r["6"]
				require.True(t, ok)
				require.Equal(t, 15, val)
				val, ok = r["7"]
				require.True(t, ok)
				require.Equal(t, 4, val)
				val, ok = r["8"]
				require.True(t, ok)
				require.Equal(t, 29, val)
				val, ok = r["9"]
				require.True(t, ok)
				require.Equal(t, 8, val)
				val, ok = r["10"]
				require.True(t, ok)
				require.Equal(t, 4, val)
				val, ok = r["11"]
				require.True(t, ok)
				require.Equal(t, 1, val)
				val, ok = r["12"]
				require.True(t, ok)
				require.Equal(t, 0, val)
				val, ok = r["13"]
				require.True(t, ok)
				require.Equal(t, 0, val)
				val, ok = r["14"]
				require.True(t, ok)
				require.Equal(t, 0, val)
				val, ok = r["15"]
				require.True(t, ok)
				require.Equal(t, 3, val)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetAchievementDistribution.php"
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
			resp, err := client.GetAchievementDistribution(test.params)
			test.assert(t, resp, err)
		})
	}
}

func TestGetGameRankAndScore(tt *testing.T) {
	latest := true
	all := false
	lastAwarded1, err := time.Parse(time.DateTime, "2022-01-28 21:18:15")
	require.NoError(tt, err)
	lastAwarded2, err := time.Parse(time.DateTime, "2022-01-29 04:19:34")
	require.NoError(tt, err)
	tests := []struct {
		name            string
		params          models.GetGameRankAndScoreParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage []models.GetGameRankAndScore
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, resp []models.GetGameRankAndScore, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetGameRankAndScoreParameters{
				GameID:        14402,
				LatestMasters: &latest,
			},
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetGameRankAndScore, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetGameRankAndScore.php?g=14402&t=1&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			params: models.GetGameRankAndScoreParameters{
				GameID:        14402,
				LatestMasters: &all,
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
			assert: func(t *testing.T, resp []models.GetGameRankAndScore, err error) {
				require.Nil(t, resp)
				require.EqualError(t, err, "parsing response list: error code 401 returned: {\"message\":\"test\",\"errors\":[{\"status\":401,\"code\":\"unauthorized\",\"title\":\"Not Authorized\"}]}")
			},
		},
		{
			name: "success",
			params: models.GetGameRankAndScoreParameters{
				GameID:        515,
				LatestMasters: &latest,
			},
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseMessage: []models.GetGameRankAndScore{
				{
					User:            "Blazekickn",
					NumAchievements: 61,
					TotalScore:      453,
					LastAward: models.DateTime{
						Time: lastAwarded1,
					},
				},
				{
					User:            "mamekin",
					NumAchievements: 61,
					TotalScore:      453,
					LastAward: models.DateTime{
						Time: lastAwarded2,
					},
				},
			},
			response: func(messageBytes []byte, errorBytes []byte) []byte {
				return messageBytes
			},
			assert: func(t *testing.T, resp []models.GetGameRankAndScore, err error) {
				require.NotNil(t, resp)
				require.Len(t, resp, 2)
				require.Equal(t, "Blazekickn", resp[0].User)
				require.Equal(t, 61, resp[0].NumAchievements)
				require.Equal(t, 453, resp[0].TotalScore)
				require.Equal(t, lastAwarded1, resp[0].LastAward.Time)
				require.Equal(t, "mamekin", resp[1].User)
				require.Equal(t, 61, resp[1].NumAchievements)
				require.Equal(t, 453, resp[1].TotalScore)
				require.Equal(t, lastAwarded2, resp[1].LastAward.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := "/API/API_GetGameRankAndScore.php"
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
			resp, err := client.GetGameRankAndScore(test.params)
			test.assert(t, resp, err)
		})
	}
}
