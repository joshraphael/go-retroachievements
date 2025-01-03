package http

import (
	"fmt"
	"strconv"
	"strings"
)

// Request holds values for an http call
type Request struct {
	Host    string
	Path    string
	Method  string
	Params  map[string]string
	Headers map[string]string
}

type RequestDetail interface {
	detail(r *Request)
}

type requestDetailInstance struct {
	fn func(r *Request)
}

func (rdi *requestDetailInstance) detail(r *Request) {
	rdi.fn(r)
}

func requestDetailFn(fn func(r *Request)) RequestDetail {
	return &requestDetailInstance{fn: fn}
}

// Method tells what http verb the request will do
func Method(method string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Method = method
	})
}

// APIToken adds an api token to the query parameters
func APIToken(token string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["y"] = token
	})
}

// BearerToken adds an authorization header with bearer token
func BearerToken(token string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	})
}

// UserAgent adds an User-Agent header with the package version
func UserAgent(userAgent string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Headers["User-Agent"] = userAgent
	})
}

// M adds a 'u' string to the query parameters
func U(u string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["u"] = u
	})
}

// M adds a 'm' number to the query parameters
func M(m int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["m"] = strconv.Itoa(m)
	})
}

// F adds a 'f' number to the query parameters
func F(f int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["f"] = strconv.Itoa(f)
	})
}

// T adds a 't' number to the query parameters
func T(t int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["t"] = strconv.Itoa(t)
	})
}

// D adds a 'd' string to the query parameters
func D(d string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["d"] = d
	})
}

// I adds a 'i' string list to the query parameters
func I(i []string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["i"] = strings.Join(i, ",")
	})
}

// K adds a 'k' string list to the query parameters
func K(k []string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["k"] = strings.Join(k, ",")
	})
}

// G adds a 'g' number to the query parameters
func G(g int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["g"] = strconv.Itoa(g)
	})
}

// C adds a 'c' number to the query parameters
func C(c int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["c"] = strconv.Itoa(c)
	})
}

// O adds a 'o' number to the query parameters
func O(o int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["o"] = strconv.Itoa(o)
	})
}

// A adds a 'a' number to the query parameters
func A(a int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["a"] = strconv.Itoa(a)
	})
}

// H adds a 'h' number to the query parameters
func H(h int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["h"] = strconv.Itoa(h)
	})
}

// Path adds a URL path to the host
func Path(path string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Path = path
	})
}

// NewRequest initializes a http request using a host
func NewRequest(host string, details ...RequestDetail) *Request {
	request := &Request{
		Host:    host,
		Params:  map[string]string{},
		Headers: map[string]string{},
	}
	for _, detail := range details {
		detail.detail(request)
	}
	return request
}
