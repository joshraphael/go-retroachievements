package models

type GetGameLeaderboardsParameters struct {
	// The target game ID
	GameID int

	// [Optional] The number of records to return (default: 100, max: 500).
	Count *int

	// [Optional] The number of entries to skip (default: 0).
	Offset *int
}

type GetGameLeaderboards struct {
	Count   int                         `json:"Count"`
	Total   int                         `json:"Total"`
	Results []GetGameLeaderboardsResult `json:"Results"`
}

type GetGameLeaderboardsResult struct {
	ID          int                          `json:"ID"`
	RankAsc     bool                         `json:"RankAsc"`
	Title       string                       `json:"Title"`
	Description string                       `json:"Description"`
	Format      string                       `json:"Format"`
	TopEntry    *GetGameLeaderboardsTopEntry `json:"TopEntry"`
}

type GetGameLeaderboardsTopEntry struct {
	User           string `json:"User"`
	Score          int    `json:"Score"`
	FormattedScore string `json:"FormattedScore"`
}
