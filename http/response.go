package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/joshraphael/go-retroachievements/models"
)

type Response struct {
	StatusCode int
	Data       []byte
}

func checkNotFoundResponse(data []byte) bool {
	l, err := unmarshalResponseList[interface{}](data)
	if err != nil {
		return false
	}
	return len(l) == 0
}

func unmarshalResponseObject[Obj interface{}](data []byte) (*Obj, error) {
	// for some reason this api returns a 200 and empty list on some endpoints when the resource is not found
	emptyResp := checkNotFoundResponse(data)
	if emptyResp {
		return nil, nil
	}
	body := io.NopCloser(bytes.NewReader(data))
	defer body.Close()
	obj := new(Obj)
	err := json.NewDecoder(body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func unmarshalResponseList[Obj interface{}](data []byte) ([]Obj, error) {
	body := io.NopCloser(bytes.NewReader(data))
	defer body.Close()
	objs := []Obj{}
	err := json.NewDecoder(body).Decode(&objs)
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func parseError(data []byte) error {
	respError, err := unmarshalResponseObject[models.ErrorResponse](data)
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
func ResponseObject[Obj interface{}](resp *Response) (*Obj, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return unmarshalResponseObject[Obj](resp.Data)
	case http.StatusNotFound:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, parseError(resp.Data)
	default:
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
}

// ResponseList parses a http response and converts it to a generic list of objects
func ResponseList[Obj interface{}](resp *Response) ([]Obj, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return unmarshalResponseList[Obj](resp.Data)
	case http.StatusUnauthorized:
		return nil, parseError(resp.Data)
	default:
		return nil, fmt.Errorf("unknown error returned: %d", resp.StatusCode)
	}
}
