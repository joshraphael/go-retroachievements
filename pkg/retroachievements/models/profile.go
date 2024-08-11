package models

import "time"

type Profile struct {
	User                string      `json:"User"`
	UserPic             string      `json:"UserPic"`
	MemberSince         RATimeStamp `json:"MemberSince"`
	RichPresenceMsg     string      `json:"RichPresenceMsg"`
	LastGameID          int         `json:"LastGameID"`
	ContribCount        int         `json:"ContribCount"`
	ContribYield        int         `json:"ContribYield"`
	TotalPoints         int         `json:"TotalPoints"`
	TotalSoftcorePoints int         `json:"TotalSoftcorePoints"`
	TotalTruePoints     int         `json:"TotalTruePoints"`
	Permissions         int         `json:"Permissions"`
	Untracked           int         `json:"Untracked"`
	ID                  int         `json:"ID"`
	UserWallActive      bool        `json:"UserWallActive"`
	Motto               string      `json:"Motto"`
}

type RATimeStamp struct {
	Time time.Time
}

func (t *RATimeStamp) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}
