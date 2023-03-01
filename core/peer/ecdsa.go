package peer

type Signer interface {
	Sign(message []byte) ([]byte, error)
}

type Verifier interface {
	Erecover(message []byte, sig []byte) (pub []byte, err error)
	Verify(message []byte, sig []byte) (bool, error)
}
