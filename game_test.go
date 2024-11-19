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
				require.EqualError(t, err, "parsing response object: error responses: [401] Not Authorized")
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
			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
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
	tests := []struct {
		name            string
		params          models.GetGameExtentedParameters
		modifyURL       func(url string) string
		responseCode    int
		responseMessage models.GetGameExtented
		responseError   models.ErrorResponse
		response        func(messageBytes []byte, errorBytes []byte) []byte
		assert          func(t *testing.T, game *models.GetGameExtented, err error)
	}{
		{
			name: "fail to call endpoint",
			params: models.GetGameExtentedParameters{
				GameID:     2991,
				Unofficial: true,
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
				require.EqualError(t, err, "parsing response object: error responses: [401] Not Authorized")
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

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
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
		assert          func(t *testing.T, game *models.GetGameHashes, err error)
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
				require.EqualError(t, err, "parsing response object: error responses: [401] Not Authorized")
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

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
			resp, err := client.GetGameHashes(test.params)
			test.assert(t, resp, err)
		})
	}
}
