package client

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

func (c *Client) GetUserRecentAchievements(username string, lookbackMinutes int) ([]models.Achievement, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserRecentAchievements.php"),
		raHttp.APIToken(c.secret),
		raHttp.Username(username),
		raHttp.LookbackMinutes(lookbackMinutes),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	defer resp.Body.Close()
	achievements, err := raHttp.ResponseList[models.Achievement](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}
