package models

import "time"

type GetCommentsType interface {
	GetCommentsType() int
}

type GetCommentsGame struct {
	GameID int
}

func (g GetCommentsGame) GetCommentsType() int {
	return 1
}

type GetCommentsAchievement struct {
	AchievementID int
}

func (a GetCommentsAchievement) GetCommentsType() int {
	return 2
}

type GetCommentsUser struct {
	Username string
}

func (u GetCommentsUser) GetCommentsType() int {
	return 3
}

type GetCommentsParameters struct {
	Type   GetCommentsType
	Count  *int
	Offset *int
}

type GetComments struct {
	Count   int                 `json:"Count"`
	Total   int                 `json:"Total"`
	Results []GetCommentsResult `json:"Results"`
}

type GetCommentsResult struct {
	User        string    `json:"User"`
	Submitted   time.Time `json:"Submitted"`
	CommentText string    `json:"CommentText"`
}
