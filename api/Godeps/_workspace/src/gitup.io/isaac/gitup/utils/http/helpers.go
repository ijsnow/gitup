package http

import "net/http"

// ErrorResponse formats an error response
func ErrorResponse(errors []string) map[string][]string {
	return map[string][]string{"errors": errors}
}

// Unauthorized is a helper func to return a 401
func Unauthorized() (int, interface{}) {
	return http.StatusUnauthorized, "Unauthorized"
}

// BadRequest sends a bad request response
func BadRequest() (int, interface{}) {
	return http.StatusBadRequest, "Bad Request"
}

// InternalError is a helper func to return an internal server error
func InternalError() (int, interface{}) {
	return http.StatusInternalServerError, "Internal Server Error"
}

// NotFound is a helper func to return a 404 not found error
func NotFound() (int, interface{}) {
	return http.StatusNotFound, "Not found"
}
