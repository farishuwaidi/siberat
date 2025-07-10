package errorhandler

import "net/http"

type NotFoundError struct {
	CodeErr  any
	MessageText string
}

type BadRequestError struct {
	CodeErr any
	MessageText string
}

type InternalServerError struct {
	MessageText string
}

type UnAuthorizedError struct {
	CodeErr any
	MessageText string
}

func (c *BadRequestError) Error() string { return c.MessageText }
func (c *BadRequestError) Code() string  { return c.CodeErr.(string) }
func (c *BadRequestError) Status() int   { return http.StatusBadRequest }
func (c *BadRequestError) Message() string { return c.MessageText }

func (c *NotFoundError) Error() string { return c.MessageText }
func (c *NotFoundError) Code() string { return c.CodeErr.(string) }
func (c *NotFoundError) Status() int { return http.StatusNotFound}
func (c *NotFoundError) Message() string { return c.MessageText}

func (c *InternalServerError) Error() string { return c.MessageText }
func (c *InternalServerError) Code() string { return "0500" }
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