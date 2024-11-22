package models

import "time"

type GetAchievementUnlocksParameters struct {
	// The target achievement ID
	AchievementID int

	// [Optional] The number of records to return (default: 50, max: 500).
	Count *int

	// [Optional] The number of entries to skip (default: 0).
	Offset *int
}

type GetAchievementUnlocks struct {
	Achievement          GetAchievementUnlocksAchievement `json:"Achievement"`
	Console              GetAchievementUnlocksConsole     `json:"Console"`
	Game                 GetAchievementUnlocksGame        `json:"Game"`
	UnlocksCount         int                              `json:"UnlocksCount"`
	UnlocksHardcoreCount int                              `json:"UnlocksHardcoreCount"`
	TotalPlayers         int                              `json:"TotalPlayers"`
	Unlocks              []GetAchievementUnlocksUnlock    `json:"Unlocks"`
}

type GetAchievementUnlocksAchievement struct {
	ID           int      `json:"ID"`
	Title        string   `json:"Title"`
	Description  string   `json:"Description"`
	Points       int      `json:"Points"`
	TrueRatio    int      `json:"TrueRatio"`
	Author       string   `json:"Author"`
	DateCreated  DateTime `json:"DateCreated"`
	DateModified DateTime `json:"DateModified"`
	Type         *string  `json:"Type"`
}

type GetAchievementUnlocksConsole struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
}

type GetAchievementUnlocksGame struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
}

type GetAchievementUnlocksUnlock struct {
	User             string    `json:"User"`
	RAPoints         int       `json:"RAPoints"`
	RASoftcorePoints int       `json:"RASoftcorePoints"`
	DateAwarded      time.Time `json:"DateAwarded"`
	HardcoreMode     int       `json:"HardcoreMode"`
}
