package structs

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}
