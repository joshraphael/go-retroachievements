package retroachievements

type Client struct {
	host   string
	secret string
}

func New(host string, secret string) *Client {
	return &Client{
		host:   host,
		secret: secret,
	}
}
