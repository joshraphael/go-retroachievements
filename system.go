package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetConsoleIDs gets the complete list of all system ID and name pairs on the site.
func (c *Client) GetConsoleIDs(params models.GetConsoleIDsParameters) ([]models.GetConsoleIDs, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetConsoleIDs.php"),
		raHttp.APIToken(c.Secret),
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
