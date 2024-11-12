package key

import (
	"crypto/ed25519"
	"encoding/base64"
)

func GenerateKeyPair() (map[string]string, error) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	keyPair := map[string]string{
		"pubkey":  base64.StdEncoding.EncodeToString(pub),
		"privkey": base64.StdEncoding.EncodeToString(priv),
	}
	return keyPair, nil
}
