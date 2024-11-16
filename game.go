package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetGame get basic metadata about a game.
func (c *Client) GetGame(params models.GetGameParameters) (*models.GetGame, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGame.php"),
		raHttp.APIToken(c.Secret),
		raHttp.IDs([]int{params.GameID}),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	game, err := raHttp.ResponseObject[models.GetGame](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return game, nil
}

// GetGameExtended get extended metadata about a game.
func (c *Client) GetGameExtended(id int) (*models.ExtentedGameInfo, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameExtended.php"),
		raHttp.APIToken(c.Secret),
		raHttp.IDs([]int{id}),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	game, err := raHttp.ResponseObject[models.ExtentedGameInfo](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return game, nil
}
