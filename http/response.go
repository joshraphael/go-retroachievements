package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/joshraphael/go-retroachievements/models"
)

func unmarshalResponseObject[A interface{}](resp *http.Response) (*A, error) {
	obj := new(A)
	err := json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func unmarshalResponseList[A interface{}](resp *http.Response) ([]A, error) {
	obj := []A{}
	err := json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
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

func ResponseObject[A interface{}](resp *http.Response) (*A, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return unmarshalResponseObject[A](resp)
	case http.StatusNotFound:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, parseError(resp)
	default:
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
}

func ResponseList[A interface{}](resp *http.Response) ([]A, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return unmarshalResponseList[A](resp)
	case http.StatusUnauthorized:
		return nil, parseError(resp)
	default:
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
}
