package models

type ErrorResponse struct {
	Message string    `json:"message"`
	Errors  []RAError `json:"errors"`
}

type RAError struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
}
