package models

type GetCodeNotesParameters struct {
	// The target game ID
	GameID int
}

type GetCodeNotes struct {
	Success   bool                   `json:"Success"`
	CodeNotes []GetCodeNotesCodeNote `json:"CodeNotes"`
}

type GetCodeNotesCodeNote struct {
	User    string `json:"User"`
	Address string `json:"Address"`
	Note    string `json:"Note"`
}
