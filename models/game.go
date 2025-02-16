package models

import "time"

// GetGameParameters contains the parameters needed for getting a games summary
type GetGameParameters struct {
	// The target game ID
	GameID int
}

type GetGame struct {
	Title        string    `json:"Title"`
	ConsoleID    int       `json:"ConsoleID"`
	ForumTopicID *int      `json:"ForumTopicID"`
	Flags        *int      `json:"Flags"`
	ImageIcon    string    `json:"ImageIcon"`
	ImageTitle   string    `json:"ImageTitle"`
	ImageIngame  string    `json:"ImageIngame"`
	ImageBoxArt  string    `json:"ImageBoxArt"`
	Publisher    string    `json:"Publisher"`
	Developer    string    `json:"Developer"`
	Genre        string    `json:"Genre"`
	Released     *DateOnly `json:"Released"`
	GameTitle    string    `json:"GameTitle"`
	ConsoleName  string    `json:"ConsoleName"`
	Console      string    `json:"Console"`
	GameIcon     string    `json:"GameIcon"`
}

// GetGameExtentedParameters contains the parameters needed for getting a games summary
type GetGameExtentedParameters struct {
	// The target game ID
	GameID int

	// [Optional] Get unofficial achievements (default: false)
	Unofficial *bool
}

type GetGameExtented struct {
	Title                      string                             `json:"Title"`
	ConsoleID                  int                                `json:"ConsoleID"`
	ForumTopicID               *int                               `json:"ForumTopicID"`
	Flags                      *int                               `json:"Flags"`
	ImageIcon                  string                             `json:"ImageIcon"`
	ImageTitle                 string                             `json:"ImageTitle"`
	ImageIngame                string                             `json:"ImageIngame"`
	ImageBoxArt                string                             `json:"ImageBoxArt"`
	Publisher                  string                             `json:"Publisher"`
	Developer                  string                             `json:"Developer"`
	Genre                      string                             `json:"Genre"`
	Released                   *DateOnly                          `json:"Released"`
	ID                         int                                `json:"ID"`
	IsFinal                    bool                               `json:"IsFinal"`
	RichPresencePatch          string                             `json:"RichPresencePatch"`
	GuideURL                   *string                            `json:"GuideURL"`
	Updated                    *time.Time                         `json:"Updated"`
	ConsoleName                string                             `json:"ConsoleName"`
	ParentGameID               *int                               `json:"ParentGameID"`
	NumDistinctPlayers         int                                `json:"NumDistinctPlayers"`
	NumAchievements            int                                `json:"NumAchievements"`
	Achievements               map[int]GetGameExtentedAchievement `json:"Achievements"`
	Claims                     []GetGameExtentedClaim             `json:"Claims"`
	NumDistinctPlayersCasual   int                                `json:"NumDistinctPlayersCasual"`
	NumDistinctPlayersHardcore int                                `json:"NumDistinctPlayersHardcore"`
}

type GetGameExtentedAchievement struct {
	Title              string    `json:"Title"`
	Description        string    `json:"Description"`
	Points             int       `json:"Points"`
	TrueRatio          int       `json:"TrueRatio"`
	Author             string    `json:"Author"`
	ID                 int       `json:"ID"`
	NumAwarded         int       `json:"NumAwarded"`
	NumAwardedHardcore int       `json:"NumAwardedHardcore"`
	DateModified       DateTime  `json:"DateModified"`
	DateCreated        DateTime  `json:"DateCreated"`
	BadgeName          string    `json:"BadgeName"`
	DisplayOrder       int       `json:"DisplayOrder"`
	MemAddr            string    `json:"MemAddr"`
	Type               string    `json:"type"`
	DateEarnedHardcore *DateTime `json:"DateEarnedHardcore"`
	DateEarned         *DateTime `json:"DateEarned"`
}

type GetGameExtentedClaim struct {
	User       string   `json:"User"`
	SetType    int      `json:"SetType"`
	GameID     int      `json:"GameID"`
	ClaimType  int      `json:"ClaimType"`
	Created    DateTime `json:"Created"`
	Expiration DateTime `json:"Expiration"`
}

type GetGameHashesParameters struct {
	// The target game ID
	GameID int
}

type GetGameHashes struct {
	Results []GetGameHashesResult `json:"Results"`
}

type GetGameHashesResult struct {
	Name     string   `json:"Name"`
	MD5      string   `json:"MD5"`
	Labels   []string `json:"Labels"`
	PatchUrl *string  `json:"PatchUrl"`
}

type GetAchievementCountParameters struct {
	// The target game ID
	GameID int
}

type GetAchievementCount struct {
	GameID         int   `json:"GameID"`
	AchievementIDs []int `json:"AchievementIDs"`
}

type GetAchievementDistributionParameters struct {
	// The target game ID
	GameID int

	// [Optional] Only query hardcore unlocks (default: false)
	Hardcore *bool

	// [Optional] Get unofficial achievements (default: false)
	Unofficial *bool
}

type GetAchievementDistribution map[string]int

type GetGameRankAndScoreParameters struct {
	// The target game ID
	GameID int

	// [Optional] Return the latest masters (dafualt: false)
	LatestMasters *bool
}

type GetGameRankAndScore struct {
	User            string   `json:"User"`
	NumAchievements int      `json:"NumAchievements"`
	TotalScore      int      `json:"TotalScore"`
	LastAward       DateTime `json:"LastAward"`
}
