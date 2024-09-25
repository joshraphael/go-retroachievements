package models

type Progress struct {
	NumPossibleAchievements int `json:"NumPossibleAchievements"`
	PossibleScore           int `json:"PossibleScore"`
	NumAchieved             int `json:"NumAchieved"`
	ScoreAchieved           int `json:"ScoreAchieved"`
	NumAchievedHardcore     int `json:"NumAchievedHardcore"`
	ScoreAchievedHardcore   int `json:"ScoreAchievedHardcore"`
}
