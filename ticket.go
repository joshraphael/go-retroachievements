package retroachievements

import (
	"fmt"
	"net/http"
	"strconv"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetTicketByID gets ticket metadata information about a single achievement ticket, targeted by its ticket ID.
func (c *Client) GetTicketByID(params models.GetTicketByIDParameters) (*models.GetTicketByID, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.Path("/API/API_GetTicketData.php"),
		raHttp.APIToken(c.Secret),
		raHttp.I([]string{
			strconv.Itoa(params.TicketID),
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetTicketByID](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
