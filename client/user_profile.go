package client

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

func (c *Client) GetUserProfile(username string) (*models.Profile, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProfile.php"),
		raHttp.APIToken(c.secret),
		raHttp.Username(username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	defer resp.Body.Close()
	profile, err := raHttp.ResponseObject[models.Profile](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return profile, nil
}
