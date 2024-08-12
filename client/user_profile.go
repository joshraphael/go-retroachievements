package client

import (
	"fmt"
	"net/http"
	"time"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

type rawProfile struct {
	User                string `json:"User"`
	UserPic             string `json:"UserPic"`
	MemberSince         string `json:"MemberSince"`
	RichPresenceMsg     string `json:"RichPresenceMsg"`
	LastGameID          int    `json:"LastGameID"`
	ContribCount        int    `json:"ContribCount"`
	ContribYield        int    `json:"ContribYield"`
	TotalPoints         int    `json:"TotalPoints"`
	TotalSoftcorePoints int    `json:"TotalSoftcorePoints"`
	TotalTruePoints     int    `json:"TotalTruePoints"`
	Permissions         int    `json:"Permissions"`
	Untracked           int    `json:"Untracked"`
	ID                  int    `json:"ID"`
	UserWallActive      bool   `json:"UserWallActive"`
	Motto               string `json:"Motto"`
}

func (rp *rawProfile) ToProfile() (*models.Profile, error) {
	t, err := time.Parse(time.DateTime, rp.MemberSince)
	if err != nil {
		return nil, err
	}
	return &models.Profile{
		User:                rp.User,
		UserPic:             rp.UserPic,
		MemberSince:         t,
		RichPresenceMsg:     rp.RichPresenceMsg,
		LastGameID:          rp.LastGameID,
		ContribCount:        rp.ContribCount,
		ContribYield:        rp.ContribYield,
		TotalPoints:         rp.TotalPoints,
		TotalSoftcorePoints: rp.TotalSoftcorePoints,
		TotalTruePoints:     rp.TotalTruePoints,
		Permissions:         rp.Permissions,
		Untracked:           rp.Untracked,
		ID:                  rp.ID,
		UserWallActive:      rp.UserWallActive,
		Motto:               rp.Motto,
	}, nil
}

func (c *Client) GetUserProfile(username string) (*models.Profile, error) {
	resp, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetUserProfile.php"),
		raHttp.APIToken(c.secret),
		raHttp.Username(username),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	defer resp.Body.Close()
	profile, err := raHttp.ResponseObject[rawProfile](resp)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}

	return profile.ToProfile()
}
