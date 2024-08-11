package models

import "time"

type Profile struct {
	User                string    `json:"User"`
	UserPic             string    `json:"UserPic"`
	MemberSince         time.Time `json:"MemberSince"`
	RichPresenceMsg     string    `json:"RichPresenceMsg"`
	LastGameID          int       `json:"LastGameID"`
	ContribCount        int       `json:"ContribCount"`
	ContribYield        int       `json:"ContribYield"`
	TotalPoints         int       `json:"TotalPoints"`
	TotalSoftcorePoints int       `json:"TotalSoftcorePoints"`
	TotalTruePoints     int       `json:"TotalTruePoints"`
	Permissions         int       `json:"Permissions"`
	Untracked           int       `json:"Untracked"`
	ID                  int       `json:"ID"`
	UserWallActive      int       `json:"UserWallActive"`
	Motto               string    `json:"Motto"`
}
