package core

import (
	"github.com/pkg/errors"

	"github.com/elacity/crypto-protocol/contracts"
	"github.com/elacity/crypto-protocol/core/license"
	"github.com/elacity/crypto-protocol/core/peer"
	util "github.com/elacity/crypto-protocol/internal"
)

type Session struct {
	peer    *peer.Peer
	r       uint64
	outcome chan Response
	payload *FrontRequest
	kids    []string
}

func createRequestSession(kids []string, payload *FrontRequest) (*Session, error) {
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
		kids:    kids,
	}, nil
}

func (s *Session) Output() <-chan Response {
	return s.outcome
}

type LicenseResponse struct {
	src *contracts.ILicenseProviderLicense
	dst license.LicenseOutput
}

func (raw *LicenseResponse) Bytes() []byte {
	return raw.dst.Bytes()
}

func (raw *LicenseResponse) Code() int {
	return 200
}

func NewLicenseResponse(kids []string, lic contracts.ILicenseProviderLicense) *LicenseResponse {
	// for now, we only support clearkey
	dst := license.BuildClearKeyResponse(kids, &lic)
	dst.Issuer = lic.Issuer.String()

	return &LicenseResponse{
		src: &lic,
		dst: dst,
	}
}
