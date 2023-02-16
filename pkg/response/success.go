package response

import (
	"github.com/rodericusifo/echo-template/pkg/response/structs"
)

func ResponseSuccess[T any](message string, data T, meta *structs.Meta) structs.ResponseData[T] {
	return structs.ResponseData[T]{
		Success: true,
		Message: message,
		Meta:    meta,
		Data:    data,
	}
}
