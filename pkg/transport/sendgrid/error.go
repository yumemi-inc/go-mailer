package sendgrid

import (
	"strings"
)

type Error struct {
	Message string  `json:"message"`
	Field   *string `json:"field,omitempty"`
	Help    *string `json:"help,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

type ErrorResponse struct {
	Errors []Error `json:"error"`
}

func (e *ErrorResponse) Error() string {
	messages := make([]string, 0, len(e.Errors))
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}

	return strings.Join(messages, "\n")
}
