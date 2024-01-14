package constants

import "net/http"

const (
	BAD_REUQEST_ERROR    = "BAD_REQUEST"
	DATA_NOT_FOUND_ERROR = "DATA_NOT_FOUND"
	METHOD_NOT_ALLOWED   = "METHOD_NOT_ALLOWED"
)

type ErrorType struct {
	ErrorCode  string
	HttpStatus int
}

var ERROR_TYPES = map[string]ErrorType{
	BAD_REUQEST_ERROR: {
		ErrorCode:  "TECH-001",
		HttpStatus: http.StatusBadRequest,
	},
	DATA_NOT_FOUND_ERROR: {
		ErrorCode:  "DATA-001",
		HttpStatus: http.StatusNotFound,
	},
	METHOD_NOT_ALLOWED: {
		ErrorCode:  "TECH-000",
		HttpStatus: http.StatusMethodNotAllowed,
	},
}
