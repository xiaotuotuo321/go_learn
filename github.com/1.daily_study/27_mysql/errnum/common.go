package errnum

import (
	"fmt"
)

// internal error
type Er struct {
	Code    int
	Message string
}

func (e *Er) Error() string {
	return e.Message
}

// api error
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func New(er *Er, err error) *Err {
	return &Err{Code: er.Code, Message: er.Message, Err: err}
}
