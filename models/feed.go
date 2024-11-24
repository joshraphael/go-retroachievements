package models

import "time"

type GetRecentGameAwardsParameters struct {
	// [Optional] Starting date (YYYY-MM-DD) (default: now).
	StartingDate *time.Time

	// [Optional] The number of records to return (default: 100, max: 500).
	Count *int

	// [Optional] The number of entries to skip (default: 0).
	Offset *int

	// [Optional] Return partial list of awards based on type (default: return everything).
	IncludePartialAwards *GetRecentGameAwardsParametersPartialAwards
}

type GetRecentGameAwardsParametersPartialAwards struct {
	// Include beaten softcore awards
	BeatenSoftcore bool

	// Include beaten hardcore awards
	BeatenHardcore bool

	// Include completed game awards
	Completed bool

	// Include mastered game awards
	Mastered bool
}

type GetRecentGameAwards struct {
	Count   int                         `json:"Count"`
	Total   int                         `json:"Total"`
	Results []GetRecentGameAwardsResult `json:"Results"`
}

type GetRecentGameAwardsResult struct {
	User        string            `json:"User"`
	AwardKind   string            `json:"AwardKind"`
	AwardDate   RFC3339NumColonTZ `json:"AwardDate"`
	GameID      int               `json:"GameID"`
	GameTitle   string            `json:"GameTitle"`
	ConsoleID   int               `json:"ConsoleID"`
	ConsoleName string            `json:"ConsoleName"`
}
type GetActiveClaimsParameters struct{}

type GetActiveClaims struct {
	ID          int      `json:"ID"`
	User        string   `json:"User"`
	GameID      int      `json:"GameID"`
	GameTitle   string   `json:"GameTitle"`
	GameIcon    string   `json:"GameIcon"`
	ConsoleID   int      `json:"ConsoleID"`
	ConsoleName string   `json:"ConsoleName"`
	ClaimType   int      `json:"ClaimType"`
	SetType     int      `json:"SetType"`
	Status      int      `json:"Status"`
	Extension   int      `json:"Extension"`
	Special     int      `json:"Special"`
	Created     DateTime `json:"Created"`
	DoneTime    DateTime `json:"DoneTime"`
	Updated     DateTime `json:"Updated"`
	UserIsJrDev int      `json:"UserIsJrDev"`
	MinutesLeft int      `json:"MinutesLeft"`
}

type GetClaimsParametersKind interface {
	GetClaimsParametersKindID() int
}

type GetClaimsParametersKindCompleted struct{}

func (c *GetClaimsParametersKindCompleted) GetClaimsParametersKindID() int {
	return 1
}

type GetClaimsParametersKindDropped struct{}

func (d *GetClaimsParametersKindDropped) GetClaimsParametersKindID() int {
	return 2
}

// NOTE: Expired claims returns a strange format different from the rest, disabling it for now as there is not many available
// type GetClaimsParametersKindExpired struct{}

// func (e *GetClaimsParametersKindExpired) GetClaimsParametersKindID() int {
// 	return 3
// }

type GetClaimsParameters struct {
	// [Optional] The desired claim kind: completed, dropped, or expired (default: completed).
	Kind GetClaimsParametersKind
}

type GetClaims struct {
	ID          int      `json:"ID"`
	User        string   `json:"User"`
	GameID      int      `json:"GameID"`
	GameTitle   string   `json:"GameTitle"`
	GameIcon    string   `json:"GameIcon"`
	ConsoleID   int      `json:"ConsoleID"`
	ConsoleName string   `json:"ConsoleName"`
	ClaimType   int      `json:"ClaimType"`
	SetType     int      `json:"SetType"`
	Status      int      `json:"Status"`
	Extension   int      `json:"Extension"`
	Special     int      `json:"Special"`
	Created     DateTime `json:"Created"`
	DoneTime    DateTime `json:"DoneTime"`
	Updated     DateTime `json:"Updated"`
	UserIsJrDev int      `json:"UserIsJrDev"`
	MinutesLeft int      `json:"MinutesLeft"`
}
