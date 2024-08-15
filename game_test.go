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

func makGame(released time.Time) models.Game {
	return models.Game{
		Title:        "Twisted Metal: Black",
		ConsoleID:    21,
		ForumTopicID: 16654,
		Flags:        0,
		ImageIcon:    "/Images/057992.png",
		ImageTitle:   "/Images/056152.png",
		ImageIngame:  "/Images/056151.png",
		ImageBoxArt:  "/Images/050832.png",
		Publisher:    "Sony Computer Entertainment",
		Developer:    "Incognito Entertainment",
		Genre:        "Vehicular Combat",
		Released: models.LongMonthDate{
			Time: released,
		},
	}
}

func TestGetGame(tt *testing.T) {
	released, err := time.Parse(models.LongMonthDateFormat, "June 18, 2001")
	require.NoError(tt, err)
	tests := []struct {
		name             string
		id               int
		modifyURL        func(url string) string
		responseCode     int
		responseGameInfo models.GameInfo
		responseError    models.ErrorResponse
		response         func(gameInfoBytes []byte, errorBytes []byte) []byte
		assert           func(t *testing.T, gameInfo *models.GameInfo, err error)
	}{
		{
			name: "fail to call endpoint",
			id:   2991,
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			responseGameInfo: models.GameInfo{
				Game:        makGame(released),
				GameTitle:   "Twisted Metal: Black",
				ConsoleName: "Playstation 2",
				Console:     "Playstation 2",
				GameIcon:    "/Images/057992.png",
			},
			response: func(gameInfoBytes []byte, errorBytes []byte) []byte {
				return gameInfoBytes
			},
			assert: func(t *testing.T, gameInfo *models.GameInfo, err error) {
				require.Nil(t, gameInfo)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetGame.php?i=2991&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			id:   2991,
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
			response: func(gameInfoBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, gameInfo *models.GameInfo, err error) {
				require.Nil(t, gameInfo)
				require.EqualError(t, err, "parsing response object: error responses: [401] Not Authorized")
			},
		},
		{
			name: "success",
			id:   2991,
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseGameInfo: models.GameInfo{
				Game:        makGame(released),
				GameTitle:   "Twisted Metal: Black",
				ConsoleName: "Playstation 2",
				Console:     "Playstation 2",
				GameIcon:    "/Images/057992.png",
			},
			response: func(gameInfoBytes []byte, errorBytes []byte) []byte {
				return gameInfoBytes
			},
			assert: func(t *testing.T, gameInfo *models.GameInfo, err error) {
				require.NotNil(t, gameInfo)
				require.Equal(t, gameInfo.Title, "Twisted Metal: Black")
				require.Equal(t, gameInfo.GameTitle, "Twisted Metal: Black")
				require.Equal(t, gameInfo.ConsoleID, 21)
				require.Equal(t, gameInfo.ConsoleName, "Playstation 2")
				require.Equal(t, gameInfo.Console, "Playstation 2")
				require.Equal(t, gameInfo.ForumTopicID, 16654)
				require.Equal(t, gameInfo.Flags, 0)
				require.Equal(t, gameInfo.GameIcon, "/Images/057992.png")
				require.Equal(t, gameInfo.ImageIcon, "/Images/057992.png")
				require.Equal(t, gameInfo.ImageTitle, "/Images/056152.png")
				require.Equal(t, gameInfo.ImageIngame, "/Images/056151.png")
				require.Equal(t, gameInfo.ImageBoxArt, "/Images/050832.png")
				require.Equal(t, gameInfo.Publisher, "Sony Computer Entertainment")
				require.Equal(t, gameInfo.Developer, "Incognito Entertainment")
				require.Equal(t, gameInfo.Genre, "Vehicular Combat")
				require.Equal(t, gameInfo.Released.Time, released)
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
				gameInfoBytes, err := json.Marshal(test.responseGameInfo)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(gameInfoBytes, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
			gameInfo, err := client.GetGame(test.id)
			test.assert(t, gameInfo, err)
		})
	}
}

func TestGetGameExtended(tt *testing.T) {
	released, err := time.Parse(models.LongMonthDateFormat, "June 18, 2001")
	require.NoError(tt, err)
	updated, err := time.Parse(time.RFC3339Nano, "2024-08-15T11:46:06.000000Z")
	require.NoError(tt, err)
	modified, err := time.Parse(time.DateTime, "2022-10-25 17:00:49")
	require.NoError(tt, err)
	created, err := time.Parse(time.DateTime, "2022-09-28 00:36:26")
	require.NoError(tt, err)
	tests := []struct {
		name                     string
		id                       int
		modifyURL                func(url string) string
		responseCode             int
		responseExtendedGameInfo models.ExtentedGameInfo
		responseError            models.ErrorResponse
		response                 func(gameBytes []byte, errorBytes []byte) []byte
		assert                   func(t *testing.T, game *models.ExtentedGameInfo, err error)
	}{
		{
			name: "fail to call endpoint",
			id:   2991,
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			responseExtendedGameInfo: models.ExtentedGameInfo{
				Game:               makGame(released),
				ID:                 2991,
				IsFinal:            0,
				RichPresencePatch:  "e7a5e12072a6c976a1146756726fdd8c",
				GuideURL:           "",
				Updated:            updated,
				ConsoleName:        "PlayStation 2",
				ParentGameID:       "",
				NumDistinctPlayers: 1287,
				NumAchievements:    93,
				Achievements: map[int]models.GameAchievement{
					252117: {
						Achievement: models.Achievement{
							Title:       "Zorko Bros. Scrap & Salvage",
							Description: "Destroy all enemies in Junkyard in Story Mode",
							Points:      5,
							TrueRatio:   5,
							Author:      "TheJediSonic",
						},
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
			response: func(extendedGameInfoBytes []byte, errorBytes []byte) []byte {
				return extendedGameInfoBytes
			},
			assert: func(t *testing.T, game *models.ExtentedGameInfo, err error) {
				require.Nil(t, game)
				require.EqualError(t, err, "calling endpoint: Get \"/API/API_GetGameExtended.php?i=2991&y=some_secret\": unsupported protocol scheme \"\"")
			},
		},
		{
			name: "error response",
			id:   2991,
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
			response: func(gameBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, game *models.ExtentedGameInfo, err error) {
				require.Nil(t, game)
				require.EqualError(t, err, "parsing response object: error responses: [401] Not Authorized")
			},
		},
		{
			name: "success",
			id:   2991,
			modifyURL: func(url string) string {
				return url
			},
			responseCode: http.StatusOK,
			responseExtendedGameInfo: models.ExtentedGameInfo{
				Game:               makGame(released),
				ID:                 2991,
				IsFinal:            0,
				RichPresencePatch:  "e7a5e12072a6c976a1146756726fdd8c",
				GuideURL:           "",
				Updated:            updated,
				ConsoleName:        "PlayStation 2",
				ParentGameID:       "",
				NumDistinctPlayers: 1287,
				NumAchievements:    93,
				Achievements: map[int]models.GameAchievement{
					252117: {
						Achievement: models.Achievement{
							Title:       "Zorko Bros. Scrap & Salvage",
							Description: "Destroy all enemies in Junkyard in Story Mode",
							Points:      5,
							TrueRatio:   5,
							Author:      "TheJediSonic",
						},
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
			response: func(extendedGameInfoBytes []byte, errorBytes []byte) []byte {
				return extendedGameInfoBytes
			},
			assert: func(t *testing.T, extendedGameInfo *models.ExtentedGameInfo, err error) {
				require.NotNil(t, extendedGameInfo)
				require.Equal(t, extendedGameInfo.Title, "Twisted Metal: Black")
				require.Equal(t, extendedGameInfo.ConsoleID, 21)
				require.Equal(t, extendedGameInfo.ConsoleName, "PlayStation 2")
				require.Equal(t, extendedGameInfo.ForumTopicID, 16654)
				require.Equal(t, extendedGameInfo.Flags, 0)
				require.Equal(t, extendedGameInfo.ImageIcon, "/Images/057992.png")
				require.Equal(t, extendedGameInfo.ImageTitle, "/Images/056152.png")
				require.Equal(t, extendedGameInfo.ImageIngame, "/Images/056151.png")
				require.Equal(t, extendedGameInfo.ImageBoxArt, "/Images/050832.png")
				require.Equal(t, extendedGameInfo.Publisher, "Sony Computer Entertainment")
				require.Equal(t, extendedGameInfo.Developer, "Incognito Entertainment")
				require.Equal(t, extendedGameInfo.Genre, "Vehicular Combat")
				require.Equal(t, extendedGameInfo.Released.Time, released)
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
				extendedGameInfoBytes, err := json.Marshal(test.responseExtendedGameInfo)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(extendedGameInfoBytes, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
			extendedGameInfo, err := client.GetGameExtended(test.id)
			test.assert(t, extendedGameInfo, err)
		})
	}
}
