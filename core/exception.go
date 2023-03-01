package core

import (
	"encoding/json"
)

type LicenseError struct {
	err error
}

func (e *LicenseError) Bytes() []byte {
	b, _ := json.Marshal(map[string]interface{}{
		"message": e.err.Error(),
		"code":    -1,
	})
	return b
}

func (e *LicenseError) Code() int {
	return 401
}

func NewLicenseError(err error) *LicenseError {
	return &LicenseError{err}
}

type UnexpectedError struct {
	err  error
	code int
}

func (e *UnexpectedError) Bytes() []byte {
	b, _ := json.Marshal(map[string]interface{}{
		"message": e.err.Error(),
		"code":    e.code,
	})
	return b
}

func (e *UnexpectedError) Code() int {
	return e.code
}

func NewUnexpectedError(code int, err error) *UnexpectedError {
	return &UnexpectedError{code: code, err: err}
}

func WrapError(err error, code ...int) <-chan Response {
	if len(code) == 0 {
		return WrapError(err, 500)
	}

	resp := make(chan Response)
	resp <- NewUnexpectedError(code[0], err)

	return resp
}
