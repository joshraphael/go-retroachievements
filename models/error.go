package models

// ErrorResponse is the generic error response from the RetroAchievement API
type ErrorResponse struct {
	// Readable problem returned from the API
	Message string `json:"message"`

	// Array of specific errors
	Errors []ErrorDetail `json:"errors"`
}

// ErrorDetail describes an error reported back from the RetroAchievement API
type ErrorDetail struct {
	// HTTP response code status
	Status int `json:"status"`

	// Readable message of the response status
	Code string `json:"code"`

	// Readable problem for this error
	Title string `json:"title"`
}

type UnprocessableErrorResponse struct {
	// Readable problem returned from the API
	Message string `json:"message"`

	// Map of specific errors
	Errors map[string][]string `json:"errors"`
}
