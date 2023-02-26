package protocol

import (
	"github.com/elacity/crypto-protocol/log"
)

// SeLLEx stands for "Secure Line for License Exchange"
type Sellex struct {
	addr string
	rpc  string
}

type SellexOptions struct{}

// NewSellex creates a new SeLLEx instance
func NewSellex(addr string, rpc string, opt ...SellexOptions) *Sellex {
	log.Tracef("Initializing SeLLEx... {addr: %s, rpc: %s, opt: %+v}", addr, rpc, opt)
	return &Sellex{
		addr: addr,
		rpc:  rpc,
	}
}
