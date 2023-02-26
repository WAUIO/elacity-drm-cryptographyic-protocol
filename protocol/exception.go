package protocol

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
