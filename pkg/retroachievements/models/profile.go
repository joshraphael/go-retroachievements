package models

import "time"

type Profile struct {
	User                string
	UserPic             string
	MemberSince         time.Time
	RichPresenceMsg     string
	LastGameID          int
	ContribCount        int
	ContribYield        int
	TotalPoints         int
	TotalSoftcorePoints int
	TotalTruePoints     int
	Permissions         int
	Untracked           int
	ID                  int
	UserWallActive      bool
	Motto               string
}
