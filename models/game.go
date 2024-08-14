package models

type Game struct {
	Title        string        `json:"Title"`
	GameTitle    string        `json:"GameTitle"`
	ConsoleID    int           `json:"ConsoleID"`
	ConsoleName  string        `json:"ConsoleName"`
	Console      string        `json:"Console"`
	ForumTopicID int           `json:"ForumTopicID"`
	Flags        int           `json:"Flags"`
	GameIcon     string        `json:"GameIcon"`
	ImageIcon    string        `json:"ImageIcon"`
	ImageTitle   string        `json:"ImageTitle"`
	ImageIngame  string        `json:"ImageIngame"`
	ImageBoxArt  string        `json:"ImageBoxArt"`
	Publisher    string        `json:"Publisher"`
	Developer    string        `json:"Developer"`
	Genre        string        `json:"Genre"`
	Released     LongMonthDate `json:"Released"`
}
