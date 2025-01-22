package retroachievements

import (
	"fmt"
	"net/http"
	"strconv"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetGameLeaderboards gets a given games's list of leaderboards.
func (c *Client) GetGameLeaderboards(params models.GetGameLeaderboardsParameters) (*models.GetGameLeaderboards, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/API/API_GetGameLeaderboards.php"),
		raHttp.Y(c.APISecret),
		raHttp.I([]string{strconv.Itoa(params.GameID)}),
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
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/API/API_GetLeaderboardEntries.php"),
		raHttp.Y(c.APISecret),
		raHttp.I([]string{strconv.Itoa(params.LeaderboardID)}),
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

// GetUserGameLeaderboards gets a user's list of leaderboards for a given game.
func (c *Client) GetUserGameLeaderboards(params models.GetUserGameLeaderboardsParameters) (*models.GetUserGameLeaderboards, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/API/API_GetUserGameLeaderboards.php"),
		raHttp.Y(c.APISecret),
		raHttp.U(params.Username),
		raHttp.I([]string{strconv.Itoa(params.GameID)}),
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
	resp, err := raHttp.ResponseObject[models.GetUserGameLeaderboards](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
