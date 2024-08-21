package models

type Claim struct {
	User       string   `json:"User"`
	SetType    int      `json:"SetType"`
	GameID     int      `json:"GameID"`
	ClaimType  int      `json:"ClaimType"`
	Created    DateTime `json:"Created"`
	Expiration DateTime `json:"Expiration"`
}

type UserClaims struct {
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
