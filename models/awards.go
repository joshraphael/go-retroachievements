package models

type UserAwards struct {
	TotalAwardsCount          int     `json:"TotalAwardsCount"`
	HiddenAwardsCount         int     `json:"HiddenAwardsCount"`
	MasteryAwardsCount        int     `json:"MasteryAwardsCount"`
	CompletionAwardsCount     int     `json:"CompletionAwardsCount"`
	BeatenHardcoreAwardsCount int     `json:"BeatenHardcoreAwardsCount"`
	BeatenSoftcoreAwardsCount int     `json:"BeatenSoftcoreAwardsCount"`
	EventAwardsCount          int     `json:"EventAwardsCount"`
	SiteAwardsCount           int     `json:"SiteAwardsCount"`
	VisibleUserAwards         []Award `json:"VisibleUserAwards"`
}

type Award struct {
	AwardedAt      RFC3339NumColonTZ `json:"AwardedAt"`
	AwardType      string            `json:"AwardType"`
	AwardData      int               `json:"AwardData"`
	AwardDataExtra int               `json:"AwardDataExtra"`
	DisplayOrder   int               `json:"DisplayOrder"`
	Title          *string           `json:"Title"`
	ConsoleID      *int              `json:"ConsoleID"`
	ConsoleName    *string           `json:"ConsoleName"`
	Flags          *int              `json:"Flags"`
	ImageIcon      *string           `json:"ImageIcon"`
}
