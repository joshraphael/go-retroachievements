package models

type GetTicketByIDParameters struct {
	// The target ticket ID
	TicketID int
}

type GetTicketByID struct {
	ID                     int       `json:"ID"`
	AchievementID          int       `json:"AchievementID"`
	AchievementTitle       string    `json:"AchievementTitle"`
	AchievementDesc        string    `json:"AchievementDesc"`
	AchievementType        *string   `json:"AchievementType"`
	Points                 int       `json:"Points"`
	BadgeName              string    `json:"BadgeName"`
	AchievementAuthor      string    `json:"AchievementAuthor"`
	GameID                 int       `json:"GameID"`
	ConsoleName            string    `json:"ConsoleName"`
	GameTitle              string    `json:"GameTitle"`
	GameIcon               string    `json:"GameIcon"`
	ReportedAt             DateTime  `json:"ReportedAt"`
	ReportType             int       `json:"ReportType"`
	ReportState            int       `json:"ReportState"`
	Hardcore               *int      `json:"Hardcore"`
	ReportNotes            string    `json:"ReportNotes"`
	ReportedBy             string    `json:"ReportedBy"`
	ResolvedAt             *DateTime `json:"ResolvedAt"`
	ResolvedBy             *string   `json:"ResolvedBy"`
	ReportStateDescription string    `json:"ReportStateDescription"`
	ReportTypeDescription  string    `json:"ReportTypeDescription"`
	URL                    string    `json:"URL"`
}

type GetMostTicketedGamesParameters struct {
	// [Optional] The number of records to return (default: 10, max: 100).
	Count *int

	// [Optional] The number of entries to skip (default: 0).
	Offset *int
}

type GetMostTicketedGames struct {
	MostReportedGames []GetMostTicketedGamesMostReportedGame `json:"MostReportedGames"`
	URL               string                                 `json:"URL"`
}

type GetMostTicketedGamesMostReportedGame struct {
	GameID      int    `json:"GameID"`
	GameTitle   string `json:"GameTitle"`
	GameIcon    string `json:"GameIcon"`
	Console     string `json:"Console"`
	OpenTickets int    `json:"OpenTickets"`
}
