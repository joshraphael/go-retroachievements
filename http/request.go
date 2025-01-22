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
	detail(req *Request)
}

type requestDetailInstance struct {
	fn func(req *Request)
}

func (rdi *requestDetailInstance) detail(req *Request) {
	rdi.fn(req)
}

func requestDetailFn(fn func(req *Request)) RequestDetail {
	return &requestDetailInstance{fn: fn}
}

// Method tells what http verb the request will do
func Method(method string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Method = method
	})
}

// BearerToken adds an authorization header with bearer token
func BearerToken(token string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	})
}

// UserAgent adds an User-Agent header with the package version
func UserAgent(userAgent string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Headers["User-Agent"] = userAgent
	})
}

// A adds a 'a' number to the query parameters
func A(a int) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["a"] = strconv.Itoa(a)
	})
}

// C adds a 'c' number to the query parameters
func C(c int) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["c"] = strconv.Itoa(c)
	})
}

// D adds a 'd' string to the query parameters
func D(d string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["d"] = d
	})
}

// F adds a 'f' number to the query parameters
func F(f int) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["f"] = strconv.Itoa(f)
	})
}

// G adds a 'g' number to the query parameters
func G(g int) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["g"] = strconv.Itoa(g)
	})
}

// H adds a 'h' number to the query parameters
func H(h int) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["h"] = strconv.Itoa(h)
	})
}

// I adds a 'i' string list to the query parameters
func I(i []string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["i"] = strings.Join(i, ",")
	})
}

// K adds a 'k' string list to the query parameters
func K(k []string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["k"] = strings.Join(k, ",")
	})
}

// M adds a 'm' number to the query parameters
func M(m int) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["m"] = strconv.Itoa(m)
	})
}

// O adds a 'o' number to the query parameters
func O(o int) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["o"] = strconv.Itoa(o)
	})
}

// R adds a 'r' string to the query parameters
func R(r string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["r"] = r
	})
}

// T adds a 't' string to the query parameters
func T(t string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["t"] = t
	})
}

// U adds a 'u' string to the query parameters
func U(u string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["u"] = u
	})
}

// Y adds a 'y' string to the query parameters
func Y(y string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Params["y"] = y
	})
}

// Path adds a URL path to the host
func Path(path string) RequestDetail {
	return requestDetailFn(func(req *Request) {
		req.Path = path
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
