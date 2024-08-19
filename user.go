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
func (c *Client) GetGameInfoAndUserProgress(username string, game int, includeAwardMetadata bool) (*models.UserGameProgress, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameInfoAndUserProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(username),
		raHttp.GameID(game),
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
