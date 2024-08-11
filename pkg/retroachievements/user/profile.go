package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/joshraphael/go-retroachievements/pkg/retroachievements/models"
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
	var profile models.Profile
	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		return nil, fmt.Errorf("decoding response body profile: %w", err)
	}
	return &profile, nil
}
