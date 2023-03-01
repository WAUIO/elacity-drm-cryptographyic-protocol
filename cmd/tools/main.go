package main

import (
	"bytes"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	util "github.com/elacity/crypto-protocol/internal"
	"github.com/elacity/crypto-protocol/log"
)

var tool = flag.String("tool", "", "The tool to use.")
var pkStr = flag.String("pk", "", "private key.")
var remoteX = flag.String("remote:X", "", "X coordinates of remote public key")
var remoteY = flag.String("remote:Y", "", "Y coordinates of remote public key")

func main() {
	flag.Parse()

	fmt.Println("EC tools for development purposes.")
	fmt.Println("--------------======--------------")

	switch *tool {
	case "keygen":
		if err := keygen(); err != nil {
			log.Errorf("failed to generate key; %s", err)
		}
	case "dh":
		dhPk, _ := hex.DecodeString(*pkStr)
		xStr, _ := hex.DecodeString(*remoteX)
		yStr, _ := hex.DecodeString(*remoteY)
		pub, _ := ecdh.P256().NewPublicKey(append([]byte{0x04}, append(xStr, yStr...)...))
		priv, _ := ecdh.P256().NewPrivateKey(dhPk)
		if err := diffieHellmanSharedKey(priv, pub); err != nil {
			log.Errorf("failed to generate key; %s", err)
		}
	case "test":
		pkBytes, _ := hex.DecodeString(*pkStr)
		if err := test(pkBytes); err != nil {
			log.Errorf("failed to generate key; %s", err)
		}
	}
}

func diffieHellmanSharedKey(pk *ecdh.PrivateKey, pub *ecdh.PublicKey) error {
	// generate shared secret
	shared, err := pk.ECDH(pub)
	if err != nil {
		return errors.Wrap(err, "failed to generate shared secret")
	}

	log.Tracef("shared secret: %x", shared)

	return nil
}

func keygen() error {
	// generate Diffie Hellman private key
	entropy := util.RandomSrc()

	// ECDSA key
	pk0, _ := ecdsa.GenerateKey(elliptic.P256(), entropy)

	// secp256r1
	pk1, err := pk0.ECDH()
	if err != nil {
		return errors.Wrap(err, "failed to generate DH key")
	}

	log.Tracef("pk: %x", pk1.Bytes())

	// ECDH
	// here public key is compressed: 0x04e5ea486c859ab4ece83f921bb2b9c5987b240cc36a1499a9159089fafdd7f4b9818fcb894d578ab290f60acbf918502aefab1bcf1466b9a6cb91cebcc87a0505
	// -> 0x04
	// -> X: 0xe5ea486c859ab4ece83f921bb2b9c5987b240cc36a1499a9159089fafdd7f4b9
	// -> Y: 0x818fcb894d578ab290f60acbf918502aefab1bcf1466b9a6cb91cebcc87a0505
	log.Tracef("ECDH pubkey coordinates: (%x, %x)", pk1.PublicKey().Bytes()[1:33], pk1.PublicKey().Bytes()[33:])

	// sign
	hash := crypto.Keccak256Hash([]byte("hello!"))
	m, err := pk0.Sign(entropy, hash.Bytes(), nil)
	if err != nil {
		return errors.Wrap(err, "failed to sign")
	}
	log.Infof("hash: %x", hash.Bytes())
	log.Infof("signature: %x", m)

	return nil
}

// this use S256 (secp256k1) curve (which is not yet supported by crypto/ecdh)
func test(pkBytes []byte) error {
	privateKey, err := crypto.ToECDSA(pkBytes)
	if err != nil {
		return errors.Wrap(err, "failed to parse private key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return err
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return err
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		return err
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true

	return nil
}
