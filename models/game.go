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
	Updated                    time.Time               `json:"Updated"`
	ConsoleName                string                  `json:"ConsoleName"`
	ParentGameID               *int                    `json:"ParentGameID"`
	NumDistinctPlayers         int                     `json:"NumDistinctPlayers"`
	NumAchievements            int                     `json:"NumAchievements"`
	Achievements               map[int]GameAchievement `json:"Achievements"`
	Claims                     []Claim                 `json:"Claims"`
	NumDistinctPlayersCasual   int                     `json:"NumDistinctPlayersCasual"`
	NumDistinctPlayersHardcore int                     `json:"NumDistinctPlayersHardcore"`
}
