// Package client contains all call handlers to retro achievements
package client

import (
	"fmt"
	"io"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
)

type Client struct {
	host       string
	secret     string
	HttpClient *http.Client
}

func New(host string, secret string) *Client {
	return &Client{
		host:   host,
		secret: secret,
		HttpClient: &http.Client{
			Transport: http.DefaultTransport,
		},
	}
}

func (c *Client) do(details ...raHttp.RequestDetail) (*raHttp.Response, error) {
	r := raHttp.NewRequest(c.host, details...)

	url := r.Host
	if r.Path != "" {
		url = fmt.Sprintf("%s%s", r.Host, r.Path)
	}

	req, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating new http request: %w", err)
	}
	q := req.URL.Query()
	for k, v := range r.Params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &raHttp.Response{
		StatusCode: resp.StatusCode,
		Data:       data,
	}, nil
}
