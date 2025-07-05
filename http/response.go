package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Data       []byte
}

func checkNotFoundResponse(data []byte) bool {
	l, err := unmarshalResponseList[any](data)
	if err != nil {
		return false
	}
	return len(l) == 0
}

func unmarshalResponseObject[Obj any](data []byte) (*Obj, error) {
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

func unmarshalResponseList[Obj any](data []byte) ([]Obj, error) {
	body := io.NopCloser(bytes.NewReader(data))
	defer body.Close()
	objs := []Obj{}
	err := json.NewDecoder(body).Decode(&objs)
	if err != nil {
		return nil, err
	}
	return objs, nil
}

// ResponseObject parses a http response and converts it to a generic object
func ResponseObject[Obj any](resp *Response) (*Obj, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return unmarshalResponseObject[Obj](resp.Data)
	case http.StatusNotFound:
		return nil, nil
	default:
		return nil, fmt.Errorf("error code %d returned: %s", resp.StatusCode, string(resp.Data))
	}
}

// ResponseList parses a http response and converts it to a generic list of objects
func ResponseList[Obj any](resp *Response) ([]Obj, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return unmarshalResponseList[Obj](resp.Data)
	default:
		return nil, fmt.Errorf("error code %d returned: %s", resp.StatusCode, string(resp.Data))
	}
}
