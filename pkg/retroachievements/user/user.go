package user

type User struct {
	Host   string
	secret string
}

func New(host string, secret string) *User {
	return &User{
		Host:   host,
		secret: secret,
	}
}
