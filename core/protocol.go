package core

import (
	_ "encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/elacity/crypto-protocol/contracts"
	"github.com/elacity/crypto-protocol/log"
)

// SeLLEx stands for "Secure Line for License Exchange"
type Sellex struct {
	client    *ethclient.Client
	authority *contracts.AuthorityGateway
}

type SellexOptions struct{}

// NewSellex creates a new SeLLEx instance
func NewSellex(addr string, rpc string, opt ...SellexOptions) (*Sellex, error) {
	log.Tracef("Initializing SeLLEx... {addr: %s, rpc: %s, opt: %+v}", addr, rpc, opt)

	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the blockchain")
	}

	authority, err := contracts.NewAuthorityGateway(common.HexToAddress(addr), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize the authority contract")
	}

	return &Sellex{
		client:    client,
		authority: authority,
	}, nil
}

func (l *Sellex) AwaitFor(p *FrontRequest) <-chan Response {
	// 0. [/] peer has generated keys
	// 1. [x] prove Authority of the remote contract
	// 2. [] build the request to send to remote contract containing
	// 		- protocol version bytes4
	// 		- peer entropy int64
	// 		- the public key bytes64 x 2
	// 		- payload for verification from wallet {.signer, .signature, .message}
	// 3. [] send the request to the remote contract .acquireLicense
	// 4. [] receive the response that should contains
	// 		- protocol version
	// 		- remote entropy
	// 		- public key of the remote
	// 		- the license
	// 5. [x] verify integrity of the response
	// 6. [] decipher the license using peer.ecdh() that should issue the same
	// shared secret used by remote to encrypt the license
	r, _ := createRequestSession(p)

	l.send(r)

	return r.Output()
}

func (l *Sellex) send(s *Session) {
	license, err := l.authority.AcquireLicense(&bind.CallOpts{
		From: s.payload.Signer,
	}, contracts.SignedLicenseRequest{
		Request:   *s.payload.Message,
		Signature: s.payload.Signature,
	})
	if err != nil {
		s.outcome <- NewLicenseError(err)
		return
	}

	s.outcome <- NewLicenseResponse(license)
}
