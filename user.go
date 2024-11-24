package retroachievements

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetUserProfile get a user's basic profile information.
func (c *Client) GetUserProfile(params models.GetUserProfileParameters) (*models.GetUserProfile, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProfile.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetUserProfile](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetUserRecentAchievements get a list of achievements recently earned by the user.
func (c *Client) GetUserRecentAchievements(params models.GetUserRecentAchievementsParameters) ([]models.GetUserRecentAchievements, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserRecentAchievements.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	}
	if params.LookbackMinutes != nil {
		details = append(details, raHttp.M(*params.LookbackMinutes))
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetUserRecentAchievements](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}

// GetAchievementsEarnedBetween get a list of achievements earned by a user between two dates.
func (c *Client) GetAchievementsEarnedBetween(params models.GetAchievementsEarnedBetweenParameters) ([]models.GetAchievementsEarnedBetween, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementsEarnedBetween.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
		raHttp.F(int(params.From.Unix())),
		raHttp.T(int(params.To.Unix())),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetAchievementsEarnedBetween](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}

// GetAchievementsEarnedOnDay get a list of achievements earned by a user on a given date.
func (c *Client) GetAchievementsEarnedOnDay(params models.GetAchievementsEarnedOnDayParameters) ([]models.GetAchievementsEarnedOnDay, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementsEarnedOnDay.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
		raHttp.D(params.Date.UTC().Format(time.DateOnly)),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetAchievementsEarnedOnDay](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}

// GetGameInfoAndUserProgress get metadata about a game as well as a user's progress on that game.
func (c *Client) GetGameInfoAndUserProgress(params models.GetGameInfoAndUserProgressParameters) (*models.GetGameInfoAndUserProgress, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetGameInfoAndUserProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
		raHttp.G(params.GameID),
	}
	if params.IncludeAwardMetadata != nil {
		a := 0
		if *params.IncludeAwardMetadata {
			a = 1
		}
		details = append(details, raHttp.A(a))
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetGameInfoAndUserProgress](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetUserCompletionProgress get metadata about all the user's played games and any awards associated with them.
func (c *Client) GetUserCompletionProgress(params models.GetUserCompletionProgressParameters) (*models.GetUserCompletionProgress, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserCompletionProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetUserCompletionProgress](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetUserAwards get a list of a user's site awards/badges.
func (c *Client) GetUserAwards(params models.GetUserAwardsParameters) (*models.GetUserAwards, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserAwards.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetUserAwards](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetUserClaims get a list of set development claims made over the lifetime of a user.
func (c *Client) GetUserClaims(params models.GetUserClaimsParameters) ([]models.GetUserClaims, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserClaims.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetUserClaims](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}

// GetUserGameRankAndScore get metadata about how a user has performed on a given game.
func (c *Client) GetUserGameRankAndScore(params models.GetUserGameRankAndScoreParameters) ([]models.GetUserGameRankAndScore, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserGameRankAndScore.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
		raHttp.G(params.GameID),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetUserGameRankAndScore](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}

// GetUserPoints get a user's total hardcore and softcore points.
func (c *Client) GetUserPoints(params models.GetUserPointsParameters) (*models.GetUserPoints, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserPoints.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetUserPoints](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetUserProgress get a user's progress on a list of specified games.
func (c *Client) GetUserProgress(params models.GetUserProgressParameters) (*map[string]models.GetUserProgress, error) {
	strIDs := []string{}
	for i := range params.GameIDs {
		strIDs = append(strIDs, strconv.Itoa(params.GameIDs[i]))
	}
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProgress.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
		raHttp.I(strIDs),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[map[string]models.GetUserProgress](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetUserRecentlyPlayedGames get a list of games a user has recently played.
func (c *Client) GetUserRecentlyPlayedGames(params models.GetUserRecentlyPlayedGamesParameters) ([]models.GetUserRecentlyPlayedGames, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserRecentlyPlayedGames.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
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
	resp, err := raHttp.ResponseList[models.GetUserRecentlyPlayedGames](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}

// GetUserSummary get summary information about a given user.
func (c *Client) GetUserSummary(params models.GetUserSummaryParameters) (*models.GetUserSummary, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserSummary.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	}
	if params.GamesCount != nil {
		details = append(details, raHttp.G(*params.GamesCount))
	}
	if params.AchievementsCount != nil {
		details = append(details, raHttp.A(*params.AchievementsCount))
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetUserSummary](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetUserCompletedGames gets completion metadata about the games a given user has played.
func (c *Client) GetUserCompletedGames(params models.GetUserCompletedGamesParameters) ([]models.GetUserCompletedGames, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserCompletedGames.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetUserCompletedGames](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}

// GetUserWantToPlayList gets a given user's "Want to Play Games" list.
func (c *Client) GetUserWantToPlayList(params models.GetUserWantToPlayListParameters) (*models.GetUserWantToPlayList, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserWantToPlayList.php"),
		raHttp.APIToken(c.Secret),
		raHttp.U(params.Username),
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
	resp, err := raHttp.ResponseObject[models.GetUserWantToPlayList](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
