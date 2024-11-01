package models

type Progress struct {
	NumPossibleAchievements int `json:"NumPossibleAchievements"`
	PossibleScore           int `json:"PossibleScore"`
	NumAchieved             int `json:"NumAchieved"`
	ScoreAchieved           int `json:"ScoreAchieved"`
	NumAchievedHardcore     int `json:"NumAchievedHardcore"`
	ScoreAchievedHardcore   int `json:"ScoreAchievedHardcore"`
}

type UserRecentlyPlayed struct {
	Progress
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
