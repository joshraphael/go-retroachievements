package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetAchievementUnlocks gets a list of users who have earned an achievement.
func (c *Client) GetAchievementUnlocks(params models.GetAchievementUnlocksParameters) (*models.GetAchievementUnlocks, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/API/API_GetAchievementUnlocks.php"),
		raHttp.Y(c.APISecret),
		raHttp.A(params.AchievementID),
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
	resp, err := raHttp.ResponseObject[models.GetAchievementUnlocks](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
