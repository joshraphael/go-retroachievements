package retroachievements

import (
	"fmt"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
)

// GetCodeNotes gets the list of code notes for a given game.
func (c *Client) GetCodeNotes(params models.GetCodeNotesParameters) (*models.GetCodeNotes, error) {
	r, err := c.do(
		raHttp.Method(http.MethodGet),
		raHttp.UserAgent(c.UserAgent),
		raHttp.Path("/dorequest.php"),
		raHttp.G(params.GameID),
		raHttp.R("codenotes2"),
	)
	if err != nil {
		return nil, fmt.Errorf("calling endpoint: %w", err)
	}
	resp, err := raHttp.ResponseObject[models.GetCodeNotes](r)
	if err != nil {
		return nil, fmt.Errorf("parsing response object: %w", err)
	}
	return resp, nil
}
