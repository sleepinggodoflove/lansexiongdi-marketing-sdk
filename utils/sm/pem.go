package sm

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func PrivateKeySM(privateKeyStr string) (*sm2.PrivateKey, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("privateKey base64 decode failed: %v", err)
	}
	privateKey, err := x509.ReadPrivateKeyFromHex(hex.EncodeToString(privateKeyBytes))
	if err != nil {
		return nil, fmt.Errorf("read private key from hex failed: %v", err)
	}
	return privateKey, nil
}

func PublicKeySM(publicKeyStr string) (*sm2.PublicKey, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return nil, fmt.Errorf("publicKeyStr base64 decode failed: %v", err)
	}

	publicKey, err := x509.ReadPublicKeyFromHex(hex.EncodeToString(pubKeyBytes))
	if err != nil {
		return nil, fmt.Errorf("read public key from hex failed: %v", err)
	}
	return publicKey, nil
}
