package validator

import (
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

type IRequestBody interface {
	CustomValidateRequestBody() error
}

type IRequestQuery interface {
	CustomValidateRequestQuery() error
}

type IRequestParams interface {
	CustomValidateRequestParams() error
}

type IRequestUser interface {
	CustomValidateRequestUser() error
}

type IPayload interface {
	CustomValidatePayload() error
}

func ValidateRequestBody(ctx echo.Context, req IRequestBody) error {
	if err := dBinder.BindBody(ctx, req); err != nil {
		log.WithFields(log.Fields{
			"message": "bind request body fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST BODY]")
		return err
	}
	if err := cValidator.Validate(req); err != nil {
		log.WithFields(log.Fields{
			"message": "validate request body fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST BODY]")
		return err
	}
	if err := req.CustomValidateRequestBody(); err != nil {
		return err
	}
	return nil
}

func ValidateRequestQuery(ctx echo.Context, req IRequestQuery) error {
	if err := dBinder.BindQueryParams(ctx, req); err != nil {
		log.WithFields(log.Fields{
			"message": "bind request query fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST QUERY]")
		return err
	}
	if err := cValidator.Validate(req); err != nil {
		log.WithFields(log.Fields{
			"message": "validate request query fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST QUERY]")
		return err
	}
	if err := req.CustomValidateRequestQuery(); err != nil {
		return err
	}
	return nil
}

func ValidateRequestParams(ctx echo.Context, req IRequestParams) error {
	if err := dBinder.BindPathParams(ctx, req); err != nil {
		log.WithFields(log.Fields{
			"message": "bind request params fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST PARAMS]")
		return err
	}
	if err := cValidator.Validate(req); err != nil {
		log.WithFields(log.Fields{
			"message": "validate request params fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST PARAMS]")
		return err
	}
	if err := req.CustomValidateRequestParams(); err != nil {
		return err
	}
	return nil
}

func ValidateRequestUser(req IRequestUser) error {
	if err := cValidator.Validate(req); err != nil {
		log.WithFields(log.Fields{
			"message": "validate request user fail",
			"detail":  err,
		}).Errorln("[VALIDATE REQUEST USER]")
		return err
	}
	if err := req.CustomValidateRequestUser(); err != nil {
		return err
	}
	return nil
}

func ValidatePayload(payload IPayload) error {
	if err := cValidator.Validate(payload); err != nil {
		log.WithFields(log.Fields{
			"message": "validate payload fail",
			"detail":  err,
		}).Errorln("[VALIDATE PAYLOAD]")
		return err
	}
	if err := payload.CustomValidatePayload(); err != nil {
		return err
	}
	return nil
}
