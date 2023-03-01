package peer

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/rand"

	util "github.com/elacity/crypto-protocol/internal"
)

type Peer struct {
	entropy *rand.Rand

	// private keys
	// pk0 represents the privte in ECDSA usage format
	pk0 *ecdsa.PrivateKey

	// pk1 represent the same private key in ECDH usage format
	pk1 *ecdh.PrivateKey
}

func NewPeer() (*Peer, error) {
	entropy := rand.New(util.RandomSrc())

	// ECDSA key (secp256k1: used for signing)
	pk0, err := ecdsa.GenerateKey(elliptic.P256(), entropy)
	if err != nil {
		return nil, err
	}

	// secp256r1: generate Diffie Hellman private key
	pk1, err := pk0.ECDH()
	if err != nil {
		return nil, err
	}

	return &Peer{
		entropy: entropy,
		pk0:     pk0,
		pk1:     pk1,
	}, nil
}

func (p *Peer) Decrypt() error {
	return nil
}

func (p *Peer) ecdh(pub *ecdh.PublicKey) ([]byte, error) {
	// generate shared secret
	shared, err := p.pk1.ECDH(pub)
	if err != nil {
		return nil, err
	}

	return shared, nil
}
