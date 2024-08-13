package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/joshraphael/go-retroachievements/models"
)

func unmarshalResponseObject[Obj interface{}](resp *http.Response) (*Obj, error) {
	obj := new(Obj)
	err := json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func unmarshalResponseList[Obj interface{}](resp *http.Response) ([]Obj, error) {
	objs := []Obj{}
	err := json.NewDecoder(resp.Body).Decode(&objs)
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func parseError(resp *http.Response) error {
	respError, err := unmarshalResponseObject[models.ErrorResponse](resp)
	if err != nil {
		return err
	}
	errText := []string{}
	for i := range respError.Errors {
		err := respError.Errors[i]
		errText = append(errText, fmt.Sprintf("[%d] %s", err.Status, err.Title))
	}
	return fmt.Errorf("error responses: %s", strings.Join(errText, ", "))
}

// ResponseObject parses a http response and converts it to a generic object
func ResponseObject[Obj interface{}](resp *http.Response) (*Obj, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return unmarshalResponseObject[Obj](resp)
	case http.StatusNotFound:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, parseError(resp)
	default:
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
}

// ResponseList parses a http response and converts it to a generic list of objects
func ResponseList[Obj interface{}](resp *http.Response) ([]Obj, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return unmarshalResponseList[Obj](resp)
	case http.StatusUnauthorized:
		return nil, parseError(resp)
	default:
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
}
