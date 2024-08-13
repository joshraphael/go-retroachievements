package models

// Profile describes elements of a users profile
type Profile struct {
	User                string   `json:"User"`
	UserPic             string   `json:"UserPic"`
	MemberSince         DateTime `json:"MemberSince"`
	RichPresenceMsg     string   `json:"RichPresenceMsg"`
	LastGameID          int      `json:"LastGameID"`
	ContribCount        int      `json:"ContribCount"`
	ContribYield        int      `json:"ContribYield"`
	TotalPoints         int      `json:"TotalPoints"`
	TotalSoftcorePoints int      `json:"TotalSoftcorePoints"`
	TotalTruePoints     int      `json:"TotalTruePoints"`
	Permissions         int      `json:"Permissions"`
	Untracked           int      `json:"Untracked"`
	ID                  int      `json:"ID"`
	UserWallActive      bool     `json:"UserWallActive"`
	Motto               string   `json:"Motto"`
}
