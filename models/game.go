package models

// GetGameParameters contains the parameters needed for getting a games summary
type GetGameParameters struct {
	// The target game ID
	GameID int
}

type GetGame struct {
	Title        string   `json:"Title"`
	ConsoleID    int      `json:"ConsoleID"`
	ForumTopicID *int     `json:"ForumTopicID"`
	Flags        *int     `json:"Flags"`
	ImageIcon    string   `json:"ImageIcon"`
	ImageTitle   string   `json:"ImageTitle"`
	ImageIngame  string   `json:"ImageIngame"`
	ImageBoxArt  string   `json:"ImageBoxArt"`
	Publisher    string   `json:"Publisher"`
	Developer    string   `json:"Developer"`
	Genre        string   `json:"Genre"`
	Released     DateOnly `json:"Released"`
	GameTitle    string   `json:"GameTitle"`
	ConsoleName  string   `json:"ConsoleName"`
	Console      string   `json:"Console"`
	GameIcon     string   `json:"GameIcon"`
}
