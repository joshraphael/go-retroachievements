package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetGame get basic metadata about a game.
func (c *Client) GetGame(params models.GetGameParameters) (*models.GetGame, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGame.php"),
		raHttp.APIToken(c.Secret),
		raHttp.IDs([]int{params.GameID}),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	game, err := raHttp.ResponseObject[models.GetGame](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return game, nil
}

// GetGameExtended get extended metadata about a game.
func (c *Client) GetGameExtended(params models.GetGameExtentedParameters) (*models.GetGameExtented, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameExtended.php"),
		raHttp.APIToken(c.Secret),
		raHttp.IDs([]int{params.GameID}),
	}
	if params.Unofficial {
		details = append(details, raHttp.From(int64(5)))
	}
	resp, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	game, err := raHttp.ResponseObject[models.GetGameExtented](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return game, nil
}

// GetGameHashes get the hashes linked to a game.
func (c *Client) GetGameHashes(params models.GetGameHashesParameters) (*models.GetGameHashes, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameHashes.php"),
		raHttp.APIToken(c.Secret),
		raHttp.IDs([]int{params.GameID}),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetGameHashes](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetAchievementCount the list of achievement IDs for a game.
func (c *Client) GetAchievementCount(params models.GetAchievementCountParameters) (*models.GetAchievementCount, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementCount.php"),
		raHttp.APIToken(c.Secret),
		raHttp.IDs([]int{params.GameID}),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetAchievementCount](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetAchievementDistribution gets how many players have unlocked how many achievements for a game.
func (c *Client) GetAchievementDistribution(params models.GetAchievementDistributionParameters) (*models.GetAchievementDistribution, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementDistribution.php"),
		raHttp.APIToken(c.Secret),
		raHttp.IDs([]int{params.GameID}),
	}
	if params.Unofficial != nil {
		if *params.Unofficial {
			details = append(details, raHttp.From(int64(5)))
		} else {
			details = append(details, raHttp.From(int64(3)))
		}
	}
	if params.Hardcore != nil {
		details = append(details, raHttp.Hardcore(*params.Hardcore))
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetAchievementDistribution](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetGameRankAndScore gets metadata about either the latest masters for a game, or the highest points earners for a game.
func (c *Client) GetGameRankAndScore(params models.GetGameRankAndScoreParameters) ([]models.GetGameRankAndScore, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameRankAndScore.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Game(params.GameID),
	}
	if params.LatestMasters != nil {
		if *params.LatestMasters {
			details = append(details, raHttp.To(1))
		} else {
			details = append(details, raHttp.To(0))
		}
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetGameRankAndScore](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}
