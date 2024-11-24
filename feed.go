package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetRecentGameAwards gets all recently granted game awards across the site's userbase.
func (c *Client) GetRecentGameAwards(params models.GetRecentGameAwardsParameters) (*models.GetRecentGameAwards, error) {
	details := []raHttp.RequestDetail{
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetRecentGameAwards.php"),
		raHttp.APIToken(c.Secret),
	}
	if params.StartingDate != nil {
		details = append(details, raHttp.D(*params.StartingDate))
	}
	if params.Count != nil {
		details = append(details, raHttp.C(*params.Count))
	}
	if params.Offset != nil {
		details = append(details, raHttp.O(*params.Offset))
	}
	if params.IncludePartialAwards != nil {
		beatenSoftcore := params.IncludePartialAwards.BeatenSoftcore
		beatenHardcore := params.IncludePartialAwards.BeatenHardcore
		completed := params.IncludePartialAwards.Completed
		mastered := params.IncludePartialAwards.Mastered
		if beatenSoftcore || beatenHardcore || completed || mastered {
			k := []string{}
			if beatenSoftcore {
				k = append(k, "beaten-softcore")
			}
			if beatenHardcore {
				k = append(k, "beaten-hardcore")
			}
			if beatenHardcore {
				k = append(k, "completed")
			}
			if beatenHardcore {
				k = append(k, "mastered")
			}
			details = append(details, raHttp.K(k))
		}
	}
	r, err := c.do(details...)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetRecentGameAwards](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}

// GetActiveClaims gets information about all active set claims (max: 1000).
func (c *Client) GetActiveClaims(params models.GetActiveClaimsParameters) ([]models.GetActiveClaims, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetActiveClaims.php"),
		raHttp.APIToken(c.Secret),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseList[models.GetActiveClaims](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response list: %w", err)
	}
	return resp, nil
}
