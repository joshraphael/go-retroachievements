package retroachievements

import (
	"fmt"
	"net/http"
	"time"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetUserProfile get a user's basic profile information.
func (c *Client) GetUserProfile(username string) (*models.Profile, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProfile.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	profile, err := raHttp.ResponseObject[models.Profile](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return profile, nil
}

// GetUserRecentAchievements get a list of achievements recently earned by the user.
func (c *Client) GetUserRecentAchievements(username string, lookbackMinutes int) ([]models.UnlockedAchievement, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserRecentAchievements.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
		raHttp.LookbackMinutes(lookbackMinutes),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	achievements, err := raHttp.ResponseList[models.UnlockedAchievement](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}

// GetAchievementsEarnedBetween get a list of achievements earned by a user between two dates.
func (c *Client) GetAchievementsEarnedBetween(username string, from time.Time, to time.Time) ([]models.UnlockedAchievement, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementsEarnedBetween.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
		raHttp.FromTime(from),
		raHttp.ToTime(to),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	achievements, err := raHttp.ResponseList[models.UnlockedAchievement](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}

// GetAchievementsEarnedOnDay get a list of achievements earned by a user on a given date.
func (c *Client) GetAchievementsEarnedOnDay(username string, date time.Time) ([]models.UnlockedAchievement, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementsEarnedOnDay.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
		raHttp.Date(date),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	achievements, err := raHttp.ResponseList[models.UnlockedAchievement](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}

// GetGameInfoAndUserProgress get metadata about a game as well as a user's progress on that game.
func (c *Client) GetGameInfoAndUserProgress(username string, gameId int, includeAwardMetadata bool) (*models.UserGameProgress, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameInfoAndUserProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
		raHttp.GameID(gameId),
		raHttp.AwardMetadata(includeAwardMetadata),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	gameProgress, err := raHttp.ResponseObject[models.UserGameProgress](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return gameProgress, nil
}

// GetUserCompletionProgress get metadata about all the user's played games and any awards associated with them.
func (c *Client) GetUserCompletionProgress(username string) (*models.UserCompletionProgress, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserCompletionProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	completionProgress, err := raHttp.ResponseObject[models.UserCompletionProgress](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return completionProgress, nil
}

// GetUserAwards get a list of a user's site awards/badges.
func (c *Client) GetUserAwards(username string) (*models.UserAwards, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserAwards.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	awards, err := raHttp.ResponseObject[models.UserAwards](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return awards, nil
}

// GetUserClaims get a list of set development claims made over the lifetime of a user.
func (c *Client) GetUserClaims(username string) ([]models.UserClaims, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserClaims.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	awards, err := raHttp.ResponseList[models.UserClaims](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return awards, nil
}

// GetUserGameRankAndScore get metadata about how a user has performed on a given game.
func (c *Client) GetUserGameRankAndScore(username string, gameId int) ([]models.UserGameRankScore, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserGameRankAndScore.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
		raHttp.GameID(gameId),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	userGameRankScore, err := raHttp.ResponseList[models.UserGameRankScore](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return userGameRankScore, nil
}

// GetUserPoints get a user's total hardcore and softcore points.
func (c *Client) GetUserPoints(username string) (*models.Points, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserPoints.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	points, err := raHttp.ResponseObject[models.Points](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return points, nil
}

// GetUserProgress get a user's progress on a list of specified games.
func (c *Client) GetUserProgress(username string, gameIDs []int) (map[string]models.Progress, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
		raHttp.IDs(gameIDs),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	progress, err := raHttp.ResponseObject[map[string]models.Progress](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return *progress, nil
}
