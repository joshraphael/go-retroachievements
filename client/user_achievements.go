package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/joshraphael/go-retroachievements/models"
)

type rawAchievement struct {
	Date          string `json:"Date"`
	HardcoreMode  int    `json:"HardcoreMode"`
	AchievementID int    `json:"AchievementID"`
	Title         string `json:"Title"`
	Description   string `json:"Description"`
	BadgeName     string `json:"BadgeName"`
	Points        int    `json:"Points"`
	TrueRatio     int    `json:"TrueRatio"`
	Type          string `json:"Type"`
	Author        string `json:"Author"`
	GameTitle     string `json:"GameTitle"`
	GameIcon      string `json:"GameIcon"`
	GameID        int    `json:"GameID"`
	ConsoleName   string `json:"ConsoleName"`
	BadgeURL      string `json:"BadgeURL"`
	GameURL       string `json:"GameURL"`
}

func (ra *rawAchievement) ToAchievement() (*models.Achievement, error) {
	t, err := time.Parse(time.DateTime, ra.Date)
	if err != nil {
		return nil, err
	}
	return &models.Achievement{
		Date:          t,
		HardcoreMode:  ra.HardcoreMode,
		AchievementID: ra.AchievementID,
		Title:         ra.Title,
		Description:   ra.Description,
		BadgeName:     ra.BadgeName,
		Points:        ra.Points,
		TrueRatio:     ra.TrueRatio,
		Type:          ra.Type,
		Author:        ra.Author,
		GameTitle:     ra.GameTitle,
		GameIcon:      ra.GameIcon,
		GameID:        ra.GameID,
		ConsoleName:   ra.ConsoleName,
		BadgeURL:      ra.BadgeURL,
		GameURL:       ra.GameURL,
	}, nil
}

func (c *Client) GetUserRecentAchievements(username string, lookbackMinutes int) ([]models.Achievement, error) {

	u, err := url.Parse(c.host + "/API/API_GetUserRecentAchievements.php")
	if err != nil {
		return nil, fmt.Errorf("parsing GetUserRecentAchievements url: %w", err)
	}
	q := u.Query()
	q.Set("y", c.secret)
	q.Set("u", username)
	q.Set("m", strconv.Itoa(lookbackMinutes))
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("calling GetUserRecentAchievements: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		var respError models.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&respError)
		if err != nil {
			return nil, fmt.Errorf("decoding response error body: %w", err)
		}
		errText := []string{}
		for i := range respError.Errors {
			err := respError.Errors[i]
			errText = append(errText, fmt.Sprintf("[%d] %s", err.Status, err.Title))
		}
		return nil, fmt.Errorf("calling get user profile: %s", strings.Join(errText, ", "))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}

	var as []rawAchievement
	err = json.NewDecoder(resp.Body).Decode(&as)
	if err != nil {
		return nil, fmt.Errorf("decoding response body profile: %w", err)
	}

	if len(as) == 0 {
		return nil, nil
	}
	achievements := []models.Achievement{}
	for i := range as {
		achievement := as[i]
		a, err := achievement.ToAchievement()
		if err != nil {
			return nil, fmt.Errorf("converting response to achievement: %w", err)
		}
		achievements = append(achievements, *a)
	}
	return achievements, nil
}
