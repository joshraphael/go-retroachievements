package user

type User struct {
	Host     string
	Username string
	secret   string
}

func New(host string, username string, secret string) *User {
	return &User{
		Host:     host,
		Username: username,
		secret:   secret,
	}
}
