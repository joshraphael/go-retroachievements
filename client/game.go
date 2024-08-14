package client

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetGame gets info of a game
func (c *Client) GetGame(id int) (*models.Game, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGame.php"),
		raHttp.APIToken(c.secret),
		raHttp.ID(id),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	game, err := raHttp.ResponseObject[models.Game](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return game, nil
}
