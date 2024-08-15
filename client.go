// Package retroachievements contains all call handlers to the retro achievements API
package retroachievements

import (
	"fmt"
	"io"
	"net/http"

	raHttp "github.com/joshraphael/go-retroachievements/http"
)

const (
	RetroAchievementHost = "https://retroachievements.org"
)

type Client struct {
	Host       string
	Secret     string
	HttpClient *http.Client
}

type ClientDetail interface {
	detail(c *Client)
}

type clientDetailInstance struct {
	fn func(c *Client)
}

func (cdi *clientDetailInstance) detail(c *Client) {
	cdi.fn(c)
}

func clientDetailFn(fn func(c *Client)) ClientDetail {
	return &clientDetailInstance{fn: fn}
}

// HttpClient overrides the default http client used
func HttpClient(httpClient *http.Client) ClientDetail {
	return clientDetailFn(func(c *Client) {
		c.HttpClient = httpClient
	})
}

// NewClient makes a new client using the default retroachievement host
func NewClient(secret string) *Client {
	return New(RetroAchievementHost, secret)
}

// New creates a new client for a given hostname
func New(host string, secret string, details ...ClientDetail) *Client {
	client := &Client{
		Host:   host,
		Secret: secret,
		HttpClient: &http.Client{
			Transport: http.DefaultTransport,
		},
	}
	for _, detail := range details {
		detail.detail(client)
	}
	return client
}

func (c *Client) do(details ...raHttp.RequestDetail) (*raHttp.Response, error) {
	r := raHttp.NewRequest(c.Host, details...)

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
