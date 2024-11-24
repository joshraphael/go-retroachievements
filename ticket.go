package retroachievements

import (
	"fmt"
	"net/http"
	"strconv"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetTicketByID gets ticket metadata information about a single achievement ticket, targeted by its ticket ID.
func (c *Client) GetTicketByID(params models.GetTicketByIDParameters) (*models.GetTicketByID, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetTicketData.php"),
		raHttp.APIToken(c.Secret),
		raHttp.I([]string{
			strconv.Itoa(params.TicketID),
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetTicketByID](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetMostTicketedGames gets the games on the site with the highest count of opened achievement tickets.
func (c *Client) GetMostTicketedGames(params models.GetMostTicketedGamesParameters) (*models.GetMostTicketedGames, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetTicketData.php"),
		raHttp.APIToken(c.Secret),
		raHttp.F(1),
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
	resp, err := raHttp.ResponseObject[models.GetMostTicketedGames](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetMostRecentTickets gets ticket metadata information about the latest opened achievement tickets on RetroAchievements.
func (c *Client) GetMostRecentTickets(params models.GetMostRecentTicketsParameters) (*models.GetMostRecentTickets, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetTicketData.php"),
		raHttp.APIToken(c.Secret),
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
	resp, err := raHttp.ResponseObject[models.GetMostRecentTickets](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetGameTicketStats gets ticket stats for a game, targeted by that game's unique ID.
func (c *Client) GetGameTicketStats(params models.GetGameTicketStatsParameters) (*models.GetGameTicketStats, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetTicketData.php"),
		raHttp.APIToken(c.Secret),
		raHttp.G(params.GameID),
	}
	if params.Unofficial != nil && *params.Unofficial {
		details = append(details, raHttp.F(5))
	}
	if params.IncludeTicketMetadata != nil && *params.IncludeTicketMetadata {
		details = append(details, raHttp.D(strconv.Itoa(1)))
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetGameTicketStats](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetDeveloperTicketStats gets ticket stats for a developer, targeted by that developer's site username.
func (c *Client) GetDeveloperTicketStats(params models.GetDeveloperTicketStatsParameters) (*models.GetDeveloperTicketStats, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetTicketData.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetDeveloperTicketStats](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
