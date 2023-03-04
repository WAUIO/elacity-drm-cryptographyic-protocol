//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/elacity/crypto-protocol/core"
	"github.com/elacity/crypto-protocol/log"
	"github.com/elacity/crypto-protocol/runtime"
)

func main() {
	var keepForever = make(chan int)

	log.Infof("Initiating DRM Cryptographic Protocol... version %s", core.Version)
	runtime.Register()

	// keep process running
	<-keepForever
	log.Warn("Exiting...")
}
