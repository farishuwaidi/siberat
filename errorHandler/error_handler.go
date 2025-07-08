package errorhandler

import (
	"Siberat/helper"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError interface {
	error
	Code() string
	Status() int
	Message() string
}

type BaseError struct {
	ErrMessage string
	ErrCode string
	HTTPStatus int
}

func (e *BaseError) Error() string {return e.ErrMessage}
func (e *BaseError) Code() string {return e.ErrCode}
func (e *BaseError) Status() int {return e.HTTPStatus}
func (e *BaseError) Message() string {return e.ErrMessage}

func HandlerError(c *gin.Context, err error, param any) {
	var customErr CustomError
	if errors.As(err, &customErr) {
		c.JSON( customErr.Status(), helper.ErrorResponseWithParam(
			customErr.Code(),
			customErr.Message(),
			param,
		))
		return
	}

	c.JSON(http.StatusInternalServerError, helper.ErrorResponseWithParam(
		"0500",
		"internal server error",
		param,
	))
}