package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetAchievementOfTheWeek gets comprehensive metadata about the current Achievement of the Week.
func (c *Client) GetAchievementOfTheWeek(params models.GetAchievementOfTheWeekParameters) (*models.GetAchievementOfTheWeek, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/API/API_GetAchievementOfTheWeek.php"),
		raHttp.Y(c.APISecret),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetAchievementOfTheWeek](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
