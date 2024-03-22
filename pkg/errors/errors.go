package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

// StatusCode extracts the HTTP status code from an error.
func StatusCode(err error) int {
	// Check if the error is an HTTP error
	if e, ok := errors.Cause(err).(HTTPError); ok {
		// Return the HTTP status code
		return e.Code
	}

	// Return the default HTTP status code
	return http.StatusInternalServerError
}

// Response extracts the error message and details from an error.
func Response(err error) map[string]interface{} {
	// Create an empty response object
	response := make(map[string]interface{})

	// Check if the error is an HTTP error
	if e, ok := errors.Cause(err).(HTTPError); ok {
		// Add the error message and details to the response
		response["message"] = e.Message
		response["details"] = e.Details
	} else {
		// Add the default error message to the response
		response["message"] = "Internal server error"
	}

	return response
}