package client

import (
	"fmt"
	"net/http"
	"time"

	raHttp "github.com/joshraphael/go-retroachievements/http"
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
	rawAchievementList, err := raHttp.ResponseList[rawAchievement](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}

	achievements := []models.Achievement{}
	for i := range rawAchievementList {
		achievement := rawAchievementList[i]
		a, err := achievement.ToAchievement()
		if err != nil {
			return nil, fmt.Errorf("converting response to achievement: %w", err)
		}
		achievements = append(achievements, *a)
	}
	return achievements, nil
}
