package core

import (
	"encoding/hex"
	"math/big"

	"github.com/elacity/crypto-protocol/contracts"
)

func FromRawToComprehensive(r *RawFrontRequest) (*FrontRequest, error) {
	ir := new(FrontRequest)

	if signer, err := hex.DecodeString(r.Signer[2:]); err == nil {
		if len(signer) == 20 {
			ir.Signer = [20]byte(signer[:20])
		}
	}

	if sig, err := hex.DecodeString(r.Signature[2:]); err == nil {
		ir.Signature = sig
	}

	m := new(contracts.LicenseRequest)
	m.Entitlement = r.Message.Entitlement
	m.Entity = contracts.IPEntity{}

	if ledgerAddr, err := hex.DecodeString(r.Message.Entity.Ledger[2:]); err == nil {
		if len(ledgerAddr) == 20 {
			m.Entity.Ledger = [20]byte(ledgerAddr[:20])
		}
	}

	if contentId, err := hex.DecodeString(r.Message.Entity.ContentId[2:]); err == nil {
		if len(contentId) == 16 {
			m.Entity.ContentId = [16]byte(contentId[:16])
		}
	}

	tokenId := new(big.Int)
	tk, ok := tokenId.SetString(r.Message.Entity.TokenId, 10)
	if ok {
		m.Entity.TokenId = tk
	}

	ir.Message = m

	return ir, nil
}
