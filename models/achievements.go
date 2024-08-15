package models

// Achievement is a representation of achievements unlocked by a user
type Achievement struct {
	// The date the achievement was unlocked
	Date DateTime `json:"Date"`

	// Mode the achievement was unlocked in: 1 if in hardcore mode, 0 if not
	HardcoreMode int `json:"HardcoreMode"`

	// The ID of the achievement
	AchievementID int `json:"AchievementID"`

	// Title of the achievement
	Title string `json:"Title"`

	// Description of the achievement
	Description string `json:"Description"`

	// Name of the padge resource
	BadgeName string `json:"BadgeName"`

	// Points awarded
	Points int `json:"Points"`

	// Ratio of points the achievemnet is worth
	TrueRatio int `json:"TrueRatio"`

	// Type of achievement (standard, missable, progression, win_condition)
	Type string `json:"Type"`

	// Username of the author of the achievement
	Author string `json:"Author"`

	// Title of the game
	GameTitle string `json:"GameTitle"`

	// URL resource of the game icon
	GameIcon string `json:"GameIcon"`

	// ID of the game
	GameID int `json:"GameID"`

	// Common name of the console
	ConsoleName string `json:"ConsoleName"`

	// URL resource to the image used for the achievment badge
	BadgeURL string `json:"BadgeURL"`

	// URL resource to the game page
	GameURL string `json:"GameURL"`
}

type GameAchievement struct {
	ID                 int      `json:"ID"`
	NumAwarded         int      `json:"NumAwarded"`
	NumAwardedHardcore int      `json:"NumAwardedHardcore"`
	Title              string   `json:"Title"`
	Description        string   `json:"Description"`
	Points             int      `json:"Points"`
	TrueRatio          int      `json:"TrueRatio"`
	Author             string   `json:"Author"`
	DateModified       DateTime `json:"DateModified"`
	DateCreated        DateTime `json:"DateCreated"`
	BadgeName          string   `json:"BadgeName"`
	DisplayOrder       int      `json:"DisplayOrder"`
	MemAddr            string   `json:"MemAddr"`
	Type               string   `json:"type"`
}
