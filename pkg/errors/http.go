package errors

// HTTPError is a custom error type that includes an HTTP status code,
// error message, and details about the error.
type HTTPError struct {
	Code    int
	Message string
	Details interface{}
}

// Error satisfies the error interface.
func (e HTTPError) Error() string {
	return e.Message
}