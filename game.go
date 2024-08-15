package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetGame gets info of a game
func (c *Client) GetGame(id int) (*models.GameInfo, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGame.php"),
		raHttp.APIToken(c.secret),
		raHttp.ID(id),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	game, err := raHttp.ResponseObject[models.GameInfo](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return game, nil
}

// GetGameExtended gets extended info of a game
func (c *Client) GetGameExtended(id int) (*models.ExtentedGameInfo, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameExtended.php"),
		raHttp.APIToken(c.secret),
		raHttp.ID(id),
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
