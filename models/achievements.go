package models

import "time"

type Achievement struct {
	Date          time.Time `json:"Date"`
	HardcoreMode  int       `json:"HardcoreMode"`
	AchievementID int       `json:"AchievementID"`
	Title         string    `json:"Title"`
	Description   string    `json:"Description"`
	BadgeName     string    `json:"BadgeName"`
	Points        int       `json:"Points"`
	TrueRatio     int       `json:"TrueRatio"`
	Type          string    `json:"Type,omitempty"`
	Author        string    `json:"Author"`
	GameTitle     string    `json:"GameTitle"`
	GameIcon      string    `json:"GameIcon"`
	GameID        int       `json:"GameID"`
	ConsoleName   string    `json:"ConsoleName"`
	BadgeURL      string    `json:"BadgeURL"`
	GameURL       string    `json:"GameURL"`
}
