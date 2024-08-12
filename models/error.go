package models

type ErrorResponse struct {
	Message string        `json:"message"`
	Errors  []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
}
