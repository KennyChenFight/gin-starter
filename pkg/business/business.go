package business

import "errors"

func NewError(businessCode int, httpStatusCode int, message string, reason error) *Error {
	if reason == nil {
		reason = errors.New(message)
	}
	return &Error{BusinessCode: businessCode, Base: Base{HTTPStatusCode: httpStatusCode}, Message: message, Reason: reason}
}

func NewSuccess(httpStatusCode int, response interface{}) *Success {
	return &Success{Base: Base{HTTPStatusCode: httpStatusCode}, Response: response}
}

type Base struct {
	HTTPStatusCode int `json:"-"`
}

type Error struct {
	Base
	BusinessCode     int               `json:"code"`
	Message          string            `json:"message"`
	ValidationErrors map[string]string `json:"validationErrors,omitempty"`
	Reason           error             `json:"-"`
}

func (b *Error) Error() string {
	return b.Reason.Error()
}

type Success struct {
	Base
	Response interface{} `json:",omitempty"`
}
