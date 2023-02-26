package protocol

import (
	"errors"
	"time"

	util "github.com/elacity/crypto-protocol/internal"
)

type Session struct {
	r       uint64
	outcome chan Response
	payload interface{}
}

func CreateRequestWith(payload interface{}) *Session {
	out := make(chan Response, 1)
	return &Session{
		r:       util.RandomNum(),
		outcome: out,
		payload: payload,
	}
}

func (s *Session) doRequest() {
	time.Sleep(3 * time.Second)
	s.outcome <- NewLicenseError(errors.New("not implemented"))
}

func (s *Session) Output() <-chan Response {
	go s.doRequest()
	return s.outcome
}
