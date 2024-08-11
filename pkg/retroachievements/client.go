package retroachievements

import "github.com/joshraphael/go-retroachievements/pkg/retroachievements/user"

type Client struct {
	host   string
	secret string
	User   *user.User
}

func New(host string, secret string) *Client {
	user := user.New(host, secret)
	return &Client{
		host:   host,
		secret: secret,
		User:   user,
	}
}
