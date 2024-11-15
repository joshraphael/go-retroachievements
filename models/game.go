package models

import "time"

type Game struct {
	Title        string        `json:"Title"`
	ConsoleID    int           `json:"ConsoleID"`
	ForumTopicID *int          `json:"ForumTopicID"`
	Flags        *int          `json:"Flags"`
	ImageIcon    string        `json:"ImageIcon"`
	ImageTitle   string        `json:"ImageTitle"`
	ImageIngame  string        `json:"ImageIngame"`
	ImageBoxArt  string        `json:"ImageBoxArt"`
	Publisher    string        `json:"Publisher"`
	Developer    string        `json:"Developer"`
	Genre        string        `json:"Genre"`
	Released     LongMonthDate `json:"Released"`
}

type GameInfo struct {
	Game
	GameTitle   string `json:"GameTitle"`
	ConsoleName string `json:"ConsoleName"`
	Console     string `json:"Console"`
	GameIcon    string `json:"GameIcon"`
}

type ExtentedGameInfo struct {
	Game
	ID                         int                     `json:"ID"`
	IsFinal                    int                     `json:"IsFinal"`
	RichPresencePatch          string                  `json:"RichPresencePatch"`
	GuideURL                   *string                 `json:"GuideURL"`
	Updated                    *time.Time              `json:"Updated,omitempty"`
	ConsoleName                string                  `json:"ConsoleName"`
	ParentGameID               *int                    `json:"ParentGameID"`
	NumDistinctPlayers         int                     `json:"NumDistinctPlayers"`
	NumAchievements            int                     `json:"NumAchievements"`
	Achievements               map[int]GameAchievement `json:"Achievements"`
	Claims                     []Claim                 `json:"Claims,omitempty"`
	NumDistinctPlayersCasual   int                     `json:"NumDistinctPlayersCasual"`
	NumDistinctPlayersHardcore int                     `json:"NumDistinctPlayersHardcore"`
}

type UserGameProgress struct {
	ExtentedGameInfo
	ReleasedAt               *time.Time         `json:"released_at"`
	ReleasedAtGranularity    *string            `json:"released_at_granularity"`
	PlayersTotal             int                `json:"players_total"`
	AchievementsPublished    int                `json:"achievements_published"`
	PointsTotal              int                `json:"points_total"`
	NumAwardedToUser         int                `json:"NumAwardedToUser"`
	NumAwardedToUserHardcore int                `json:"NumAwardedToUserHardcore"`
	UserCompletion           string             `json:"UserCompletion"`
	UserCompletionHardcore   string             `json:"UserCompletionHardcore"`
	HighestAwardKind         *string            `json:"HighestAwardKind"`
	HighestAwardDate         *RFC3339NumColonTZ `json:"HighestAwardDate"`
}

type UserGameRankScore struct {
	User       string    `json:"User"`
	UserRank   int       `json:"UserRank"`
	TotalScore int       `json:"TotalScore"`
	LastAward  *DateTime `json:"LastAward"`
}
