package custom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/pkg/response"

	log "github.com/sirupsen/logrus"
)

func CustomHTTPErrorHandler(err error, ctx echo.Context) {
	log.WithFields(log.Fields{
		"type":   fmt.Sprintf("%T", err),
		"detail": err,
	}).Errorln("[CUSTOM HTTP ERROR HANDLER]")
	he, ok := err.(*echo.HTTPError)
	if ok {
		ctx.JSON(he.Code, response.ResponseFail(fmt.Sprint(he.Message), he.Internal))
		return
	}
	ve, ok := err.(validator.ValidationErrors)
	if ok {
		type ErrorResponse struct {
			FailedField string `json:"failed_field"`
			Tag         string `json:"tag"`
			Error       string `json:"error"`
		}
		var errors []*ErrorResponse
		for _, err := range ve {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Error = err.Error()
			errors = append(errors, &element)
		}
		ctx.JSON(http.StatusBadRequest, response.ResponseFail("validation error", errors))
		return
	}
	me, ok := err.(*json.MarshalerError)
	if ok {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResponseFail(me.Error(), me.Unwrap()))
		return
	}
	re, ok := err.(runtime.Error)
	if ok {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResponseFail(re.Error(), re))
		return
	}
	ctx.JSON(http.StatusInternalServerError, response.ResponseFail(err.Error(), err))
}
