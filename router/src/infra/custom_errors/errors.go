package custom_errors

import (
	"encoding/json"
	"errors"
)

type CustomError struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}

func New(source string, message string) CustomError {
	return CustomError{
		Source:  source,
		Message: message,
	}
}

func (h CustomError) Bytes() []byte {
	bytes, _ := json.Marshal(h)
	return bytes
}

func (h CustomError) Err() error {
	return errors.New(h.Message)
}

func (h CustomError) Error() string {
	return h.Message
}
