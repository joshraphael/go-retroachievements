// Package retroachievements contains all call handlers to the retro achievements API
package retroachievements

import (
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"sync"

	raHttp "github.com/joshraphael/go-retroachievements/http"
)

const (
	RetroAchievementHost = "https://retroachievements.org"
)

type ClientConfig struct {
	Host          string
	UserAgent     string
	APISecret     string
	ConnectConfig *ClientConnectConfig
}

type ClientConnectConfig struct {
	ConnectSecret   string
	ConnectUsername string
}

type Client struct {
	UserAgent       string
	Host            string
	APISecret       string
	ConnectSecret   string
	ConnectUsername string
	HttpClient      *http.Client
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

var version = sync.OnceValue(func() string {
	libraryVersion := "v0.0.0"
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range buildInfo.Deps {
			if dep.Path == "github.com/joshraphael/go-retroachievements" {
				libraryVersion = dep.Version
				break
			}
		}
	}

	return "go-retroachievements/" + libraryVersion
})

// NewClient makes a new client using the default retroachievement host
func NewClient(secret string) *Client {
	return New(ClientConfig{
		Host:      RetroAchievementHost,
		UserAgent: version(),
		APISecret: secret,
	})
}

// New creates a new client for a given hostname
func New(config ClientConfig, details ...ClientDetail) *Client {
	client := &Client{
		UserAgent: config.UserAgent,
		Host:      config.Host,
		APISecret: config.APISecret,
		HttpClient: &http.Client{
			Transport: http.DefaultTransport,
		},
	}
	if config.ConnectConfig != nil && len(config.ConnectConfig.ConnectSecret) > 0 && len(config.ConnectConfig.ConnectUsername) > 0 {
		client.ConnectSecret = config.ConnectConfig.ConnectSecret
		client.ConnectUsername = config.ConnectConfig.ConnectUsername
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
