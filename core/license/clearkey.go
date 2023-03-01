package license

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/elacity/crypto-protocol/contracts"
)

func strPtr(str string) *string {
	return &str
}

type JsonWebKey struct {
	Key string  `json:"k"`
	Kty string  `json:"kty"`
	Kid string  `json:"kid"`
	Alg *string `json:"alg,omitempty"`
}

type ClearKeyResponse struct {
	Keys   []JsonWebKey `json:"keys"`
	Type   string       `json:"type"`
	Issuer string       `json:"issuer"`
	Exp    uint64       `json:"exp,omitempty"`
}

func (c *ClearKeyResponse) Bytes() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		return nil
	}

	return b
}

func BuildClearKeyResponse(kids []string, lic *contracts.ILicenseProviderLicense) *ClearKeyResponse {
	keyset := &ClearKeyResponse{
		Keys: []JsonWebKey{},
		Type: "temporary",
		// expires withing 4 hours
		Exp: uint64(time.Now().Add(4 * time.Hour).Unix()),
	}

	for k, kid := range kids {
		key := base64.StdEncoding.EncodeToString(lic.Key[k][:])
		key = strings.Replace(key, "+", "-", -1) // 62nd char of encoding
		key = strings.Replace(key, "/", "_", -1) // 63rd char of encoding
		key = strings.Replace(key, "=", "", -1)  // Remove any trailing '='s

		ck := JsonWebKey{
			Key: key,
			Kty: "oct",
			Kid: kid,
		}
		ck.Alg = strPtr("A128KW")

		keyset.Keys = append(keyset.Keys, ck)
	}

	return keyset
}
