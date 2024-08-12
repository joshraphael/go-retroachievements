package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/joshraphael/go-retroachievements/models"
)

func UnmarshalResponseObject[A interface{}](resp *http.Response) (*A, error) {
	obj := new(A)
	err := json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func UnmarshalResponseList[A interface{}](resp *http.Response) ([]A, error) {
	obj := []A{}
	err := json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func ResponseObject[A interface{}](resp *http.Response) (*A, error) {
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		return UnmarshalResponseObject[A](resp)
	case http.StatusNotFound:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, parseError(resp)
	default:
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
}

func ResponseList[A interface{}](resp *http.Response) ([]A, error) {
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		return UnmarshalResponseList[A](resp)
	case http.StatusUnauthorized:
		return nil, parseError(resp)
	default:
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
}

func parseError(resp *http.Response) error {
	respError, err := UnmarshalResponseObject[models.ErrorResponse](resp)
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
