package client

import (
	"fmt"
	"net/http"
	"time"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetUserProfile gets the profile of a given username
func (c *Client) GetUserProfile(username string) (*models.Profile, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProfile.php"),
		raHttp.APIToken(c.secret),
		raHttp.Username(username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	defer resp.Body.Close()
	profile, err := raHttp.ResponseObject[models.Profile](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return profile, nil
}

// GetUserRecentAchievements gets all achievements within the last specified amount of minutes for a given username
func (c *Client) GetUserRecentAchievements(username string, lookbackMinutes int) ([]models.Achievement, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserRecentAchievements.php"),
		raHttp.APIToken(c.secret),
		raHttp.Username(username),
		raHttp.LookbackMinutes(lookbackMinutes),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	defer resp.Body.Close()
	achievements, err := raHttp.ResponseList[models.Achievement](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}

// GetAchievementsEarnedBetween gets all achievements earned within a time frame for a given username
func (c *Client) GetAchievementsEarnedBetween(username string, from time.Time, to time.Time) ([]models.Achievement, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementsEarnedBetween.php"),
		raHttp.APIToken(c.secret),
		raHttp.Username(username),
		raHttp.FromTime(from),
		raHttp.ToTime(to),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	defer resp.Body.Close()
	achievements, err := raHttp.ResponseList[models.Achievement](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}

// GetAchievementsEarnedOnDay gets all achievements earned on a specific day for a given username
func (c *Client) GetAchievementsEarnedOnDay(username string, date time.Time) ([]models.Achievement, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetAchievementsEarnedOnDay.php"),
		raHttp.APIToken(c.secret),
		raHttp.Username(username),
		raHttp.Date(date),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	defer resp.Body.Close()
	achievements, err := raHttp.ResponseList[models.Achievement](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return achievements, nil
}
