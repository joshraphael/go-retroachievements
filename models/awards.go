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
