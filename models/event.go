package models

import "time"

type GetAchievementOfTheWeekParameters struct{}

type GetAchievementOfTheWeek struct {
	Achievement          GetAchievementOfTheWeekAchievement `json:"Achievement"`
	Console              GetAchievementOfTheWeekConsole     `json:"Console"`
	Game                 GetAchievementOfTheWeekGame        `json:"Game"`
	UnlocksCount         int                                `json:"UnlocksCount"`
	UnlocksHardcoreCount int                                `json:"UnlocksHardcoreCount"`
	TotalPlayers         int                                `json:"TotalPlayers"`
	Unlocks              []GetAchievementOfTheWeekUnlock    `json:"Unlocks"`
}

type GetAchievementOfTheWeekAchievement struct {
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

type GetAchievementOfTheWeekConsole struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
}

type GetAchievementOfTheWeekGame struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
}

type GetAchievementOfTheWeekUnlock struct {
	User             string    `json:"User"`
	RAPoints         int       `json:"RAPoints"`
	RASoftcorePoints int       `json:"RASoftcorePoints"`
	DateAwarded      time.Time `json:"DateAwarded"`
	HardcoreMode     int       `json:"HardcoreMode"`
}
