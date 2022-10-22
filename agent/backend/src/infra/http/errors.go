package http

type Error struct {
	Code    string `json:"internal_code"`
	Message string `json:"message"`
}

func newError(code string, message string) Error {
	return Error{Code: code, Message: message}
}
