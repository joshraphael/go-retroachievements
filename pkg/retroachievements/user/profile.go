package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/joshraphael/go-retroachievements/pkg/retroachievements/user/models"
)

func (user *User) GetUserProfile(username string) (*models.Profile, error) {
	u, err := url.Parse(user.Host + "/API_GetUserProfile.php")
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

	var profile models.Profile
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		return nil, fmt.Errorf("decoding response body profile: %w", err)
	}
	return &profile, nil
}
