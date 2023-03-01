package peer

type Encrypter interface {
	Encrypt(plaintext []byte) ([]byte, error)
}

type Decrypter interface {
	Decrypt(ciphertext []byte) ([]byte, error)
}
