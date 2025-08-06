package errorhandler

import "net/http"

type NotFoundError struct {
	Code  string
	MessageText string
}

type BadRequestError struct {
	Code string
	MessageText string
}

type InternalServerError struct {
	Code string
	MessageText string
}

type UnAuthorizedError struct {
	CodeErr any
	MessageText string
}

func (c *BadRequestError) Error() string { return c.MessageText }
func (c *BadRequestError) CodeErr() string  { return c.Code}
func (c *BadRequestError) Status() int   { return http.StatusBadRequest }
func (c *BadRequestError) Message() string { return c.MessageText }

func (c *NotFoundError) Error() string { return c.MessageText }
func (c *NotFoundError) CodeErr() string { return c.Code}
func (c *NotFoundError) Status() int { return http.StatusNotFound}
func (c *NotFoundError) Message() string { return c.MessageText}

func (c *InternalServerError) Error() string { return c.MessageText }
func (c *InternalServerError) CodeErr() string { return c.Code}
func (c *InternalServerError) Status() int { return http.StatusNotFound}
func (c *InternalServerError) Message() string { return c.MessageText}

func (c *UnAuthorizedError) Error() string { return c.MessageText }
func (c *UnAuthorizedError) Code() string {
	if c.CodeErr == nil {
		return "0040"
	}
	if codeStr, ok := c.CodeErr.(string); ok {
		return codeStr
	}
	return "0040"
}
func (c *UnAuthorizedError) Status() int { return http.StatusNotFound}
func (c *UnAuthorizedError) Message() string { return c.MessageText}