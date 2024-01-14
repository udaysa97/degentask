package types

type ErrorResponse struct {
	Code      int    `json:"code,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type ResponseDTO[T any] struct {
	Status  Status         `json:"status"`
	Success bool           `json:"success"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Result  *T             `json:"data,omitempty"`
}
