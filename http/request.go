package http

import (
	"fmt"
	"strconv"
)

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

func Method(method string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Method = method
	})
}

func APIToken(token string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["y"] = token
	})
}

func BearerToken(token string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	})
}

func Username(username string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["u"] = username
	})
}

func LookbackMinutes(minutes int) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Params["m"] = strconv.Itoa(minutes)
	})
}

func Path(path string) RequestDetail {
	return requestDetailFn(func(r *Request) {
		r.Path = path
	})
}

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
