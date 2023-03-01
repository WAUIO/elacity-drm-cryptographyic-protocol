package core

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/elacity/crypto-protocol/contracts"
)

type PsshBody struct {
	Type    string `json:"protectionType"`
	Variant string `json:"variant"`
	Data    struct {
		Authority string `json:"authority"`
		ChainId   uint16 `json:"chainId"`
		RpcURL    string `json:"rpc"`
	} `json:"data"`
}

type LicensePayload struct {
	Kids []string `json:"kids"`
	Type string   `json:"type"` // "temporary" or "permanent"
}

type RawFrontRequest struct {
	Signer    string `json:"signer"`
	Signature string `json:"signature"`
	Message   struct {
		Entitlement string `json:"entitlement"`
		Entity      struct {
			ContentId string `json:"contentId"`
			Ledger    string `json:"ledger"`
			TokenId   string `json:"tokenId"`
		} `json:"entity"`
	} `json:"message"`
}

type FrontRequest struct {
	Signer    common.Address
	Signature []byte
	Message   *contracts.LicenseRequest
}

type Request struct {
	Pssh    PsshBody
	Payload LicensePayload
	Request *FrontRequest
}

func RequestLicense(r *Request) <-chan Response {
	sl, err := NewSellex(r.Pssh.Data.Authority, r.Pssh.Data.RpcURL)
	if err != nil {
		return WrapError(err, 500)
	}

	return sl.AwaitFor(r.Payload.Kids, r.Request)
}
