package models

type UserRecentlyPlayed struct {
	GetUserProgress
	GameID            int      `json:"GameID"`
	ConsoleID         int      `json:"ConsoleID"`
	ConsoleName       string   `json:"ConsoleName"`
	Title             string   `json:"Title"`
	ImageIcon         string   `json:"ImageIcon"`
	ImageTitle        string   `json:"ImageTitle"`
	ImageIngame       string   `json:"ImageIngame"`
	ImageBoxArt       string   `json:"ImageBoxArt"`
	LastPlayed        DateTime `json:"LastPlayed"`
	AchievementsTotal int      `json:"AchievementsTotal"`
}
