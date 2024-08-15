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
	released, err := time.Parse(models.LongMonthDateFormat, "June 18, 2001")
	require.NoError(tt, err)
	tests := []struct {
		name          string
		id            int
		modifyURL     func(url string) string
		responseCode  int
		responseGame  models.Game
		responseError models.ErrorResponse
		response      func(gameBytes []byte, errorBytes []byte) []byte
		assert        func(t *testing.T, game *models.Game, err error)
	}{
		{
			name: "fail to call endpoint",
			id:   2991,
			modifyURL: func(url string) string {
				return ""
			},
			responseCode: http.StatusOK,
			responseGame: models.Game{
				CommonGame: models.CommonGame{
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
				},
				GameTitle:   "Twisted Metal: Black",
				ConsoleName: "Playstation 2",
				Console:     "Playstation 2",
				GameIcon:    "/Images/057992.png",
			},
			response: func(gameBytes []byte, errorBytes []byte) []byte {
				return gameBytes
			},
			assert: func(t *testing.T, game *models.Game, err error) {
				require.Nil(t, game)
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
			response: func(gameBytes []byte, errorBytes []byte) []byte {
				return errorBytes
			},
			assert: func(t *testing.T, game *models.Game, err error) {
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
			responseGame: models.Game{
				CommonGame: models.CommonGame{
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
				},
				GameTitle:   "Twisted Metal: Black",
				ConsoleName: "Playstation 2",
				Console:     "Playstation 2",
				GameIcon:    "/Images/057992.png",
			},
			response: func(gameBytes []byte, errorBytes []byte) []byte {
				return gameBytes
			},
			assert: func(t *testing.T, game *models.Game, err error) {
				require.NotNil(t, game)
				require.Equal(t, game.Title, "Twisted Metal: Black")
				require.Equal(t, game.GameTitle, "Twisted Metal: Black")
				require.Equal(t, game.ConsoleID, 21)
				require.Equal(t, game.ConsoleName, "Playstation 2")
				require.Equal(t, game.Console, "Playstation 2")
				require.Equal(t, game.ForumTopicID, 16654)
				require.Equal(t, game.Flags, 0)
				require.Equal(t, game.GameIcon, "/Images/057992.png")
				require.Equal(t, game.ImageIcon, "/Images/057992.png")
				require.Equal(t, game.ImageTitle, "/Images/056152.png")
				require.Equal(t, game.ImageIngame, "/Images/056151.png")
				require.Equal(t, game.ImageBoxArt, "/Images/050832.png")
				require.Equal(t, game.Publisher, "Sony Computer Entertainment")
				require.Equal(t, game.Developer, "Incognito Entertainment")
				require.Equal(t, game.Genre, "Vehicular Combat")
				require.Equal(t, game.Released.Time, released)
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
				gameBytes, err := json.Marshal(test.responseGame)
				require.NoError(t, err)
				errBytes, err := json.Marshal(test.responseError)
				require.NoError(t, err)
				resp := test.response(gameBytes, errBytes)
				num, err := w.Write(resp)
				require.NoError(t, err)
				require.Equal(t, num, len(resp))
			}))
			defer server.Close()

			client := retroachievements.New(test.modifyURL(server.URL), "some_secret")
			game, err := client.GetGame(test.id)
			test.assert(t, game, err)
		})
	}
}

func TestGetGameExtended(tt *testing.T) {
}
