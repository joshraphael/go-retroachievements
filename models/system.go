package models

type GetConsoleIDsParameters struct {
	// [Optional] Get only active systems, use false for all systems (default: false)
	OnlyActive *bool

	// [Optional] Get only active systems, use false for all system types (Events, Hubs, etc) (default: false)
	OnlyGameSystems *bool
}

type GetConsoleIDs struct {
	ID           int    `json:"ID"`
	Name         string `json:"Name"`
	IconURL      string `json:"IconURL"`
	Active       bool   `json:"Active"`
	IsGameSystem bool   `json:"IsGameSystem"`
}

type GetGameListParameters struct {
	// The target system ID
	SystemID int

	// [Optional] Only return games that have achievements (default: false)
	HasAchievements *bool

	// [Optional] Also return supported hashes for games (default: false)
	IncludeHashes *bool

	// [Optional] The number of records to return
	Count *int

	// [Optional] The number of entries to skip
	Offset *int
}

type GetGameList struct {
	Title           string    `json:"Title"`
	ID              int       `json:"ID"`
	ConsoleID       int       `json:"ConsoleID"`
	ConsoleName     string    `json:"ConsoleName"`
	ImageIcon       string    `json:"ImageIcon"`
	NumAchievements int       `json:"NumAchievements"`
	NumLeaderboards int       `json:"NumLeaderboards"`
	Points          int       `json:"Points"`
	DateModified    *DateTime `json:"DateModified"`
	ForumTopicID    *int      `json:"ForumTopicID"`
	Hashes          []string  `json:"Hashes"`
}
