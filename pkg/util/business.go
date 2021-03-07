package util

func NewBusinessError(businessCode int, httpStatusCode int, message string, reason error) *BusinessError {
	return &BusinessError{BusinessCode: businessCode, HTTPStatusCode: httpStatusCode, Message: message, Reason: reason}
}

type BusinessError struct {
	BusinessCode     int               `json:"code"`
	HTTPStatusCode   int               `json:"-"`
	Message          string            `json:"message"`
	ValidationErrors map[string]string `json:"validationErrors,omitempty"`
	Reason           error             `json:"-"`
}

func (b *BusinessError) Error() string {
	return b.Reason.Error()
}
