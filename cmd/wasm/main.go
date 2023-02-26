package main

import (
	"github.com/elacity/crypto-protocol/log"
	"github.com/elacity/crypto-protocol/protocol"
	"github.com/elacity/crypto-protocol/runtime"
)

func main() {
	var keepForever = make(chan int)

	log.Infof("Initiating DRM Cryptographic Protocol... version %s", protocol.Version)
	runtime.Register()

	// keep process running
	<-keepForever
	log.Warn("Exiting...")
}
