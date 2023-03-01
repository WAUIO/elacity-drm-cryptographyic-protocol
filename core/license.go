package core

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/elacity/crypto-protocol/contracts"
	"github.com/elacity/crypto-protocol/core/peer"
	util "github.com/elacity/crypto-protocol/internal"
)

type Session struct {
	peer    *peer.Peer
	r       uint64
	outcome chan Response
	payload *FrontRequest
}

func createRequestSession(payload *FrontRequest) (*Session, error) {
	out := make(chan Response, 1)

	peer, err := peer.NewPeer()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to initialize the peer")
	}

	return &Session{
		peer:    peer,
		r:       util.RandomNum(),
		outcome: out,
		payload: payload,
	}, nil
}

func (s *Session) Output() <-chan Response {
	return s.outcome
}

type LicenseResponse struct {
	contracts.ILicenseProviderLicense
}

func (raw *LicenseResponse) Bytes() []byte {
	b, err := json.Marshal(raw)
	if err != nil {
		return nil
	}

	return b
}

func (raw *LicenseResponse) Code() int {
	return 200
}

func NewLicenseResponse(license contracts.ILicenseProviderLicense) *LicenseResponse {
	return &LicenseResponse{
		ILicenseProviderLicense: license,
	}
}
