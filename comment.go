package retroachievements

import (
	"fmt"
	"net/http"
	"strconv"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetComments gets comments of a specified kind: game, achievement, or user.
func (c *Client) GetComments(params models.GetCommentsParameters) (*models.GetComments, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/API/API_GetComments.php"),
		raHttp.Y(c.APISecret),
		raHttp.T(strconv.Itoa(params.Type.GetCommentsType())),
	}
	switch params.Type.(type) {
	case models.GetCommentsGame:
		game := params.Type.(models.GetCommentsGame)
		details = append(details, raHttp.I([]string{strconv.Itoa(game.GameID)}))
	case models.GetCommentsAchievement:
		achievement := params.Type.(models.GetCommentsAchievement)
		details = append(details, raHttp.I([]string{strconv.Itoa(achievement.AchievementID)}))
	case models.GetCommentsUser:
		user := params.Type.(models.GetCommentsUser)
		details = append(details, raHttp.I([]string{user.Username}))
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
	resp, err := raHttp.ResponseObject[models.GetComments](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
