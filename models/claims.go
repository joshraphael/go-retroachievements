package models

type Claim struct {
	User       string   `json:"User"`
	SetType    int      `json:"SetType"`
	GameID     int      `json:"GameID"`
	ClaimType  int      `json:"ClaimType"`
	Created    DateTime `json:"Created"`
	Expiration DateTime `json:"Expiration"`
}
