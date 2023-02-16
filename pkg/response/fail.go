package response

import (
	"github.com/rodericusifo/echo-template/pkg/response/structs"
)

func ResponseFail(message string, err any) structs.ResponseError {
	return structs.ResponseError{
		Success: false,
		Message: message,
		Error:   err,
	}
}
