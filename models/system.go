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
