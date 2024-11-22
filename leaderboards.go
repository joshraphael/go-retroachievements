package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetGameLeaderboards gets a given games's list of leaderboards.
func (c *Client) GetGameLeaderboards(params models.GetGameLeaderboardsParameters) (*models.GetGameLeaderboards, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameLeaderboards.php"),
		raHttp.APIToken(c.Secret),
		raHttp.I([]int{params.GameID}),
	}
	if params.Count != nil {
		details = append(details, raHttp.C(*params.Count))
	}
	if params.Offset != nil {
		details = append(details, raHttp.O(*params.Offset))
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetGameLeaderboards](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetLeaderboardEntries gets a given leaderboards's entries.
func (c *Client) GetLeaderboardEntries(params models.GetLeaderboardEntriesParameters) (*models.GetLeaderboardEntries, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetLeaderboardEntries.php"),
		raHttp.APIToken(c.Secret),
		raHttp.I([]int{params.LeaderboardID}),
	}
	if params.Count != nil {
		details = append(details, raHttp.C(*params.Count))
	}
	if params.Offset != nil {
		details = append(details, raHttp.O(*params.Offset))
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetLeaderboardEntries](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
