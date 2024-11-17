package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetUserProfile get a user's basic profile information.
func (c *Client) GetUserProfile(params models.GetUserProfileParameters) (*models.GetUserProfile, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProfile.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	profile, err := raHttp.ResponseObject[models.GetUserProfile](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return profile, nil
}

// GetUserRecentAchievements get a list of achievements recently earned by the user.
func (c *Client) GetUserRecentAchievements(params models.GetUserRecentAchievementsParameters) ([]models.GetUserRecentAchievements, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserRecentAchievements.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	}
	if params.LookbackMinutes != nil {
		details = append(details, raHttp.LookbackMinutes(*params.LookbackMinutes))
	}
	resp, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	achievements, err := raHttp.ResponseList[models.GetUserRecentAchievements](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}

// GetAchievementsEarnedBetween get a list of achievements earned by a user between two dates.
func (c *Client) GetAchievementsEarnedBetween(params models.GetAchievementsEarnedBetweenParameters) ([]models.GetAchievementsEarnedBetween, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementsEarnedBetween.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
		raHttp.From(params.From.Unix()),
		raHttp.To(params.To.Unix()),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	achievements, err := raHttp.ResponseList[models.GetAchievementsEarnedBetween](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}

// GetAchievementsEarnedOnDay get a list of achievements earned by a user on a given date.
func (c *Client) GetAchievementsEarnedOnDay(params models.GetAchievementsEarnedOnDayParameters) ([]models.GetAchievementsEarnedOnDay, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementsEarnedOnDay.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
		raHttp.Date(params.Date),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	achievements, err := raHttp.ResponseList[models.GetAchievementsEarnedOnDay](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}

// GetGameInfoAndUserProgress get metadata about a game as well as a user's progress on that game.
func (c *Client) GetGameInfoAndUserProgress(params models.GetGameInfoAndUserProgressParameters) (*models.GetGameInfoAndUserProgress, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameInfoAndUserProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
		raHttp.Game(params.GameID),
	}
	if params.IncludeAwardMetadata != nil && *params.IncludeAwardMetadata {
		details = append(details, raHttp.Achievement(1))
	}
	resp, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	gameProgress, err := raHttp.ResponseObject[models.GetGameInfoAndUserProgress](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return gameProgress, nil
}

// GetUserCompletionProgress get metadata about all the user's played games and any awards associated with them.
func (c *Client) GetUserCompletionProgress(params models.GetUserCompletionProgressParameters) (*models.GetUserCompletionProgress, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserCompletionProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	completionProgress, err := raHttp.ResponseObject[models.GetUserCompletionProgress](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return completionProgress, nil
}

// GetUserAwards get a list of a user's site awards/badges.
func (c *Client) GetUserAwards(params models.GetUserAwardsParameters) (*models.GetUserAwards, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserAwards.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	awards, err := raHttp.ResponseObject[models.GetUserAwards](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return awards, nil
}

// GetUserClaims get a list of set development claims made over the lifetime of a user.
func (c *Client) GetUserClaims(params models.GetUserClaimsParameters) ([]models.GetUserClaims, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserClaims.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	claims, err := raHttp.ResponseList[models.GetUserClaims](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return claims, nil
}

// GetUserGameRankAndScore get metadata about how a user has performed on a given game.
func (c *Client) GetUserGameRankAndScore(params models.GetUserGameRankAndScoreParameters) ([]models.GetUserGameRankAndScore, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserGameRankAndScore.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
		raHttp.Game(params.GameID),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	userGameRankScore, err := raHttp.ResponseList[models.GetUserGameRankAndScore](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return userGameRankScore, nil
}

// GetUserPoints get a user's total hardcore and softcore points.
func (c *Client) GetUserPoints(params models.GetUserPointsParameters) (*models.GetUserPoints, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserPoints.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	points, err := raHttp.ResponseObject[models.GetUserPoints](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return points, nil
}

// GetUserProgress get a user's progress on a list of specified games.
func (c *Client) GetUserProgress(params models.GetUserProgressParameters) (*map[string]models.GetUserProgress, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
		raHttp.IDs(params.GameIDs),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	progress, err := raHttp.ResponseObject[map[string]models.GetUserProgress](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return progress, nil
}

// GetUserRecentlyPlayedGames get a list of games a user has recently played.
func (c *Client) GetUserRecentlyPlayedGames(params models.GetUserRecentlyPlayedGamesParameters) ([]models.GetUserRecentlyPlayedGames, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserRecentlyPlayedGames.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	}
	if params.Count != nil {
		details = append(details, raHttp.Count(*params.Count))
	}
	if params.Offset != nil {
		details = append(details, raHttp.Offset(*params.Offset))
	}
	resp, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	recentlyPlayed, err := raHttp.ResponseList[models.GetUserRecentlyPlayedGames](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return recentlyPlayed, nil
}

// GetUserSummary get summary information about a given user.
func (c *Client) GetUserSummary(params models.GetUserSummaryParameters) (*models.GetUserSummary, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserSummary.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	}
	if params.GamesCount != nil {
		details = append(details, raHttp.Game(*params.GamesCount))
	}
	if params.AchievementsCount != nil {
		details = append(details, raHttp.Achievement(*params.AchievementsCount))
	}
	resp, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	recentlyPlayed, err := raHttp.ResponseObject[models.GetUserSummary](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return recentlyPlayed, nil
}

// GetUserCompletedGames gets completion metadata about the games a given user has played.
func (c *Client) GetUserCompletedGames(params models.GetUserCompletedGamesParameters) ([]models.GetUserCompletedGames, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserCompletedGames.php"),
		raHttp.APIToken(c.Secret),
		raHttp.Username(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	games, err := raHttp.ResponseList[models.GetUserCompletedGames](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return games, nil
}
