package retroachievements

import (
	"fmt"
	"net/http"
	"strconv"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetConsoleIDs gets the complete list of all system ID and name pairs on the site.
func (c *Client) GetConsoleIDs(params models.GetConsoleIDsParameters) ([]models.GetConsoleIDs, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/API/API_GetConsoleIDs.php"),
		raHttp.Y(c.APISecret),
	}
	if params.OnlyActive != nil {
		active := 0
		if *params.OnlyActive {
			active = 1
		}
		details = append(details, raHttp.A(active))
	}
	if params.OnlyGameSystems != nil {
		g := 0
		if *params.OnlyGameSystems {
			g = 1
		}
		details = append(details, raHttp.G(g))
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetConsoleIDs](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}

// GetGameList gets the complete list of games for a specified console on the site.
func (c *Client) GetGameList(params models.GetGameListParameters) ([]models.GetGameList, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/API/API_GetGameList.php"),
		raHttp.Y(c.APISecret),
		raHttp.I([]string{strconv.Itoa(params.SystemID)}),
	}
	if params.HasAchievements != nil {
		f := 0
		if *params.HasAchievements {
			f = 1
		}
		details = append(details, raHttp.F(f))
	}
	if params.IncludeHashes != nil {
		h := 0
		if *params.IncludeHashes {
			h = 1
		}
		details = append(details, raHttp.H(h))
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
	resp, err := raHttp.ResponseList[models.GetGameList](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}
