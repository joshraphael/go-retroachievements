package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

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

func (user *User) GetUserProfile(username string) (*models.Profile, error) {
	u, err := url.Parse(user.Host + "/API/API_GetUserProfile.php")
	if err != nil {
		return nil, fmt.Errorf("parsing GetUserProfile url: %w", err)
	}
	q := u.Query()
	q.Set("y", user.secret)
	q.Set("u", username)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("calling GetUserProfile: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode == http.StatusUnauthorized {
		var respError models.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&respError)
		if err != nil {
			return nil, fmt.Errorf("decoding response error body: %w", err)
		}
		errText := []string{}
		for i := range respError.Errors {
			err := respError.Errors[i]
			errText = append(errText, fmt.Sprintf("[%d] %s", err.Status, err.Title))
		}
		return nil, fmt.Errorf("calling get user profile: %s", strings.Join(errText, ", "))
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
	var profile rawProfile
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		return nil, fmt.Errorf("decoding response body profile: %w", err)
	}
	p, err := profile.ToProfile()
	if err != nil {
		return nil, fmt.Errorf("converting response to profile: %w", err)
	}
	return p, nil
}
