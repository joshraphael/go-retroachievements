package http

import (
	"fmt"
	"strconv"
	"time"
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

// Username adds the username to the query parameters
func Username(username string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["u"] = username
	})
}

// LookbackMinutes adds the lookback minutes to the query parameters
func LookbackMinutes(minutes int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["m"] = strconv.Itoa(minutes)
	})
}

// FromTime adds a start time to the query parameters in unix seconds
func FromTime(t time.Time) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["f"] = strconv.Itoa(int(t.Unix()))
	})
}

// ToTime adds a end time to the query parameters in unix seconds
func ToTime(t time.Time) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["t"] = strconv.Itoa(int(t.Unix()))
	})
}

// Date adds a string date (ignoring timestamp) to the query parameters
func Date(t time.Time) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["d"] = t.Format(time.DateOnly)
	})
}

// ID adds a target game id to the query parameters
func ID(id int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["i"] = strconv.Itoa(id)
	})
}

// GameID adds a target game id to the query parameters
func GameID(gameId int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["g"] = strconv.Itoa(gameId)
	})
}

// AwardMetadata adds a target game id to the query parameters
func AwardMetadata(awardMetadata bool) RequestDetail {
	return requestDetailFn(func(r *Request) {
		a := "0"
		if awardMetadata {
			a = "1"
		}
		r.Params["a"] = a
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
